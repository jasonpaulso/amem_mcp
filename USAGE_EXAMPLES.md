# A-MEM Usage Examples

## Real-World Integration Examples

This guide shows practical examples of using A-MEM with Claude Code and Claude Desktop for enhanced coding productivity.

## Example 1: Building a React Component Library

### Scenario
You're building a React component library and want to maintain memory of patterns, components, and solutions across development sessions.

### Step 1: Store Component Patterns

**You**: "Store this reusable button component in memory"

```jsx
import React from 'react';
import styled from 'styled-components';

const StyledButton = styled.button`
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  
  ${props => props.variant === 'primary' && `
    background: #007bff;
    color: white;
    &:hover { background: #0056b3; }
  `}
  
  ${props => props.variant === 'secondary' && `
    background: #6c757d;
    color: white;
    &:hover { background: #545b62; }
  `}
`;

export const Button = ({ variant = 'primary', children, ...props }) => (
  <StyledButton variant={variant} {...props}>
    {children}
  </StyledButton>
);
```

**Claude Response**:
```
✅ Memory stored successfully!

Memory ID: btn-comp-001
Keywords: react, button, styled-components, variant, reusable, component
Tags: react, component, ui, styled-components, reusable
Links Created: 0

The memory has been analyzed and stored. This reusable button component pattern 
with variant support and styled-components is now available for future reference.
```

### Step 2: Store Related Hook Pattern

**You**: "Store this custom hook for button loading states"

```jsx
import { useState, useCallback } from 'react';

export const useAsyncButton = (asyncAction) => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const execute = useCallback(async (...args) => {
    try {
      setIsLoading(true);
      setError(null);
      const result = await asyncAction(...args);
      return result;
    } catch (err) {
      setError(err.message);
      throw err;
    } finally {
      setIsLoading(false);
    }
  }, [asyncAction]);

  return { execute, isLoading, error };
};
```

**Claude Response**:
```
✅ Memory stored successfully!

Memory ID: hook-async-btn-001
Keywords: react, hook, async, loading, error, button, state
Tags: react, hook, async, state-management, error-handling
Links Created: 1

The memory has been stored and linked to your button component pattern.
This creates a connection between UI components and state management patterns.
```

### Step 3: Retrieve Related Patterns Later

**You**: "I need to create a form submission button. Find relevant patterns from my memory."

**Claude Response**:
```
I'll search your memory for relevant button and form patterns.

[Uses retrieve_relevant_memories tool]

Found 2 relevant memories:

**Memory 1** (Relevance: 92.4%)
- Context: Reusable button component with variant support
- Keywords: react, button, styled-components, variant, reusable
- Match Reason: Keyword match: button, react

**Memory 2** (Relevance: 87.1%)  
- Context: Custom hook for async button loading states
- Keywords: react, hook, async, loading, error, button
- Match Reason: Keyword match: button, async

Based on your memory, I can help you create a form submission button that combines 
these patterns. Here's a solution using your existing components:

```jsx
import { Button } from './Button';
import { useAsyncButton } from './useAsyncButton';

export const SubmitButton = ({ onSubmit, children, ...props }) => {
  const { execute, isLoading, error } = useAsyncButton(onSubmit);

  return (
    <Button 
      variant="primary" 
      onClick={execute}
      disabled={isLoading}
      {...props}
    >
      {isLoading ? 'Submitting...' : children}
    </Button>
  );
};
```

This leverages your stored button component and async hook patterns!
```

## Example 2: API Integration Patterns

### Scenario
You're working on multiple projects with different API patterns and want to remember best practices.

### Store API Client Pattern

**You**: "Store this API client pattern for future projects"

