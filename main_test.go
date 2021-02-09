package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iancoleman/strcase"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	files, err := ioutil.ReadDir("testdata")
	require.NoError(t, err)

	var tests []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".go" {
			continue
		}
		if strings.HasPrefix(filepath.Base(file.Name()), "_") {
			continue
		}
		tests = append(tests, file.Name())
	}

	for _, test := range tests {
		t.Run(test, testFile(test).run)
	}
}

type testFile string

func (f testFile) run(t *testing.T) {
	c := config{
		Output:           "_ignore.go",
		Input:            "testdata",
		OptionTypename:   "Option",
		OptionFuncPrefix: "Option",
		ignoreOutput:     true,
		StructName:       strcase.ToCamel(strings.TrimSuffix(string(f), ".go")),
	}

	g := generator{
		config: c,
	}

	err := g.Run()
	require.NoError(t, err)
}
