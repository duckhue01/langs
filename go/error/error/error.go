package error

import (
	"bytes"
	"fmt"
)

const (
	E_INTERNAL_ERROR_CODE    = "internal"
	E_INTERNAL_ERROR_MESSAGE = "An internal error has occurred"
)

type (
	Code    string
	Op      string
	Message string

	Error struct {
		Code    Code
		Message Message
		Op      Op
		Err     error
	}
)

func E(args ...interface{}) *Error {
	err := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case error:
			err.Err = arg
		case Code:
			err.Code = arg
		case Message:
			err.Message = arg
		case Op:
			err.Op = arg

		}
	}
	return err
}

// ErrorCode returns the code of the root error, if available. Otherwise returns EINTERNAL.
func ErrorCode(err error) Code {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return E_INTERNAL_ERROR_CODE
}

// ErrorMessage
func ErrorMessage(err error) Message {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return E_INTERNAL_ERROR_MESSAGE
}
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(string(e.Message))
	}
	return buf.String()
}
