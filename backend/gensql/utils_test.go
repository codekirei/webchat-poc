package gensql

import (
	"bytes"
	"reflect"
	"testing"
)

func gotWant(t *testing.T, got any, want any) {
	if got != want {
		t.Fatalf("\ngot: %v\nwant: %v", got, want)
	}
}

func inputGotWant(t *testing.T, input any, got any, want any) {
	if got != want {
		t.Fatalf("\ninput: %v\ngot: %v\nwant: %v", input, got, want)
	}
}

func isNil(val any) bool {
	if val == nil {
		return true
	}

	nillableKinds := map[reflect.Kind]bool{
		reflect.Chan:      true,
		reflect.Func:      true,
		reflect.Interface: true,
		reflect.Map:       true,
		reflect.Ptr:       true,
		reflect.Slice:     true,
	}

	rval := reflect.ValueOf(val)
	kind := rval.Kind()
	if nillableKinds[kind] {
		return rval.IsNil()
	}

	return false
}

func gotNil(t *testing.T, got any) {
	if !isNil(got) {
		t.Fatalf("\ngot: %v\nwant: <nil>", got)
	}
}

func TestWriteString(t *testing.T) {
	b := new(bytes.Buffer)
	s := "hello world"
	writeString(b, s)
	gotWant(t, b.String(), s)
}

var camelizeCases = []struct {
	input string
	want  string
}{
	{"afilename", "afilename"},
	{"a_file_name", "aFileName"},
	{"a-file-name", "aFileName"},
	{"a_file-name", "aFileName"},
}

func TestCamelize(t *testing.T) {
	for _, c := range camelizeCases {
		inputGotWant(t, c.input, camelize(c.input), c.want)
	}
}

func TestPathToFile(t *testing.T) {
	input := "some/file/path/file_name.go"
	want := "file_name"
	got := pathToFile(input)
	gotWant(t, got, want)
}

func TestFileToConst(t *testing.T) {
	input := "some/file_name.go"
	want := "fileNameStmt"
	got := fileToConst(input)
	gotWant(t, got, want)
}

func TestMutate(t *testing.T) {
	input := "select * from table where field = '?' and other_field = '?'"
	want := "select * from table where field = ? and other_field = ?"
	got := mutate(input)
	gotWant(t, got, want)
}
