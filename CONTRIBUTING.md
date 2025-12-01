# Contributing to go-structurarium

Thank you for your interest in contributing to go-structurarium! This document provides guidelines for contributing to this project.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/go-structurarium.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`

## Development Setup

### Prerequisites

- Go 1.18 or higher
- Git

### Setting Up Your Environment

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/go-structurarium.git
cd go-structurarium

# Install dependencies
go mod download

# Run tests to ensure everything works
go test ./...
```

## Making Changes

### Branch Naming Convention

- `feature/` - New features (e.g., `feature/add-linked-list`)
- `fix/` - Bug fixes (e.g., `fix/stack-overflow`)
- `docs/` - Documentation updates (e.g., `docs/update-readme`)
- `test/` - Test additions/improvements (e.g., `test/add-hashmap-tests`)
- `refactor/` - Code refactoring (e.g., `refactor/simplify-queue`)

### Code Style

- Follow standard Go conventions
- Run `go fmt` before committing
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and concise

### Testing

- Write tests for all new features
- Ensure all existing tests pass
- Aim for high test coverage
- Include edge cases in tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run tests with race detection
go test -race ./...
```

### Commit Messages

Write clear, concise commit messages:

```
Good examples:
✅ Add linked list implementation
✅ Fix stack overflow in Pop operation
✅ Update README with new examples
✅ Add tests for HashMap edge cases

Bad examples:
❌ Update
❌ Fix bug
❌ Changes
❌ WIP
```

## Pull Request Process

1. **Update your branch** with the latest main:
   ```bash
   git checkout main
   git pull upstream main
   git checkout your-branch
   git rebase main
   ```

2. **Run all tests** and ensure they pass:
   ```bash
   go test ./...
   ```

3. **Push your changes**:
   ```bash
   git push origin your-branch
   ```

4. **Create Pull Request**:
   - Go to the repository on GitHub
   - Click "Compare & pull request"
   - Fill in the PR template with:
     - Clear description of changes
     - Why the change is needed
     - Any breaking changes
     - How to test the changes

5. **Address review comments**:
   - Be responsive to feedback
   - Make requested changes
   - Push updates to the same branch

6. **Wait for approval**:
   - At least one approval is required
   - All tests must pass
   - No unresolved conversations

## Pull Request Guidelines

### Good Pull Requests

- ✅ Focused on a single feature/fix
- ✅ Include tests
- ✅ Update documentation if needed
- ✅ Have clear commit messages
- ✅ Pass all tests
- ✅ Follow code style guidelines

### What to Avoid

- ❌ Multiple unrelated changes in one PR
- ❌ Missing tests
- ❌ Breaking existing functionality
- ❌ Ignoring code review feedback
- ❌ Force-pushing after review has started

## Code Review

All submissions require review. We use GitHub pull requests for this purpose.

**As a Contributor:**
- Be patient and respectful
- Accept feedback gracefully
- Explain your reasoning when appropriate
- Be willing to make changes

**As a Reviewer:**
- Be constructive and respectful
- Focus on the code, not the person
- Explain why changes are needed
- Approve when satisfied

## Testing Checklist

Before submitting a PR, ensure:

- [ ] All tests pass: `go test ./...`
- [ ] Code is formatted: `go fmt ./...`
- [ ] No linting errors: `go vet ./...`
- [ ] Race detector passes: `go test -race ./...`
- [ ] New code has tests
- [ ] Documentation is updated

## Adding New Data Structures

When adding a new data structure:

1. Create a new package directory
2. Implement private methods with channel-based synchronization
3. Create public wrapper methods
4. Write comprehensive tests (both internal and wrapper)
5. Add documentation comments
6. Update README.md with the new data structure
7. Add usage examples to USAGE.md

### Example Structure

```
newstructure/
├── newstructure.go           # Core implementation
├── newstructure_wrapper.go   # Public wrapper
├── newstructure_test.go      # Internal tests
└── newstructure_wrapper_test.go  # Wrapper tests
```

## Reporting Issues

### Bug Reports

Include:
- Go version
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Code sample (if applicable)

### Feature Requests

Include:
- Clear description of the feature
- Use cases
- Why it would be beneficial
- Possible implementation approach (optional)

## Questions?

If you have questions:
- Check existing issues and pull requests
- Review the documentation
- Open a new issue with your question

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (MIT License).

## Thank You!

Your contributions help make go-structurarium better for everyone. We appreciate your time and effort! 🎉

