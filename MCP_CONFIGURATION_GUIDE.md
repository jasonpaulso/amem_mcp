# A-MEM MCP Configuration Guide

## Overview

This guide provides detailed instructions for configuring the A-MEM MCP Server with Claude Code and Claude Desktop. The Model Context Protocol (MCP) enables Claude to use A-MEM's memory capabilities directly within your coding environment.

## Configuration File Locations

### Claude Code (VS Code Extension)

| Operating System | Configuration File Path |
|------------------|-------------------------|
| **macOS** | `~/Library/Application Support/Code/User/settings.json` | //TODO: Incorrect
| **Linux** | `~/.config/Code/User/settings.json` | //TODO: Not checked
| **Windows** | `%APPDATA%\Code\User\settings.json` |//TODO: Not checked

### Claude Desktop

| Operating System | Configuration File Path |
|------------------|-------------------------|
| **macOS** | `~/Library/Application Support/Claude/claude_desktop_config.json` | 
| **Linux** | `~/.config/claude/claude_desktop_config.json` | //TODO: Not checked
| **Windows** | `%APPDATA%\Claude\claude_desktop_config.json` | //TODO: Not checked

## Configuration Templates

### 1. Claude Code Configuration

Add this to your VS Code `settings.json`:

```json
{
  "claude.mcpServers": {
    "amem-augmented": {
      "command": "/absolute/path/to/amem-server",
      "args": ["-config", "/absolute/path/to/config/production.yaml"],
      "env": {
        "OPENAI_API_KEY": "your_openai_api_key_here"
      }
    }
  }
}
```

### 2. Claude Desktop Configuration

Create or update `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "amem-augmented": {
      "command": "/absolute/path/to/amem-server",
      "args": ["-config", "/absolute/path/to/config/production.yaml"],
      "env": {
        "OPENAI_API_KEY": "your_openai_api_key_here"
      }
    }
  }
}
```

## Environment-Specific Configurations

### Development Configuration

For development with enhanced logging and local services:

```json
{
  "mcpServers": {
    "amem-augmented": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/config/development.yaml"],
      "env": {
        "OPENAI_API_KEY": "your_api_key",
        "AMEM_ENV": "development",
        "AMEM_LOG_LEVEL": "debug",
        "CHROMADB_HOST": "http://localhost:8002",
        "EMBEDDING_SERVICE_URL": "http://localhost:8003"
      }
    }
  }
}
```

### Production Configuration

For production with optimized settings:

```json
{
  "mcpServers": {
    "amem-augmented": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/config/production.yaml"],
      "env": {
        "OPENAI_API_KEY": "your_api_key",
        "AMEM_ENV": "production",
        "AMEM_LOG_LEVEL": "info"
      }
    }
  }
}
```

## Configuration Options

### Required Parameters

| Parameter | Description | Example |
|-----------|-------------|---------|
| `command` | Absolute path to amem-server executable | `/home/user/amem/amem-server` |
| `args` | Command line arguments for the server | `["-config", "/path/to/config.yaml"]` |

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `OPENAI_API_KEY` | OpenAI API key for enhanced features | - | Recommended |
| `AMEM_ENV` | Environment (development/production) | development | No |
| `AMEM_LOG_LEVEL` | Logging level (debug/info/warn/error) | info | No |
| `CHROMADB_HOST` | ChromaDB connection URL | http://localhost:8002 | No |
| `EMBEDDING_SERVICE_URL` | Embedding service URL | http://localhost:8003 | No |

### Optional Advanced Settings

```json
{
  "mcpServers": {
    "amem-augmented": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/config/production.yaml"],
      "env": {
        "OPENAI_API_KEY": "your_api_key",
        "AMEM_ENV": "production",
        "AMEM_LOG_LEVEL": "info",
        "AMEM_EVOLUTION_ENABLED": "true",
        "AMEM_METRICS_ENABLED": "true",
        "LITELLM_DEFAULT_MODEL": "gpt-4-turbo",
        "EMBEDDING_SERVICE": "openai"
      },
      "timeout": 30000,
      "retries": 3
    }
  }
}
```

## Step-by-Step Configuration

### For Claude Code (VS Code)

1. **Open VS Code Settings**:
   ```bash
   # macOS/Linux
   code ~/.config/Code/User/settings.json
   
   # Or use VS Code UI: Cmd/Ctrl + , → Open Settings (JSON)
   ```

2. **Add A-MEM Configuration**:
   - If the file is empty, start with `{}`
   - Add the `claude.mcpServers` section
   - Replace `/path/to/amem-server` with your actual path
   - Replace `your_openai_api_key_here` with your API key

3. **Find Your A-MEM Path**:
   ```bash
   cd /path/to/amem-server
   pwd  # Copy this absolute path
   ```

4. **Complete Configuration Example**:
   ```json
   {
     "editor.fontSize": 14,
     "claude.mcpServers": {
       "amem-augmented": {
         "command": "/Users/username/amem-server/amem-server",
         "args": ["-config", "/Users/username/amem-server/config/production.yaml"],
         "env": {
           "OPENAI_API_KEY": "sk-..."
         }
       }
     }
   }
   ```

5. **Restart VS Code**

### For Claude Desktop

1. **Locate Configuration Directory**:
   ```bash
   # macOS
   mkdir -p ~/Library/Application\ Support/Claude
   
   # Linux
   mkdir -p ~/.config/claude
   
   # Windows
   mkdir %APPDATA%\Claude
   ```

