package utils

func FilterMap(original map[string]interface{}, keysToKeep []string) map[string]interface{} {
	filtered := make(map[string]interface{})
	for _, key := range keysToKeep {
		if value, exists := original[key]; exists {
			filtered[key] = value
		}
	}
	return filtered
}
