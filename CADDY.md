# Caddyserver Configuration

**This is a recommendation, use at your own risk!**

## Custom 404

As a prerequisite for the `Caddyfile` below, place this `404.html` in the matching root folder:

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Not Found</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style>
    body {
      font-family: system-ui, sans-serif;
      background: #fff;
      color: #222;
      margin: 0;
      padding: 3rem;
    }
    h1 {
      font-size: 1.25rem;
      margin-bottom: 0.5rem;
    }
    p {
      font-size: 0.95rem;
      color: #555;
    }
  </style>
</head>
<body>
  <h1>404 – Not Found</h1>
  <p>The requested resource could not be found.</p>
</body>
</html>
```

## Exemplified `Caddyfile`

What it does?

- Blocks common exploit paths
- Reduces fingerprinting value
- Adds security headers (with strong but safe defaults)

```
(blocked_paths) {
	@blocked_paths {
		path /.env*
		path /.git/*
		path /wp-admin*
		path /wp-includes*
		path /wp-content*
		path /server*
		path /console*
	}
	respond @blocked_paths "Access Denied" 403
}

(common_errors) {
    handle_errors {
        @404 expression {err.status_code} == 404
        
        handle @404 {
            root * /path/to/caddy/common
            rewrite * /404.html
            file_server
        }
    }
}

(security_headers) {
    header {
        X-Content-Type-Options nosniff
        X-Frame-Options DENY
        Referrer-Policy strict-origin-when-cross-origin
        Permissions-Policy "interest-cohort=()"
        Content-Security-Policy `
            default-src 'self';
            script-src 'self';
            style-src 'self';
            img-src 'self' data:;
            font-src 'self';
            connect-src 'self';
            object-src 'none';
            base-uri 'none';
            frame-ancestors 'none';
        `
    }
}

cv.example.com {
    import blocked_paths
    import security_headers
    import common_errors
    log {
        format console
        output file /path/to/caddy/logs/site.log {
            roll_size 100mb
            roll_keep 20
            roll_keep_for 7d
        }
    }

    root * /path/to/caddy/sites/cv.example.com
    file_server
}
```

## Further Improvements

### Rate-limit scanners & bots

Requires
- [caddy-ratelimit](https://github.com/mholt/caddy-ratelimit)
- [xcaddy](https://github.com/caddyserver/xcaddy) 
- `caddy` to be updated separately (due to custom `xcaddy` build)

```
@bots {
    header User-Agent *l9scan*
    header User-Agent *Go-http-client*
    header User-Agent *HeadlessChrome*
}
rate_limit @bots {
    zone bots
    events 10
    window 1m
}
```
