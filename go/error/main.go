package main

import (
	"fmt"
	"os"

	"github.com/duckhue01/lang/go/error/k_error"
)

func main() {

	_, err := os.ReadFile("./wrong")
	err = &k_error.Error{
		Code:    "1001",
		Message: "error to read data from",
		Op:      "FileOpen",
		Err:     err,
	}

	newErr := &k_error.Error{
		Code:    "100",
		Message: "asda",
		Op:      "ReadFile",
		Err:     err,
	}

	fmt.Println(k_error.ErrorCode(newErr))

}
