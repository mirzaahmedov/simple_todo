package errors

import (
	"errors"
	"fmt"
)

func InternalServer() error {
	return errors.New("Internal server error")
}
func DuplicateValue(column string) error {
	return errors.New(fmt.Sprintf("duplicate %v", column))
}
func Validation(field string, err string) error {
	return errors.New(fmt.Sprintf("%v: %v", field, err))
}
