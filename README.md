# Unsplash MCP Server

[![smithery badge](https://smithery.ai/badge/@douglarek/unsplash-mcp-server)](https://smithery.ai/server/@douglarek/unsplash-mcp-server)

A rewrite of the [Unsplash MCP Server](https://github.com/hellokaton/unsplash-mcp-server) using the [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) library.

## Usage

Before building, you must install go 1.24+ first.

```bash
git clone https://github.com/douglarek/unsplash-mcp-server.git
cd unsplash-mcp-server
make build
```

### Cursor Editor Integration

To use this server in Cursor, you can add the following to your `mcp.json` file:

```json
{
  "mcpServers": {
    "unsplash": {
      "command": "<source_dir>/cmd/server/unsplash-mcp-server",
      "args": [],
      "env": {
        "UNSPLASH_ACCESS_KEY": "<your_unsplash_access_key>"
      }
    }
  }
}
```
