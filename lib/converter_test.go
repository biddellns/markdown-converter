package lib

import (
	"bytes"
	"strings"
	"testing"
)

func TestBase(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedErr    bool
	}{
		{
			name: "First, simple sample",
			input: `# Sample Document
Hello!

This is sample markdown for the [Mailchimp](https://www.mailchimp.com) homework assignment.`,
			expectedOutput: `
<h1>Sample Document</h1>

<p>Hello</p>

<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment</p>`,
			expectedErr: false,
		},
		{
			name: "Sample 2, nested, inline link",
			input: `
# Header one

Hello there

How are you?
What's going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)`,
			expectedOutput: `
<h1>Header one</h1>

<p>Hello there</p>

<p>How are you?
What's going on?</p>

<h2>Another Header</h2>

<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>

<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>`,
			expectedErr: false,
		},
	}

	for testNum, test := range tests {
		var b bytes.Buffer
		err := MarkdownToHtml(strings.NewReader(test.input), &b, RegExpLineConverter{})

		if (test.expectedErr && err == nil) || (!test.expectedErr && err != nil) {
			t.Errorf("expected err? %t. got=%v", test.expectedErr, err)
		}

		if b.String() != test.expectedOutput {
			t.Errorf("%d:%s - incorrect output. wanted=%s, got=%s",
				testNum, test.name, test.expectedOutput, b.String())
		}
	}
}
