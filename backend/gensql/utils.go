package gensql

import (
	"io"
	"path/filepath"
	"strings"
)

func writeString(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		panic(err)
	}
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

func pathToFile(path string) string {
	basename := filepath.Base(path)
	filename := strings.TrimSuffix(basename, filepath.Ext(basename))
	return filename
}

func fileToConst(path string) string {
	filename := pathToFile(path)
	return camelize(filename) + "Stmt"
}

func mutate(str string) string {
	return strings.ReplaceAll(str, "'?'", "?")
}
