# Markdown Converter

Convert markdown files to HTML in your terminal!

Currently supports headers and anchor tags.

## Running the program

### Native Binary
`./target/bin/markdown-converter -in myMarkdownFile.md -out myMarkdownFile.html` 

HTML doc and body headers will surround your output by default.

If you only want the strict conversion of your markdown into an html fragment, 
specify `-no-wrap=true`.
``

## Building the program
### Native Binary
Run `make build-binary`

The binary will be created under the `./target/bin/` directory.
