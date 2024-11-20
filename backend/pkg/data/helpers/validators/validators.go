package validators

import (
	validatorErrors "BetterPC_2.0/pkg/data/helpers/validators/errors"
	"github.com/sirupsen/logrus"
	"reflect"
)

func ValidateStruct(input interface{}) error {
	v := reflect.ValueOf(input)

	logrus.Info(reflect.TypeOf(input))

	// Ensure we are working with a pointer to a struct, or a pointer to a pointer, etc.
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return validatorErrors.ErrInvalidInput
		}
		v = v.Elem() // Dereference the pointer
	}

	// Iterate over the struct fields
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		// Skip if the field is nil
		if field.IsNil() {
			continue
		}

		// Check if the field is a pointer to a struct
		if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
			// Recursively check the struct for any non-zero values
			if hasNonZeroField(field.Elem()) {
				return nil // Found a non-zero field within a nested struct
			}
		} else {
			// For pointers to basic types, check if the value is non-zero
			if !isZeroValue(field.Elem()) {
				return nil
			}
		}
	}

	// If all fields are nil or zero, return an error
	return validatorErrors.ErrEmptyStruct
}

// Helper function to check if a struct has at least one non-zero field
func hasNonZeroField(v reflect.Value) bool {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !isZeroValue(field) {
			return true
		}
	}
	return false
}

// Helper function to determine if a value is the zero value for its type
func isZeroValue(v reflect.Value) bool {
	return v.Interface() == reflect.Zero(v.Type()).Interface()
}
