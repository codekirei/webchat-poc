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
// TODO: allow passing in mutation func as arg
func Generate(pkg string, inputPath string, outputPath string, opts Opts) {
	log.Printf("Generating %s", outputPath)
	gen := CreateGenerator(pkg, inputPath, outputPath, opts)

	gen.CreateFileAndWriteHeader()

	gen.OpenOutFile()
	defer gen.OutFile.Close()

	gen.GetInputFiles()
	gen.ParseInputFiles()

	log.Printf("Finished generating %s", outputPath)
}
