# A-MEM MCP Server Installation Guide

## Overview

A-MEM (Agentic Memory) is an AI-powered memory system that enables Claude Code and Claude Desktop to maintain contextual awareness across coding sessions. This guide will walk you through installing and configuring A-MEM to work with your Claude environment.

## What You'll Get

- **Persistent Memory**: Store code snippets with AI-generated analysis
- **Smart Retrieval**: Find relevant memories using vector similarity search
- **Intelligent Evolution**: Automatic memory network optimization
- **Seamless Integration**: Works directly within Claude Code and Claude Desktop

## Recent Updates

- ✅ **Fixed OpenAI API Integration**: Resolved critical issue where API key was not being properly passed to server
- ✅ **Enhanced Container Cleanup**: Install script now automatically cleans up existing containers and processes
- ✅ **Fixed MCP Protocol Compliance**: Resolved JSON-RPC notification handling for Claude Desktop integration
- ✅ **Improved Process Management**: Added comprehensive cleanup of amem-server processes started by Claude Desktop
- ✅ **Enhanced Claude Desktop Integration**: Proper MCP server configuration with correct API key handling
- ✅ **Renamed Server**: Changed from 'amem' to 'amem-augmented' for better namespace clarity

## Prerequisites

### System Requirements
- **Operating System**: macOS, Linux, or Windows with WSL2
- **Memory**: 4GB RAM minimum, 8GB recommended
- **Storage**: 2GB free space for models and data
- **Network**: Internet connection for initial setup and LLM API calls

### Required Software
- **Docker & Docker Compose**: For running the A-MEM services
- **Git**: For cloning the repository
- **Claude Code** or **Claude Desktop**: The AI assistant you want to enhance

### API Keys (Optional but Recommended)
- **OpenAI API Key**: For enhanced embeddings and analysis
- **Anthropic API Key**: For Claude-based analysis (alternative to OpenAI)

## Quick Installation

### Option 1: Automated Installation (Recommended)

```bash
# Clone the repository
git clone https://github.com/amem/mcp-server.git
cd mcp-server

# Run the automated installer
./scripts/install.sh

# Follow the interactive prompts to configure for your Claude environment
```

### Option 2: Manual Installation

If you prefer to install manually or need custom configuration, follow the detailed steps below.

## Detailed Installation Steps

### Step 1: Clone and Setup

```bash
# Clone the repository
git clone https://github.com/amem/mcp-server.git
cd mcp-server

# Copy environment template
cp .env.example .env

# Edit the environment file with your API keys
nano .env  # or use your preferred editor
```

### Step 2: Configure Environment

Edit the `.env` file and add your API keys:

```bash
# Required for enhanced functionality
OPENAI_API_KEY=your_openai_api_key_here

# Optional: Alternative to OpenAI
# ANTHROPIC_API_KEY=your_anthropic_api_key_here

# The rest can remain as defaults for most users
AMEM_ENV=production
AMEM_LOG_LEVEL=info
```

### Step 3: Start A-MEM Services

```bash
# Start all required services
docker-compose up -d

# Verify services are running
docker-compose ps

# Check service health
curl http://localhost:8004/api/v1/heartbeat  # ChromaDB
curl http://localhost:8005/health            # Sentence Transformers
```

Expected output:
```
NAME                                    STATUS
amem-chromadb-1                        Up (healthy)
amem-redis-1                           Up (healthy)  
amem-sentence-transformers-1           Up (healthy)
```

### Step 4: Build A-MEM Server

```bash
# Build the MCP server
make build

# Test the server
make test
```

### Step 5: Configure Claude Integration

Choose your Claude environment:

#### For Claude Code (VS Code Extension)

1. **Locate Claude Code Configuration**:
   ```bash
   # macOS
   ~/Library/Application Support/Code/User/settings.json
   
   # Linux
   ~/.config/Code/User/settings.json
   
   # Windows
   %APPDATA%\Code\User\settings.json
   ```

2. **Add A-MEM MCP Configuration**:
   ```json
   {
     "claude.mcpServers": {
       "amem-augmented": {
         "command": "/path/to/amem-server/amem-server",
         "args": ["-config", "/path/to/amem-server/config/production.yaml"],
         "env": {
           "OPENAI_API_KEY": "your_openai_api_key_here"
         }
       }
     }
   }
   ```

#### For Claude Desktop

