package unittest

import "testing"

func TestValidatePassword(t *testing.T) {
	testcases := []struct {
		Name   string
		Input  string
		Wanted bool
	}{
		{
			Name:   "test_password_meet_condition",
			Input:  "Password123_",
			Wanted: true,
		},
		{
			Name:   "test_password_failed_by_less_than_10_character",
			Input:  "P@ssword9",
			Wanted: false,
		},
		{
			Name:   "test_password_failed_by_missing_special_character",
			Input:  "Password123",
			Wanted: false,
		},
		{
			Name:   "test_password_failed_by_missing_uppercase",
			Input:  "password123@",
			Wanted: false,
		},
		{
			Name:   "test_password_failed_by_missing_lowercase",
			Input:  "PASSWORD123@",
			Wanted: false,
		},
		{
			Name:   "test_password_failed_by_missing_number",
			Input:  "Password@@",
			Wanted: false,
		},
		{
			Name:   "test_password_failed_by_with_white_spaces",
			Input:  "P@ssword 9",
			Wanted: false,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			matched := ValidatePassword(tc.Input)
			if matched != tc.Wanted {
				t.Fail()
			}
		})
	}
}
