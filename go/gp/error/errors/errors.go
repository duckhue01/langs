package errors

// Error defines a standard application error
type Errors struct {
	// machine readale error code
	Code string

	// human readable error code
	Message string

	// Logical operation and nestede error
	Op  string
	Err error
}
