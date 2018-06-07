// jsonfmt reads JSON from standard input, formats it, and writes it to standard output.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("jsonfmt: ")

	if len(os.Args) > 1 {
		fmt.Fprintf(os.Stderr, "usage: jsonfmt <input >output\n")
		os.Exit(2)
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	json.Indent(&buf, b, "", "    ")
	buf.WriteByte('\n')

	os.Stdout.Write(buf.Bytes())
}
