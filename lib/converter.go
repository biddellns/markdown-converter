package lib

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"io"
	"regexp"
)

const headingToken = '#'
const linkToken = '['

var (
	h  = regexp.MustCompile(`^(#{1,6})(\s|)(.*?)$`)
	h1 = regexp.MustCompile(`^#(\s|)(.*?)$`)
	h2 = regexp.MustCompile(`^##(\s|)(.*?)$`)
	h3 = regexp.MustCompile(`^###(\s|)(.*?)$`)
	h4 = regexp.MustCompile(`^####(\s|)(.*?)$`)
	h5 = regexp.MustCompile(`^#####(\s|)(.*?)$`)
	h6 = regexp.MustCompile(`^######(\s|)(.*?)$`)

	a = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	p = regexp.MustCompile(`(.*)`)
)

func isFormattingToken(char byte) bool {
	return char == headingToken ||
		char == linkToken
}

func MarkdownToHtml(input io.Reader, output io.Writer) error {
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

			} else {
				line = h6.ReplaceAll(line, []byte(`<h6>$2</h6>`))
				line = h5.ReplaceAll(line, []byte(`<h5>$2</h5>`))
				line = h4.ReplaceAll(line, []byte(`<h4>$2</h4>`))
				line = h3.ReplaceAll(line, []byte(`<h3>$2</h3>`))
				line = h2.ReplaceAll(line, []byte(`<h2>$2</h2>`))
				line = h1.ReplaceAll(line, []byte(`<h1>$2</h1>`))

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
