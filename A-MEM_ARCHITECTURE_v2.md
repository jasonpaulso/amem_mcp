# A-MEM MCP Server Architecture & Implementation Guide v2

## 🎯 Current Status (v1.0.1)

**Fully Operational** ✅ - All systems working correctly

**Recent Achievements:**
- ✅ **Complete Memory System**: Storage, retrieval, and search fully functional
- ✅ **Fixed Relevance Calculation**: Proper L2 distance to similarity conversion
- ✅ **Resolved Embedding Issues**: 384-dimensional embeddings working correctly
- ✅ **OpenAI API Integration**: Fixed authentication and direct API calls
- ✅ **MCP Protocol Compliance**: Full JSON-RPC 2.0 support with notification handling
- ✅ **Robust Installation**: Enhanced container and process cleanup
- ✅ **Claude Desktop Integration**: Proper configuration and error handling

## Table of Contents

### Part 1: Architecture Overview
1. [System Architecture](#1-system-architecture)
2. [Core Components](#2-core-components)
3. [Data Flow Patterns](#3-data-flow-patterns)
4. [Integration Points](#4-integration-points)

### Part 2: Core Workflows
5. [Memory Creation Flow](#5-memory-creation-flow)
6. [Memory Retrieval Flow](#6-memory-retrieval-flow)
7. [Memory Evolution Process](#7-memory-evolution-process)

### Part 3: Implementation Details
8. [Data Models & Schemas](#8-data-models--schemas)
9. [Prompt Engineering](#9-prompt-engineering)
10. [Service Integration](#10-service-integration)
11. [Error Handling & Resilience](#11-error-handling--resilience)

### Part 4: Operations
12. [Configuration Management](#12-configuration-management)
13. [Deployment Architecture](#13-deployment-architecture)
14. [Monitoring & Observability](#14-monitoring--observability)
15. [API Reference](#15-api-reference)

---

## Part 1: Architecture Overview

### 1. System Architecture

The A-MEM (Agentic Memory) MCP Server implements an AI-powered memory system that enables Claude Code to maintain contextual awareness across coding sessions.

```
┌─────────────────────────────────────────────────────────────────┐
│                    Claude Code Environment                      │
│                  (MCP Client Application)                       │
└─────────────────────┬───────────────────────────────────────────┘
                      │ MCP Protocol (JSON-RPC)
                      │
┌─────────────────────▼───────────────────────────────────────────┐
│                   MCP Server Layer                              │
│  ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐   │
│  │   Tool Handler  │ │   Tool Handler  │ │   Tool Handler  │   │
│  │ store_memory    │ │ retrieve_memory │ │ evolve_network  │   │
│  └─────────────────┘ └─────────────────┘ └─────────────────┘   │
└─────────────────────┬───────────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────────┐
│                Core Memory System                               │
│                                                                 │
│  ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐   │
│  │ Note Constructor│ │ Link Generator  │ │ Memory Evolver  │   │
│  │                 │ │                 │ │                 │   │
│  │ • Keywords      │ │ • Similarity    │ │ • Context Update│   │
│  │ • Tags          │ │ • Connections   │ │ • Relationship  │   │
│  │ • Context       │ │ • Embeddings    │ │ • Evolution     │   │
│  └─────────────────┘ └─────────────────┘ └─────────────────┘   │
└─────────────────────┬───────────────────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────────────────┐
│                External Services Layer                          │
│                                                                 │
│  ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐   │
│  │   LiteLLM API   │ │   ChromaDB      │ │  Embedding      │   │
│  │                 │ │                 │ │   Service       │   │
│  │ • Multi-LLM     │ │ • Vector Store  │ │ • SentenceTransf│   │
│  │ • Rate Limiting │ │ • Similarity    │ │ • HTTP Client   │   │
│  │ • Fallbacks     │ │ • Metadata      │ │ • Batch Process │   │
│  └─────────────────┘ └─────────────────┘ └─────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### 2. Core Components

#### 2.1 MCP Server Layer
- Implements Model Context Protocol for Claude Code integration
- Handles JSON-RPC communication
- Exposes memory tools as callable functions
- Manages request/response lifecycle

#### 2.2 AgenticMemorySystem
```
┌─────────────────────────────────────────────────────────────────┐
│                    AgenticMemorySystem                          │
│                                                                 │
│  ┌─────────────┐   ┌─────────────┐   ┌─────────────┐           │
│  │   Content   │   │   Memory    │   │   Memory    │           │
│  │  Analysis   │◄──┤  Evolution  │◄──┤ Consolidation│           │
│  │             │   │             │   │             │           │
│  └─────────────┘   └─────────────┘   └─────────────┘           │
│         │                 │                 │                   │
│         ▼                 ▼                 ▼                   │
│  ┌─────────────────────────────────────────────────┐           │
│  │              ChromaDB Storage                   │           │
│  │                                                 │           │
│  │  • Vector Embeddings  • Metadata Storage        │           │
│  │  • Similarity Search  • JSON Schema Support     │           │
│  └─────────────────────────────────────────────────┘           │
└─────────────────────────────────────────────────────────────────┘
```

#### 2.3 External Service Manager
```
┌─────────────────────────────────────────────────────────────────┐
│                External Service Manager                         │
│                                                                 │
│  ┌─────────────┐   ┌─────────────┐   ┌─────────────┐           │
│  │   LiteLLM   │   │  ChromaDB   │   │ Sentence    │           │
│  │ Controller  │   │  Client     │   │ Transformer │           │
│  │             │   │             │   │             │           │
│  │ • OpenAI    │   │ • Embeddings│   │ • Text      │           │
│  │ • Ollama    │   │ • Similarity│   │ • Encoding  │           │
│  │ • Anthropic │   │ • Metadata  │   │ • Vectors   │           │
│  └─────────────┘   └─────────────┘   └─────────────┘           │
└─────────────────────────────────────────────────────────────────┘
```

### 3. Data Flow Patterns

```
┌───────────────┐    ┌──────────────┐    ┌──────────────┐
│   Claude      │───▶│  MCP Server  │───▶│  Memory      │
│   Code        │    │              │    │  System      │
└───────────────┘    └──────────────┘    └──────┬───────┘
        ▲                    ▲                   │
        │                    │                   ▼
        │            ┌──────────────┐    ┌──────────────┐
        │            │   Response   │    │   LiteLLM    │
        └────────────│  Formatter   │◄───│   Analysis   │
                     └──────────────┘    └──────┬───────┘
                                                │
                                                ▼
                                        ┌──────────────┐
                                        │  ChromaDB    │
                                        │ Vector Store │
                                        └──────────────┘
```

### 4. Integration Points

#### 4.1 MCP Protocol Integration
- JSON-RPC 2.0 over stdio
- Tool registration and discovery
- Request/response handling
- Error propagation

#### 4.2 LLM Integration (via LiteLLM)
- Unified interface for multiple providers
- Automatic retry and fallback
- Rate limiting and quota management
- Response format validation

#### 4.3 Vector Database Integration
- ChromaDB for persistent storage
- Embedding generation and indexing
- Similarity search with metadata filtering
- Batch operations for performance

---

## Part 2: Core Workflows

### 5. Memory Creation Flow

```
    ┌─────────────────┐
    │   User Input    │
    │ (Code/Question) │
    └─────────┬───────┘
              │
              ▼
    ┌─────────────────┐
    │   MCP Tool      │
    │ store_memory    │
    └─────────┬───────┘
              │
              ▼
    ┌─────────────────┐
    │ Note Constructor│
    │                 │
    │ 1. Extract text │
    │ 2. Call LiteLLM │
    │ 3. Generate     │
    │    embeddings   │
    └─────────┬───────┘
              │
              ▼
    ┌─────────────────┐
    │ Link Generator  │
    │                 │
    │ 1. Find similar │
    │ 2. Analyze      │
    │    connections  │
    │ 3. Create links │
    └─────────┬───────┘
              │
              ▼
    ┌─────────────────┐
    │ Memory Storage  │
    │                 │
    │ 1. Save memory  │
    │ 2. Update links │
    │ 3. Index vector │
    └─────────┬───────┘
              │
              ▼
    ┌─────────────────┐
    │ Event Publisher │
    │                 │
    │ Emit:           │
    │ MemoryCreated   │
    └─────────────────┘
```

### 6. Memory Retrieval Flow

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   User Query    │────▶│ retrieve_memory │────▶│Query Processing │
│                 │     │   MCP Tool      │     │                 │
└─────────────────┘     └─────────────────┘     │1. Parse context │
                                                │2. Extract keywords│
                                                │3. Generate embedding│
                                                └─────────┬───────┘
                                                          │
        ┌─────────────────────────────────────────────────┘
        │
        ▼
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│Vector Similarity│────▶│   Candidate     │────▶│    Ranking      │
│    Search       │     │   Filtering     │     │  & Selection    │
│                 │     │                 │     │                 │
│1. Cosine sim    │     │1. Code context  │     │1. Relevance     │
│2. Top-K results │     │2. Project path  │     │2. Recency       │
│3. Threshold     │     │3. Language type │     │3. Link strength │
└─────────────────┘     └─────────────────┘     └─────────┬───────┘
                                                          │
                                                          ▼
                                                ┌─────────────────┐
                                                │ Return Results  │
                                                │                 │
                                                │1. Memory content│
                                                │2. Context info  │
                                                │3. Relevance     │
                                                └─────────────────┘
```

### 7. Memory Evolution Process

```
                        ┌─────────────────┐
                        │ Evolution Event │
                        │                 │
                        │ • MemoryCreated │
                        │ • LinkCreated   │
                        │ • Scheduled     │
                        └─────────┬───────┘
                                  │
                                  ▼
                        ┌─────────────────┐
                        │ Evolution Worker│
                        │ (Separate Svc)  │
                        └─────────┬───────┘
                                  │
                                  ▼
                        ┌─────────────────┐
                        │  Get Context    │
                        │                 │
                        │ • Trigger memory│
                        │ • Linked memories│
                        │ • Neighbors     │
                        └─────────┬───────┘
                                  │
                                  ▼
┌─────────────────────────────────────────────────────────────────┐
│                    LiteLLM Analysis                             │
│                                                                 │
│  INPUT:                           OUTPUT:                      │
│  • New memory context            • should_evolve: bool         │
│  • Historical memories           • actions: []string           │
│  • Current relationships         • suggested_connections       │
│                                  • context_updates             │
│                                  • tag_updates                 │
└─────────────────┬───────────────────────────────────────────────┘
                  │
                  ▼
        ┌─────────────────┐
        │ Apply Evolution │
        └─────────┬───────┘
                  │
    ┌─────────────┼─────────────┐
    │             │             │
    ▼             ▼             ▼
┌─────────┐ ┌─────────┐ ┌─────────────┐
│ Update  │ │ Create  │ │   Update    │
│Context  │ │ Links   │ │ Neighbors   │
│         │ │         │ │             │
│• Enrich │ │• Connect│ │• Propagate  │
│• Refine │ │• Weight │ │• Strengthen │
└─────────┘ └─────────┘ └─────────────┘
```

---

## Part 3: Implementation Details

### 8. Data Models & Schemas

#### 8.1 Core Memory Object

```go
type Memory struct {
    ID          string              `json:"id"`
    Content     string              `json:"content"`
    Context     string              `json:"context"`
    Keywords    []string            `json:"keywords"`
    Tags        []string            `json:"tags"`
    ProjectPath string              `json:"project_path"`
    CodeType    string              `json:"code_type"`
    Embedding   []float32           `json:"embedding"`
    Links       []MemoryLink        `json:"links"`
    CreatedAt   time.Time           `json:"created_at"`
    UpdatedAt   time.Time           `json:"updated_at"`
    Metadata    map[string]interface{} `json:"metadata"`
}

type MemoryLink struct {
    TargetID string  `json:"target_id"`
    LinkType string  `json:"link_type"` // solution|pattern|technology|debugging|progression
    Strength float32 `json:"strength"`  // 0.0-1.0
    Reason   string  `json:"reason"`
}
```

#### 8.2 Request/Response Schemas

```go
// Store Memory Request
type StoreMemoryRequest struct {
    Content     string `json:"content" required:"true"`
    ProjectPath string `json:"project_path"`
    CodeType    string `json:"code_type"`
    Context     string `json:"context"`
}

// Store Memory Response
type StoreMemoryResponse struct {
    MemoryID      string   `json:"memory_id"`
    Keywords      []string `json:"keywords"`
    Tags          []string `json:"tags"`
    LinksCreated  int      `json:"links_created"`
    EventEmitted  bool     `json:"event_emitted"`
}

// Retrieve Memory Request
type RetrieveMemoryRequest struct {
    Query         string   `json:"query" required:"true"`
    MaxResults    int      `json:"max_results" default:"5"`
    ProjectFilter string   `json:"project_filter"`
    CodeTypes     []string `json:"code_types"`
    MinRelevance  float32  `json:"min_relevance" default:"0.7"`
}

// Retrieve Memory Response
type RetrieveMemoryResponse struct {
    Memories []RetrievedMemory `json:"memories"`
    TotalFound int            `json:"total_found"`
}

type RetrievedMemory struct {
    Memory
    RelevanceScore float32 `json:"relevance_score"`
    MatchReason    string  `json:"match_reason"`
}
```

### 9. Prompt Engineering

#### 9.1 Prompt Template Management

```yaml
# prompts/note_construction.yaml
name: note_construction
version: 1.0
model_config:
  temperature: 0.1
  max_tokens: 1000
template: |
  Generate a structured analysis of the following coding content by:
  1. Identifying the most salient keywords (focus on technical terms, functions, concepts)
  2. Extracting core programming themes and contextual elements
  3. Creating relevant categorical tags for coding classification
  
  For coding context, consider:
  - Programming language and frameworks used
  - Problem domain (web dev, algorithms, data structures, etc.)
  - Solution patterns and techniques
  - Error types and debugging context
  - Libraries and dependencies mentioned
  
  Format the response as a JSON object:
  {
    "keywords": [// 3-7 specific technical keywords, ordered by importance],
    "context": // one sentence summarizing the coding problem/solution/concept,
    "tags": [// 3-6 broad categories: language, domain, pattern type, difficulty]
  }
  
  Content for analysis: {{.Content}}
  Project Path: {{.ProjectPath}}
  Code Type: {{.CodeType}}
```

#### 9.2 Prompt Loading System

```go
type PromptManager struct {
    promptDir string
    cache     map[string]*PromptTemplate
    mu        sync.RWMutex
}

func NewPromptManager(dir string) *PromptManager {
    return &PromptManager{
        promptDir: dir,
        cache:     make(map[string]*PromptTemplate),
    }
}

func (pm *PromptManager) LoadPrompt(name string) (*PromptTemplate, error) {
    pm.mu.RLock()
    if cached, ok := pm.cache[name]; ok {
        pm.mu.RUnlock()
        return cached, nil
    }
    pm.mu.RUnlock()
    
    // Load from file
    path := filepath.Join(pm.promptDir, name+".yaml")
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read prompt: %w", err)
    }
    
    var prompt PromptTemplate
    if err := yaml.Unmarshal(data, &prompt); err != nil {
        return nil, fmt.Errorf("failed to parse prompt: %w", err)
    }
    
    // Cache for future use
    pm.mu.Lock()
    pm.cache[name] = &prompt
    pm.mu.Unlock()
    
    return &prompt, nil
}
```

### 10. Service Integration

#### 10.1 LiteLLM Service

```go
type LiteLLMService struct {
    client        *litellm.Client
    promptManager *PromptManager
    config        LiteLLMConfig
}

type LiteLLMConfig struct {
    DefaultModel string
    MaxRetries   int
    Timeout      time.Duration
    RateLimit    int // requests per minute
}

func (s *LiteLLMService) CallWithRetry(prompt string, retryOnJSON bool) (string, error) {
    var lastErr error
    
    for i := 0; i < s.config.MaxRetries; i++ {
        resp, err := s.client.Complete(litellm.CompletionRequest{
            Model:       s.config.DefaultModel,
            Messages:    []litellm.Message{{Role: "user", Content: prompt}},
            Temperature: 0.1,
            ResponseFormat: litellm.ResponseFormat{Type: "json_object"},
        })
        
        if err != nil {
            lastErr = err
            time.Sleep(time.Second * time.Duration(i+1))
            continue
        }
        
        // Validate JSON if required
        if retryOnJSON {
            var test json.RawMessage
            if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &test); err != nil {
                lastErr = fmt.Errorf("invalid JSON response: %w", err)
                continue
            }
        }
        
        return resp.Choices[0].Message.Content, nil
    }
    
    return "", fmt.Errorf("all retries failed: %w", lastErr)
}
```

#### 10.2 ChromaDB Integration

```go
type ChromaDBService struct {
    client     *chromadb.Client
    collection string
    embedder   *EmbeddingService
}

func (c *ChromaDBService) StoreMemory(ctx context.Context, memory *Memory) error {
    // Generate embedding if not provided
    if len(memory.Embedding) == 0 {
        embedding, err := c.embedder.Embed(memory.Content)
        if err != nil {
            return fmt.Errorf("failed to generate embedding: %w", err)
        }
        memory.Embedding = embedding
    }
    
    // Prepare metadata
    metadata := map[string]interface{}{
        "context":      memory.Context,
        "keywords":     strings.Join(memory.Keywords, ","),
        "tags":         strings.Join(memory.Tags, ","),
        "project_path": memory.ProjectPath,
        "code_type":    memory.CodeType,
        "created_at":   memory.CreatedAt.Unix(),
    }
    
    // Store in ChromaDB
    return c.client.Add(ctx, chromadb.AddRequest{
        Collection: c.collection,
        IDs:        []string{memory.ID},
        Embeddings: [][]float32{memory.Embedding},
        Metadatas:  []map[string]interface{}{metadata},
        Documents:  []string{memory.Content},
    })
}

func (c *ChromaDBService) SearchSimilar(ctx context.Context, query string, limit int) ([]*Memory, error) {
    // Generate query embedding
    queryEmbedding, err := c.embedder.Embed(query)
    if err != nil {
        return nil, fmt.Errorf("failed to generate query embedding: %w", err)
    }
    
    // Search ChromaDB
    results, err := c.client.Query(ctx, chromadb.QueryRequest{
        Collection:     c.collection,
        QueryEmbeddings: [][]float32{queryEmbedding},
        NResults:       limit,
        Include:        []string{"metadatas", "documents", "distances"},
    })
    
    if err != nil {
        return nil, fmt.Errorf("ChromaDB query failed: %w", err)
    }
    
    // Convert results to Memory objects
    memories := make([]*Memory, 0, len(results.IDs[0]))
    for i, id := range results.IDs[0] {
        memory := &Memory{
            ID:      id,
            Content: results.Documents[0][i],
            // Reconstruct from metadata...
        }
        memories = append(memories, memory)
    }
    
    return memories, nil
}
```

### 11. Error Handling & Resilience

#### 11.1 Resilient JSON Parsing

```go
func ParseJSONWithRepair(llmService *LiteLLMService, response string, target interface{}) error {
    // First attempt: direct parsing
    if err := json.Unmarshal([]byte(response), target); err == nil {
        return nil
    }
    
    // Log malformed response for analysis
    log.Printf("Malformed JSON response: %s", response)
    
    // Second attempt: use LLM to repair
    repairPrompt := fmt.Sprintf(`Fix the following JSON to match the expected schema.
Return ONLY valid JSON, no explanations.

Malformed JSON:
%s

Expected structure example:
%s`, response, getExampleJSON(target))
    
    repairedJSON, err := llmService.CallWithRetry(repairPrompt, true)
    if err != nil {
        return fmt.Errorf("failed to repair JSON: %w", err)
    }
    
    // Third attempt: parse repaired JSON
    if err := json.Unmarshal([]byte(repairedJSON), target); err != nil {
        return fmt.Errorf("repaired JSON still invalid: %w", err)
    }
    
    return nil
}
```

#### 11.2 Circuit Breaker Pattern

```go
type CircuitBreaker struct {
    maxFailures  int
    resetTimeout time.Duration
    failures     int
    lastFailTime time.Time
    mu           sync.Mutex
    state        string // "closed", "open", "half-open"
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    cb.mu.Lock()
    defer cb.mu.Unlock()
    
    // Check if circuit is open
    if cb.state == "open" {
        if time.Since(cb.lastFailTime) > cb.resetTimeout {
            cb.state = "half-open"
            cb.failures = 0
        } else {
            return fmt.Errorf("circuit breaker is open")
        }
    }
    
    // Attempt the call
    err := fn()
    
    if err != nil {
        cb.failures++
        cb.lastFailTime = time.Now()
        
        if cb.failures >= cb.maxFailures {
            cb.state = "open"
            return fmt.Errorf("circuit breaker opened: %w", err)
        }
        return err
    }
    
    // Success - reset failures
    cb.failures = 0
    cb.state = "closed"
    return nil
}
```

---

## Part 4: Operations

### 12. Configuration Management

#### 12.1 Configuration Structure

```yaml
# config/production.yaml
server:
  port: 8080
  log_level: info
  max_request_size: 10MB

chromadb:
  url: "http://chromadb:8000"
  collection: "amem_memories"
  batch_size: 100

litellm:
  default_model: "gpt-4-turbo"
  fallback_models:
    - "gpt-3.5-turbo"
    - "claude-2"
  max_retries: 3
  timeout: 30s
  rate_limit: 60  # per minute

embedding:
  service: "sentence-transformers"
  model: "all-MiniLM-L6-v2"
  batch_size: 32

evolution:
  enabled: true
  schedule: "0 2 * * *"  # 2 AM daily
  batch_size: 50
  worker_count: 3

prompts:
  directory: "/app/prompts"
  cache_enabled: true
  hot_reload: true

monitoring:
  metrics_port: 9090
  enable_tracing: true
  sample_rate: 0.1
```

#### 12.2 Environment Variables

```bash
# Required
AMEM_ENV=production
OPENAI_API_KEY=sk-...
CHROMADB_HOST=chromadb

# Optional with defaults
AMEM_PORT=8080
AMEM_LOG_LEVEL=info
AMEM_CONFIG_PATH=/app/config
AMEM_PROMPTS_PATH=/app/prompts

# Feature flags
AMEM_EVOLUTION_ENABLED=true
AMEM_METRICS_ENABLED=true
```

### 13. Deployment Architecture

#### 13.1 Docker Compose Setup

```yaml
version: '3.8'

services:
  amem-server:
    build: .
    ports:
      - "8080:8080"
    environment:
      - AMEM_ENV=production
      - OPENAI_API_KEY=${OPENAI_API_KEY}
    volumes:
      - ./config:/app/config
      - ./prompts:/app/prompts
    depends_on:
      - chromadb
      - redis

  amem-evolution-worker:
    build: .
    command: ["./amem-worker", "--mode=evolution"]
    environment:
      - AMEM_ENV=production
      - OPENAI_API_KEY=${OPENAI_API_KEY}
    depends_on:
      - chromadb
      - redis
      - rabbitmq
    deploy:
      replicas: 3

  chromadb:
    image: chromadb/chroma:latest
    ports:
      - "8000:8000"
    volumes:
      - chromadb_data:/chroma/chroma

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

  rabbitmq:
    image: rabbitmq:3-management
    ports:
      - "5672:5672"
      - "15672:15672"
    environment:
      - RABBITMQ_DEFAULT_USER=amem
      - RABBITMQ_DEFAULT_PASS=secure_password

volumes:
  chromadb_data:
  redis_data:
```

#### 13.2 Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: amem-server
spec:
  replicas: 3
  selector:
    matchLabels:
      app: amem-server
  template:
    metadata:
      labels:
        app: amem-server
    spec:
      containers:
      - name: amem-server
        image: amem/server:latest
        ports:
        - containerPort: 8080
        env:
        - name: AMEM_ENV
          value: "production"
        - name: OPENAI_API_KEY
          valueFrom:
            secretKeyRef:
              name: amem-secrets
              key: openai-api-key
        resources:
          requests:
            memory: "512Mi"
            cpu: "500m"
          limits:
            memory: "1Gi"
            cpu: "1000m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
---
apiVersion: v1
kind: Service
metadata:
  name: amem-server
spec:
  selector:
    app: amem-server
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

### 14. Monitoring & Observability

#### 14.1 Metrics Collection

```go
// Prometheus metrics
var (
    memoryOperations = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "amem_memory_operations_total",
            Help: "Total number of memory operations",
        },
        []string{"operation", "status"},
    )
    
    llmLatency = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "amem_llm_request_duration_seconds",
            Help: "LLM request latency",
            Buckets: prometheus.DefBuckets,
        },
        []string{"model", "operation"},
    )
    
    vectorSearchLatency = prometheus.NewHistogram(
        prometheus.HistogramOpts{
            Name: "amem_vector_search_duration_seconds",
            Help: "Vector search latency",
            Buckets: prometheus.DefBuckets,
        },
    )
)
```

#### 14.2 Logging Strategy

```go
// Structured logging with context
type Logger struct {
    *zap.Logger
}

func (l *Logger) WithMemory(memoryID string) *Logger {
    return &Logger{l.With(zap.String("memory_id", memoryID))}
}

func (l *Logger) WithOperation(op string) *Logger {
    return &Logger{l.With(zap.String("operation", op))}
}

// Usage
logger.WithMemory(memory.ID).
    WithOperation("store_memory").
    Info("Memory stored successfully",
        zap.Int("links_created", len(memory.Links)),
        zap.Duration("duration", time.Since(start)),
    )
```

#### 14.3 Distributed Tracing

```go
// OpenTelemetry integration
func (s *MCPServer) storeMemoryWithTracing(ctx context.Context, req StoreMemoryRequest) (*StoreMemoryResponse, error) {
    ctx, span := tracer.Start(ctx, "store_memory",
        trace.WithAttributes(
            attribute.String("project_path", req.ProjectPath),
            attribute.String("code_type", req.CodeType),
        ),
    )
    defer span.End()
    
    // Note construction phase
    ctx, constructSpan := tracer.Start(ctx, "note_construction")
    note, err := s.constructNote(ctx, req)
    constructSpan.End()
    if err != nil {
        span.RecordError(err)
        return nil, err
    }
    
    // Link generation phase
    ctx, linkSpan := tracer.Start(ctx, "link_generation")
    links, err := s.generateLinks(ctx, note)
    linkSpan.End()
    
    // Continue with other phases...
}
```

### 15. API Reference

#### 15.1 MCP Tool: store_coding_memory

**Request:**
```json
{
  "tool": "store_coding_memory",
  "arguments": {
    "content": "function fibonacci(n) { return n <= 1 ? n : fibonacci(n-1) + fibonacci(n-2); }",
    "project_path": "/projects/algorithms",
    "code_type": "javascript",
    "context": "Recursive implementation of Fibonacci sequence"
  }
}
```

**Response:**
```json
{
  "memory_id": "mem_abc123",
  "keywords": ["fibonacci", "recursion", "algorithm", "function"],
  "tags": ["javascript", "algorithms", "recursion", "intermediate"],
  "links_created": 2,
  "event_emitted": true
}
```

#### 15.2 MCP Tool: retrieve_relevant_memories

**Request:**
```json
{
  "tool": "retrieve_relevant_memories",
  "arguments": {
    "query": "How to implement fibonacci efficiently?",
    "max_results": 5,
    "project_filter": "/projects/algorithms",
    "min_relevance": 0.7
  }
}
```

**Response:**
```json
{
  "memories": [
    {
      "id": "mem_abc123",
      "content": "function fibonacci(n) { return n <= 1 ? n : fibonacci(n-1) + fibonacci(n-2); }",
      "context": "Recursive implementation of Fibonacci sequence",
      "keywords": ["fibonacci", "recursion", "algorithm"],
      "tags": ["javascript", "algorithms", "recursion"],
      "relevance_score": 0.95,
      "match_reason": "Direct keyword and concept match"
    },
    {
      "id": "mem_def456",
      "content": "function fibonacciDP(n) { const dp = [0, 1]; for(let i = 2; i <= n; i++) { dp[i] = dp[i-1] + dp[i-2]; } return dp[n]; }",
      "context": "Dynamic programming solution for Fibonacci",
      "keywords": ["fibonacci", "dynamic programming", "optimization"],
      "tags": ["javascript", "algorithms", "dp", "efficient"],
      "relevance_score": 0.92,
      "match_reason": "Efficient implementation of same algorithm"
    }
  ],
  "total_found": 2
}
```

#### 15.3 MCP Tool: evolve_memory_network

**Request:**
```json
{
  "tool": "evolve_memory_network",
  "arguments": {
    "trigger_type": "manual",
    "scope": "recent",
    "max_memories": 100
  }
}
```

**Response:**
```json
{
  "memories_analyzed": 100,
  "memories_evolved": 15,
  "links_created": 23,
  "links_strengthened": 45,
  "contexts_updated": 8,
  "duration_ms": 3456
}
```

---

## Appendix: Migration Guide

### From v1 to v2

1. **Database Migration**
   - No changes to ChromaDB schema
   - Update any PostgreSQL references in code

2. **Configuration Updates**
   - Move prompts from code to `/app/prompts` directory
   - Update config files to new YAML format
   - Set evolution worker deployment

3. **Code Changes**
   - Replace hardcoded prompt strings with PromptManager
   - Update JSON parsing to use resilient parser
   - Implement event publishing for memory creation

4. **Deployment Changes**
   - Deploy separate evolution worker service
   - Add RabbitMQ or alternative message queue
   - Update monitoring dashboards

---

This document serves as the authoritative guide for the A-MEM MCP Server implementation. It will be maintained and versioned alongside the codebase.