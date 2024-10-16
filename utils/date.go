package utils

import "time"

func ParseDate(date string) (*time.Time, error) {
	t, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
