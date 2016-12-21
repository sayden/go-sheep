package go_sheep

import "fmt"

type Error struct {
	Package string
	File    string
	Method  string
	Reason  string
	Err     error
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s\nReason: %s\nPackage: %s\nFile: %s\nMethod: %s", e.Err.Error(), e.Reason, e.Package, e.File, e.Method)
}

func (er *Error) NewError(r, m string, e error) error {
	return &Error{
		Err:    e,
		Reason: r,
		Method: m,
	}
}
