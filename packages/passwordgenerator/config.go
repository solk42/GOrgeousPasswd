package passwordgenerator

// Available string characters used for basic password generation
const DefaultCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Available characters for number replacements
const DefaultDigits = "0123456789"

// Available special characters for replacements
const DefaultSpecialChars = "!§$%&/()=?#+*~-_.:,"

// Definition of minimum length of a password generated
const DefaultMinLength = 4

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
