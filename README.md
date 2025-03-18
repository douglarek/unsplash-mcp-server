# Unsplash MCP Server

A rewrite of the [Unsplash MCP Server](https://github.com/hellokaton/unsplash-mcp-server) using the [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go) library.

## Usage

```bash
go build -o unsplash-mcp-server cmd/server/main.go
```

### Cursor Editor Integration

To use this server in Cursor, you can add the following to your `mcp.json` file:

```json
{
  "mcpServers": {
    "unsplash": {
      "command": "<source_dir>/unsplash-mcp-server",
      "args": [],
      "env": {
        "UNSPLASH_ACCESS_KEY": "<your_unsplash_access_key>"
      }
    }
  }
}
```