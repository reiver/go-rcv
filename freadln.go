package rcv

import (
	"github.com/reiver/go-utf8s"
	"github.com/reiver/go-whitespace"

	"fmt"
	"io"
	"strings"
)

// Readln reads in a line from ‘reader’ and puts it at ‘dst’.
//
// rcv.Freadln() does what you probably thought fmt.Fscanln() should do (but doesn't).
func Freadln(reader io.Reader, dst interface{}) (n int, err error) {
	var buffer strings.Builder

	for {
		r, size, err := utf8s.ReadRune(reader)
		n += size
		if nil != err {
			return n, err
		}

		if whitespace.IsMandatoryBreak(r) {
			break
		}

		buffer.WriteRune(r)
	}

	line := buffer.String()

	switch casted := dst.(type) {
	case *string:
		*casted = line
	case scanner:
		err = casted.Scan(line)
	default:
		err = fmt.Errorf("rcv: cannot put read line into type %T", dst)
	}

	return n, err
}
