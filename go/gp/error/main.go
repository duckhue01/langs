package main

import (
	"errors"
	"fmt"
)

func main() {

	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2 %w", err1)
	fmt.Println(errors.As(err2, err1))
}
