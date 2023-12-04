package gensql

import (
	"log"
	"os"
)

type (
	Mutator func(string) (string, error)
	Opts    struct {
		MutateSql Mutator
	}
)

// TODO: allow header override
// TODO: add varNameSuffix
// TODO: add silent option to hide logs
func Generate(pkg string, inputPath string, outputPath string, opts Opts) {
	log.Printf("Generating %s", outputPath)

	g := CreateGenerator(pkg, inputPath, outputPath, opts)

	f, err := os.OpenFile(
		g.OutputPath,
		os.O_CREATE|os.O_TRUNC|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	truncateFile(f)
	g.WriteHeader(f)

	g.GetInputFiles()
	g.ParseInputFiles(f)

	log.Printf("Finished generating %s", outputPath)
}
