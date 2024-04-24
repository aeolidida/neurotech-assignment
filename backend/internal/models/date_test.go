package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected Date
		wantErr  bool
	}{
		{
			name:     "valid date",
			data:     []byte(`"2024-04-21"`),
			expected: Date(time.Date(2024, time.April, 21, 0, 0, 0, 0, time.UTC)),
			wantErr:  false,
		},
		{
			name:     "invalid date format",
			data:     []byte(`"21-04-2024"`),
			expected: Date{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d Date
			err := json.Unmarshal(tt.data, &d)
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(d, tt.expected) {
				t.Errorf("Date.UnmarshalJSON() got = %v, want %v", d, tt.expected)
			}
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected []byte
		wantErr  bool
	}{
		{
			name:     "valid date",
			date:     Date(time.Date(2024, time.April, 21, 0, 0, 0, 0, time.UTC)),
			expected: []byte(`"2024-04-21"`),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("Date.MarshalJSON() got = %s, want %s", data, tt.expected)
			}
		})
	}
}
