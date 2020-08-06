package rcv

type scanner interface {
	Scan(src interface{}) error
}
