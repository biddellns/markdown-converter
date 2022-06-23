package lib

//
////func TestTags(t *testing.T) {
////	tests := []struct {
////		name           string
////		input          string
////		expectedOutput string
////	}{
////		{
////			name:           "H1",
////			input:          "# Header1",
////			expectedOutput: "<h1>Header1</h1>",
////		},
////		{
////			name:           "H2",
////			input:          "## Header2",
////			expectedOutput: "<h2>Header2</h2>",
////		},
////		{
////			name:           "H3",
////			input:          "### Header3",
////			expectedOutput: "<h3>Header3</h3>",
////		},
////		{
////			name:           "H4",
////			input:          "#### Header4",
////			expectedOutput: "<h4>Header4</h4>",
////		},
////		{
////			name:           "H5",
////			input:          "##### Header5",
////			expectedOutput: "<h5>Header5</h5>",
////		},
////		{
////			name:           "H6",
////			input:          "###### Header6",
////			expectedOutput: "<h6>Header6</h6>",
////		},
////		{
////			name:           "A",
////			input:          "[myBlog](https://aws.nicbiddell.com/blog)",
////			expectedOutput: `<a href="https://aws.nicbiddell.com/blog">myBlog</a>`,
////		},
////		{
////			name:           "P",
////			input:          "something something something",
////			expectedOutput: "<p>something something something</p>",
////		},
////	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			converter := RegExpLineConverter{}
//
//			output := string(converter.ConvertLine([]byte(test.input)))
//			if output != test.expectedOutput {
//				t.Errorf("incorrect output. wanted=%s, got=%s",
//					test.expectedOutput, output)
//			}
//		})
//	}
//}
