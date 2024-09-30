package utils

import "github.com/go-metaverse/zeri/validate"

// DefaultIfEmpty returns the fallback value if the provided value is considered empty or zero.
// It uses the IsZero function to determine if the value is empty.
//
// Parameters:
// - value: The value to check for emptiness.
// - fallback: The value to return if the provided value is empty.
//
// Returns:
// - The original value if it is not empty; otherwise, the fallback value.
func DefaultIfEmpty[T any](value any, fallback T) T {
	if validate.IsZero(value) {
		return fallback
	}

	return value.(T)
}

// OptionalKey returns the provided key if the disabled flag is false.
// If the disabled flag is true, it returns an empty string.
//
// Parameters:
// - disabled: A boolean indicating whether the key should be considered disabled.
// - key: The key to return if not disabled.
//
// Returns:
// - The key if not disabled; otherwise, an empty string.
func OptionalKey(disabled bool, key string) string {
	if disabled {
		return ""
	}
	return key
}

// GetFromInterface retrieves a value from a map using the provided key.
// If the key does not exist or if the value is zero (and checkZeroValue is true),
// it returns the default value.
//
// Parameters:
// - src: A map of string keys to any values.
// - key: The key to look up in the map.
// - defaultValue: The value to return if the key does not exist or if the value is zero.
// - checkZeroValue: An optional boolean flag to check if zero values should be considered invalid.
//
// Returns:
// - The value from the map if it exists and is not zero; otherwise, the default value.
func GetFromInterface[T any](src map[string]any, key string, defaultValue T, checkZeroValue ...bool) T {
	value, exists := src[key]
	if !exists || (validate.IsZero(value) && len(checkZeroValue) > 0 && checkZeroValue[0]) {
		return defaultValue
	}
	return value.(T)
}
