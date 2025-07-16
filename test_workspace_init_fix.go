package main

import (
	"context"
	"fmt"
	"log"

	"github.com/amem/mcp-server/pkg/config"
	"github.com/amem/mcp-server/pkg/memory"
	"github.com/amem/mcp-server/pkg/services"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig("config/production.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize services
	ctx := context.Background()

	// Initialize ChromaDB service
	chromaService := services.NewChromaDBService(cfg.ChromaDB, logger.Named("chromadb"))
	if err := chromaService.Initialize(ctx); err != nil {
		log.Fatalf("Failed to initialize ChromaDB: %v", err)
	}

	// Initialize workspace service
	workspaceService := services.NewWorkspaceService(chromaService, logger.Named("workspace"))

	// Test workspace_init tool
	fmt.Println("🧪 Testing workspace_init Fix")
	fmt.Println("=" * 40)

	// Create workspace_init tool
	workspaceInitTool := memory.NewWorkspaceInitTool(workspaceService, logger.Named("workspace_init"))

	fmt.Printf("✅ Created workspace_init tool: %s\n", workspaceInitTool.Name())
	fmt.Printf("📝 Description: %s\n", workspaceInitTool.Description())

	// Test 1: workspace_init with no parameters (should use default)
	fmt.Println("\n1. Testing workspace_init with no parameters...")
	initArgs := map[string]interface{}{}
	initResult, err := workspaceInitTool.Execute(ctx, initArgs)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Success! IsError: %v\n", initResult.IsError)
		fmt.Printf("📄 Content count: %d\n", len(initResult.Content))
		for i, content := range initResult.Content {
			fmt.Printf("   Content %d - Type: %s\n", i+1, content.Type)
			fmt.Printf("   Content %d - Text: %s\n", i+1, content.Text[:min(200, len(content.Text))])
			if len(content.Text) > 200 {
				fmt.Printf("   ... (truncated)\n")
			}
		}
	}

	// Test 2: workspace_init with custom identifier
	fmt.Println("\n2. Testing workspace_init with custom identifier...")
	customArgs := map[string]interface{}{
		"identifier": "test-workspace-fix",
		"name":       "Test Workspace for Fix Validation",
	}
	customResult, err := workspaceInitTool.Execute(ctx, customArgs)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Success! IsError: %v\n", customResult.IsError)
		fmt.Printf("📄 Content count: %d\n", len(customResult.Content))
		for i, content := range customResult.Content {
			fmt.Printf("   Content %d - Type: %s\n", i+1, content.Type)
			fmt.Printf("   Content %d - Text: %s\n", i+1, content.Text[:min(200, len(content.Text))])
			if len(content.Text) > 200 {
				fmt.Printf("   ... (truncated)\n")
			}
		}
	}

	// Test 3: Validate content type is "text"
	fmt.Println("\n3. Validating content type fix...")
	if len(initResult.Content) > 0 {
		contentType := initResult.Content[0].Type
		if contentType == "text" {
			fmt.Printf("✅ Content type is correct: %s\n", contentType)
		} else {
			fmt.Printf("❌ Content type is incorrect: %s (should be 'text')\n", contentType)
		}
	}

	fmt.Println("\n🎉 workspace_init fix validation completed!")
	fmt.Println("=" * 40)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
