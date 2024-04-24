package models

import (
	"encoding/json"
	"neurotech-assignment/backend/internal/errs"
)

type Gender int

const (
	Male   Gender = 0
	Female Gender = 1
)

func (g Gender) Valid() bool {
	return (g == Male || g == Female)
}

func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(g))
}

func (g *Gender) UnmarshalJSON(data []byte) error {
	var value int
	if err := json.Unmarshal(data, &value); err != nil {
		return errs.ErrInvalidGender
	}

	switch value {
	case 0:
		*g = Male
	case 1:
		*g = Female
	default:
		return errs.ErrInvalidGender
	}

	return nil
}
