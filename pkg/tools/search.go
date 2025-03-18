package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/douglarek/unsplash-mcp-server/internal/api"
	"github.com/douglarek/unsplash-mcp-server/internal/config"
	"github.com/mark3labs/mcp-go/mcp"
)

// NewSearchPhotosTool creates and returns the search photos tool definition
func NewSearchPhotosTool() mcp.Tool {
	return mcp.NewTool("search_photos",
		mcp.WithDescription("Search for Unsplash photos"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search keyword"),
		),
		mcp.WithNumber("page",
			mcp.Required(),
			mcp.DefaultNumber(1),
			mcp.Description("Page number (1-based)"),
		),
		mcp.WithNumber("per_page",
			mcp.Required(),
			mcp.DefaultNumber(5),
			mcp.Description("Results per page (1-30)"),
		),
		mcp.WithString("order_by",
			mcp.Required(),
			mcp.DefaultString("relevant"),
			mcp.Description("Sort method (relevant or latest)"),
		),
		mcp.WithString("color",
			mcp.Description("Color filter (black_and_white, black, white, yellow, orange, red, purple, magenta, green, teal, blue)"),
		),
		mcp.WithString("orientation",
			mcp.Description("Orientation filter (landscape, portrait, squarish)"),
		),
	)
}

// HandleSearchPhotos returns a handler function for search photos requests
func HandleSearchPhotos(cfg *config.Config) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Create Unsplash client
		client := api.NewClient(cfg)

		// Extract and validate parameters
		query, ok := request.Params.Arguments["query"].(string)
		if !ok || query == "" {
			return mcp.NewToolResultError("Search keyword must be provided"), nil
		}

		// Build query parameters
		params := buildSearchParams(request.Params.Arguments)

		// Search photos
		photos, err := client.SearchPhotos(ctx, params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to search photos: %v", err)), nil
		}

		// Format results
		b, _ := json.Marshal(photos)
		return mcp.NewToolResultText(string(b)), nil
	}
}

// buildSearchParams builds URL parameters from request arguments
func buildSearchParams(args map[string]interface{}) url.Values {
	params := url.Values{}

	// Required parameters
	params.Add("query", args["query"].(string))
	params.Add("page", fmt.Sprintf("%d", int(args["page"].(float64))))
	params.Add("per_page", fmt.Sprintf("%d", int(args["per_page"].(float64))))
	params.Add("order_by", args["order_by"].(string))

	// Optional parameters
	if color, ok := args["color"]; ok && color != nil {
		params.Add("color", color.(string))
	}

	if orientation, ok := args["orientation"]; ok && orientation != nil {
		params.Add("orientation", orientation.(string))
	}

	return params
}
