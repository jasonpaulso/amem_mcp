# A-MEM MCP Server
version: alpha
--- 
This is an experiment in AI Assisted Software Engineering.
A human wrote [NONE] of the code directly (not yet, anyway).
USE AT OWN RISK :)

---

## What problem does this solve:

This Zettlekasten-based Model Context Protocol memory server addresses the challenge of maintaining a persistent, evolving understanding of complex codebases across multiple sessions and projects when working with tools like Claude Code and Claude Desktop. Traditional approaches often result in fragmented, non-persistent memories that reset with each session, making it difficult to build and search a comprehensive knowledge base. This server solves that by creating a "living" memory system that self-updates as new notes and information are added, automatically discovering relationships and connections to foster deeper insights and seamless continuity.

## HOW TO RUN

See QUICK_START.md for a quick start guide.

```bash
./scripts/install.sh

```
This will start the docker containers and build the server. 
`Should` append MCP config to Claude Desktop (takes backup of existing config, there is a restore script in case you need it) - "It works on my machine" (TM)

Restart Claude Desktop and it should be in the list of MCP servers. Any issues, raise them here :)

---
Technical description is in SYSTEM_DOCUMENTATION.md
---


An AI-powered memory system for Claude Code that enables persistent, contextual awareness across coding sessions.

## 🎯 Current Status

**Latest Release**: v1.1.0 - Workspace Management ✅

**New Features**:
- ✅ **Workspace Management**: Logical grouping of memories by filesystem path or user-defined name
- ✅ **Three New MCP Tools**: workspace_init, workspace_create, workspace_retrieve
- ✅ **Backward Compatibility**: Seamless migration from project-based organization
- ✅ **Smart Defaults**: Automatic workspace detection and initialization

**Recent Fixes**:
- ✅ **Memory Storage & Retrieval**: Fixed false error reporting and relevance calculation
- ✅ **Embedding Integration**: Resolved dimension mismatch and service connectivity
- ✅ **OpenAI API Integration**: Fixed critical authentication issue
- ✅ **MCP Protocol Compliance**: Resolved JSON-RPC notification handling
- ✅ **Container Management**: Enhanced cleanup and process management
- ✅ **Claude Desktop Integration**: Improved configuration handling

## Features

### Core Memory System
- **Memory Creation**: Store code snippets with AI-generated keywords, tags, and embeddings
- **Memory Retrieval**: Vector similarity search with ranking and filtering
- **Memory Evolution**: AI-driven analysis to update and optimize memory networks
- **MCP Integration**: JSON-RPC 2.0 server compatible with Claude Code

### Advanced Capabilities (Phase 2)
- **Real Embeddings**: Sentence-transformers and OpenAI embedding services
- **Intelligent Evolution**: Automated memory network optimization
- **Prompt Engineering**: Template-based LLM prompt management
- **Monitoring & Metrics**: Comprehensive Prometheus observability
- **Task Scheduling**: Cron-based automated maintenance
- **Multi-LLM Support**: LiteLLM proxy for fallback and model flexibility
- **Vector Storage**: ChromaDB for scalable similarity search

## Quick Start

### 🚀 One-Command Installation

```bash
git clone git@github.com:nixlim/amem_mcp.git
cd amem_mcp
./scripts/install.sh
```

The installer will automatically:
- ✅ Check prerequisites and dependencies
- ✅ Detect your Claude installation (Code/Desktop)
- ✅ Clean up existing containers and processes
- ✅ Start A-MEM services with Docker
- ✅ Configure Claude MCP integration with proper API key handling
- ✅ Test the installation and verify connectivity
- ✅ Ensure all memory operations work correctly

### Manual Installation

If you prefer manual setup:

1. **Prerequisites**: Docker, Docker Compose, Go 1.23+, OpenAI API key
2. **Setup**: `cp .env.example .env` and add your API key
3. **Start**: `docker-compose up -d && make build`
4. **Configure**: Follow the [Installation Guide](INSTALLATION_GUIDE.md)

### 📚 Documentation

- **[Quick Start Guide](QUICK_START.md)** - Get running in 5 minutes
- **[Installation Guide](INSTALLATION_GUIDE.md)** - Comprehensive setup instructions
- **[MCP Configuration Guide](MCP_CONFIGURATION_GUIDE.md)** - Claude integration details

### Verification

After installation, verify A-MEM is working:

```bash
# Check services
docker-compose ps

# Validate installation
./scripts/validate_installation.sh

# Test in Claude
# Ask Claude: "What tools do you have available?"
# You should see: store_coding_memory, retrieve_relevant_memories, evolve_memory_network
```

## MCP Tools

The server exposes three main tools for Claude Code:

### 1. store_coding_memory

