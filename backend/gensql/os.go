package gensql

import (
	"os"
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
