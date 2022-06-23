# Markdown Converter

## Table of contents

---
Convert markdown files to HTML in your terminal!

Currently supports headers and anchor tags.

## Quickstart

Run `make build-binary`

Run `./target/bin/markdown-converter -in ./test/sample_inputs/sample1.md -out tada.html`

## Running the program

---
### Native Binary
`./target/bin/markdown-converter -in <myMarkdownFile.md> -out myConvertedMarkdownFile.html` 

HTML doc and body headers will surround your output by default.

If you only want the strict conversion of your markdown into an html fragment, 
specify `-no-wrap=true`. 

### Docker
To use docker at the moment is tricky.

We'll need to attach volumes to the container for the input and output files

Run 
```shell
inFile=yourfile.md
outFile=yourfile.html

docker run --rm \
  -v "$(pwd)/path/to/$inFile:/input/$inFile" \
  -v $(pwd)/output:/output \
  interview/markdown-converter -in "/input/$inFile" \
  -out "/output/$outFile"
```

Pre-written example:
```shell
inFile=sample1.md
outFile=sample1out.html

docker run --rm \
  -v "$(pwd)/test/sample_inputs/$inFile:/input/$inFile" \
  -v $(pwd)/output:/output \
  interview/markdown-converter -in "/input/$inFile" \
  -out "/output/$outFile"
```

## Building the program

---
### Native Binary
Run `make build-binary`

The binary will be created under the `./target/bin/` directory.

### Docker container
Run `make container`

## Design Considerations

---

There were a couple design tradeoffs I had to make:

#### #1: Regex vs Compiler
This problem fit into the weird space where I definitely could have written a lexer & parser,
then used that resulting output to generate an AST and turn that into HTML.

I chose Regex due to time constraints and being able to parse initial input quickly. 

If we think about supporting the full set of Markdown (and potentially extending), we'd want to move towards a compiler.

#### #2: Streaming data vs Full file in memory
I could have written certain pieces to assume a full dataset (aka string in memory). 

One of the prompts talked about how we might deal with large datasets.

One way to reduce the memory footprint is to "stream" input using io.Reader and io.Writer.

This lets us open a file, use STDIN, or just provide a string wrapped with `strings.NewReader()`.

We get lots of flexibility, and in fact, we aren't shut off from just reading an entire file and memory and just passing it.

The beauty of this setup is:
- Uses fewer resources -> leads to cost savings
- Easily handles a variety of use-cases -> E.g., CLI, Server application, Lambda
- Interoperable with many libraries -> Doesn't constrain library consumers

#### #3: Manually writing out base HTML vs Using [Templates](https://pkg.go.dev/text/template)
When we generate html, we can "wrap" it in boilerplate HTML. 

We used string constants and wrote them with our writer. 

This ended up working much better with our "streaming" workflow and design philosophy. Had we used a template, we'd have had to load the entire output into memory.

#### #4: Parsing with bytes instead of strings
It's easier to change bytes in place and reduces any extra allocations - one of the biggest variables for Go's GC performance.

## What's next?
- Accepting stdin and stdout for further flexibility
- Demo with an AWS Lambda
- Using a compiler approach to see if we can get the same outputs