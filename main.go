package main

import (
	"fmt"
	"math/rand"
)

func main() {
	pwlength := 10
	pwchars := 2
	pwdigits := 2
	fmt.Printf("Password for length %v, digits %v, special chars %v: %v\n\n", pwlength, pwdigits, pwchars, generatePassword(pwlength, pwchars, pwdigits))
}

func generatePassword(minLength int, numSpecialChars int, numDigits int) string {
	// generate random string of length minLength
	var generated = getRandomString(minLength)
	return generated
}

func getRandomString(length int) string {
	var basicChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var bacisCharsLength = len(basicChars)
	generated := make([]rune, length)
	for i := range generated {
		generated[i] = basicChars[rand.Intn(bacisCharsLength)]
	}
	return string(generated)
}
