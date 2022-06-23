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
	sourceFilename := flag.String("in", "", "provide a file to convert")
	destinationFilename := flag.String("out", "", "save conversion to a file")
	helpFlag := flag.Bool("h", false, "print help")
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return nil
	}

	isFlagValueProvided := map[string]bool{
		"in":  false,
		"out": false,
	}

	flag.CommandLine.Visit(func(f *flag.Flag) {
		if _, ok := isFlagValueProvided[f.Name]; ok {
			isFlagValueProvided[f.Name] = true
		}
	})

	for flagName, isSet := range isFlagValueProvided {
		if !isSet {
			return fmt.Errorf("'%s' is required", flagName)
		}
	}

	srcFile, err := os.Open(*sourceFilename)
	if err != nil {
		return errors.Wrap(err, "opening source file")
	}
	defer srcFile.Close()

	destFile, err := os.Create(*destinationFilename)
	if err != nil {
		srcFile.Close()
		return errors.Wrap(err, "creating the destination file")
	}
	defer destFile.Close()

	err = lib.MarkdownToHtml(srcFile, destFile)
	if err != nil {
		srcFile.Close()
		destFile.Close()
		return errors.Wrap(err, "converting markdown to html")
	}

	return nil
}
