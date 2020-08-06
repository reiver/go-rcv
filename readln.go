package rcv

import (
	"os"
)

// Readln reads in a line from os.Stdin and puts it at ‘dst’.
//
// rcv.Readln() does what you probably thought fmt.Scanln() should do (but doesn't).
//
// Here is an example:
//
//	var line string
//	
//	rcv.Readlne(&dst)
//	
//	fmt.Printf("LINE: %q\n", line)
func Readln(dst interface{}) (n int, err error) {
	return Freadln(os.Stdin, dst)
}
