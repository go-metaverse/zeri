package logger

import (
	"log"

	"go.uber.org/zap"
)

// ZeriLogger is a global instance of a SugaredLogger for convenient logging throughout the application.
var ZeriLogger *zap.SugaredLogger

// InitLogger initializes the global ZeriLogger instance based on the provided configuration.
// It sets up a new zap.Logger according to the specified settings and options. If the logger
// initialization fails, it logs a fatal error and terminates the application.
//
// Parameters:
// - cfg: A pointer to a Config struct containing the configuration settings for the logger.
// - opts: Optional zap options for customizing the logger further.
//
// Returns:
// - A pointer to the initialized SugaredLogger instance.
// - A function that can be called to undo the global logger replacement.
//
// Usage example:
//
//	cfg := &Config{EnableDevMode: true}
//	zeriLogger, undo := InitLogger(cfg)
//	defer undo() // Revert the global logger replacement when done
//
//	zeriLogger.Info("Logger initialized successfully")
func InitLogger(cfg *Config, opts ...zap.Option) (*zap.SugaredLogger, func()) {
	logger, err := createLogger(cfg, opts...)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	undo := zap.ReplaceGlobals(logger)
	ZeriLogger = zap.S()

	return ZeriLogger, undo
}

// NewLoggerWithAttributes creates a new SugaredLogger with additional attributes added from the provided map.
// If the global ZeriLogger is not initialized, it will initialize it with default settings.
//
// Parameters:
// - attributes: A map of key-value pairs representing additional attributes to include in the logger.
//
// Returns:
// - A pointer to the SugaredLogger instance with the specified attributes.
func NewLoggerWithAttributes(attributes Attributes) *zap.SugaredLogger {
	if ZeriLogger == nil {
		InitLogger(&Config{EnableDevMode: true})
	}

	logger := ZeriLogger

	// Add any additional attributes from the map
	for key, value := range attributes {
		logger = logger.With(zap.Any(key, value))
	}

	return logger
}
