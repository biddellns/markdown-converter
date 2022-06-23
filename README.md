# Markdown Converter

Convert markdown files to HTML in your terminal!

Currently supports headers and anchor tags.

## Running the program

### Native Binary
`./target/bin/markdown-converter -in myMarkdownFile.md -out myMarkdownFile.html` 

HTML doc and body headers will surround your output by default.

If you only want the strict conversion of your markdown into an html fragment, 
specify `-no-wrap=true`. 

### Docker
To use docker at the moment is tricky.

We'll need to attach volumes to the container for the input and output files

Run
```shell
 docker run --rm \
   -v $(pwd)/test/sample_inputs/sample1.md:input/sample1.md \
   -v $(pwd)/output:/output \
  interview/markdown-converter -in '/input/sample1.md' 
  -out output/output.html
```

**Note: At time of writing hadn't fully sorted container volume issue.**

## Building the program
### Native Binary
Run `make build-binary`

The binary will be created under the `./target/bin/` directory.

### Docker container
Run `make container`
