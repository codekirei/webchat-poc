package gensql

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Generator struct {
	Package    string
	InputGlob  string
	OutputPath string
	OutFile    *os.File
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

func (g *Generator) CreateFileAndWriteHeader() {
	f, err := os.Create(g.OutputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	headerStr := fmt.Sprintf(headerTemplate, g.Package)
	_, err = f.WriteString(headerStr)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) OpenOutFile() {
	f, err := os.OpenFile(g.OutputPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// Not calling defer f.close() here to keep reference to f open for other
	// funcs. Remember to call defer f.close() next to callsite of this func!

	g.OutFile = f
}

func (g *Generator) AppendToOut(str string) {
	_, err := g.OutFile.WriteString(str)
	if err != nil {
		panic(err)
	}
}

func (g *Generator) GetInputFiles() {
	files, _ := filepath.Glob(g.InputGlob)
	g.FilePaths = files
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

func (g *Generator) ParseInputFiles() {
	for _, v := range g.FilePaths {
		log.Printf("Processing input file: %s", v)
		constName := fileToConst(v)
		content := readFile(v)
		mContent := g.Mutate(content)
		output := fmt.Sprintf(constTemplate, constName, mContent)
		g.AppendToOut(output)
		log.Printf("Finished processing %s", v)
	}
}
