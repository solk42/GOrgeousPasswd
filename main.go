package main

import (
	pwgen "GOrgeousPasswd/packages/passwordgenerator"
	"fmt"
)

func main() {
	pwlength := 8
	pwchars := 2
	pwdigits := 2

	var pwGenerator = pwgen.NewPasswordGenerator(pwlength, pwdigits, pwchars)
	fmt.Printf("Password for length %v, digits %v, special chars %v:\n", pwlength, pwdigits, pwchars)
	fmt.Println(pwGenerator.GeneratePassword())
	fmt.Println(pwGenerator.GeneratePassword())
	fmt.Println(pwGenerator.GeneratePassword())
}
