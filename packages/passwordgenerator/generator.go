package passwordgenerator

import (
	dlh "GOrgeousPasswd/packages/dirtylittlehelper"
	"math/rand"
	"time"
)

// PasswordGenerator struct holding general functions and basic password generation information
type PasswordGenerator struct {
	config        passwordConfig
	password      []rune
	passwordIndex []int
}

// NewPasswordGenerator - constructor
// create a new password generator for a defined password structure, e. g.:
//  NewPasswordGenerator(8, 2, 2)
// will create and configure a generator for a 8 character long password including 2 digits and 2 special chars
func NewPasswordGenerator(minLength int, numDigits int, numSpecialChars int) PasswordGenerator {
	var pwgen = PasswordGenerator{
		config: newPasswordConfig(minLength, numDigits, numSpecialChars),
	}
	return pwgen
}

// PasswordGenerator.GeneratePassword creates a new random password matching the defined password generator configuration, e. g.:
//  NewPasswordGenerator(8, 2, 2).GeneratePassword()
// could return a password like Le7§X(8P
func (pwgen *PasswordGenerator) GeneratePassword() string {
	rand.Seed(time.Now().UTC().UnixNano())
	pwgen.generateBaseString()
	pwgen.replaceDigitChars()
	pwgen.replaceSpecialChars()
	return string(pwgen.password)
}

func (pwgen *PasswordGenerator) generateBaseString() {
	generated := make([]rune, pwgen.config.Length)
	var generatedIndex []int
	for i := range generated {
		generated[i] = pwgen.config.Chars[rand.Intn(len(pwgen.config.Chars))]
		generatedIndex = append(generatedIndex, i)
	}
	pwgen.password = generated
	pwgen.passwordIndex = generatedIndex
}

func (pwgen *PasswordGenerator) replaceSpecialChars() {
	pwgen.replaceRandomChar(pwgen.config.Special, pwgen.config.NumSpecial)
}

func (pwgen *PasswordGenerator) replaceDigitChars() {
	pwgen.replaceRandomChar(pwgen.config.Digits, pwgen.config.NumDigits)
}

func (pwgen *PasswordGenerator) replaceRandomChar(charList []rune, count int) {
	// Loop for replacing multi characters (count)
	for i := 0; i < count; i++ {
		// random char to set in password
		var replaceValue = charList[rand.Intn(len(charList))]
		// random index from unchanged password chars
		var listIndex = rand.Intn(len(pwgen.passwordIndex))
		var replaceIndexValue = pwgen.passwordIndex[listIndex]
		// replace password character
		pwgen.password[replaceIndexValue] = replaceValue
		// remove replaced char from list of replaceable password charindex
		pwgen.passwordIndex = dlh.RemoveFromSlice(pwgen.passwordIndex, listIndex)
	}
}
