package decomposers

import (
	"reflect"
	"strings"
)

func DecomposeWithTag[T any](input *T, tag string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := flattenStruct(reflect.ValueOf(input).Elem(), "", tag, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func flattenStruct(v reflect.Value, prefix, tag string, result map[string]interface{}) error {
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tagValue := getTagName(field, tag) // Use your tag parsing logic here
		if tagValue == "" || tagValue == "-" {
			continue
		}

		fullKey := tagValue
		if prefix != "" {
			fullKey = prefix + "." + tagValue
		}

		fieldValue := v.Field(i)
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				continue // Skip nil pointers
			}
			fieldValue = fieldValue.Elem()
		} else {
			if fieldValue.IsZero() {
				continue
			}
		}

		// Handle nested structs
		if fieldValue.Kind() == reflect.Struct {
			if err := flattenStruct(fieldValue, fullKey, tag, result); err != nil {
				return err
			}
			continue
		}

		result[fullKey] = fieldValue.Interface()
	}
	return nil
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
