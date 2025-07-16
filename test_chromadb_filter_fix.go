package main

import (
	"context"
	"fmt"
	"log"

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

	fmt.Println("🧪 Testing ChromaDB Filter Fix")
	fmt.Println("=" * 40)

	// Test the exact query that was failing
	fmt.Println("\n1. Testing the failing query with multiple filters...")

	// This is the exact request that was causing the error
	retrieveReq := models.RetrieveMemoryRequest{
		Query:        "collision detection spawnSplitterChildren",
		CodeTypes:    []string{"javascript"},
		MaxResults:   3,
		WorkspaceID:  "/",
		MinRelevance: 0.3,
	}

	fmt.Printf("Request details:\n")
	fmt.Printf("  Query: %s\n", retrieveReq.Query)
	fmt.Printf("  CodeTypes: %v\n", retrieveReq.CodeTypes)
	fmt.Printf("  WorkspaceID: %s\n", retrieveReq.WorkspaceID)
	fmt.Printf("  MaxResults: %d\n", retrieveReq.MaxResults)
	fmt.Printf("  MinRelevance: %.1f\n", retrieveReq.MinRelevance)

	// Execute the query
	result, err := memorySystem.RetrieveMemories(ctx, retrieveReq)
	if err != nil {
		fmt.Printf("❌ Error (this should be fixed now): %v\n", err)
	} else {
		fmt.Printf("✅ Success! Retrieved %d memories\n", len(result.Memories))
		for i, mem := range result.Memories {
			fmt.Printf("   Memory %d: Relevance=%.3f, WorkspaceID='%s', CodeType='%s'\n",
				i+1, mem.RelevanceScore, mem.Memory.WorkspaceID, mem.Memory.CodeType)
		}
	}

	// Test 2: Single filter condition (workspace only)
	fmt.Println("\n2. Testing single filter condition (workspace only)...")
	singleFilterReq := models.RetrieveMemoryRequest{
		Query:        "test query",
		WorkspaceID:  "/",
		MaxResults:   3,
		MinRelevance: 0.3,
	}

	singleResult, err := memorySystem.RetrieveMemories(ctx, singleFilterReq)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Success! Retrieved %d memories with single filter\n", len(singleResult.Memories))
	}

	// Test 3: Single filter condition (code type only)
	fmt.Println("\n3. Testing single filter condition (code type only)...")
	codeOnlyReq := models.RetrieveMemoryRequest{
		Query:        "test query",
		CodeTypes:    []string{"javascript"},
		MaxResults:   3,
		MinRelevance: 0.3,
	}

	codeOnlyResult, err := memorySystem.RetrieveMemories(ctx, codeOnlyReq)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Success! Retrieved %d memories with code type filter\n", len(codeOnlyResult.Memories))
	}

	// Test 4: No filters (should work)
	fmt.Println("\n4. Testing no filters...")
	noFilterReq := models.RetrieveMemoryRequest{
		Query:        "test query",
		MaxResults:   3,
		MinRelevance: 0.3,
	}

	noFilterResult, err := memorySystem.RetrieveMemories(ctx, noFilterReq)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Success! Retrieved %d memories with no filters\n", len(noFilterResult.Memories))
	}

	// Test 5: Multiple code types
	fmt.Println("\n5. Testing multiple code types with workspace...")
	multiCodeReq := models.RetrieveMemoryRequest{
		Query:        "test query",
		CodeTypes:    []string{"javascript", "python", "go"},
		WorkspaceID:  "/",
		MaxResults:   3,
		MinRelevance: 0.3,
	}

	multiCodeResult, err := memorySystem.RetrieveMemories(ctx, multiCodeReq)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	} else {
		fmt.Printf("✅ Success! Retrieved %d memories with multiple code types\n", len(multiCodeResult.Memories))
	}

	fmt.Println("\n🎉 ChromaDB filter fix validation completed!")
	fmt.Println("=" * 40)
}
