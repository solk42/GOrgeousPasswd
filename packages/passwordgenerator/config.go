package passwordgenerator

const DefaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DefaultDigits = "0123456789"
const DefaultSpecialChars = "!§$%&/()=?#+*~-_.:,"

const DefaultMinLength = 4

type PasswordConfig struct {
	Chars      []rune
	Digits     []rune
	Special    []rune
	Length     int
	NumDigits  int
	NumSpecial int
}

func NewPasswordConfig(minLength int, numDigits int, numSpecialChars int) PasswordConfig {
	var passwordLength = minLength
	if passwordLength < DefaultMinLength {
		passwordLength = DefaultMinLength
	}
	var sumSpecial = numDigits + numSpecialChars
	if passwordLength < sumSpecial {
		passwordLength = sumSpecial
	}

	var currentConfig = PasswordConfig{
		Chars:      []rune(DefaultChars),
		Digits:     []rune(DefaultDigits),
		Special:    []rune(DefaultSpecialChars),
		Length:     passwordLength,
		NumDigits:  numDigits,
		NumSpecial: numSpecialChars,
	}
	return currentConfig
}
