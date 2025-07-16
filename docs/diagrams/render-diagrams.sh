#!/bin/bash
# A-MEM Diagram Rendering Script
# Renders all Mermaid diagrams to PNG and SVG formats

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Print colored output
print_info() { echo -e "${BLUE}ℹ️  $1${NC}"; }
print_success() { echo -e "${GREEN}✅ $1${NC}"; }
print_warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }
print_error() { echo -e "${RED}❌ $1${NC}"; }

# Check if mermaid-cli is installed
check_dependencies() {
    print_info "Checking dependencies..."
    
    if ! command -v mmdc &> /dev/null; then
        print_error "mermaid-cli (mmdc) is not installed"
        echo "Install it with: npm install -g @mermaid-js/mermaid-cli"
        exit 1
    fi
    
    print_success "Dependencies check passed"
}

# Create output directories
create_directories() {
    print_info "Creating output directories..."
    
    mkdir -p png svg
    
    print_success "Output directories created"
}

# Render a single diagram
render_diagram() {
    local input_file="$1"
    local base_name=$(basename "$input_file" .mmd)
    
    print_info "Rendering $base_name..."
    
    # Render to PNG
    if mmdc -i "$input_file" -o "png/${base_name}.png" -w 1200 -H 800 --backgroundColor white; then
        print_success "PNG: png/${base_name}.png"
    else
        print_error "Failed to render PNG for $base_name"
        return 1
    fi
    
    # Render to SVG
    if mmdc -i "$input_file" -o "svg/${base_name}.svg" --backgroundColor white; then
        print_success "SVG: svg/${base_name}.svg"
    else
        print_error "Failed to render SVG for $base_name"
        return 1
    fi
}

# Main rendering function
render_all_diagrams() {
    print_info "Starting diagram rendering..."
    
    local diagrams=(
        "system-architecture-overview.mmd"
        "memory-storage-flow.mmd"
        "memory-retrieval-flow.mmd"
        "docker-infrastructure.mmd"
        "mcp-integration-architecture.mmd"
        "system-status-overview.mmd"
    )
    
    local success_count=0
    local total_count=${#diagrams[@]}
    
    for diagram in "${diagrams[@]}"; do
        if [ -f "$diagram" ]; then
            if render_diagram "$diagram"; then
                ((success_count++))
            fi
        else
            print_warning "Diagram file not found: $diagram"
        fi
    done
    
    echo ""
    print_info "Rendering Summary:"
    print_success "Successfully rendered: $success_count/$total_count diagrams"
    
    if [ $success_count -eq $total_count ]; then
        print_success "All diagrams rendered successfully!"
    else
        print_warning "Some diagrams failed to render"
    fi
}

# Generate index HTML file
generate_index() {
    print_info "Generating index.html..."
    
    cat > index.html << 'EOF'
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>A-MEM System Diagrams</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #333; border-bottom: 3px solid #007acc; padding-bottom: 10px; }
        h2 { color: #555; margin-top: 30px; }
        .diagram { margin: 20px 0; padding: 20px; border: 1px solid #ddd; border-radius: 5px; background: #fafafa; }
        .diagram img { max-width: 100%; height: auto; border: 1px solid #ccc; border-radius: 4px; }
        .links { margin: 10px 0; }
        .links a { margin-right: 15px; color: #007acc; text-decoration: none; }
        .links a:hover { text-decoration: underline; }
        .status { background: #e8f5e8; padding: 10px; border-radius: 4px; margin: 20px 0; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎨 A-MEM System Diagrams</h1>
        
        <div class="status">
            <strong>Status:</strong> Fully Operational ✅ | 
            <strong>Version:</strong> 1.0.1 | 
            <strong>Generated:</strong> $(date)
        </div>
        
        <h2>📊 Available Diagrams</h2>
        
        <div class="diagram">
            <h3>System Architecture Overview</h3>
            <p>High-level system architecture showing all components and their relationships.</p>
            <div class="links">
                <a href="png/system-architecture-overview.png">PNG</a>
                <a href="svg/system-architecture-overview.svg">SVG</a>
                <a href="system-architecture-overview.mmd">Source (.mmd)</a>
            </div>
            <img src="png/system-architecture-overview.png" alt="System Architecture Overview">
        </div>
        
        <div class="diagram">
            <h3>Memory Storage Flow</h3>
            <p>Step-by-step process of storing a memory with AI enhancement and vector embedding.</p>
            <div class="links">
                <a href="png/memory-storage-flow.png">PNG</a>
                <a href="svg/memory-storage-flow.svg">SVG</a>
                <a href="memory-storage-flow.mmd">Source (.mmd)</a>
            </div>
            <img src="png/memory-storage-flow.png" alt="Memory Storage Flow">
        </div>
        
        <div class="diagram">
            <h3>Memory Retrieval Flow</h3>
            <p>Process of retrieving memories based on vector similarity search and relevance scoring.</p>
            <div class="links">
                <a href="png/memory-retrieval-flow.png">PNG</a>
                <a href="svg/memory-retrieval-flow.svg">SVG</a>
                <a href="memory-retrieval-flow.mmd">Source (.mmd)</a>
            </div>
            <img src="png/memory-retrieval-flow.png" alt="Memory Retrieval Flow">
        </div>
        
        <div class="diagram">
            <h3>Docker Infrastructure</h3>
            <p>Container architecture, networking, and volume management for the A-MEM system.</p>
            <div class="links">
                <a href="png/docker-infrastructure.png">PNG</a>
                <a href="svg/docker-infrastructure.svg">SVG</a>
                <a href="docker-infrastructure.mmd">Source (.mmd)</a>
            </div>
            <img src="png/docker-infrastructure.png" alt="Docker Infrastructure">
        </div>
        
        <div class="diagram">
            <h3>MCP Integration Architecture</h3>
            <p>Model Context Protocol integration with Claude Desktop and server configuration.</p>
            <div class="links">
                <a href="png/mcp-integration-architecture.png">PNG</a>
                <a href="svg/mcp-integration-architecture.svg">SVG</a>
                <a href="mcp-integration-architecture.mmd">Source (.mmd)</a>
            </div>
            <img src="png/mcp-integration-architecture.png" alt="MCP Integration Architecture">
        </div>
        
        <div class="diagram">
            <h3>System Status Overview</h3>
            <p>Current operational status with recent fixes and quality metrics.</p>
            <div class="links">
                <a href="png/system-status-overview.png">PNG</a>
                <a href="svg/system-status-overview.svg">SVG</a>
                <a href="system-status-overview.mmd">Source (.mmd)</a>
            </div>
            <img src="png/system-status-overview.png" alt="System Status Overview">
        </div>
        
        <h2>📚 Additional Resources</h2>
        <ul>
            <li><a href="../../README.md">Project README</a></li>
            <li><a href="../../SYSTEM_DOCUMENTATION.md">System Documentation</a></li>
            <li><a href="../../A-MEM_ARCHITECTURE_v2.md">Architecture Guide</a></li>
            <li><a href="README.md">Diagram Documentation</a></li>
        </ul>
    </div>
</body>
</html>
EOF
    
    print_success "Generated index.html"
}

# Main execution
main() {
    echo "🎨 A-MEM Diagram Rendering Script"
    echo "=================================="
    echo ""
    
    # Change to diagrams directory
    cd "$(dirname "$0")"
    
    check_dependencies
    create_directories
    render_all_diagrams
    generate_index
    
    echo ""
    print_success "Diagram rendering complete!"
    print_info "Open index.html in your browser to view all diagrams"
}

# Run main function
main "$@"
