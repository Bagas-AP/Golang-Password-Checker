package config

import passwordchecker "github.com/adityarizkyramadhan/password-checker"

func MakePasswordCheckerConfig() *passwordchecker.PasswordChecker {
	return passwordchecker.NewPasswordChecker(passwordchecker.PasswordCheckerConfig{
		MinLength:      8,
		MaxLength:      20,
		AllowSymbol:    true,
		AllowSpace:     false,
		AllowNumber:    true,
		AllowLowerCase: true,
		AllowUpperCase: true,
		MustUnique:     true,
	})
}
