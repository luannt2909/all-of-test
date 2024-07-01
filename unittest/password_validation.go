package unittest

import "regexp"

func ValidatePassword(password string) bool {
	// Check for no white space
	hasWhitespace, _ := regexp.MatchString(`\s`, password)

	// Check for at least one digit
	hasDigit, _ := regexp.MatchString(`\d`, password)

	// Check for at least one lowercase letter
	hasLower, _ := regexp.MatchString(`[a-z]`, password)

	// Check for at least one uppercase letter
	hasUpper, _ := regexp.MatchString(`[A-Z]`, password)

	// Check for at least one special character
	hasSpecial, _ := regexp.MatchString(`[\W_]`, password)

	// Check for overall length
	hasMinLength := len(password) >= 10

	// Combine all checks
	return !hasWhitespace && hasDigit && hasLower && hasUpper && hasSpecial && hasMinLength
}
