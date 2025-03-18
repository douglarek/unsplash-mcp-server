package main

import (
	"log/slog"
	"os"

	"github.com/douglarek/unsplash-mcp-server/internal/config"
	"github.com/douglarek/unsplash-mcp-server/pkg/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Create MCP server
	s := server.NewMCPServer(
		"Unsplash MCP Server",
		"0.1.0",
	)

	// Add search tool
	searchTool := tools.NewSearchPhotosTool()

	// Register tool handler
	s.AddTool(searchTool, tools.HandleSearchPhotos(cfg))

	// Start stdio server
	if err := server.ServeStdio(s); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}
}
