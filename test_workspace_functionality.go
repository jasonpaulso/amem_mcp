package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/amem/mcp-server/pkg/config"
	"github.com/amem/mcp-server/pkg/memory"
	"github.com/amem/mcp-server/pkg/models"
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

	// Initialize LiteLLM service
	llmService := services.NewLiteLLMService(cfg.LiteLLM, logger.Named("litellm"))

	// Initialize embedding service
	embeddingService := services.NewEmbeddingService(cfg.Embedding, logger.Named("embedding"))

	// Initialize ChromaDB service
	chromaService := services.NewChromaDBService(cfg.ChromaDB, logger.Named("chromadb"))
	if err := chromaService.Initialize(ctx); err != nil {
		log.Fatalf("Failed to initialize ChromaDB: %v", err)
	}

	// Initialize workspace service
	workspaceService := services.NewWorkspaceService(chromaService, logger.Named("workspace"))

	// Initialize memory system
	memorySystem := memory.NewSystem(logger.Named("memory"), llmService, chromaService, embeddingService, workspaceService)

	// Test workspace functionality
	fmt.Println("🧪 Testing A-MEM Workspace Functionality")
	fmt.Println("=" * 50)

	// Test 1: Create workspace tools
	fmt.Println("\n1. Testing Workspace Tool Creation...")
	workspaceInitTool := memory.NewWorkspaceInitTool(workspaceService, logger.Named("workspace_init"))
	workspaceCreateTool := memory.NewWorkspaceCreateTool(workspaceService, logger.Named("workspace_create"))
	workspaceRetrieveTool := memory.NewWorkspaceRetrieveTool(workspaceService, logger.Named("workspace_retrieve"))

	fmt.Printf("✅ Created workspace_init tool: %s\n", workspaceInitTool.Name())
	fmt.Printf("✅ Created workspace_create tool: %s\n", workspaceCreateTool.Name())
	fmt.Printf("✅ Created workspace_retrieve tool: %s\n", workspaceRetrieveTool.Name())

	// Test 2: Test workspace_init with default workspace
	fmt.Println("\n2. Testing workspace_init with default workspace...")
	initArgs := map[string]interface{}{}
	initResult, err := workspaceInitTool.Execute(ctx, initArgs)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ workspace_init result: %s\n", initResult.Content[0].Text)
		if len(initResult.Content) > 1 {
			fmt.Printf("📄 JSON Response: %s\n", initResult.Content[1].Text)
		}
	}

	// Test 3: Test workspace_create with custom workspace
	fmt.Println("\n3. Testing workspace_create with custom workspace...")
	createArgs := map[string]interface{}{
		"identifier":  "test-project",
		"name":        "Test Project Workspace",
		"description": "A test workspace for validation",
	}
	createResult, err := workspaceCreateTool.Execute(ctx, createArgs)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ workspace_create result: %s\n", createResult.Content[0].Text)
		if len(createResult.Content) > 1 {
			fmt.Printf("📄 JSON Response: %s\n", createResult.Content[1].Text)
		}
	}

	// Test 4: Test workspace_retrieve
	fmt.Println("\n4. Testing workspace_retrieve...")
	retrieveArgs := map[string]interface{}{
		"identifier": "test-project",
	}
	retrieveResult, err := workspaceRetrieveTool.Execute(ctx, retrieveArgs)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ workspace_retrieve result: %s\n", retrieveResult.Content[0].Text)
		if len(retrieveResult.Content) > 1 {
			fmt.Printf("📄 JSON Response: %s\n", retrieveResult.Content[1].Text)
		}
	}

	// Test 5: Test memory storage with workspace_id
	fmt.Println("\n5. Testing memory storage with workspace_id...")
	storeReq := models.StoreMemoryRequest{
		Content:     "function calculateDistance(x1, y1, x2, y2) { return Math.sqrt((x2-x1)**2 + (y2-y1)**2); }",
		WorkspaceID: "test-project",
		CodeType:    "javascript",
		Context:     "Distance calculation function for game development",
	}

	storeResult, err := memorySystem.CreateMemory(ctx, storeReq)
	if err != nil {
		fmt.Printf("❌ Memory storage error: %v\n", err)
	} else {
		fmt.Printf("✅ Memory stored successfully: %s\n", storeResult.MemoryID)
		fmt.Printf("   Keywords: %v\n", storeResult.Keywords)
		fmt.Printf("   Tags: %v\n", storeResult.Tags)
	}

	// Test 6: Test memory retrieval with workspace_id
	fmt.Println("\n6. Testing memory retrieval with workspace_id...")
	retrieveReq := models.RetrieveMemoryRequest{
		Query:        "distance calculation",
		WorkspaceID:  "test-project",
		MaxResults:   5,
		MinRelevance: 0.3,
	}

	retrieveMemResult, err := memorySystem.RetrieveMemories(ctx, retrieveReq)
	if err != nil {
		fmt.Printf("❌ Memory retrieval error: %v\n", err)
	} else {
		fmt.Printf("✅ Retrieved %d memories\n", len(retrieveMemResult.Memories))
		for i, mem := range retrieveMemResult.Memories {
			fmt.Printf("   Memory %d: WorkspaceID='%s', Content='%s...'\n",
				i+1, mem.WorkspaceID, mem.Content[:min(50, len(mem.Content))])
		}
	}

	// Test 7: Test backward compatibility with project_path
	fmt.Println("\n7. Testing backward compatibility with project_path...")
	legacyStoreReq := models.StoreMemoryRequest{
		Content:     "class Player { constructor(name) { this.name = name; this.health = 100; } }",
		ProjectPath: "/legacy/game-project",
		CodeType:    "javascript",
		Context:     "Player class for legacy game project",
	}

	legacyStoreResult, err := memorySystem.CreateMemory(ctx, legacyStoreReq)
	if err != nil {
		fmt.Printf("❌ Legacy memory storage error: %v\n", err)
	} else {
		fmt.Printf("✅ Legacy memory stored successfully: %s\n", legacyStoreResult.MemoryID)
	}

	// Test 8: Test retrieval with project_filter (backward compatibility)
	fmt.Println("\n8. Testing retrieval with project_filter (backward compatibility)...")
	legacyRetrieveReq := models.RetrieveMemoryRequest{
		Query:         "player class",
		ProjectFilter: "/legacy/game-project",
		MaxResults:    5,
		MinRelevance:  0.3,
	}

	legacyRetrieveResult, err := memorySystem.RetrieveMemories(ctx, legacyRetrieveReq)
	if err != nil {
		fmt.Printf("❌ Legacy memory retrieval error: %v\n", err)
	} else {
		fmt.Printf("✅ Retrieved %d legacy memories\n", len(legacyRetrieveResult.Memories))
		for i, mem := range legacyRetrieveResult.Memories {
			fmt.Printf("   Memory %d: WorkspaceID='%s', ProjectPath='%s'\n",
				i+1, mem.WorkspaceID, mem.ProjectPath)
		}
	}

	fmt.Println("\n🎉 Workspace functionality testing completed!")
	fmt.Println("=" * 50)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
