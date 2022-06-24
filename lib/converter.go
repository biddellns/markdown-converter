package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"regexp"
)

const (
	beginningHtmlBoilerplate = `<!DOCTYPE html>
<html lang="en">
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<body>
`

	endingHtmlBoilerplate = `
</body>
</html>`

	// #
	headerToken byte = 35

	// [
	linkToken byte = 91

	// '\n'
	newLine byte = 10
)

var (
	header = regexp.MustCompile(`^#{1,6}\s+(.*?)$`)

	a = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	p = regexp.MustCompile(`(.*)`)
)

func MarkdownToHtmlAndWrap(input io.Reader, output io.Writer) error {
	return markdownToHtml(input, output, true)
}

func MarkdownToHtmlNoWrap(input io.Reader, output io.Writer) error {
	return markdownToHtml(input, output, false)
}

// This wrapping logic is probably easier and cleaner if we used a template.
// The reason we went with a manual write is to handle a "large" or streaming input
//
// In most cases, a template should suffice.
func markdownToHtml(input io.Reader, output io.Writer, wrapWithHtmlSkeleton bool) error {
	if !wrapWithHtmlSkeleton {
		err := convertInput(input, output)
		if err != nil {
			return errors.Wrap(err, "converting input to html")
		}

		return nil
	}

	_, err := output.Write([]byte(beginningHtmlBoilerplate))
	if err != nil {
		return errors.Wrap(err, "writing beginning html boilerplate")
	}

	err = convertInput(input, output)
	if err != nil {
		return errors.Wrap(err, "converting input to html")
	}

	_, err = output.Write([]byte(endingHtmlBoilerplate))
	if err != nil {
		return errors.Wrap(err, "writing ending html boilerplate")
	}

	return nil
}

func convertInput(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)

	paragraphOpen := false

	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())

		if len(line) != 0 {
			firstChar := line[0]
			// If it isn't a formatting token, it's plaintext, ready for a <p> tag
			if !startsWithFormattedText(line) {
				if !paragraphOpen {
					line = p.ReplaceAll(line, []byte(`<p>$1`))
					paragraphOpen = true
				}

				// If the paragraph is open, ensure that we have a new line for the next text block.
				_, err := output.Write([]byte{newLine})
				if err != nil {
					return errors.Wrap(err, "writing output")
				}
			} else if firstChar == headerToken {
				line = convertHeader(line)
			}
			line = a.ReplaceAll(line, []byte(`<a href="$2">$1</a>`))
		} else {
			if paragraphOpen {
				line = []byte(`</p>`)
				line = append(line, newLine)
				paragraphOpen = false
			} else {
				line = []byte{newLine}
			}
		}

		_, err := output.Write(line)
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	// If last item is a paragraph, ensure it gets closed
	if paragraphOpen {
		_, err := output.Write([]byte(`</p>`))
		if err != nil {
			return errors.Wrap(err, "writing output")
		}
	}

	return nil
}

func convertHeader(line []byte) []byte {
	headerSize := 1
	for _, t := range line[1:6] {
		if t != headerToken {
			break
		}

		headerSize += 1
	}
	headerSizeStr := fmt.Sprint(headerSize)

	return header.ReplaceAll(line, []byte(`<h`+headerSizeStr+`>`+`$1`+`</h`+headerSizeStr+`>`))
}

// Used to determine if the content should be surrounded in <p> tags.
// According to the spec,
// 		A line of `[Link text](https://www.example.com)`
//		should be `<a href="https://www.example.com">Link text</a>`
func startsWithFormattedText(line []byte) bool {
	return line[0] == linkToken || header.Match(line)
}
