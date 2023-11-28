package gensql

import (
	"reflect"
	"testing"
)

func gotWant(t *testing.T, got any, want any) {
	if got != want {
		t.Fatalf("\ngot: %v\nwant: %v", got, want)
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

func TestCamelizeNoSpChars(t *testing.T) {
	input := "afilename"
	want := input
	got := camelize(input)
	gotWant(t, got, want)
}

func TestCamelizeWithUnder(t *testing.T) {
	input := "a_file_name"
	want := "aFileName"
	got := camelize(input)
	gotWant(t, got, want)
}

func TestCamelizeWithDash(t *testing.T) {
	input := "a-file-name"
	want := "aFileName"
	got := camelize(input)
	gotWant(t, got, want)
}

func TestCamelizeWithMixed(t *testing.T) {
	input := "a_file-name"
	want := "aFileName"
	got := camelize(input)
	gotWant(t, got, want)
}

func TestPathToFile(t *testing.T) {
	input := "some/file/path/file_name.go"
	want := "file_name"
	got := pathToFile(input)
	gotWant(t, got, want)
}