1. **Locate Claude Desktop Configuration**:
   ```bash
   # macOS
   ~/Library/Application Support/Claude/claude_desktop_config.json
   
   # Linux
   ~/.config/claude/claude_desktop_config.json
   
   # Windows
   %APPDATA%\Claude\claude_desktop_config.json
   ```

2. **Add A-MEM MCP Configuration**:
   ```json
   {
     "mcpServers": {
       "amem-augmented": {
         "command": "/path/to/amem-server/amem-server",
         "args": ["-config", "/path/to/amem-server/config/production.yaml"],
         "env": {
           "OPENAI_API_KEY": "your_openai_api_key_here"
         }
       }
     }
   }
   ```

### Step 6: Restart Claude

- **Claude Code**: Restart VS Code or reload the Claude extension
- **Claude Desktop**: Restart the Claude Desktop application

### Step 7: Verify Installation

1. **Check MCP Tools**: In Claude, you should see new tools available:
   - `store_coding_memory`
   - `retrieve_relevant_memories` 
   - `evolve_memory_network`

2. **Test Basic Functionality**:
   ```
   Ask Claude: "Can you store this code snippet in memory?"
   
   Then provide any code snippet. Claude should respond with confirmation
   that the memory was stored with AI-generated keywords and tags.
   ```

## Configuration Options

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `OPENAI_API_KEY` | OpenAI API key for enhanced features | - | Recommended |
| `AMEM_ENV` | Environment (development/production) | development | No |
| `AMEM_LOG_LEVEL` | Logging level (debug/info/warn/error) | info | No |
| `CHROMADB_HOST` | ChromaDB connection URL | http://localhost:8002 | No |
| `EMBEDDING_SERVICE` | Embedding service (openai/sentence-transformers) | sentence-transformers | No |

### Service Ports

| Service | Port | Purpose |
|---------|------|---------|
| ChromaDB | 8004 | Vector database |
| Sentence Transformers | 8005 | Local embeddings |
| Redis | 6382 | Caching |
| A-MEM Metrics | 9092 | A-MEM server metrics |
| Prometheus | 9091 | Prometheus monitoring (optional) |

### Configuration Files

- `config/production.yaml`: Production settings
- `config/development.yaml`: Development settings
- `.env`: Environment variables and API keys

## Usage Examples

### Storing Code Memories

```
You: "Store this function in memory"

function fibonacci(n) {
  if (n <= 1) return n;
  return fibonacci(n-1) + fibonacci(n-2);
}

Claude: I'll store this code snippet in your A-MEM system.

[Uses store_coding_memory tool]

✅ Memory stored successfully!
- Memory ID: abc123
- Keywords: fibonacci, recursion, algorithm, javascript
- Tags: javascript, algorithm, recursive, mathematical
- Links Created: 2

The memory has been analyzed and stored with AI-generated keywords and tags.
```

### Retrieving Relevant Memories

```
You: "Find any memories related to recursive algorithms"

Claude: I'll search your memory system for relevant content.

[Uses retrieve_relevant_memories tool]

Found 3 relevant memories:

**Memory 1** (Relevance: 95.2%)
- Context: Recursive Fibonacci implementation
- Keywords: fibonacci, recursion, algorithm
- Code: function fibonacci(n) { ... }

**Memory 2** (Relevance: 87.3%)
- Context: Binary tree traversal using recursion
- Keywords: tree, recursion, traversal
- Code: function inorderTraversal(node) { ... }
```

### Memory Network Evolution

```
You: "Optimize my memory network"

Claude: I'll analyze and evolve your memory network to improve connections.

[Uses evolve_memory_network tool]

✅ Memory network evolution completed!

Results:
- Memories Analyzed: 45
- Memories Evolved: 12
- Links Created: 8
- Contexts Updated: 5
- Duration: 2.3 seconds

Your memory network has been optimized with better connections and improved context descriptions.
```

## Troubleshooting

### Common Issues

#### 1. "MCP server not found" Error

**Problem**: Claude can't find the A-MEM server executable.

**Solution**:
```bash
# Verify the server was built
ls -la amem-server

# If missing, rebuild
make build

# Update the path in your Claude configuration to the absolute path
which amem-server  # Use this full path in config
```

#### 2. "Connection refused" Error

**Problem**: A-MEM services aren't running.

**Solution**:
```bash
# Check service status
docker-compose ps

# Restart services if needed
docker-compose down
docker-compose up -d

# Check logs for errors
docker-compose logs
```

