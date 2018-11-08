package utils

import "fmt"

type NoSuchCheckError string
func NewNoSuchCheckError(name string) NoSuchCheckError {
	return NoSuchCheckError(fmt.Sprintf("No such check: %s", name))
}
func (e NoSuchCheckError) Error() string {
	return string(e)
}
