package utils

import passwordchecker "github.com/adityarizkyramadhan/password-checker"

func CheckPassword(pw string, passwordChecker *passwordchecker.PasswordChecker) bool {
	return passwordChecker.Check(pw)
}