2. **Create Configuration File**:
   ```bash
   # macOS
   touch ~/Library/Application\ Support/Claude/claude_desktop_config.json
   
   # Linux
   touch ~/.config/claude/claude_desktop_config.json
   ```

3. **Add Configuration**:
   ```json
   {
     "mcpServers": {
       "amem-augmented": {
         "command": "/Users/username/amem-server/amem-server",
         "args": ["-config", "/Users/username/amem-server/config/production.yaml"],
         "env": {
           "OPENAI_API_KEY": "sk-..."
         }
       }
     }
   }
   ```

4. **Restart Claude Desktop**

## Verification

### 1. Check Configuration Syntax

```bash
# Validate JSON syntax (macOS)
cat ~/Library/Application\ Support/Claude/claude_desktop_config.json | jq .

# Linux
cat ~/.config/claude/claude_desktop_config.json | jq .

# Should output formatted JSON without errors
```

### 2. Test A-MEM Server

```bash
# Test server directly
echo '{"jsonrpc": "2.0", "id": 1, "method": "initialize", "params": {}}' | \
  /path/to/amem-server -config /path/to/config/production.yaml
```

### 3. Verify in Claude

After restarting Claude, you should see these tools available:
- `store_coding_memory`
- `retrieve_relevant_memories`
- `evolve_memory_network`

Test with: "What tools do you have available?"

## Troubleshooting

### Common Issues

#### 1. "Command not found" Error

**Problem**: Claude can't find the amem-server executable.

**Solution**:
```bash
# Check if file exists and is executable
ls -la /path/to/amem-server
chmod +x /path/to/amem-server

# Use absolute path in configuration
which amem-server  # If in PATH
pwd  # If in current directory
```

#### 2. "Permission denied" Error

**Problem**: amem-server doesn't have execute permissions.

**Solution**:
```bash
chmod +x /path/to/amem-server
```

#### 3. "Configuration file not found" Error

**Problem**: Config file path is incorrect.

**Solution**:
```bash
# Check if config file exists
ls -la /path/to/config/production.yaml

# Use absolute path
realpath config/production.yaml
```

#### 4. "API key not configured" Warning

**Problem**: OpenAI API key is missing or invalid.

**Solution**:
```bash
# Test API key
curl -H "Authorization: Bearer your_api_key" \
     https://api.openai.com/v1/models

# Update configuration with correct key
```

#### 5. "Connection refused" Error

**Problem**: A-MEM services aren't running.

**Solution**:
```bash
# Start services
docker-compose up -d

# Check status
docker-compose ps

# Check logs
docker-compose logs
```

### Debug Mode

Enable debug logging for troubleshooting:

```json
{
  "mcpServers": {
    "amem-augmented": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/config/development.yaml"],
      "env": {
        "OPENAI_API_KEY": "your_api_key",
        "AMEM_LOG_LEVEL": "debug"
      }
    }
  }
}
```

### Log Files

Check these locations for logs:
- A-MEM Server: `./amem-server.log` (if configured)
- Docker Services: `docker-compose logs`
- Claude Code: VS Code Developer Console
- Claude Desktop: Application logs in system directories

## Advanced Configuration

### Multiple Environments

You can configure different A-MEM instances for different projects:

```json
{
  "mcpServers": {
    "amem-augmented-work": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/work-config.yaml"],
      "env": {
        "OPENAI_API_KEY": "work_api_key",
        "CHROMADB_COLLECTION": "work_memories"
      }
    },
    "amem-augmented-personal": {
      "command": "/path/to/amem-server",
      "args": ["-config", "/path/to/personal-config.yaml"],
      "env": {
        "OPENAI_API_KEY": "personal_api_key",
        "CHROMADB_COLLECTION": "personal_memories"
      }
    }
  }
}
```

### Custom Embedding Models

Configure different embedding services:

```json
{
  "env": {
    "EMBEDDING_SERVICE": "openai",
    "EMBEDDING_MODEL": "text-embedding-ada-002"
  }
}
```

Or:

```json
{
  "env": {
    "EMBEDDING_SERVICE": "sentence-transformers",
    "EMBEDDING_MODEL": "all-MiniLM-L6-v2"
  }
}
```

### Resource Limits

For resource-constrained environments:

```json
{
  "env": {
    "AMEM_EVOLUTION_ENABLED": "false",
    "EMBEDDING_BATCH_SIZE": "16",
    "CHROMADB_BATCH_SIZE": "50"
  }
}
```

## Security Best Practices

1. **API Key Security**:
   - Never commit API keys to version control
   - Use environment variables or secure key management
   - Rotate keys regularly

2. **File Permissions**:
   ```bash
   # macOS
   chmod 600 ~/Library/Application\ Support/Claude/claude_desktop_config.json
   chmod 700 ~/Library/Application\ Support/Claude

   # Linux
   chmod 600 ~/.config/claude/claude_desktop_config.json
   chmod 700 ~/.config/claude
   ```

3. **Network Security**:
   - Run A-MEM services on localhost only
   - Use firewall rules to restrict access
   - Consider VPN for remote access

## Support

If you encounter issues:

1. **Check Logs**: Always start with log files
2. **Validate Configuration**: Use JSON validators
3. **Test Components**: Test each part individually
4. **Community Help**: GitHub issues, Discord, documentation

For more help, see:
- [Installation Guide](INSTALLATION_GUIDE.md)
- [Troubleshooting Guide](TROUBLESHOOTING.md)
- [GitHub Issues](https://github.com/amem/mcp-server/issues)
