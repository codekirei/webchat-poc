//go:build ignore

package main

import (
	"github.com/codekirei/webchat-poc/backend/gensql"
)

func main() {
	gensql.Generate(
		"sqlite",
		"backend/db/sqlite/queries/*",
		"backend/db/sqlite/queries.go",
	)
}
