package gensql

import (
	"testing"
)

const (
	pkg        = "helloworld"
	inputGlob  = "some/files/*"
	outputPath = "some/file.go"
)

func TestCreateGeneratorNoMut(t *testing.T) {
	out := CreateGenerator(pkg, inputGlob, outputPath, Opts{})

	gotWant(t, out.Package, pkg)
	gotWant(t, out.InputGlob, inputGlob)
	gotWant(t, out.OutputPath, outputPath)
	gotNil(t, out.Mutator)
}

func TestCreateGeneratorWithMut(t *testing.T) {
}
