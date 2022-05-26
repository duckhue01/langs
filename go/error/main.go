package main

import (
	"errors"
	"fmt"

	error "github.com/duckhue01/lang/go/error/error"
)

func main() {
	user := UserService{}

	fmt.Println(myError.ErrorMessage(user.createUser()))

}

type (
	UserService struct {
	}
	Persistance struct {
	}
)

func (us *UserService) createUser() error {
	op := "UserService.createUser"
	persist := &Persistance{}

	err := persist.insertUser()
	err = persist.insertRole()

	newErr := error.E( error.Op(op), error.Message("message") err)
	return newErr
}

func (p *Persistance) insertUser() error {
	op := "Persistance.insertUser"
	err := errors.New("fail to insert user")
	return myError.E( myError.Op(op), err)

}

func (p *Persistance) insertRole() error {
	op := "Persistance.insertRole"
	err := errors.New("fail to insert role")
	return myError.E( myError.Op(op), err)

}
