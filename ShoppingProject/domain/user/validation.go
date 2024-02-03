package user

import "regexp"

var usernameRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9]{7,29}$")
var passwordRegex = regexp.MustCompile("^[A-Za-z][A-Za-z0-9]{7,29}$")

func ValidateUsername(name string) bool {
	return usernameRegex.MatchString(name)
}
func ValidatePassword(password string) bool {
	return passwordRegex.MatchString(password)
}
