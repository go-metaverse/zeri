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
