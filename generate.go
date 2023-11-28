//go:build ignore

package main

import (
	"strings"

	"github.com/codekirei/webchat-poc/backend/gensql"
)

func mutate(str string) (string, error) {
	return strings.ReplaceAll(str, "'?'", "?"), nil
}

func main() {
	gensql.Generate(
		"sqlite",
		"backend/db/sqlite/queries/*",
		"backend/db/sqlite/queries.go",
		gensql.Opts{
			MutateSql: mutate,
		},
	)
}
