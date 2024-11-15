package typeValidators

import (
	"github.com/pkg/errors"
	"reflect"
)

func ValidateType(input interface{}, targetType reflect.Type) error {
	inputType := reflect.TypeOf(input)

	for inputType.Kind() == reflect.Ptr {
		inputType = inputType.Elem()
	}

	if inputType != targetType {
		return errors.New("types mismatch: " + inputType.String() + " vs " + targetType.String())
	}
	return nil
}
