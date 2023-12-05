package gensql

import (
	"log"
	"os"
)

func Generate(pkg string, inputPath string, outputPath string) {
	log.Printf("Generating %s", outputPath)

	g := &Generator{pkg, inputPath, outputPath}

	w := g.getWriter()

	f, isFile := w.(*os.File)
	if isFile {
		defer f.Close()
	}

	header := g.buildHeader()
	writeString(w, header)

	paths := g.getInputFiles()
	g.parseInputFiles(w, paths)

	log.Printf("Finished generating %s", outputPath)
}
