package dicon

import (
	"bytes"
	"testing"

	"github.com/andreyvit/diff"
)

func TestGenerator_AppendHeader(t *testing.T) {
	ex := pretty(t, []byte(`// Code generated by "dicon"; DO NOT EDIT.

	package main

	import (
		"log"
		"github.com/pkg/errors"
	)
`))
	g := &Generator{PackageName: "main"}
	it := &InterfaceType{
		PackageName: "main",
	}
	g.AppendHeader(it)
	act := pretty(t, g.buf.Bytes())
	if !bytes.Equal(act, ex) {
		t.Errorf("Not Matched: \n%v", diff.LineDiff(string(ex), string(act)))
	}
}
