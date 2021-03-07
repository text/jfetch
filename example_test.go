package main

import (
	"os"
	"strings"
)

func Example() {
	in := strings.NewReader(`{
		"kind": "metasyntactic",
		"variables": [
			"foo",
			"bar",
			"baz"
		]
	}`)
	fetch(".variables[:]", in, os.Stdout, os.Stdout)
	// Output:
	// ["foo","bar","baz"]
}
