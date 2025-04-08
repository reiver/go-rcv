package rcv

import (
	"github.com/reiver/go-utf8"
	"github.com/reiver/go-whitespace"

	"fmt"
	"io"
	"strconv"
	"strings"
)

// Readln reads in a line from ‘reader’ and puts it at ‘dst’.
//
// rcv.Freadln() does what you probably thought fmt.Fscanln() should do (but doesn't).
func Freadln(reader io.Reader, dst interface{}) (n int, err error) {
	var buffer strings.Builder

	for {
		r, size, err := utf8.ReadRune(reader)
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
	case scanner:
		err = casted.Scan(line)

	case fmt.Scanner:
		_, err = fmt.Sscan(line, casted)

	case *string:
		*casted = line


	case *int64:
		converted, err := strconv.ParseInt(line, 0, 64)
		if nil != err {
			return n, err
		}
		*casted = converted
	case *int32:
		converted, err := strconv.ParseInt(line, 0, 32)
		if nil != err {
			return n, err
		}
		*casted = int32(converted)
	case *int16:
		converted, err := strconv.ParseInt(line, 0, 16)
		if nil != err {
			return n, err
		}
		*casted = int16(converted)
	case *int8:
		converted, err := strconv.ParseInt(line, 0, 8)
		if nil != err {
			return n, err
		}
		*casted = int8(converted)


	case *uint64:
		converted, err := strconv.ParseUint(line, 0, 64)
		if nil != err {
			return n, err
		}
		*casted = converted
	case *uint32:
		converted, err := strconv.ParseUint(line, 0, 32)
		if nil != err {
			return n, err
		}
		*casted = uint32(converted)
	case *uint16:
		converted, err := strconv.ParseUint(line, 0, 16)
		if nil != err {
			return n, err
		}
		*casted = uint16(converted)
	case *uint8:
		converted, err := strconv.ParseUint(line, 0, 8)
		if nil != err {
			return n, err
		}
		*casted = uint8(converted)


	default:
		err = fmt.Errorf("rcv: cannot put read line into type %T", dst)
	}

	return n, err
}
