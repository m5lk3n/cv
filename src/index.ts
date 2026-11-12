import { handleMCP } from "./mcp";

interface Env {
  ASSETS: Fetcher;
}

export default {
  async fetch(request: Request, env: Env): Promise<Response> {
    const url = new URL(request.url);

    if (url.pathname === "/mcp") {
      return handleMCP(request);
    }

    if (url.pathname === "/.well-known/mcp.json") {
      return new Response(
        JSON.stringify({
          mcpServers: {
            cv: {
              url: `${url.origin}/mcp`,
              type: "http",
            },
          },
        }),
        {
          headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
          },
        },
      );
    }

    if (
      url.pathname.startsWith("/.well-known/") ||
      url.pathname === "/register" ||
      url.pathname === "/authorize" ||
      url.pathname === "/token"
    ) {
      return new Response(null, { status: 404 });
    }

    return env.ASSETS.fetch(request);
  },
} satisfies ExportedHandler<Env>;
