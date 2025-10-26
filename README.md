# gn

User-friendly messaging and error handling for Go with colorized output.

## Installation

```bash
go get github.com/gnames/gn
```

## Usage

### Messages

```go
import "github.com/gnames/gn"

gn.Message("Processing data...")           // no icon
gn.Info("Server started on port %d", 8080) // ℹ️
gn.Warn("Connection timeout")              // ⚠️
gn.Success("All tests passed!")            // ✅
gn.Progress("Downloading files...")        // ⏳
```

### Inline Formatting

```go
gn.Info("Starting <title>Production Server</title>")
gn.Message("Found <em>42</em> records")
gn.Info("Status: <warn>degraded</warn>")
gn.Info("Errors: <err>3</err>")
```

Available tags:
- `<title>...</title>` - green with `**` emphasis
- `<em>...</em>` - green
- `<warn>...</warn>` - yellow
- `<err>...</err>` - red

### Custom Errors

```go
const (
    ErrDatabase gn.ErrorCode = 1000
    ErrNetwork  gn.ErrorCode = 2000
)

err := &gn.Error{
    Code: ErrDatabase,
    Err:  errors.New("connection failed"),
    Msg:  "Could not connect to database: %s",
    Vars: []any{"postgres"},
}

gn.PrintErrorMessage(err) // ❌ Could not connect to database: postgres

// Works with standard error handling
var gnErr *gn.Error
if errors.As(err, &gnErr) {
    fmt.Printf("Error code: %d\n", gnErr.Code)
}
```

## License

See LICENSE file.
