package decomposers

import (
	"fmt"
	"reflect"
	"strings"
)

func DecomposeWithTag(instance interface{}, tagKey string) (map[string]interface{}, error) {
	fieldsValues := make(map[string]interface{})

	val := reflect.ValueOf(instance).Elem() // Get the value of the UpdateCpuInput instance
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		tag := getTagName(fieldType, tagKey)

		// Only process non-nil fields
		if !field.IsNil() {
			fieldValue := field.Elem().Interface() // Get the actual value inside the pointer

			// Handle nested structs
			if field.Kind() == reflect.Ptr && field.Elem().Kind() == reflect.Struct {
				// Decompose nested struct fields if they are non-zero
				nestedFields := decomposeStructWithTag(fieldValue, tag, tagKey)
				for k, v := range nestedFields {
					fieldsValues[k] = v
				}
			} else {
				// Add the simple field to fieldsValues
				fieldsValues[tag] = fieldValue
			}
		}
	}
	return fieldsValues, nil
}

func getTagName(field reflect.StructField, tagKey string) string {
	tag := field.Tag.Get(tagKey)
	if idx := len(tag); idx > 0 {
		// Split the tag by comma and take only the first part
		tagName := strings.Split(tag, ",")[0]
		return tagName
	}
	return ""
}

// Helper function to decompose nested structs
func decomposeStructWithTag(data interface{}, prefix string, tagKey string) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		tag := getTagName(fieldType, tagKey)
		fieldName := tag
		if prefix != "" {
			fieldName = fmt.Sprintf("%s.%s", prefix, tag)
		}

		// Check for zero value; only add non-zero values
		if !isZeroValue(field) {
			result[fieldName] = field.Interface()
		}
	}
	return result
}

// Check if a reflect.Value is the zero value for its type
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Struct:
		// For structs, check if each field is the zero value
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		// Compare basic comparable types directly
		return v.Interface() == reflect.Zero(v.Type()).Interface()
	}
}
