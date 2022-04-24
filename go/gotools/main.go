package main

type Test struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

func main() {
	test := &Test{
		X: 0,
		Y: 0,
	}

	thisFunc(10, 23.2, "string")
}

func thisFunc(a int, b float32, c string) {
	return
}

type Reader interface {
	Read(int)
}

type MyReader struct {
}

fffunc (m *MyReader) Read(_ int) {
	panic("not implemented") // TODO: Implement
}
