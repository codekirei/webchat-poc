package gensql

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
)

type Generator struct {
	Package    string
	InputGlob  string
	OutputPath string
	FilePaths  []string
	Mutator    Mutator
}

func CreateGenerator(pkg string, inputGlob string, outputPath string, opts Opts) *Generator {
	gen := &Generator{
		Package:    pkg,
		InputGlob:  inputGlob,
		OutputPath: outputPath,
	}

	if opts.MutateSql != nil {
		gen.Mutator = opts.MutateSql
	}

	return gen
}

func (g *Generator) WriteHeader(w io.Writer) string {
	header := fmt.Sprintf(headerTemplate, g.Package)
	writeString(w, header)
	return header
}

func (g *Generator) GetInputFiles() []string {
	files, _ := filepath.Glob(g.InputGlob)
	g.FilePaths = files
	return files
}

func (g *Generator) Mutate(str string) string {
	if g.Mutator == nil {
		return str
	}

	mStr, err := g.Mutator(str)
	if err != nil {
		panic(err)
	}

	return mStr
}

func (g *Generator) ParseInputFiles(w io.Writer) {
	for _, v := range g.FilePaths {
		log.Printf("Processing input file: %s", v)
		constName := fileToConst(v)
		content := readFile(v)
		mContent := g.Mutate(content)
		output := fmt.Sprintf(constTemplate, constName, mContent)
		writeString(w, output)
		log.Printf("Finished processing %s", v)
	}
}
