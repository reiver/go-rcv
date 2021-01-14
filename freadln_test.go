package rcv_test

import (
	"github.com/reiver/go-rcv"

	"io"
	"strings"

	"testing"
)

func TestFreadln_string(t *testing.T) {
	tests := []struct{
		Input     string
		Expected  string
		ExpectedN int
	}{
		{
			//                                     line feed
			Input:          "apple banana cherry"+"\u000A"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u000A"),
		},
		{
			//                                     vertical tab
			Input:          "apple banana cherry"+"\u000B"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u000B"),
		},
		{
			//                                     form feed
			Input:          "apple banana cherry"+"\u000C"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u000C"),
		},
		{
			//                                     carriage return
			Input:          "apple banana cherry"+"\u000D"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u000D"),
		},
		{
			//                                     next line
			Input:          "apple banana cherry"+"\u0085"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u0085"),
		},
		{
			//                                     line separator
			Input:          "apple banana cherry"+"\u2028"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u2028"),
		},
		{
			//                                     paragraph separator
			Input:          "apple banana cherry"+"\u2029"+"one two three",
			Expected:       "apple banana cherry",
			ExpectedN : len("apple banana cherry"+"\u2029"),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Input)

		var dst string
		n, err := rcv.Freadln(reader, &dst)

		if nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}
		if expected, actual := test.ExpectedN, n; expected != actual {
			t.Errorf("For test #%d, the actual number of bytes read is not what was expected.", testNumber)
			t.Logf("INPUT:     %q", test.Input)
			t.Logf("READ LINE: %q", dst)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
		if expected, actual := test.Expected, dst; expected != actual {
			t.Errorf("For test #%d, the read line is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}

func TestFreadln_int64(t *testing.T) {
	tests := []struct{
		Input     string
		Expected  int64
		ExpectedN int
	}{
		{
			//                                     line feed
			Input:          "-123"+"\u000A"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u000A"),
		},
		{
			//                                     vertical tab
			Input:          "-123"+"\u000B"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u000B"),
		},
		{
			//                                     form feed
			Input:          "-123"+"\u000C"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u000C"),
		},
		{
			//                                     carriage return
			Input:          "-123"+"\u000D"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u000D"),
		},
		{
			//                                     next line
			Input:          "-123"+"\u0085"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u0085"),
		},
		{
			//                                     line separator
			Input:          "-123"+"\u2028"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u2028"),
		},
		{
			//                                     paragraph separator
			Input:          "-123"+"\u2029"+"one two three",
			Expected:        -123,
			ExpectedN : len("-123"+"\u2029"),
		},
	}

	for testNumber, test := range tests {

		var reader io.Reader = strings.NewReader(test.Input)

		var dst int64
		n, err := rcv.Freadln(reader, &dst)

		if nil != err {
			t.Errorf("For test #%d, received an error, but did not actually expect one.", testNumber)
			t.Logf("ERROR: (%T) %q", err, err)
			continue
		}
		if expected, actual := test.ExpectedN, n; expected != actual {
			t.Errorf("For test #%d, the actual number of bytes read is not what was expected.", testNumber)
			t.Logf("INPUT:     %q", test.Input)
			t.Logf("READ LINE: %q", dst)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
		if expected, actual := test.Expected, dst; expected != actual {
			t.Errorf("For test #%d, the read line is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
