package utils

import (
	"regexp"
	"strings"
)

// NormalizeUsername converts input to lowercase, removes special chars, and replaces spaces with underscores
func NormalizeUsername(name string) string {
	// 1. trim spaces
	name = strings.TrimSpace(name)

	// 2. lowercase
	name = strings.ToLower(name)

	// 3. replace multiple spaces with a single underscore
	space := regexp.MustCompile(`\s+`)
	name = space.ReplaceAllString(name, "_")

	// 4. remove non-alphanumeric chars except underscore
	clean := regexp.MustCompile(`[^a-z0-9_]`)
	name = clean.ReplaceAllString(name, "")

	return name
}
