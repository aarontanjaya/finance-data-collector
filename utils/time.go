package util

import "time"

func ParseTimeInLocationOrDefault(format string, date string, loc string) (*time.Time, error) {
	var t *time.Time
	l, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return t, err
	}

	dt, err := time.ParseInLocation(format, date, l)
	if err != nil {
		if date == "" {
			return &time.Time{}, nil
		} else {
			return nil, err
		}
	}
	return &dt, err
}
