package product

import "strings"

type Error struct {
	message []string
}

func (e *Error) Error() string {
	return strings.Join(e.message, ", ")
}

func (e *Error) AddError(message string) {
	e.message = append(e.message, message)
}
