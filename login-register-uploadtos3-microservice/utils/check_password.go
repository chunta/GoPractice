package utils

import "errors"

// ValidatePasswordFormat checks if the password meets the required format
func ValidatePasswordFormat(password string) error {
	if len(password) < 8 || len(password) > 20 {
		return errors.New("password must be between 8 and 20 characters")
	}
	// Add more validation rules as needed (e.g., special characters, numbers)
	return nil
}
