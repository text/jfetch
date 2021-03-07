package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/text/jpath"
)

func main() {
	flag.Parse()
	expression := strings.Join(flag.Args(), " ")
	fetch(expression, os.Stdin, os.Stdout, os.Stderr)
}

func fetch(expression string, in io.Reader, out io.Writer, errOut io.Writer) {
	ch, err := jpath.Evaluate(expression, in)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return
	}
	result := make([]interface{}, 0)
	for v := range ch {
		if v.Error != nil {
			fmt.Fprintln(errOut, v.Error)
		} else {
			result = append(result, v.Value)
		}
	}
	enc := json.NewEncoder(out)
	err = enc.Encode(result)
	if err != nil {
		fmt.Fprintln(errOut, err)
	}
}
