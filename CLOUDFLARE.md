# Cloudflare Configuration

**This is a recommendation, use at your own risk!**

This project deploys to Cloudflare Workers (free tier) via the
[Makefile](Makefile)'s `publish-to-web` target and [wrangler.jsonc](wrangler.jsonc).
The rules below add defense-in-depth on the Cloudflare side against the
attack patterns observed in my access logs (October 2026).

All rules are configured under the managed domain's **Security >
Security Rules > Custom rules** in the Cloudflare dashboard.

## Rule 1 — Block Empty User-Agent From Cloud Providers

The largest attack cluster is an automated botnet running on Azure VMs:
57 IPs, 11,000+ requests, all with an empty `User-Agent`, probing for
PHP webshells and WordPress installations. Block empty-UA requests from
major cloud ASNs outright — no legitimate browser omits the header.

| Field | Value |
| :--- | :--- |
| **Rule name** | Block empty UA from cloud ASNs |
| **Action** | Block |

**Expression (Edit expression):**

```text
(not http.user_agent ne "" and ip.geoip.asnum in {8075 13335 14618 15169 16509 36351 396982})
```

The expression reads: UA is empty **and** the source IP belongs to one
of these ASNs:

| ASN | Provider |
| :--- | :--- |
| 8075 | Microsoft Azure |
| 13335 | Cloudflare |
| 14618 | IBM Cloud |
| 15169 | Google Cloud |
| 16509 | Amazon AWS |
| 36351 | SoftLayer (IBM) |
| 396982 | Google Cloud (secondary) |

> **Why not block empty UA globally?** Some monitoring agents (such as
> health checks from your own infrastructure) or IoT clients send
> requests without a `User-Agent`. Limiting the block to cloud ASNs
> avoids false positives while catching the botnet pattern.

## Rule 2 — Block PHP and WordPress Probing

The site serves no PHP and has no WordPress installation. Requests for
`.php` files or `wp-*` paths are always malicious in this context.

| Field | Value |
| :--- | :--- |
| **Rule name** | Block PHP and WordPress probes |
| **Action** | Block |

**Expression:**

```text
(http.request.uri.path contains ".php" or http.request.uri.path contains "/wp-admin" or http.request.uri.path contains "/wp-content" or http.request.uri.path contains "/wp-includes" or http.request.uri.path contains "/wp-login" or http.request.uri.path contains "/xmlrpc")
```

## Rule 3 — Block Credential and Config File Probing

Attackers enumerate common secret and configuration file paths. None of
these files exist, but blocking them at the edge avoids unnecessary
origin traffic and keeps these paths out of the logs.

| Field | Value |
| :--- | :--- |
| **Rule name** | Block credential and config probes |
| **Action** | Block |

**Expression:**

```text
(http.request.uri.path contains "/.env" or http.request.uri.path contains "/.git" or http.request.uri.path contains "/.aws" or http.request.uri.path contains "/.docker" or http.request.uri.path contains "/.ssh" or http.request.uri.path contains "/.npmrc" or http.request.uri.path contains "/credentials.json" or http.request.uri.path contains "/config.json" or http.request.uri.path contains "/service-account" or http.request.uri.path contains "/docker-compose" or http.request.uri.path contains "/.gitlab-ci")
```

## Rule 4 — Block Path Traversal Attempts

Catch `../` sequences that attempt to escape the web root.

| Field | Value |
| :--- | :--- |
| **Rule name** | Block path traversal |
| **Action** | Block |

**Expression:**

```text
(http.request.uri.path contains ".." or http.request.uri contains "%2e%2e" or http.request.uri contains "%252e")
```

## Rule 5 — Challenge Suspicious Methods

The site is entirely static. Only `GET` and `HEAD` are legitimate.

| Field | Value |
| :--- | :--- |
| **Rule name** | Challenge non-GET/HEAD methods |
| **Action** | Managed Challenge |

**Expression:**

```text
(not http.request.method in {"GET" "HEAD"})
```

> Using a managed challenge rather than a hard block allows `OPTIONS`
> preflight requests from browsers to pass after solving.

## Rule Priority

Cloudflare evaluates custom rules top to bottom. Place the rules in the
order listed above so the cheap ASN + UA check fires first and the
method challenge is a catch-all at the bottom.

## WAF Verification

After deploying the rules, verify they work as expected:

```sh
# Should be blocked (empty UA from a cloud ASN — run from a cloud VM):
curl -s -o /dev/null -w "%{http_code}" -A "" https://cv.michael-klein.info/

# Should be blocked (PHP probe):
curl -s -o /dev/null -w "%{http_code}" https://cv.michael-klein.info/test.php

# Should be blocked (env file probe):
curl -s -o /dev/null -w "%{http_code}" https://cv.michael-klein.info/.env

# Should be blocked (path traversal):
curl -s -o /dev/null -w "%{http_code}" https://cv.michael-klein.info/assets../.env

# Should succeed (normal request):
curl -s -o /dev/null -w "%{http_code}" https://cv.michael-klein.info/
```

Blocked requests return a `403` from Cloudflare and never reach the
origin.

## Security Headers

The [Caddyfile](CADDY.md) sets security headers via its
`(security_headers)` snippet. When serving through Cloudflare instead
of Caddy, replicate them with a rule under the managed domain's
**Rules > Overview > Response Header Transform Rules**.

| Field | Value |
| :--- | :--- |
| **Rule name** | Security headers |
| **When** | All incoming requests (or: `(http.host eq "cv.michael-klein.info")`) |

Add each header as a **Set static** action:

| Header | Value |
| :--- | :--- |
| `X-Content-Type-Options` | `nosniff` |
| `X-Frame-Options` | `DENY` |
| `Referrer-Policy` | `strict-origin-when-cross-origin` |
| `Permissions-Policy` | `interest-cohort=()` |
| `Content-Security-Policy` | `default-src 'self'; script-src 'self' 'unsafe-inline' 'wasm-unsafe-eval'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; connect-src 'self'; object-src 'none'; base-uri 'none'; frame-ancestors 'none';` |

> **Set** (not **Add**) ensures each header appears exactly once, even
> if the origin already sends one.

### What each header does

- **`X-Content-Type-Options: nosniff`** — prevents browsers from
  MIME-sniffing a response away from its declared `Content-Type`,
  blocking attacks that trick the browser into executing a file as
  script.
- **`X-Frame-Options: DENY`** — prevents the site from being embedded
  in an `<iframe>`, mitigating clickjacking. Reinforced by the
  `frame-ancestors 'none'` directive in CSP.
- **`Referrer-Policy: strict-origin-when-cross-origin`** — sends the
  full URL as referrer for same-origin requests but only the origin
  (no path) for cross-origin ones, limiting information leakage.
- **`Permissions-Policy: interest-cohort=()`** — opts out of
  Federated Learning of Cohorts (FLoC), Google's deprecated
  interest-based tracking proposal. Harmless to keep as a signal.
- **`Content-Security-Policy`** — allowlists exactly what the browser
  may load. `'unsafe-inline'` is needed for the inline terminal
  scripts; `'wasm-unsafe-eval'` allows the Go WASM module to run.
  Everything else is restricted to same-origin or blocked entirely.

### Verification

```sh
curl -sI https://cv.michael-klein.info/ | grep -iE "x-content-type|x-frame|referrer-policy|permissions-policy|content-security"
```

All five headers should appear in the response.