```typescript
class ApiClient {
  private baseURL: string;
  private defaultHeaders: Record<string, string>;

  constructor(baseURL: string, apiKey?: string) {
    this.baseURL = baseURL;
    this.defaultHeaders = {
      'Content-Type': 'application/json',
      ...(apiKey && { 'Authorization': `Bearer ${apiKey}` })
    };
  }

  private async request<T>(
    endpoint: string, 
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseURL}${endpoint}`;
    const config: RequestInit = {
      ...options,
      headers: {
        ...this.defaultHeaders,
        ...options.headers,
      },
    };

    const response = await fetch(url, config);
    
    if (!response.ok) {
      throw new Error(`API Error: ${response.status} ${response.statusText}`);
    }

    return response.json();
  }

  async get<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: 'GET' });
  }

  async post<T>(endpoint: string, data: any): Promise<T> {
    return this.request<T>(endpoint, {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }
}
```

**Claude Response**:
```
✅ Memory stored successfully!

Memory ID: api-client-001
Keywords: typescript, api, client, fetch, error-handling, bearer-token
Tags: typescript, api, http, client, error-handling, authentication
Links Created: 0

This API client pattern with TypeScript generics and error handling 
is now stored for reuse across projects.
```

### Later: Retrieve for New Project

**You**: "I'm starting a new project that needs to call a REST API. What patterns do I have?"

**Claude Response**:
```
[Uses retrieve_relevant_memories tool]

Found 1 relevant memory:

**Memory 1** (Relevance: 94.7%)
- Context: TypeScript API client with error handling and authentication
- Keywords: typescript, api, client, fetch, error-handling, bearer-token

Here's your stored API client pattern that you can adapt for your new project:

[Shows the stored code]

You can customize this for your new project by:
1. Setting the appropriate baseURL
2. Adding any project-specific headers
3. Extending with additional HTTP methods if needed
4. Adding request/response interceptors if required

Would you like me to help you adapt this pattern for your specific API requirements?
```

## Example 3: Algorithm and Data Structure Patterns

### Store Algorithm Solutions

**You**: "Store this binary search implementation"

```python
def binary_search(arr, target):
    """
    Binary search implementation with detailed comments.
    Time complexity: O(log n), Space complexity: O(1)
    """
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = left + (right - left) // 2  # Avoid overflow
        
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return -1  # Target not found

# Example usage and test cases
def test_binary_search():
    arr = [1, 3, 5, 7, 9, 11, 13]
    assert binary_search(arr, 7) == 3
    assert binary_search(arr, 1) == 0
    assert binary_search(arr, 13) == 6
    assert binary_search(arr, 4) == -1
```

### Memory Evolution in Action

**You**: "Optimize my memory network to find better connections"

**Claude Response**:
```
[Uses evolve_memory_network tool]

✅ Memory network evolution completed!

Results:
- Memories Analyzed: 15
- Memories Evolved: 8
- Links Created: 5
- Contexts Updated: 3
- Duration: 1.8 seconds

Key improvements made:
1. Connected your binary search algorithm to your API client error handling 
   (both use similar boundary checking patterns)
2. Linked your React button component to your async patterns 
   (both handle loading states)
3. Updated context descriptions to be more searchable
4. Created cross-language pattern connections (Python algorithms ↔ TypeScript APIs)

Your memory network is now better organized for discovering related patterns 
across different technologies and problem domains.
```

## Example 4: Debugging and Error Patterns

### Store Error Solutions

**You**: "Store this solution for a tricky React useEffect dependency issue"

```jsx
// Problem: useEffect running on every render due to object dependency
// Solution: Use useCallback and useMemo to stabilize dependencies

import { useEffect, useCallback, useMemo } from 'react';

const DataFetcher = ({ userId, filters }) => {
  // Stabilize the filters object to prevent unnecessary re-renders
  const stableFilters = useMemo(() => ({
    ...filters,
    userId
  }), [filters.category, filters.dateRange, userId]); // Only specific props

  // Stabilize the fetch function
  const fetchData = useCallback(async () => {
    try {
      const response = await api.getData(stableFilters);
      setData(response);
    } catch (error) {
      setError(error.message);
    }
  }, [stableFilters]);

  useEffect(() => {
    fetchData();
  }, [fetchData]); // Now this dependency is stable

  // ... rest of component
};
```

### Retrieve Debugging Solutions

**You**: "I'm having issues with useEffect running too often. Any solutions in my memory?"

**Claude Response**:
```
[Uses retrieve_relevant_memories tool]

Found 1 highly relevant memory:

**Memory 1** (Relevance: 96.3%)
- Context: React useEffect dependency stabilization with useCallback and useMemo
- Keywords: react, useEffect, dependency, useCallback, useMemo, optimization
- Match Reason: Keyword match: useEffect, dependency

This is exactly the pattern you need! Your stored solution shows how to:

1. Use `useMemo` to stabilize object dependencies
2. Use `useCallback` to stabilize function dependencies  
3. Be specific about which object properties trigger re-computation

The key insight from your memory: instead of passing entire objects as dependencies,
destructure and specify only the specific properties that should trigger the effect.

Would you like me to help you apply this pattern to your current useEffect issue?
```

## Example 5: Cross-Project Learning

### Scenario
You've worked on multiple projects and want to leverage learnings across them.

**You**: "I'm starting a new e-commerce project. What relevant patterns do I have from previous work?"

**Claude Response**:
```
[Uses retrieve_relevant_memories tool with project scope]

Found 8 relevant memories across your projects:

**Authentication Patterns** (3 memories)
- JWT token handling with refresh logic
- Protected route components  
- User session management hooks

**API Integration** (2 memories)
- TypeScript API client with error handling
- Retry logic for failed requests

**UI Components** (2 memories)
- Reusable button components with loading states
- Form validation patterns

**State Management** (1 memory)
- Redux toolkit patterns for async actions

Based on your memory, here's a suggested architecture for your e-commerce project:

1. **Authentication**: Use your stored JWT + protected routes pattern
2. **API Layer**: Adapt your TypeScript API client for product/cart endpoints
3. **UI Components**: Leverage your button and form patterns for checkout
4. **State**: Apply your Redux patterns for cart and user state

Would you like me to help you adapt any of these specific patterns for e-commerce use cases?
```

## Best Practices for A-MEM Usage

### 1. Store Complete Context
```
❌ "Store this function"
✅ "Store this React hook for managing form validation with error handling"
```

### 2. Include Problem Context
```
❌ Just store code
✅ Store code + the problem it solves + why this approach was chosen
```

### 3. Use Descriptive Queries
```
❌ "Find React stuff"
✅ "Find React patterns for handling async data fetching with error states"
```

### 4. Regular Evolution
```
Run memory evolution weekly to optimize connections:
"Optimize my memory network focusing on recent React patterns"
```

### 5. Project-Specific Retrieval
```
"Find authentication patterns from my last e-commerce project"
"Show me error handling approaches I've used in Node.js APIs"
```

## Integration Tips

### VS Code Integration
- Use A-MEM when starting new files: "What patterns do I have for this type of component?"
- Before refactoring: "Find similar refactoring solutions I've used"
- When debugging: "Search for solutions to this type of error"

### Claude Desktop Integration  
- During architecture planning: "What architectural patterns have I used successfully?"
- For code reviews: "Find examples of how I've handled this pattern before"
- When learning: "Show me how I've implemented similar features"

## Troubleshooting Common Usage Issues

### Memory Not Found
- Use broader search terms
- Check if the memory was stored with different keywords
- Try searching by technology or problem domain

### Poor Relevance Scores
- Store more context with your code
- Use consistent terminology
- Run memory evolution to improve connections

### Too Many Results
- Use more specific queries
- Filter by project or technology
- Specify the exact problem you're solving

---

**Next Steps**: Start building your memory database by storing your most useful code patterns, then use retrieval to accelerate your development workflow!
