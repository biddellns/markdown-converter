package lib

import (
	"fmt"
	"regexp"
)

var (
	h1 = regexp.MustCompile(`^#(\s|)(.*?)$`)
	h2 = regexp.MustCompile(`^##(\s|)(.*?)$`)
	h3 = regexp.MustCompile(`^###(\s|)(.*?)$`)
	h4 = regexp.MustCompile(`^####(\s|)(.*?)$`)
	h5 = regexp.MustCompile(`^#####(\s|)(.*?)$`)
	h6 = regexp.MustCompile(`^######(\s|)(.*?)$`)

	a = regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
)

type RegExpLineConverter struct{}

func (relc RegExpLineConverter) ConvertLine(input []byte) []byte {
	if len(input) <= 0 {
		return []byte{'\n'}
	}

	output := h6.ReplaceAll(input, []byte(`<h6>$2</h6>`))
	output = h5.ReplaceAll(output, []byte(`<h5>$2</h5>`))
	output = h4.ReplaceAll(output, []byte(`<h4>$2</h4>`))
	output = h3.ReplaceAll(output, []byte(`<h3>$2</h3>`))
	output = h2.ReplaceAll(output, []byte(`<h2>$2</h2>`))
	output = h1.ReplaceAll(output, []byte(`<h1>$2</h1>`))

	fmt.Println(string(output))
	output = a.ReplaceAll(output, []byte(`<a href="$2">$1</a>`))
	fmt.Println(string(output))
	return output
}
