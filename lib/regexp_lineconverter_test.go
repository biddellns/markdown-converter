package lib

import (
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
			name:           "A",
			input:          `[myBlog](https://aws.nicbiddell.com/blog)`,
			expectedOutput: `<a href="https://aws.nicbiddell.com/blog">myBlog</a>`,
		},
	}

	for testNum, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			converter := RegExpLineConverter{}

			output := string(converter.ConvertLine([]byte(test.input)))
			if output != test.expectedOutput {
				t.Errorf("%d:%s - incorrect output. wanted=%s, got=%s",
					testNum, test.name, test.expectedOutput, output)
			}
		})
	}
}

func TestConversions(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
	}{
		//		{
		//			name: "First, simple sample",
		//			input: `# Sample Document
		//Hello!
		//
		//This is sample markdown for the [Mailchimp](https://www.mailchimp.com) homework assignment.`,
		//			expectedOutput: `
		//<h1>Sample Document</h1>
		//
		//<p>Hello</p>
		//
		//<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment</p>`,
		//		},
		//		{
		//			name: "Sample 2, nested, inline link",
		//			input: `
		//# Header one
		//
		//Hello there
		//
		//How are you?
		//What's going on?
		//
		//## Another Header
		//
		//This is a paragraph [with an inline link](http://google.com). Neat, eh?
		//
		//## This is a header [with a link](http://yahoo.com)`,
		//			expectedOutput: `
		//<h1>Header one</h1>
		//
		//<p>Hello there</p>
		//
		//<p>How are you?
		//What's going on?</p>
		//
		//<h2>Another Header</h2>
		//
		//<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>
		//
		//<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>`,
		//		},
	}

	for testNum, test := range tests {
		converter := RegExpLineConverter{}

		output := string(converter.ConvertLine([]byte(test.input)))
		t.Log(output)
		if output != test.expectedOutput {
			t.Errorf("%d:%s - incorrect output. wanted=%s, got=%s",
				testNum, test.name, test.expectedOutput, output)
		}
	}
}
