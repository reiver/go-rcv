# go-rcv

Package **rcv** provides tools for _receiving_ input; including reading a line from the CLI.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-rcv

[![GoDoc](https://godoc.org/github.com/reiver/go-rcv?status.svg)](https://godoc.org/github.com/reiver/go-rcv)

## Example

Here is an example usage:
```Go
var line string

rcv.Readln(&dst)

fmt.Printf("LINE: %q\n", line)
```
