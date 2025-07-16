# A-MEM System Documentation

**Version**: 1.0.1  
**Status**: Fully Operational  
**Last Updated**: [2025]

## 📋 Table of Contents

1. [System Overview](#system-overview)
2. [Architecture Components](#architecture-components)
3. [Data Flow](#data-flow)
4. [Configuration](#configuration)
5. [API Reference](#api-reference)
6. [Deployment](#deployment)
7. [Monitoring](#monitoring)
8. [Troubleshooting](#troubleshooting)

## 🎯 System Overview

A-MEM (Augmented Memory) is an AI-powered memory system that provides persistent, contextual awareness for Claude Desktop through the Model Context Protocol (MCP). The system enables storage, retrieval, and evolution of coding memories with vector similarity search and AI-driven analysis.

### Key Features

- **Workspace Management**: Logical grouping of memories by filesystem path or user-defined name
- **Memory Storage**: AI-enhanced storage with automatic keyword/tag generation
- **Vector Search**: 384-dimensional embedding-based similarity search
- **Memory Evolution**: AI-driven memory network optimization
- **MCP Integration**: Seamless Claude Desktop integration via JSON-RPC 2.0
- **Real-time Processing**: Immediate memory storage and retrieval
- **Automatic Linking**: Intelligent connections between related memories
- **Backward Compatibility**: Seamless migration from project-based to workspace-based organization

## 🏗️ Architecture Components

### Client Layer
- **Claude Desktop**: Primary user interface and MCP client
- **Claude Code**: Alternative client interface (future support)

### MCP Integration Layer
- **MCP Protocol**: JSON-RPC 2.0 communication standard
- **A-MEM Server**: `amem-augmented` MCP server process

### Core Services
- **Memory System**: Central orchestration and business logic
- **Embedding Service**: 384-dimensional vector generation
- **ChromaDB Service**: Vector database operations with UUID handling
- **LiteLLM Service**: OpenAI API integration for AI analysis

### External Dependencies
- **OpenAI API**: GPT-4.1 for memory analysis and enhancement
- **Sentence Transformers**: all-MiniLM-L6-v2 for embedding generation

### Data Storage
- **ChromaDB**: Vector database for embeddings and metadata
- **Redis**: Caching and session management

### Monitoring & Events
- **Prometheus**: Metrics collection and monitoring
- **RabbitMQ**: Event streaming and message queuing

## 🔄 Data Flow

### Memory Storage Flow

1. **Input Validation**: Validate content, context, and metadata
2. **AI Analysis**: Generate keywords, tags, and enhanced context using GPT-4.1
3. **Embedding Generation**: Create 384-dimensional vectors using sentence-transformers
4. **Vector Storage**: Store in ChromaDB with proper UUID handling
5. **Link Generation**: Create connections to similar existing memories
6. **Response**: Return memory ID, metadata, and link count

### Memory Retrieval Flow

1. **Query Processing**: Parse search query and apply filters
2. **Embedding Generation**: Convert query to 384-dimensional vector
3. **Similarity Search**: Perform L2 distance search in ChromaDB
4. **Relevance Calculation**: Convert distances to scores using `1.0 / (1.0 + distance)`
5. **Filtering & Ranking**: Apply relevance threshold and sort results
6. **Response**: Return ranked memories with similarity scores

## ⚙️ Configuration

### Environment Variables
```bash
# Core Configuration
OPENAI_API_KEY=sk-proj-...                    # OpenAI API key
EMBEDDING_SERVICE_URL=http://localhost:8005   # Sentence transformers service
CHROMADB_URL=http://localhost:8004           # ChromaDB vector database
REDIS_URL=redis://localhost:6380             # Redis cache

# Service Configuration
LITELLM_DEFAULT_MODEL=gpt-4.1                # OpenAI model
EMBEDDING_MODEL=all-MiniLM-L6-v2             # Embedding model
CHROMADB_COLLECTION=amem_memories             # Collection name
```

### Configuration Files
- **production.yaml**: Production environment settings
- **development.yaml**: Development environment settings
- **docker.yaml**: Docker container settings
- **.env**: Environment variables and API keys

### Claude Desktop Configuration
```json
{
  "claude.mcpServers": {
    "amem-augmented": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/config/production.yaml"],
      "env": {
        "OPENAI_API_KEY": "actual-api-key-value"
      }
    }
  }
}
```

## 📡 API Reference

### MCP Tools

#### Workspace Management Tools

#### workspace_init
Smart workspace initialization - creates new workspace or retrieves existing one.

**Parameters**:
- `identifier` (string, optional): Path or name for the workspace. If not provided, uses current working directory
- `name` (string, optional): Human-readable name for the workspace

**Response**:
- `workspace` (object): Workspace information including ID, name, description, and memory count
- `created` (boolean): True if workspace was created, false if retrieved

#### workspace_create
Explicit workspace creation - fails if workspace already exists.

**Parameters**:
- `identifier` (string, required): Path or name for the workspace
- `name` (string, optional): Human-readable name for the workspace
- `description` (string, optional): Description of the workspace

**Response**:
- `workspace` (object): Created workspace information
- `created` (boolean): Always true for successful creation

#### workspace_retrieve
Explicit workspace retrieval - fails if workspace doesn't exist.

**Parameters**:
- `identifier` (string, required): Path or name of the workspace to retrieve

**Response**:
- `workspace` (object): Retrieved workspace information including memory count
- `created` (boolean): Always false for retrieval

#### Memory Management Tools

#### store_coding_memory
Stores a new coding memory with AI enhancement.

**Parameters**:
- `content` (string, required): Code or text content
- `workspace_id` (string): Workspace identifier (path or name) for organizing memories
- `context` (string): Contextual description
- `project_path` (string): Project location (deprecated: use workspace_id)
- `code_type` (string): Programming language or content type

**Response**:
- `memory_id` (string): Unique identifier
- `keywords` (array): AI-generated keywords
- `tags` (array): AI-generated tags
- `links_created` (number): Number of links to similar memories

#### retrieve_relevant_memories
Retrieves memories based on similarity search.

**Parameters**:
- `query` (string, required): Search query
- `workspace_id` (string): Workspace identifier to filter results
- `max_results` (number, default: 5): Maximum results to return
- `min_relevance` (number, default: 0.7): Minimum relevance threshold
- `project_filter` (string): Filter by project path (deprecated: use workspace_id)
- `code_types` (array): Filter by code types

**Response**:
- `memories` (array): Retrieved memories with relevance scores
- `total_found` (number): Total number of matching memories

#### evolve_memory_network
Evolves and optimizes the memory network using AI analysis.

**Parameters**:
- `trigger_type` (string): manual|scheduled|event
- `scope` (string): recent|all|project
- `max_memories` (number): Maximum memories to analyze
- `project_path` (string): Scope to specific project

**Response**:
- `memories_analyzed` (number): Number of memories processed
- `memories_evolved` (number): Number of memories updated
- `links_created` (number): New connections created
- `duration_ms` (number): Processing time

## 🚀 Deployment

### Docker Deployment
```bash
# Start all services
docker-compose up -d

# Check service status
docker-compose ps

# View logs
docker-compose logs amem-server
```

### Manual Deployment
```bash
# Build server
go build -o amem-server cmd/server/main.go

# Start server
./amem-server -config config/production.yaml
```

### Installation Script
```bash
# Automated installation
./scripts/install.sh

# Features:
# - Container cleanup
# - Process management
# - Configuration generation
# - Claude Desktop integration
```

## 📊 Monitoring

### Metrics (Prometheus)
- **Memory Operations**: Storage/retrieval rates and latencies
- **API Performance**: Response times and error rates
- **System Health**: Service availability and resource usage
- **Vector Operations**: Embedding generation and search performance

### Health Checks
- **HTTP Endpoints**: `/health`, `/metrics`, `/ready`
- **Service Dependencies**: ChromaDB, Redis, Sentence Transformers
- **External APIs**: OpenAI API connectivity

### Logging
- **Structured Logging**: JSON format with contextual information
- **Log Levels**: DEBUG, INFO, WARN, ERROR
- **Request Tracing**: Unique request IDs for debugging

## 🔧 Troubleshooting

### Common Issues

#### Memory Storage Failures
- **Check OpenAI API key**: Verify environment variable is set correctly
- **Verify embedding service**: Ensure sentence-transformers is running on port 8005
- **ChromaDB connectivity**: Check collection exists and is accessible

#### Memory Retrieval Issues
- **Relevance threshold**: Lower min_relevance if no results returned
- **Embedding dimensions**: Verify 384-dimensional embeddings are being used
- **Collection state**: Ensure ChromaDB collection has correct dimensionality

#### MCP Integration Problems
- **Claude Desktop config**: Verify JSON syntax and file paths
- **Process conflicts**: Use install script to clean up existing processes
- **API key handling**: Ensure actual key value, not shell variable syntax

### Debug Commands
```bash
# Check service status
curl http://localhost:8080/health

# Test embedding service
curl http://localhost:8005/embeddings -d '{"sentences":["test"]}'

# Query ChromaDB directly
curl http://localhost:8004/api/v1/collections

# Check server logs
docker-compose logs -f amem-server
```

### Performance Optimization
- **Embedding Cache**: Sentence transformers model caching
- **Connection Pooling**: HTTP client connection reuse
- **Batch Processing**: Multiple embeddings in single request
- **Memory Limits**: Configure appropriate container resource limits

---

## 📚 Additional Resources

- **Installation Guide**: [INSTALLATION_GUIDE.md](INSTALLATION_GUIDE.md)
- **Architecture Details**: [A-MEM_ARCHITECTURE_v2.md](A-MEM_ARCHITECTURE_v2.md)
- **Project Status**: [PROJECT_STATUS.md](PROJECT_STATUS.md)
- **Change Log**: [CHANGELOG.md](CHANGELOG.md)

For support and issues, refer to the troubleshooting section or check the project documentation.
