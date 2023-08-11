# Slog-implementation
## This package provides a simple wrapper around log/slog for structured logging in Go.

## Features
- Initializes a slog Logger with default options
- Exposes slog logging methods through a custom Logger interface
- Provides context-aware logging methods
- Allows retrieving a Logger from context

## The Logger interface exposes the following methods:
- `Debug`
- `DebugContext`
- `Error`
- `ErrorContext`
- `Info`
- `InfoContext`
- `Warn`
- `WarnContext`

`The context-aware methods allow associating log messages with a context.`

`The package also provides ContextWithLogger and GetLogger functions for managing loggers in context.`

## Implementation Details
- The logger struct wraps a slog Logger and implements the custom Logger interface.
- Context values are associated with the ctxLogger{} key.
- The default logger outputs to stdout with LevelInfo.
