package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGender_Valid(t *testing.T) {
	tests := []struct {
		name     string
		gender   Gender
		expected bool
	}{
		{
			name:     "valid male",
			gender:   Male,
			expected: true,
		},
		{
			name:     "valid female",
			gender:   Female,
			expected: true,
		},
		{
			name:     "invalid gender",
			gender:   Gender(2),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.gender.Valid()
			if valid != tt.expected {
				t.Errorf("Gender.Valid() got = %v, want %v", valid, tt.expected)
			}
		})
	}
}

func TestGender_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		gender   Gender
		expected []byte
		wantErr  bool
	}{
		{
			name:     "male",
			gender:   Male,
			expected: []byte("0"),
			wantErr:  false,
		},
		{
			name:     "female",
			gender:   Female,
			expected: []byte("1"),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.gender)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gender.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(data, tt.expected) {
				t.Errorf("Gender.MarshalJSON() got = %s, want %s", data, tt.expected)
			}
		})
	}
}

func TestGender_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected Gender
		wantErr  bool
	}{
		{
			name:     "valid male",
			data:     []byte("0"),
			expected: Male,
			wantErr:  false,
		},
		{
			name:     "valid female",
			data:     []byte("1"),
			expected: Female,
			wantErr:  false,
		},
		{
			name:     "invalid gender",
			data:     []byte("2"),
			expected: Gender(0),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g Gender
			err := json.Unmarshal(tt.data, &g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gender.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if g != tt.expected {
				t.Errorf("Gender.UnmarshalJSON() got = %v, want %v", g, tt.expected)
			}
		})
	}
}
