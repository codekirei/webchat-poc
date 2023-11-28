package gensql

import (
	"os"
	"path/filepath"
	"strings"
)

// This isn't efficient for very large files, but that should be okay for the
// intended use case. As a tradeoff, this implementation is much simpler than
// parsing each file in chunks.
func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func camelize(str string) string {
	var out strings.Builder
	shouldUpper := false

	if !strings.ContainsAny(str, "_-") {
		return str
	}

	for i, char := range str {
		if i == 0 {
			out.WriteRune(char)
			continue
		}

		if char == '-' || char == '_' {
			shouldUpper = true
			continue
		}

		if shouldUpper {
			out.WriteString(strings.ToUpper(string(char)))
			shouldUpper = false
			continue
		}

		out.WriteRune(char)
	}

	return out.String()
}

func fileToConst(path string) string {
	basename := filepath.Base(path)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))
	return camelize(filename)
}
