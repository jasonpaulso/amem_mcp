package memory

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/amem/mcp-server/pkg/models"
	"github.com/amem/mcp-server/pkg/services"
	"go.uber.org/zap"
)

// WorkspaceInitTool implements smart workspace initialization
type WorkspaceInitTool struct {
	workspaceService *services.WorkspaceService
	logger           *zap.Logger
}

// NewWorkspaceInitTool creates a new workspace init tool
func NewWorkspaceInitTool(workspaceService *services.WorkspaceService, logger *zap.Logger) *WorkspaceInitTool {
	return &WorkspaceInitTool{
		workspaceService: workspaceService,
		logger:           logger,
	}
}

func (t *WorkspaceInitTool) Name() string {
	return "workspace_init"
}

func (t *WorkspaceInitTool) Description() string {
	return "Smart workspace initialization - creates new workspace or retrieves existing one. If no identifier provided, uses current working directory."
}

func (t *WorkspaceInitTool) InputSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"identifier": map[string]interface{}{
				"type":        "string",
				"description": "Path or name for the workspace. If not provided, uses current working directory",
			},
			"name": map[string]interface{}{
				"type":        "string",
				"description": "Human-readable name for the workspace (optional)",
			},
		},
		"required": []string{},
	}
}

func (t *WorkspaceInitTool) Execute(ctx context.Context, args map[string]interface{}) (*models.MCPToolResult, error) {
	// Parse arguments
	var req models.WorkspaceRequest

	if identifier, ok := args["identifier"].(string); ok {
		req.Identifier = identifier
	}

	if name, ok := args["name"].(string); ok {
		req.Name = name
	}

	// Initialize workspace
	workspace, created, err := t.workspaceService.InitializeWorkspace(ctx, &req)
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error initializing workspace: %v", err),
			}},
		}, nil
	}

	// Create response
	response := models.WorkspaceResponse{
		Workspace: *workspace,
		Created:   created,
	}

	action := "Retrieved"
	if created {
		action = "Created"
	}

	// Serialize response to JSON
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error serializing response: %v", err),
			}},
		}, nil
	}

	return &models.MCPToolResult{
		IsError: false,
		Content: []models.MCPContent{
			{
				Type: "text",
				Text: fmt.Sprintf("%s workspace '%s' (%s)\n\nWorkspace Details:\n```json\n%s\n```",
					action, workspace.Name, workspace.ID, string(responseJSON)),
			},
		},
	}, nil
}

// WorkspaceCreateTool implements explicit workspace creation
type WorkspaceCreateTool struct {
	workspaceService *services.WorkspaceService
	logger           *zap.Logger
}

// NewWorkspaceCreateTool creates a new workspace create tool
func NewWorkspaceCreateTool(workspaceService *services.WorkspaceService, logger *zap.Logger) *WorkspaceCreateTool {
	return &WorkspaceCreateTool{
		workspaceService: workspaceService,
		logger:           logger,
	}
}

func (t *WorkspaceCreateTool) Name() string {
	return "workspace_create"
}

func (t *WorkspaceCreateTool) Description() string {
	return "Explicit workspace creation - fails if workspace already exists. Supports both filesystem paths and logical names."
}

func (t *WorkspaceCreateTool) InputSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"identifier": map[string]interface{}{
				"type":        "string",
				"description": "Path or name for the workspace (required)",
			},
			"name": map[string]interface{}{
				"type":        "string",
				"description": "Human-readable name for the workspace (optional)",
			},
			"description": map[string]interface{}{
				"type":        "string",
				"description": "Description of the workspace (optional)",
			},
		},
		"required": []string{"identifier"},
	}
}

func (t *WorkspaceCreateTool) Execute(ctx context.Context, args map[string]interface{}) (*models.MCPToolResult, error) {
	// Parse arguments
	var req models.WorkspaceRequest

	if identifier, ok := args["identifier"].(string); ok {
		req.Identifier = identifier
	} else {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: "Error: 'identifier' parameter is required",
			}},
		}, nil
	}

	if name, ok := args["name"].(string); ok {
		req.Name = name
	}

	if description, ok := args["description"].(string); ok {
		req.Description = description
	}

	// Create workspace
	workspace, err := t.workspaceService.CreateWorkspace(ctx, &req)
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error creating workspace: %v", err),
			}},
		}, nil
	}

	// Create response
	response := models.WorkspaceResponse{
		Workspace: *workspace,
		Created:   true,
	}

	// Serialize response to JSON
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error serializing response: %v", err),
			}},
		}, nil
	}

	return &models.MCPToolResult{
		IsError: false,
		Content: []models.MCPContent{
			{
				Type: "text",
				Text: fmt.Sprintf("Created workspace '%s' (%s)\n\nWorkspace Details:\n```json\n%s\n```",
					workspace.Name, workspace.ID, string(responseJSON)),
			},
		},
	}, nil
}

// WorkspaceRetrieveTool implements explicit workspace retrieval
type WorkspaceRetrieveTool struct {
	workspaceService *services.WorkspaceService
	logger           *zap.Logger
}

// NewWorkspaceRetrieveTool creates a new workspace retrieve tool
func NewWorkspaceRetrieveTool(workspaceService *services.WorkspaceService, logger *zap.Logger) *WorkspaceRetrieveTool {
	return &WorkspaceRetrieveTool{
		workspaceService: workspaceService,
		logger:           logger,
	}
}

func (t *WorkspaceRetrieveTool) Name() string {
	return "workspace_retrieve"
}

func (t *WorkspaceRetrieveTool) Description() string {
	return "Explicit workspace retrieval - fails if workspace doesn't exist. Returns comprehensive workspace metadata including memory count."
}

func (t *WorkspaceRetrieveTool) InputSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"identifier": map[string]interface{}{
				"type":        "string",
				"description": "Path or name of the workspace to retrieve (required)",
			},
		},
		"required": []string{"identifier"},
	}
}

func (t *WorkspaceRetrieveTool) Execute(ctx context.Context, args map[string]interface{}) (*models.MCPToolResult, error) {
	// Parse arguments
	var req models.WorkspaceRequest

	if identifier, ok := args["identifier"].(string); ok {
		req.Identifier = identifier
	} else {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: "Error: 'identifier' parameter is required",
			}},
		}, nil
	}

	// Normalize workspace ID
	workspaceID := t.workspaceService.NormalizeWorkspaceID(req.Identifier)

	// Check if workspace exists
	exists, err := t.workspaceService.WorkspaceExists(ctx, workspaceID)
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error checking workspace existence: %v", err),
			}},
		}, nil
	}

	if !exists {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Workspace '%s' does not exist", workspaceID),
			}},
		}, nil
	}

	// Get workspace info
	workspace, err := t.workspaceService.GetWorkspaceInfo(ctx, workspaceID)
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error retrieving workspace info: %v", err),
			}},
		}, nil
	}

	// Create response
	response := models.WorkspaceResponse{
		Workspace: *workspace,
		Created:   false,
	}

	// Serialize response to JSON
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return &models.MCPToolResult{
			IsError: true,
			Content: []models.MCPContent{{
				Type: "text",
				Text: fmt.Sprintf("Error serializing response: %v", err),
			}},
		}, nil
	}

	return &models.MCPToolResult{
		IsError: false,
		Content: []models.MCPContent{
			{
				Type: "text",
				Text: fmt.Sprintf("Retrieved workspace '%s' (%s) with %d memories\n\nWorkspace Details:\n```json\n%s\n```",
					workspace.Name, workspace.ID, workspace.MemoryCount, string(responseJSON)),
			},
		},
	}, nil
}