Store a coding memory with AI analysis.

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

### 2. retrieve_relevant_memories

Search for relevant memories using vector similarity.

```json
{
  "tool": "retrieve_relevant_memories",
  "arguments": {
    "query": "How to implement fibonacci efficiently?",
    "max_results": 5,
    "min_relevance": 0.7
  }
}
```

### 3. evolve_memory_network

Trigger memory network evolution (Phase 2 feature).

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

## Configuration

Configuration is managed through YAML files and environment variables:

- `config/development.yaml` - Development settings
- `config/production.yaml` - Production settings
- `.env` - Environment variables (API keys, overrides)

Key configuration sections:

- **server**: Port, logging, request limits
- **chromadb**: Vector database connection
- **litellm**: LLM proxy settings and fallbacks
- **evolution**: Memory evolution scheduling
- **monitoring**: Metrics and tracing

## Architecture

```
┌─────────────────┐    ┌──────────────┐    ┌──────────────┐
│   Claude Code   │───▶│  MCP Server  │───▶│  Memory      │
│                 │    │              │    │  System      │
└─────────────────┘    └──────────────┘    └──────┬───────┘
                                                  │
                                                  ▼
                                        ┌──────────────┐
                                        │   LiteLLM    │
                                        │   Analysis   │
                                        └──────┬───────┘
                                               │
                                               ▼
                                        ┌──────────────┐
                                        │  ChromaDB    │
                                        │ Vector Store │
                                        └──────────────┘
```

## Development

### Project Structure

```
├── cmd/server/          # Main server entry point
├── pkg/
│   ├── config/         # Configuration management
│   ├── mcp/            # MCP protocol handlers
│   ├── memory/         # Core memory system
│   ├── models/         # Data models and schemas
│   └── services/       # External service integrations
├── config/             # Configuration files
├── prompts/            # LLM prompt templates
└── docker/             # Docker configurations
```

### Running Tests

```bash
go test ./...
```

### Building

```bash
# Development build
go build -o amem-server cmd/server/main.go

# Production build
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o amem-server cmd/server/main.go
```

## Monitoring

The server exposes Prometheus metrics on port 9090:

- Memory operation counts
- LLM request latency
- Vector search duration
- Error rates

Access metrics at: `http://localhost:9090/metrics`

## Troubleshooting

### Common Issues

1. **ChromaDB connection failed**:
   - Ensure ChromaDB is running: `docker-compose ps chromadb`
   - Check URL in config: `chromadb.url`

2. **LLM API errors**:
   - Verify API key in `.env` file
   - Check rate limits and quotas
   - Review fallback models in config

3. **Memory storage errors**:
   - Check ChromaDB logs: `docker-compose logs chromadb`
   - Verify collection initialization

### Logs

View server logs:
```bash
# Docker deployment
docker-compose logs amem-server

# Direct execution
./amem-server -log-level debug
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes with tests
4. Submit a pull request

## License

MIT License - see LICENSE file for details.

## Roadmap

- **Phase 1 (Current)**: MVP with core memory operations
- **Phase 2**: Memory evolution and optimization
- **Phase 3**: Advanced scheduling and monitoring // TODO: update
- **Phase 4**: Multi-user support and scaling


# The What and The Why

I have recently been interested in the problem of persistent context of AI Coding Agents and their context windows.
Gemini Pro has 2M tokens of context. Claude - 200K. Claude is hands down better coder at specific tasks. But big picture view - Gemini Pro is your best bet.

Augment Code has the best context management system I worked with so far. But, I wanted to give the agent things like semantic search, evolving memory, persistent context across sessions. 

I went to Arxiv, found the A-MEM paper and built an A-MEM MCP Server with AI.

In 10 hours I delivered a project that would take me, working solo without the AI, about 2-3 weeks of 5 days a week, 8 hours a day. 

It is by no means perfect. It is good enough though. I will continue to improve it. 

---
ACKNOWLEDGEMENTS:
This MCP Server was built on the basis of the following paper:

```
@article{xu2025mem,
title={A-mem: Agentic memory for llm agents},
author={Xu, Wujiang and Liang, Zujie and Mei, Kai and Gao, Hang and Tan, Juntao and Zhang, Yongfeng},
journal={arXiv preprint arXiv:2502.12110},
year={2025}
}
```
Link to pdf of paper: https://arxiv.org/pdf/2502.12110v1
Link to paper's github: https://github.com/WujiangXu/A-mem

The authors of the paper also have their own implementation of the system (don't think it's an MCP Server):
https://github.com/WujiangXu/A-mem-sys
___

It works. It has tests, startup scripts, local docker. Claude Desktop integration works, as should Claude Code.

I did not write a line of code. I paired with AI, I navigated and it drove. 

