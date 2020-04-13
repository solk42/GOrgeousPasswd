package passwordgenerator

const DefaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DefaultDigits = "0123456789"
const DefaultSpecialChars = "!§$%&/()=?#+*~-_.:,"

const DefaultMinLength = 4

type PasswordConfig struct {
	Chars      string
	Digits     string
	Special    string
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
		Chars:      DefaultChars,
		Digits:     DefaultDigits,
		Special:    DefaultSpecialChars,
		Length:     passwordLength,
		NumDigits:  numDigits,
		NumSpecial: numSpecialChars,
	}
	return currentConfig
}
