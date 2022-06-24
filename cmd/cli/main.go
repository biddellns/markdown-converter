package main

import (
	"flag"
	"fmt"
	"github.com/biddellns/markdown-converter/lib"
	"github.com/pkg/errors"
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
	sourceFilenameFlag := flag.String("in", "", "provide a file to convert")
	destinationFilenameFlag := flag.String("out", "", "save conversion to a file")
	noWrapFlag := flag.Bool("no-wrap", false, "don't wrap converted html in doc/body tags")
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

	srcFile, err := os.Open(*sourceFilenameFlag)
	if err != nil {
		return errors.Wrap(err, "opening source file")
	}
	defer srcFile.Close()

	srcStat, err := srcFile.Stat()
	if err != nil {
		srcFile.Close()
		return errors.Wrap(err, "getting srcFile stat")
	}

	if srcStat.IsDir() {
		srcFile.Close()
		return errors.New("source input cannot be a directory")
	}

	if srcStat.Size() == 0 {
		srcFile.Close()
		return errors.New("source file is empty")
	}

	destFile, err := os.Create(*destinationFilenameFlag)
	if err != nil {
		srcFile.Close()
		return errors.Wrap(err, "creating the destination file")
	}
	defer destFile.Close()

	if *noWrapFlag {
		err = lib.MarkdownToHtmlNoWrap(srcFile, destFile)
		if err != nil {
			srcFile.Close()
			destFile.Close()
			return errors.Wrap(err, "converting markdown to html")
		}

		return nil
	}

	err = lib.MarkdownToHtmlAndWrap(srcFile, destFile)
	if err != nil {
		srcFile.Close()
		destFile.Close()
		return errors.Wrap(err, "converting markdown to html")
	}

	return nil
}
