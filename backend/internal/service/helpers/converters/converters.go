package converters

import (
	"github.com/pkg/errors"
	"time"
)

func ConvertDateFromString(dateString string) (time.Time, error) {
	convertedDate, err := time.Parse("2006-01-02", dateString) //Parse dob string to time.Time
	if err != nil {
		return convertedDate, errors.New("invalid date of birth format")
	}
	return convertedDate, nil
}
