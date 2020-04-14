// Package passwordgenerator provides all required functions to generate random passwords by configuration options
package passwordgenerator

const (
	// Available string characters used for basic password generation
	DefaultCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Available characters for number replacements
	DefaultDigits = "0123456789"
	// Available special characters for replacements
	DefaultSpecialChars = "!§$%&/()=?#+*~-_.:,"
	// Definition of minimum length of a password generated
	DefaultMinLength = 4
)

type passwordConfig struct {
	Chars      []rune
	Digits     []rune
	Special    []rune
	Length     int
	NumDigits  int
	NumSpecial int
}

func newPasswordConfig(minLength int, numDigits int, numSpecialChars int) passwordConfig {
	var passwordLength = minLength
	if passwordLength < DefaultMinLength {
		passwordLength = DefaultMinLength
	}
	var sumSpecial = numDigits + numSpecialChars
	if passwordLength < sumSpecial {
		passwordLength = sumSpecial
	}

	var currentConfig = passwordConfig{
		Chars:      []rune(DefaultCharacters),
		Digits:     []rune(DefaultDigits),
		Special:    []rune(DefaultSpecialChars),
		Length:     passwordLength,
		NumDigits:  numDigits,
		NumSpecial: numSpecialChars,
	}
	return currentConfig
}
