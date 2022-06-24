package lib

import (
	"bytes"
	"strings"
	"testing"
)

func TestTags(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		{
			name:           "H1",
			input:          "# Header1",
			expectedOutput: "<h1>Header1</h1>",
		},
		{
			name:           "H2",
			input:          "## Header2",
			expectedOutput: "<h2>Header2</h2>",
		},
		{
			name:           "H3",
			input:          "### Header3",
			expectedOutput: "<h3>Header3</h3>",
		},
		{
			name:           "H4",
			input:          "#### Header4",
			expectedOutput: "<h4>Header4</h4>",
		},
		{
			name:           "H5",
			input:          "##### Header5",
			expectedOutput: "<h5>Header5</h5>",
		},
		{
			name:           "H6",
			input:          "###### Header6",
			expectedOutput: "<h6>Header6</h6>",
		},
		{
			name:           "Header must have space between pound sign and content",
			input:          "#content",
			expectedOutput: "<p>#content</p>",
		},
		{
			name:           "A",
			input:          "[myBlog](https://aws.nicbiddell.com/blog)",
			expectedOutput: `<a href="https://aws.nicbiddell.com/blog">myBlog</a>`,
		},
		{
			name:           "P",
			input:          "something something something",
			expectedOutput: "<p>something something something</p>",
		},
	}

	for testNum, test := range tests {
		var b bytes.Buffer
		MarkdownToHtmlNoWrap(strings.NewReader(test.input), &b)

		actual := strings.TrimSpace(b.String())
		if actual != test.expectedOutput {
			t.Errorf("%d:%s - incorrect output. wanted=%s, got=%s",
				testNum, test.name, test.expectedOutput, actual)
		}
	}
}

func TestSampleDocs(t *testing.T) {
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
			expectedOutput: `<h1>Sample Document</h1>
<p>Hello!</p>
<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment.</p>`,
			expectedErr: false,
		},
		{
			name: "Sample 2, nested, inline link",
			input: `# Header one

Hello there

How are you?
What's going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)`,
			expectedOutput: `<h1>Header one</h1>
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
		err := MarkdownToHtmlNoWrap(strings.NewReader(test.input), &b)

		if (test.expectedErr && err == nil) || (!test.expectedErr && err != nil) {
			t.Errorf("expected err? %t. got=%v", test.expectedErr, err)
		}

		if b.String() != test.expectedOutput {
			t.Errorf("%d:%s - incorrect output. wanted=%s, got=%s",
				testNum, test.name, test.expectedOutput, b.String())
		}
	}
}

func TestWrapping(t *testing.T) {

	input := "hello world"
	expectedOutput := `<!DOCTYPE html>
<html lang="en">
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<body>
<p>hello world</p>
</body>
</html>`
	var b bytes.Buffer
	err := MarkdownToHtmlAndWrap(strings.NewReader(input), &b)

	if err != nil {
		t.Errorf("failed to wrap and convert input")
	}

	if b.String() != expectedOutput {
		t.Errorf("incorrect output. wanted=%s, got=%s",
			expectedOutput, b.String())
	}
}
