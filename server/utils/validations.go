package utils

import "regexp"

func PasswordMatchRegex(pwd string) bool {
	if len(pwd) < 6 {
		return false
	}

	hasNumber, _ := regexp.MatchString("[0-9]", pwd)
	hasLower, _ := regexp.MatchString("[a-z]", pwd)
	hasUpper, _ := regexp.MatchString("[A-Z]", pwd)

	return hasNumber && hasLower && hasUpper
}
