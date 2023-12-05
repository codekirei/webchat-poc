package gensql

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Generator struct {
	Package    string
	InputGlob  string
	OutputPath string
}

func (g *Generator) getWriter() io.Writer {
	if g.OutputPath == "" {
		return new(bytes.Buffer)
	}
	return g.openOutputFile()
}

func (g *Generator) openOutputFile() *os.File {
	f, err := os.OpenFile(
		g.OutputPath,
		os.O_CREATE|os.O_TRUNC|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		panic(err)
	}

	err = f.Truncate(0)
	if err != nil {
		f.Close()
		panic(err)
	}

	return f
}

func (g *Generator) buildHeader() string {
	header := fmt.Sprintf(headerTemplate, g.Package)
	return header
}

func (g *Generator) getInputFiles() []string {
	paths, _ := filepath.Glob(g.InputGlob)
	return paths
}

func (g *Generator) parseInputFiles(w io.Writer, paths []string) {
	for _, v := range paths {
		log.Printf("Processing input file: %s", v)
		constName := fileToConst(v)
		content := readFile(v)
		mContent := mutate(content)
		output := fmt.Sprintf(constTemplate, constName, mContent)
		writeString(w, output)
		log.Printf("Finished processing %s", v)
	}
}
