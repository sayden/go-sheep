package go_sheep

import "fmt"

type Error struct {
	Package string
	File    string
	Method  string
	Reason  string
	Errs    []error
}

func (e *Error) Error() (res string) {
	res = fmt.Sprintf("Package: %s\nFile: %s\nMethod: %s\n", e.Package, e.File, e.Method)

	for _, err := range e.Errs {
		res = fmt.Sprintf("%sError: %s\nReason: %s\n", res, err.Error(), e.Reason)
	}
	return
}

func (er *Error) NewError(r, m string, e ...error) error {
	return &Error{
		Errs:   e,
		Reason: r,
		Method: m,
	}
}

type CheckError struct {
	Source Node
	Target Node
	Err error
}

func (c *CheckError) Error() string {
	return fmt.Sprintf("Error checking node %s using node %s:\n%s", c.Target.Address, c.Source.Address, c.Err)
}

func NewCheckError(s, t Node, err error) error {
	return &CheckError{
		Err:err,
		Source:s,
		Target:t,
	}
}