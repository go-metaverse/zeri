package tag

import (
	"strings"
)

// ParseTag parses a string input with key-value pairs separated by a custom separator.
// The key-value pairs are extracted and returned in a map. If a key has no value, an empty string is assigned.
//
// Parameters:
// - input: The string containing key-value pairs (e.g., "key1=value1;key2=value2").
// - separator: The character or string used to separate the key-value pairs in the input.
//
// Returns:
//   - map[string]string: A map where each key corresponds to a key in the input, and each value corresponds
//     to its associated value. If a key has no value, the value in the map is an empty string.
//
// Example usage:
//
//	settings := ParseTag("key1=value1;key2=value2", ";")
//	fmt.Println(settings)
//	// Output: map[key1:value1 key2:value2]
func ParseTag(input string, separator string) map[string]string {
	settings := make(map[string]string)
	parts := splitEscaped(input, separator)

	for _, part := range parts {
		key, value := extractKeyValue(part)
		if key != "" {
			settings[key] = value
		}
	}

	return settings
}

// splitEscaped splits the input string by the specified separator, handling escaped separators.
func splitEscaped(input string, separator string) []string {
	var result []string
	parts := strings.Split(input, separator)

	for i := 0; i < len(parts); i++ {
		current := parts[i]

		// Handle escaped separators
		for len(current) > 0 && current[len(current)-1] == '\\' {
			if i+1 < len(parts) {
				current = current[:len(current)-1] + separator + parts[i+1]
				i++
			} else {
				break
			}
		}
		result = append(result, current)
	}

	return result
}

// extractKeyValue extracts the key and value from a string in the format "key:value".
func extractKeyValue(part string) (string, string) {
	kv := strings.SplitN(part, ":", 2)
	key := strings.TrimSpace(strings.ToUpper(kv[0]))
	value := ""

	if len(kv) > 1 {
		value = strings.TrimSpace(kv[1])
	}

	if value == "" {
		value = key // Default value is the key itself if the value is empty
	}

	return key, value
}
