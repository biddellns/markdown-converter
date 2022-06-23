package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"regexp"
)

var (

	// #
	headerToken byte = 35

	// [
	linkToken byte = 91

	header = regexp.MustCompile(`^#{1,6}(\s|)(.*?)$`)

	a = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	p = regexp.MustCompile(`(.*)`)
)

func isFormattingToken(char byte) bool {
	return char == headerToken ||
		char == linkToken
}

func MarkdownToHtmlAndWrap(input io.Reader, output io.Writer) error {
	return markdownToHtml(input, output, true)
}

func MarkdownToHtmlNoWrap(input io.Reader, output io.Writer) error {
	return markdownToHtml(input, output, false)
}

func convertInput(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	ParagraphOpen := false

	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) != 0 {
			// Currently the only non-paragraph opening are headers and anchors
			if !isFormattingToken(line[0]) {
				if !ParagraphOpen {
					line = p.ReplaceAll(line, []byte(`<p>$1`))
					ParagraphOpen = true
				}

				// If the paragraph is open, ensure that we have a new line for the next text block.
				_, err := output.Write([]byte{'\n'})
				if err != nil {
					return errors.Wrap(err, "writing output")
				}
			} else if line[0] == headerToken {
				headerSize := 1
				for _, t := range line[1:6] {
					if t != headerToken {
						break
					}

					headerSize += 1
				}
				headerSizeStr := fmt.Sprint(headerSize)

				line = header.ReplaceAll(line, []byte(`<h`+headerSizeStr+`>`+`$2`+`</h`+headerSizeStr+`>`))
			}
			line = a.ReplaceAll(line, []byte(`<a href="$2">$1</a>`))
		} else {
			if ParagraphOpen {
				line = []byte(`</p>`)
				line = append(line, '\n')
				ParagraphOpen = false
			} else {
				line = []byte{'\n'}
			}
		}

		_, err := output.Write(line)
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	// If last item is a paragraph, ensure it gets closed
	if ParagraphOpen {
		_, err := output.Write([]byte(`</p>`))
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	return nil
}

func markdownToHtml(input io.Reader, output io.Writer, wrapWithHtmlSkeleton bool) error {
	if wrapWithHtmlSkeleton {
		_, err := output.Write([]byte(
			`<!DOCTYPE html>
<html lang="en">
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<body>`))
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	err := convertInput(input, output)
	if err != nil {
		return errors.Wrap(err, "converting input to html")
	}

	if wrapWithHtmlSkeleton {
		_, err := output.Write([]byte(
			`
</body>
</html>`))
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	return nil
}
