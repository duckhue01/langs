package pkg2

import (
	"fmt"

	// "github.com/duckhue01/learn/go/cicular-import/pkg1"
)

type PP2 struct{}

func New() *PP2 {
	return &PP2{}
}

func (p *PP2) HelloFromP2() {
	fmt.Println("Hello from package p2")
}

// func (p *PP2) HelloFromP1Side() {
// 	pp1 := pkg1.New()
// 	pp1.HelloFromP1()
// }
