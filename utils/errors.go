package utils

import "fmt"

type NoSuchCheckError string
func NewNoSuchCheckError(name string) NoSuchCheckError {
	return NoSuchCheckError(fmt.Sprintf("No such check: %s", name))
}
func (e NoSuchCheckError) Error() string {
	return string(e)
}

type CheckNotRunError string
func NewCheckNotRunError(name string) CheckNotRunError {
	return CheckNotRunError(fmt.Sprintf("Check %s have not been executed yet", name))
}
func (e CheckNotRunError) Error() string {
	return string(e)
}
