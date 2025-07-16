# Changelog

All notable changes to the A-MEM MCP Server project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2025-07-15

### 🚀 Major Feature Release - Workspace Management

This release introduces comprehensive workspace management functionality, enabling logical organization of memories by filesystem path or user-defined name.

### ✨ New Features

#### Workspace Management System
- **Workspace Service**: Comprehensive workspace management with ID validation, normalization, and smart defaults
- **Dual ID Support**: Supports both filesystem paths (`/Users/john/project`) and logical names (`my-game-project`)
- **Smart Initialization**: Automatic workspace detection using current working directory as default
- **Metadata Filtering**: Efficient workspace separation using ChromaDB metadata

#### New MCP Tools
- **workspace_init**: Smart initialization - creates new workspace or retrieves existing one
- **workspace_create**: Explicit workspace creation (fails if workspace already exists)
- **workspace_retrieve**: Explicit workspace retrieval (fails if workspace doesn't exist)

#### Enhanced Existing Tools
- **store_coding_memory**: Now supports `workspace_id` parameter for memory organization
- **retrieve_relevant_memories**: Enhanced with `workspace_id` filtering for scoped searches
- **evolve_memory_network**: Workspace-aware memory evolution and linking

### 🔄 Backward Compatibility
- **Seamless Migration**: Existing `project_path` functionality preserved and mapped to workspace concepts
- **Legacy Support**: All existing memories continue to work without modification
- **Graceful Fallbacks**: Smart defaults ensure existing integrations remain functional

### 🏗️ Technical Improvements
- **Data Model Extensions**: Memory struct enhanced with WorkspaceID field
- **Service Architecture**: Clean workspace service integration with dependency injection
- **ChromaDB Integration**: Enhanced metadata storage and filtering capabilities
- **OR Logic Implementation**: Proper backward compatibility for legacy project filters

---

## [1.0.1] - 2025-07-15

### 🔧 Critical Fixes - Memory System Fully Operational

This patch release resolves the final critical issues with memory storage and retrieval, making the A-MEM system fully functional.

### ✅ Fixed

#### Memory Storage and Retrieval Issues
- **Fixed false error reporting**: HTTP 201 (Created) responses from ChromaDB now correctly recognized as successful operations
- **Fixed memory retrieval failure**: Corrected relevance calculation formula from `(1.0 - distance)` to `(1.0 / (1.0 + distance))` for proper L2 distance handling
- **Resolved embedding dimension mismatch**: Recreated ChromaDB collection with correct 384-dimensional embeddings
- **Fixed embedding service connection**: Added missing URL field to EmbeddingConfig struct

#### Technical Improvements
- **Enhanced error handling**: Proper HTTP status code validation for ChromaDB operations
- **Improved mathematical accuracy**: Relevance scores now correctly reflect similarity for Euclidean distances
- **Better configuration management**: Embedding service now uses configured URLs instead of hardcoded values

### 🎯 Impact
- **Memory storage**: Now works without false error reports
- **Memory retrieval**: Returns relevant results with accurate similarity scores
- **Search functionality**: Properly finds and ranks stored memories
- **System reliability**: All core memory operations functioning correctly

---

## [1.0.0] - 2025-07-15

### 🎉 Major Release - Production Ready

This release marks the first stable version of A-MEM with full Claude Desktop integration and robust error handling.

### ✅ Fixed

#### Critical OpenAI API Integration Issue
- **Fixed authentication failure**: Resolved issue where install script was writing shell variable syntax `${OPENAI_API_KEY:-}` to Claude Desktop config instead of actual API key value
- **Root cause**: Claude Desktop doesn't expand shell variables, was passing literal string to server
- **Solution**: Install script now reads actual API key from .env file and substitutes real value into configuration
- **Impact**: Eliminates 401 authentication errors, enables proper OpenAI API integration

#### MCP Protocol Compliance
- **Fixed JSON-RPC notification handling**: Added support for messages without ID fields (like `notifications/initialized`)
- **Root cause**: Server was trying to validate all messages as requests, but notifications don't have ID fields
- **Solution**: Implemented separate `MCPNotification` struct and `handleJSONRPCNotification` method
- **Impact**: Resolves Zod validation errors in Claude Desktop integration

#### Container and Process Management
- **Enhanced container cleanup**: Install script now automatically detects and removes existing A-MEM containers
- **Added process cleanup**: Kills existing amem-server processes (including those started by Claude Desktop)
- **Port conflict resolution**: Comprehensive cleanup of processes using ports 8080 and 9092
- **Impact**: Prevents port conflicts during installation, enables clean restarts

### 🚀 Enhanced

#### Installation Experience
- **Comprehensive cleanup**: Install script now handles both containers and processes
- **Better error handling**: Graceful fallbacks and clear error messages
- **User feedback**: Detailed progress reporting during installation
- **Configuration preservation**: Install script merges with existing Claude Desktop config instead of overwriting

#### Server Configuration
- **Direct OpenAI API calls**: Changed from localhost:4000 proxy to https://api.openai.com/v1/chat/completions
- **Correct model names**: Updated to use `gpt-4.1` model as specified in OpenAI API
- **Proper authentication**: Added Bearer token authentication with OpenAI API key
- **Environment variable handling**: Improved .env file processing and validation

#### MCP Server Naming
- **Renamed server**: Changed from 'amem' to 'amem-augmented' throughout codebase
- **Better namespace clarity**: Improved identification in Claude Desktop integration
- **Consistent naming**: Updated all references, configurations, and documentation

### 🔧 Technical Improvements

#### Code Quality
- **Separated response types**: Created `MCPSuccessResponse` and `MCPErrorResponse` for proper JSON-RPC compliance
- **Enhanced error handling**: Better error messages and graceful degradation
- **Improved logging**: All logs now go to stderr, keeping stdout clean for JSON-RPC

#### Configuration Management
- **Multiple config files**: Separate configs for development, production, and Docker environments
- **Port management**: Updated port assignments to avoid conflicts (metrics on 9092, Prometheus on 9091)
- **Environment flexibility**: Better handling of different deployment scenarios

### 📚 Documentation

#### Updated Guides
- **Installation Guide**: Added troubleshooting sections for new fixes
- **README**: Updated with current status and enhanced installation steps
- **Architecture Documentation**: Reflects current implementation and fixes

#### New Troubleshooting
- **API Key Issues**: Comprehensive guide for authentication problems
- **Port Conflicts**: Solutions for container and process conflicts
- **MCP Integration**: Claude Desktop specific troubleshooting

### 🧪 Testing

#### Validation
- **Container cleanup tested**: Verified detection and removal of existing containers
- **Process cleanup tested**: Confirmed killing of amem-server processes
- **API integration tested**: Validated OpenAI API calls with correct authentication
- **MCP protocol tested**: Verified JSON-RPC compliance with both requests and notifications

### 🔄 Migration Notes

For existing installations:
1. **Re-run install script**: Recommended to get all fixes and improvements
2. **Manual config fix**: If needed, replace `${OPENAI_API_KEY:-}` with actual API key in Claude Desktop config
3. **Restart Claude Desktop**: Required to pick up configuration changes

### 🎯 Next Steps

- Monitor for any remaining integration issues
- Gather user feedback on installation experience
- Plan for additional MCP tools and capabilities
- Consider automated testing for installation scenarios

---

## [0.9.0] - 2025-07-14

### Initial Development
- Basic MCP server implementation
- Docker containerization
- Initial Claude integration
- Core memory functionality

---

**Legend:**
- 🎉 Major releases
- ✅ Bug fixes
- 🚀 Enhancements
- 🔧 Technical improvements
- 📚 Documentation
- 🧪 Testing
- 🔄 Migration notes
- 🎯 Future plans
