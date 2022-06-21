package lib

import (
	"bufio"
	"bytes"
	"io"
)

func MarkdownToHtml(input io.Reader) string {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		inputLine := bytes.TrimSpace(scanner.Bytes())
		_ = len(inputLine)
	}

	return ""
}
