package logger

import (
	"time"

	"github.com/go-metaverse/zeri/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// createLogger initializes a new zap.Logger based on the provided configuration.
// It sets up the logger according to the specified settings for development or production
// mode, configures the log level, and defines the encoder settings for formatting log messages.
//
// Parameters:
// - config: A pointer to a Config struct containing the configuration settings for the logger.
// - opts: Optional zap options for further customization of the logger.
//
// Returns:
// - A pointer to the initialized zap.Logger instance.
// - An error if the logger could not be initialized.
func createLogger(config *Config, opts ...zap.Option) (*zap.Logger, error) {
	// Get the appropriate zap config (development or production)
	zapConfig := getZapConfigByMode(config.EnableDevMode)

	// Set log level
	if !config.EnableDevMode {
		zapConfig.Level = zap.NewAtomicLevelAt(logLevelMap[utils.DefaultIfEmpty(config.Level, LevelInfo)])
	}

	// Caller and stacktrace configuration
	zapConfig.DisableCaller = config.DisableCaller
	zapConfig.DisableStacktrace = config.DisableStacktrace

	// Encode configuration
	zapConfig.Encoding = string(getEncodeByMode(config.EnableDevMode, config.Encoding))
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:    "message",
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "log",
		CallerKey:     utils.OptionalKey(config.DisableCaller, "caller"),
		StacktraceKey: utils.OptionalKey(config.DisableStacktrace, "stacktrace"),
		EncodeLevel:   getEncodeLevelByMode(config.EnableDevMode),
		EncodeTime:    getTimeEncoder,
		EncodeCaller:  zapcore.FullCallerEncoder,
	}

	return zapConfig.Build(opts...)
}

// getZapConfigByMode returns the appropriate zap.Config based on the specified mode.
// It configures the logger for either development or production use.
//
// Parameters:
// - devMode: A boolean indicating whether to use development mode (true) or production mode (false).
//
// Returns:
// - A zap.Config instance configured for the specified mode.
func getZapConfigByMode(devMode bool) zap.Config {
	if devMode {
		return zap.NewDevelopmentConfig()
	}
	return zap.NewProductionConfig()
}

func getEncodeByMode(devMode bool, encode EncodingType) EncodingType {
	if devMode {
		return EncodingConsole
	}
	return utils.DefaultIfEmpty(encode, EncodingJSON)
}

// getEncodeLevelByMode returns the appropriate level encoder based on the specified mode.
// It configures the logging level encoding to be either colored (for development mode)
// or standard (for production mode).
//
// Parameters:
// - devMode: A boolean indicating whether to use development mode (true) or production mode (false).
//
// Returns:
// - A zapcore.LevelEncoder configured for the specified mode.
func getEncodeLevelByMode(devMode bool) zapcore.LevelEncoder {
	if devMode {
		return zapcore.CapitalColorLevelEncoder
	}
	return zapcore.CapitalLevelEncoder
}

// getTimeEncoder is a custom time encoder used for logging timestamps in a specific format.
// It formats the given time as a string using RFC3339 format and appends it to the
// provided PrimitiveArrayEncoder.
//
// Parameters:
// - t: The time to be formatted and encoded.
// - enc: The zapcore.PrimitiveArrayEncoder to which the formatted time string will be appended.
func getTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}
