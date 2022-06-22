package lib

//
//import (
//	"bytes"
//	"strings"
//	"testing"
//)
//
//func TestBase(t *testing.T) {
//	tests := []struct {
//		name           string
//		input          string
//		expectedOutput string
//	}{
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
//	}
//
//	for testNum, test := range tests {
//		var b bytes.Buffer
//		MarkdownToHtml(strings.NewReader(test.expectedOutput), &b)
//
//		if b.String() != test.expectedOutput {
//			t.Errorf("%d:%s - incorrect output. wanted=%s, got=%s",
//				testNum, test.name, test.expectedOutput, b.String())
//		}
//	}
//}
