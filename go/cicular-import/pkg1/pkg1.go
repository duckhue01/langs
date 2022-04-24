// package pkg1

// import (
// 	"fmt"

// 	"github.com/duckhue01/learn/go/cicular-import/pkg2"
// )

// type PP1 struct{}

// func New() *PP1 {
// 	return &PP1{}
// }

// func (p *PP1) HelloFromP1() {
// 	fmt.Println("Hello from package p1")
// }

// func (p *PP1) HelloFromP2Side() {
// 	pp2 := pkg2.New()
// 	pp2.HelloFromP2()
// }
