package util

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var errInvalidEmail error = errors.New("email inválido")
var errPasswordLength error = errors.New("senha precisa ter no mínimo 7 dígitos")
var errCnpjCpfLength error = errors.New("CNPJ ou CPF com tamanho inválido")
var errCnpjCpfDigits error = errors.New("CNPJ ou CPF deve conter apenas números")

func ValidateEmail(email string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(email) {
		return nil
	}
	return errInvalidEmail
}

func ValidatePassword(password string) error {
	if len(password) < 7 {
		return errPasswordLength
	}
	return nil
}

func ValidateCnpjOrCpf(cnpjOrCpf string) error {
	if len(cnpjOrCpf) < 11 {
		return errCnpjCpfLength
	}
	if !OnlyDigits(cnpjOrCpf) {
		return errCnpjCpfDigits
	}

	return nil
}

func HasWhiteSpace(s string) bool {
	return len(strings.Fields(s)) > 1
}

func HasOnlyWhiteSpaces(s string) bool {
	return len(strings.Fields(s)) == 0
}

func OnlyDigits(cnpjCpf string) bool {
	isNotDigit := func(c rune) bool { return c < '0' || c > '9' }
	return strings.IndexFunc(cnpjCpf, isNotDigit) == -1
}
