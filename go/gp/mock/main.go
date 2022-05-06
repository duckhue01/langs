package main

type Stub interface {
	toString() string
}

type StubObject struct {
	data string
}

func (s *StubObject) Stub() string {
	return s.data
}
