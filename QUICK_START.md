# A-MEM Quick Start Guide

## 🚀 Get A-MEM Running in 5 Minutes

A-MEM (Agentic Memory) gives Claude persistent memory across your coding sessions. Store code snippets, retrieve relevant examples, and let AI optimize your memory network automatically.

## Prerequisites

- **Claude Code** or **Claude Desktop**
- **Docker** and **Docker Compose**
- **OpenAI API Key**

## One-Command Installation

```bash
# Clone and install
git clone https://github.com/amem/mcp-server.git
cd mcp-server
./scripts/install.sh
```

The installer will:
- ✅ Check prerequisites
- ✅ Detect your Claude installation
- ✅ Start A-MEM services
- ✅ Configure Claude integration
- ✅ Test the installation

## Manual Installation (Alternative)

### 1. Setup A-MEM

```bash
# Clone repository
git clone https://github.com/amem/mcp-server.git
cd mcp-server

# Configure environment
cp .env.example .env
# Edit .env and add your OPENAI_API_KEY

# Start services
docker-compose up -d

# Build server
make build
```

### 2. Configure Claude

#### For Claude Code (VS Code):
Add to `~/Library/Application Support/Code/User/settings.json` (macOS):
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

#### For Claude Desktop:
Create `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS):
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

### 3. Restart Claude

Restart VS Code or Claude Desktop to load the new configuration.

## First Steps with A-MEM

### 1. Store Your First Memory

Ask Claude:
```
"Store this function in memory:

function fibonacci(n) {
  if (n <= 1) return n;
  return fibonacci(n-1) + fibonacci(n-2);
}
```

Claude will respond with:
```
✅ Memory stored successfully!
- Memory ID: abc123
- Keywords: fibonacci, recursion, algorithm, javascript
- Tags: javascript, algorithm, recursive, mathematical
- Links Created: 0

The memory has been analyzed and stored with AI-generated keywords and tags.
```

### 2. Retrieve Relevant Memories

Ask Claude:
```
"Find any memories related to recursive algorithms"
```

Claude will search and return:
```
Found 1 relevant memory:

**Memory 1** (Relevance: 95.2%)
- Context: Recursive Fibonacci implementation
- Keywords: fibonacci, recursion, algorithm, javascript
- Code: function fibonacci(n) { ... }
```

### 3. Evolve Your Memory Network

Ask Claude:
```
"Optimize my memory network"
```

Claude will analyze and improve connections between your memories.

## What A-MEM Tools Do

| Tool | Purpose | Example Use |
|------|---------|-------------|
| `store_coding_memory` | Save code with AI analysis | Store functions, snippets, solutions |
| `retrieve_relevant_memories` | Find related code | Search for patterns, examples, solutions |
| `evolve_memory_network` | Optimize memory connections | Improve organization and discoverability |

## Configuration Paths

### macOS
- **Claude Code**: `~/Library/Application Support/Code/User/settings.json`
- **Claude Desktop**: `~/Library/Application Support/Claude/claude_desktop_config.json`

### Linux
- **Claude Code**: `~/.config/Code/User/settings.json`
- **Claude Desktop**: `~/.config/claude/claude_desktop_config.json`

### Windows
- **Claude Code**: `%APPDATA%\Code\User\settings.json`
- **Claude Desktop**: `%APPDATA%\Claude\claude_desktop_config.json`

## Verification

### Check Services
```bash
docker-compose ps
# Should show: chromadb, redis, sentence-transformers (all Up)
```

### Test A-MEM Server
```bash
echo '{"jsonrpc": "2.0", "id": 1, "method": "initialize", "params": {}}' | \
  ./amem-server -config config/production.yaml
# Should return JSON response
```

### Verify in Claude
Ask Claude: "What tools do you have available?"
You should see the three A-MEM tools listed.

## Troubleshooting

### "Command not found"
- Use absolute paths in configuration
- Check file permissions: `chmod +x amem-server`

### "Connection refused"
- Start services: `docker-compose up -d`
- Check ports aren't in use: `lsof -i :8004`

### "API key not configured"
- Add OpenAI API key to `.env` file
- Update Claude configuration with the key

### Services won't start
- Check port conflicts
- View logs: `docker-compose logs`
- Free up disk space

## Advanced Usage

### Multiple Projects
Configure separate A-MEM instances for different projects by using different configuration files and ChromaDB collections.

### Custom Models
Switch between OpenAI and local sentence-transformers for embeddings based on your needs and privacy requirements.

### Monitoring
Access metrics at `http://localhost:9090/metrics` when Prometheus is enabled.

## Getting Help

- **Documentation**: [Full Installation Guide](INSTALLATION_GUIDE.md)
- **Configuration**: [MCP Configuration Guide](MCP_CONFIGURATION_GUIDE.md)
- **Issues**: [GitHub Issues](https://github.com/amem/mcp-server/issues)
- **Community**: [Discord Server](https://discord.gg/amem)

## What's Next?

1. **Store More Code**: Build up your memory database with functions, patterns, and solutions
2. **Use Retrieval**: Ask Claude to find relevant examples when coding
3. **Let It Evolve**: Run memory evolution to optimize your knowledge network
4. **Customize**: Adjust configuration for your workflow and preferences

A-MEM learns and improves as you use it, creating a personalized coding knowledge base that grows with your projects.

---

**Need help?** Check the [troubleshooting section](INSTALLATION_GUIDE.md#troubleshooting) or [open an issue](https://github.com/amem/mcp-server/issues).
