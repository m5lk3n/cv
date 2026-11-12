import resume from "../resume/resume.json";

interface JsonRpcRequest {
  jsonrpc: "2.0";
  id?: string | number | null;
  method: string;
  params?: Record<string, unknown>;
}

interface JsonRpcResponse {
  jsonrpc: "2.0";
  id: string | number | null;
  result?: unknown;
  error?: { code: number; message: string; data?: unknown };
}

interface Tool {
  name: string;
  description: string;
  inputSchema: { type: "object"; properties: Record<string, never> };
}

const CORS_HEADERS: Record<string, string> = {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "POST, OPTIONS",
  "Access-Control-Allow-Headers": "Content-Type, Accept, Authorization, Mcp-Session-Id",
};

const PROTOCOL_VERSION = "2025-03-26";

const SERVER_INFO = {
  name: `${resume.basics.name}'s CV`,
  version: "1.0.1",
};

const tools: Tool[] = [
  {
    name: "get_basics",
    description:
      "Get basic profile information: name, professional title, summary, contact details, location, and links to online profiles.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_work_experience",
    description:
      "Get work history including company names, job titles, dates, summaries, and key highlights for each role.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_skills",
    description:
      "Get technical and professional skills grouped by category, with keywords.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_education",
    description:
      "Get educational background including institutions, degrees, fields of study, and dates.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_certificates",
    description:
      "Get professional certifications with issuing organizations and dates.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_languages",
    description: "Get spoken languages and fluency levels.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_projects",
    description:
      "Get side projects and personal work including descriptions, highlights, and links.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_publications",
    description:
      "Get published works including titles, publishers, dates, and summaries.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_references",
    description: "Get professional references and testimonials.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_volunteering",
    description:
      "Get volunteer experience including organizations, roles, and contributions.",
    inputSchema: { type: "object", properties: {} },
  },
  {
    name: "get_faqs",
    description:
      "Get answers to frequently asked questions about this person's professional background and preferences.",
    inputSchema: { type: "object", properties: {} },
  },
];

interface Env {
  AGENT_EMAIL?: string;
}

function callTool(name: string, env: Env): { content: Array<{ type: "text"; text: string }> } | undefined {
  const text = (data: unknown) => ({
    content: [{ type: "text" as const, text: JSON.stringify(data, null, 2) }],
  });

  switch (name) {
    case "get_basics":
      return text({ ...resume.basics, email: env.AGENT_EMAIL ?? resume.basics.email });
    case "get_work_experience":
      return text(resume.work);
    case "get_skills":
      return text(resume.skills);
    case "get_education":
      return text(resume.education);
    case "get_certificates":
      return text(resume.certificates);
    case "get_languages":
      return text(resume.languages);
    case "get_projects":
      return text(resume.projects);
    case "get_publications":
      return text(resume.publications);
    case "get_references":
      return text(resume.references);
    case "get_volunteering":
      return text(resume.volunteer);
    case "get_faqs":
      return text(resume["x-cv"].faqs);
    default:
      return undefined;
  }
}

function handleRequest(req: JsonRpcRequest, env: Env): JsonRpcResponse | null {
  if (!("id" in req) || req.id === undefined) {
    return null;
  }

  switch (req.method) {
    case "initialize":
      return {
        jsonrpc: "2.0",
        id: req.id!,
        result: {
          protocolVersion: PROTOCOL_VERSION,
          capabilities: { tools: {} },
          serverInfo: SERVER_INFO,
        },
      };

    case "ping":
      return { jsonrpc: "2.0", id: req.id!, result: {} };

    case "tools/list":
      return { jsonrpc: "2.0", id: req.id!, result: { tools } };

    case "tools/call": {
      const toolName = (req.params as { name?: string })?.name;
      if (!toolName) {
        return {
          jsonrpc: "2.0",
          id: req.id!,
          error: { code: -32602, message: "Missing tool name" },
        };
      }
      const result = callTool(toolName, env);
      if (!result) {
        return {
          jsonrpc: "2.0",
          id: req.id!,
          error: { code: -32602, message: `Unknown tool: ${toolName}` },
        };
      }
      return { jsonrpc: "2.0", id: req.id!, result };
    }

    default:
      return {
        jsonrpc: "2.0",
        id: req.id!,
        error: { code: -32601, message: `Method not found: ${req.method}` },
      };
  }
}

function jsonResponse(body: unknown, status = 200): Response {
  return new Response(JSON.stringify(body), {
    status,
    headers: { ...CORS_HEADERS, "Content-Type": "application/json" },
  });
}

export async function handleMCP(request: Request, env: Env): Promise<Response> {
  if (request.method === "OPTIONS") {
    return new Response(null, { status: 204, headers: CORS_HEADERS });
  }

  if (request.method === "GET") {
    return jsonResponse({
      ...SERVER_INFO,
      protocol: "MCP",
      protocolVersion: PROTOCOL_VERSION,
      description:
        "This is a Model Context Protocol (MCP) server. POST JSON-RPC requests to this endpoint to interact with it.",
      tools: tools.map((t) => t.name),
    });
  }

  if (request.method !== "POST") {
    return jsonResponse({ error: "Method not allowed" }, 405);
  }

  const contentType = request.headers.get("Content-Type");
  if (!contentType?.includes("application/json")) {
    return jsonResponse(
      {
        jsonrpc: "2.0",
        id: null,
        error: { code: -32700, message: "Content-Type must be application/json" },
      },
      400,
    );
  }

  let body: unknown;
  try {
    body = await request.json();
  } catch {
    return jsonResponse(
      { jsonrpc: "2.0", id: null, error: { code: -32700, message: "Parse error" } },
      400,
    );
  }

  if (Array.isArray(body)) {
    const responses = (body as JsonRpcRequest[])
      .map((req) => handleRequest(req, env))
      .filter((r): r is JsonRpcResponse => r !== null);
    if (responses.length === 0) {
      return new Response(null, { status: 204, headers: CORS_HEADERS });
    }
    return jsonResponse(responses);
  }

  const response = handleRequest(body as JsonRpcRequest, env);
  if (response === null) {
    return new Response(null, { status: 204, headers: CORS_HEADERS });
  }
  return jsonResponse(response);
}
