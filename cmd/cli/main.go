package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"interview-markdown-converter/lib"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error during execution: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}
}

func run() error {
	filename := flag.String("filename", "", "used to provide a file to convert")
	flag.Parse()

	areAllRequiredFlagsSet := false
	flag.CommandLine.Visit(func(f *flag.Flag) {
		if f.Name == "filename" {
			areAllRequiredFlagsSet = true
		}
	})

	if !areAllRequiredFlagsSet {
		return errors.New("required parameters not passed in")
	}

	file, err := os.Open(*filename)
	if err != nil {
		errors.Wrap(err, "opening file")
	}

	lib.MarkdownToHtml(file)
	return nil
}
