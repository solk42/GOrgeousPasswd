package main

import (
	"fmt"
	"math/rand"
	"time"

	dlh "GOrgeousPasswd/packages/dirtylittlehelper"
)

func main() {
	pwlength := 8
	pwchars := 2
	pwdigits := 2
	generatedPassword := generatePassword(pwlength, pwchars, pwdigits)
	fmt.Printf("Password for length %v, digits %v, special chars %v: %v\n\n", pwlength, pwdigits, pwchars, generatedPassword)
}

// main password generator function
func generatePassword(minLength int, numSpecialChars int, numDigits int) string {
	// generate random string of length minLength
	// password Index is used to prevend override of already replaced chars
	var generatedPassword, passwordIndex = getRandomString(minLength)

	// replace digits
	const replaceDigitChars = "0123456789"
	generatedPassword, passwordIndex = replaceRandomString(generatedPassword, replaceDigitChars, numDigits, passwordIndex)

	// replace special chars
	const replaceSpecialChars = "!§$%&/()=?#+*~-_.:,"
	generatedPassword, passwordIndex = replaceRandomString(generatedPassword, replaceSpecialChars, numSpecialChars, passwordIndex)

	return string(generatedPassword)
}

// random replacement of given options
func replaceRandomString(basicValue []rune, characters string, numReplace int, replaceIndex []int) ([]rune, []int) {
	// replace value setting
	var replaceChars = []rune(characters)
	var replaceCharsLength = len(replaceChars)

	// general value handling
	var updateIndex = replaceIndex
	var updatePassword = basicValue

	// Loop 0 to < numReplace
	for i := 0; i < numReplace; i++ {
		// random char from replacer
		var replaceValue = replaceChars[rand.Intn(replaceCharsLength)]
		// ramdom index from updateIndex
		var listIndex = rand.Intn(len(updateIndex))
		var replaceIndex = updateIndex[listIndex]
		// replace updatePassword on random index
		updatePassword[replaceIndex] = replaceValue
		// remove value from update index
		updateIndex = dlh.RemoveFromSlice(updateIndex, listIndex)
	}

	return updatePassword, updateIndex
}

// generation of basic string value (random)
func getRandomString(length int) ([]rune, []int) {
	// rand seed setting on each call
	rand.Seed(time.Now().UTC().UnixNano())

	// generate random string of defined charlist
	var basicChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var bacisCharsLength = len(basicChars)
	generated := make([]rune, length)
	var generatedIndex []int
	for i := range generated {
		generated[i] = basicChars[rand.Intn(bacisCharsLength)]
		generatedIndex = append(generatedIndex, i)
	}
	return generated, generatedIndex
}
