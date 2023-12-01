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
	callCount := 0
	mut := func(str string) (string, error) {
		callCount += 1
		return str, nil
	}

	g := CreateGenerator(pkg, inputGlob, outputPath, Opts{
		MutateSql: mut,
	})

	gotWant(t, callCount, 0)
	g.Mutator("hello world")
	gotWant(t, callCount, 1)
}

func TestCreateFileAndWriteHeader(t *testing.T) {
}

func TestMutate(t *testing.T) {
}

func TestOpenAndAppend(t *testing.T) {
}