#### 3. "Incorrect API key provided: ${OPENAI_API_KEY:-}" Error

**Problem**: The install script wrote shell variable syntax to Claude config instead of the actual API key.

**Solution**:
```bash
# Re-run the install script to fix the configuration
./scripts/install.sh

# Or manually edit Claude Desktop config
# Replace "${OPENAI_API_KEY:-}" with your actual API key in:
# ~/Library/Application Support/Claude/claude_desktop_config.json
```

#### 4. "Port already in use" Error

**Problem**: Existing A-MEM processes or containers are using required ports.

**Solution**:
```bash
# The install script now automatically handles this, but you can manually clean up:
# Kill processes on A-MEM ports
lsof -ti:8080 | xargs kill -9 2>/dev/null || true
lsof -ti:9092 | xargs kill -9 2>/dev/null || true

# Kill any amem-server processes
pkill -f "amem-server" 2>/dev/null || true

# Clean up containers
docker-compose down --remove-orphans
```

#### 5. "API key not configured" Warning

**Problem**: OpenAI API key is missing or invalid.

**Solution**:
```bash
# Verify API key in .env file
grep OPENAI_API_KEY .env

# Test API key
curl -H "Authorization: Bearer your_api_key" \
     https://api.openai.com/v1/models
```

#### 4. Port Conflicts

**Problem**: Required ports are already in use.

**Solution**:
```bash
# Check what's using the ports
lsof -i :8004  # ChromaDB
lsof -i :8005  # Sentence Transformers
lsof -i :6382  # Redis

# Either stop conflicting services or change ports in docker-compose.yml
```

#### 5. Memory Storage Fails

**Problem**: Memories aren't being stored properly.

**Solution**:
```bash
# Check ChromaDB health
curl http://localhost:8004/api/v1/heartbeat

# Check server logs
docker-compose logs amem-server

# Verify disk space
df -h
```

### Getting Help

1. **Check Logs**: Always start by checking the logs
   ```bash
   docker-compose logs amem-server
   docker-compose logs chromadb
   ```

2. **Verify Configuration**: Ensure all paths and API keys are correct
   ```bash
   ./amem-server -config config/production.yaml --validate
   ```

3. **Test Components**: Test each component individually
   ```bash
   make test-chromadb
   make test-embeddings
   make test-mcp
   ```

4. **Community Support**: 
   - GitHub Issues: Report bugs and get help
   - Documentation: Check the latest docs for updates
   - Discord: Join the community for real-time help

## Advanced Configuration

### Custom Embedding Models

To use different embedding models:

```yaml
# config/production.yaml
embedding:
  service: "openai"  # or "sentence-transformers"
  model: "text-embedding-ada-002"  # OpenAI model
  # model: "all-MiniLM-L6-v2"     # Sentence transformers model
```

### Memory Evolution Scheduling

Configure automatic memory optimization:

```yaml
# config/production.yaml
evolution:
  enabled: true
  schedule: "0 2 * * *"  # Daily at 2 AM
  batch_size: 50
  worker_count: 3
```

### Monitoring and Metrics

Enable Prometheus monitoring:

```bash
# Start with monitoring
docker-compose --profile monitoring up -d

# Access metrics
open http://localhost:9090  # Prometheus
```

## Security Considerations

1. **API Keys**: Store securely, never commit to version control
2. **Network**: Run on private networks in production
3. **Access**: Limit access to the MCP server
4. **Updates**: Keep dependencies updated regularly

## Performance Tuning

### For Large Codebases

```yaml
# config/production.yaml
chromadb:
  batch_size: 200  # Increase for better throughput

embedding:
  batch_size: 64   # Increase for faster processing

evolution:
  batch_size: 100  # Process more memories at once
```

### For Resource-Constrained Systems

```yaml
# config/production.yaml
embedding:
  service: "sentence-transformers"  # Use local embeddings

evolution:
  worker_count: 1  # Reduce parallel processing
```

## Next Steps

1. **Explore Features**: Try different types of code storage and retrieval
2. **Customize**: Adjust configuration for your workflow
3. **Monitor**: Set up monitoring for production use
4. **Contribute**: Help improve A-MEM by reporting issues and suggestions

## Support

- **Documentation**: [Full documentation](README.md)
- **Issues**: [GitHub Issues](https://github.com/amem/mcp-server/issues)
- **Community**: [Discord Server](https://discord.gg/amem)
- **Email**: support@amem.dev
