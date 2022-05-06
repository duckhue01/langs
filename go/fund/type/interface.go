package main

type (
	Sender interface {
		send() string
	}
	ShipperA struct {
	}
	ShipperB struct {
		ShipperA
	}
)

func (sa *ShipperA) send() {

}

func a() {
	b := &ShipperB{}
	b.send()
}
