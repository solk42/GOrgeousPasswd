package passwordgenerator

import (
	dlh "GOrgeousPasswd/packages/dirtylittlehelper"
	"math/rand"
	"time"
)

type PasswordGenerator struct {
	config        PasswordConfig
	password      []rune
	passwordIndex []int
}

func NewPasswordGenerator(minLength int, numDigits int, numSpecialChars int) PasswordGenerator {
	var pwgen = PasswordGenerator{
		config: NewPasswordConfig(minLength, numDigits, numSpecialChars),
	}
	return pwgen
}

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
	// Loop 0 to < numReplace
	for i := 0; i < count; i++ {
		// random char from replacer
		var replaceValue = charList[rand.Intn(len(charList))]
		// ramdom index from updateIndex
		var listIndex = rand.Intn(len(pwgen.passwordIndex))
		var replaceIndexValue = pwgen.passwordIndex[listIndex]
		// replace updatePassword on random index
		pwgen.password[replaceIndexValue] = replaceValue
		// remove value from update index
		pwgen.passwordIndex = dlh.RemoveFromSlice(pwgen.passwordIndex, listIndex)
	}
}
