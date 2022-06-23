package lib

import (
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
	p = regexp.MustCompile(`(.*)`)
)

type RegExpLineConverter struct{
	paragraphStarted bool
}

func (relc RegExpLineConverter) ConvertLine(input []byte) []byte {
	if len(input) == 0 {
		return []byte{'\n'}
	}

	if input[0] != '#' {
		input = p.ReplaceAll(input, []byte(`<p>$1`))
		relc.paragraphStarted = true
	} else {
		if relc.paragraphStarted = 
		input = h6.ReplaceAll(input, []byte(`<h6>$2</h6>`))
		input = h5.ReplaceAll(input, []byte(`<h5>$2</h5>`))
		input = h4.ReplaceAll(input, []byte(`<h4>$2</h4>`))
		input = h3.ReplaceAll(input, []byte(`<h3>$2</h3>`))
		input = h2.ReplaceAll(input, []byte(`<h2>$2</h2>`))
		input = h1.ReplaceAll(input, []byte(`<h1>$2</h1>`))
	}

	input = a.ReplaceAll(input, []byte(`<a href="$2">$1</a>`))

	return input
}
