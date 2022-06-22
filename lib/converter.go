package lib

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io"
)

type Converter interface {
	ConvertLine(input []byte) error
}

func MarkdownToHtml(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		inputLine := bytes.TrimSpace(scanner.Bytes())

		inputLine = append(inputLine, '\n')

		_, err := output.Write(inputLine)
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	return nil
}
