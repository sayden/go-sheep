package go_sheep

import "fmt"

type Errors []error

type CheckError struct {
	Source *Node
	Target *Node
	Err    error
}

func (c *CheckError) Error() string {
	return fmt.Sprintf("Error checking node %s using node %s:\n%s", c.Target.Address, c.Source.Address, c.Err)
}

type NetworkError struct {
	error
}

func NewErrors(n int) *Errors {
	var errors Errors = make([]error, n)
	return &errors
}

func (e Errors) Error() (err string) {
	for _, error := range e {
		err = fmt.Sprintf("%s\n%s\n", err, error.Error())
	}

	return
}

func NewCheckError(s, t *Node, err error) error {
	return &CheckError{
		Err:    err,
		Source: s,
		Target: t,
	}
}
