# Markdown Converter

Convert markdown files to HTML in your terminal!

Currently supports headers and anchor tags.

## Running the program
`./target/bin/markdown-converter -in myMarkdownFile.md -out myMarkdownFile.html` 

HTML doc and body headers will surround your output by default.

If you only want the strict conversion of your markdown into an html fragment, 
Can specify `--no-wrap`. 

## Building the program

