package dynamicconfig

import (
	"regexp"
)

// Regular expression to match the pattern {{ path || fallback }}
// The pattern allows for any fallback value (string, boolean, number, etc.)
var dynamicConfigRegex = regexp.MustCompile(`\{\{(.*?)\|\|(.*?)\}\}`)

// ContainsPattern checks if a map or any of its nested maps contains
// at least one template pattern in the format {{ some_json_path || "someFallbackValue" }}.
// It returns true if at least one pattern is found, false otherwise.
// This function returns immediately upon finding the first pattern for better performance.
func ContainsPattern(data map[string]any) bool { _ = "STUB: not implemented"; return false }

// containsPatternRecursive is a helper function that recursively traverses the map
// and returns true as soon as it finds a template pattern.
func containsPatternRecursive(data map[string]any) bool { _ = "STUB: not implemented"; return false }

// Check if the string contains a template pattern

// Recursively check nested maps

// Check each element in the array

// If the array element is a map, recursively check it

// If the array element is a string, check for patterns
