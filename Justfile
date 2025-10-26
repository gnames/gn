# Justfile for gn project

# Run all tests
test:
    go test -v -race -coverprofile=coverage.out ./...

# Run tests with coverage report
test-coverage: test
    go tool cover -html=coverage.out -o coverage.html
    @echo "Coverage report generated: coverage.html"

# Run tests and show coverage percentage
test-cover:
    go test -cover ./...

# Run tests without verbose output
test-quiet:
    go test ./...

# Run tests with specific timeout
test-timeout timeout="5m":
    go test -timeout {{timeout}} -v ./...

# Run benchmarks
bench:
    go test -bench=. -benchmem ./...

# Clean test cache and coverage files
clean:
    go clean -testcache
    rm -f coverage.out coverage.html

# Run go vet
vet:
    go vet ./...

# Run go fmt
fmt:
    go fmt ./...

# Run all checks (fmt, vet, test)
check: fmt vet test

# Show help
help:
    @just --list
