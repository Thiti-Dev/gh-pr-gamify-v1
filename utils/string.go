package utils

import "strings"

// TruncateString shortens a string to the specified length and appends "..." if truncated.
func TruncateString(s string, maxLength int) string {
	if len(s) <= maxLength || maxLength <= 3 {
		return s
	}

	// Ensure we have space for "..."
	return strings.TrimSpace(s[:maxLength-3]) + "..."
}
