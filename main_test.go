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
		basename := filepath.Base(file.Name())
		if filepath.Ext(basename) != ".go" ||
			strings.HasPrefix(basename, "_") ||
			strings.HasSuffix(basename, "_options.go") {
			continue
		}

		tests = append(tests, file.Name())
	}

	for _, test := range tests {
		t.Run(test, testFile(test).Helper)
	}
}

type testFile string

func (f testFile) Helper(t *testing.T) {
	t.Helper()
	c := config{
		Input:            "testdata",
		OptionTypename:   "Option",
		OptionFuncPrefix: "Option",
		// concreteOutput:   os.Stdout,
		StructName: strcase.ToCamel(strings.TrimSuffix(string(f), ".go")),
	}

	g := generator{
		config: c,
	}

	err := g.Run()
	require.NoError(t, err)
}
