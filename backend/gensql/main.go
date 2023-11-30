package gensql

import (
	"log"
)

type (
	Mutator func(string) (string, error)
	Opts    struct {
		MutateSql Mutator
	}
)

// TODO: allow header override
// TODO: add silent option to hide logs
func Generate(pkg string, inputPath string, outputPath string, opts Opts) {
	log.Printf("Generating %s", outputPath)

	g := CreateGenerator(pkg, inputPath, outputPath, opts)

	g.CreateFileAndWriteHeader()

	g.OpenOutFile()
	defer g.OutFile.Close()

	g.GetInputFiles()
	g.ParseInputFiles()

	log.Printf("Finished generating %s", outputPath)
}
