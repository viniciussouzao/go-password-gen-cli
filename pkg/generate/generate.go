package generate

import "github.com/viniciussouzao/go-password-gen/password"

func Password(length int, numbers bool, uppercase bool, symbols bool) (string, error) {
	return password.Generate(length, numbers, uppercase, symbols)
}

func CustomPassword(length, cNumbers, cUppercase, cSymbols int) (string, error) {
	return password.GenerateCustom(length, cNumbers, cSymbols, cUppercase)
}
