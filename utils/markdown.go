package utils

import "regexp"

// ExtractFrontMatter Extract markdown header information and return clean text
func ExtractFrontMatter(data []byte) ([]byte, []byte) {
	re := regexp.MustCompile(`(?s)^\s*---\s*\n.*?\n\s*---\s*\n`)
	return re.FindSubmatch(data)[0], re.ReplaceAll(data, nil)
}
