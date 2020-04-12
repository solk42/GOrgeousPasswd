package main

import (
	"fmt"
	"math/rand"
	"time"
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
	// rand seed setting on each call
	rand.Seed(time.Now().UTC().UnixNano())

	// generate random string of defined charlist
	var basicChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var bacisCharsLength = len(basicChars)
	generated := make([]rune, length)
	for i := range generated {
		generated[i] = basicChars[rand.Intn(bacisCharsLength)]
	}
	return string(generated)
}
