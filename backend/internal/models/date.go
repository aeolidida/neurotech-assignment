package models

import (
	"neurotech-assignment/backend/internal/errs"
	"time"
)

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return errs.ErrInvalidDate
	}
	*d = Date(parsedTime)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format(`"2006-01-02"`)), nil
}
