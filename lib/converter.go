package lib

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io"
)

type Converter interface {
	ConvertLine(input []byte) []byte
}

func MarkdownToHtml(input io.Reader, output io.Writer, converter Converter) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		inputLine := bytes.TrimSpace(scanner.Bytes())
		inputLine = converter.ConvertLine(inputLine)
		_, err := output.Write(inputLine)
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	return nil
}
