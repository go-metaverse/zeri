# <center>Zeri - A Comprehensive Golang SDK</center>

## Overview of Zeri

Zeri is a versatile software development kit (SDK) written in Golang, designed to streamline development by offering a comprehensive set of utilities. It includes robust support for logging, concurrency management (goroutines), JWT handling, database connections, and many other essential packages. Zeri continues to grow, incorporating new features to meet evolving development needs while remaining focused on providing specific, reliable tools for Golang projects.

## Installation

To integrate Zeri into your project, use the following command:

```
go get github.com/go-metaverse/zeri
```

## Usage Guide

For detailed usage examples, refer to the example folder within each package. Each folder contains a `main.go` file demonstrating how to effectively implement the specific functionality of that package.

- Example of using logger

```go
package main

import (
    "github.com/go-metaverse/zeri/logger"
)

func main() {
    logInstance, undo := logger.InitLogger(&logger.Config{
        // Disables caller info in logs (default: false, accepts: bool)
        DisableCaller: true,
        // Disables stack trace in logs (default: false, accepts: bool)
        DisableStacktrace: true,
        // Enables development mode (default: false, accepts: bool)
        EnableDevMode: false,
        // Sets log level (default: Info; defaults to Debug if EnableDevMode is true;
        // accepts: logger.LevelInfo, logger.LevelWarn, logger.LevelError, logger.LevelDebug)
        Level: logger.LevelDebug,
        // Sets log output format (default: JSON; defaults to CONSOLE if EnableDevMode is true;
        // accepts: logger.EncodingConsole, logger.EncodingJSON)
        Encoding: logger.EncodingConsole,
    })

    defer func() {
        _ = logInstance.Sync() // Flush any buffered log entries
    }()
    defer undo()

    // Log with the global logger; ensure InitLogger is called beforehand.
    logger.ZeriLogger.Info("Log with the global logger")

    log := logger.NewLoggerWithAttributes(logger.Attributes{
        "app_name": "zeri",
        "version":  "1.0.0",
    })
    log.Debug("Debug message...") // Logged only in development mode or when log level is set to Debug
    log.Info("Info message...")
    log.Warn("Warn message...")
    log.Error("Error message...")
}
```

## License

Zeri is open-source and available under the MIT license. See the LICENSE file for more details.
