---
name: Bug Report
about: Create a report to help us improve
title: '[BUG] '
labels: ['bug']
assignees: ''
---

## Bug Description
A clear and concise description of what the bug is.

## Steps to Reproduce
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

## Expected Behavior
A clear and concise description of what you expected to happen.

## Actual Behavior
A clear and concise description of what actually happened.

## Environment Information
- **OS**: [e.g. macOS, Windows, Linux]
- **Go Version**: [e.g. 1.21.0]
- **SDK Version**: [e.g. v1.0.0]
- **Dify Platform Version**: [e.g. v0.6.0]

## Code Example
```go
// Please provide a minimal code example that reproduces the issue
package main

import (
    "github.com/kingfs/godify"
)

func main() {
    client := dify.NewServiceClient("your-token", "https://api.dify.ai")
    // Your code here
}
```

## Error Messages
```
// Please paste any error messages or stack traces here
```

## Additional Context
Add any other context about the problem here, such as:
- Screenshots
- Log files
- Configuration files

## Checklist
- [ ] I have searched existing issues to avoid duplicates
- [ ] I have provided a minimal reproduction example
- [ ] I have included all relevant environment information
- [ ] I have tested with the latest version of the SDK