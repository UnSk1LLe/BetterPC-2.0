package propertyFilters

import (
	"github.com/pkg/errors"
	"strconv"
)

func getInterval[T int | float64](from string, to string, zero T, max T) (T, T, error) {
	var minValue, maxValue T

	if from != "" {
		// Convert string to the generic type
		if _, isInt := any(zero).(int); isInt {
			val, parseErr := strconv.Atoi(from)
			if parseErr != nil {
				return minValue, maxValue, parseErr
			}
			minValue = T(val)
		} else {
			val, parseErr := strconv.ParseFloat(from, 64)
			if parseErr != nil {
				return minValue, maxValue, parseErr
			}
			minValue = T(val)
		}
	} else {
		minValue = zero
	}

	if to != "" {
		if _, isInt := any(zero).(int); isInt {
			val, parseErr := strconv.Atoi(to)
			if parseErr != nil {
				return minValue, maxValue, parseErr
			}
			maxValue = T(val)
		} else {
			val, parseErr := strconv.ParseFloat(to, 64)
			if parseErr != nil {
				return minValue, maxValue, parseErr
			}
			maxValue = T(val)
		}
	} else {
		maxValue = max
	}

	if minValue > maxValue {
		minValue, maxValue = maxValue, minValue
	}

	return minValue, maxValue, nil
}

func getIntervalPrice(propertyFilters map[string]interface{}) (int, int, error) {
	keyFromFilter := propertyFilters["price_min"]
	keyToFilter := propertyFilters["price_max"]

	keyFrom := 0
	keyTo := 1000000

	keyFrom, ok := keyFromFilter.(int)
	if !ok {
		return 0, 0, errors.New("invalid property filter type")
	}

	keyTo, ok = keyToFilter.(int)
	if !ok {
		return 0, 0, errors.New("invalid property filter type")
	}
	return keyFrom, keyTo, nil
}
