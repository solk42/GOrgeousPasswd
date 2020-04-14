# passwordgenerator
--
    import "GOrgeousPasswd/packages/passwordgenerator"

Package passwordgenerator provides all required functions to generate random
passwords by configuration options

## Usage

```go
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
```

#### type PasswordGenerator

```go
type PasswordGenerator struct {
}
```

PasswordGenerator struct holding general functions and basic password generation
information

#### func  NewPasswordGenerator

```go
func NewPasswordGenerator(minLength int, numDigits int, numSpecialChars int) PasswordGenerator
```
NewPasswordGenerator - constructor create a new password generator for a defined
password structure, e. g.:

    NewPasswordGenerator(8, 2, 2)

will create and configure a generator for a 8 character long password including
2 digits and 2 special chars

#### func (*PasswordGenerator) GeneratePassword

```go
func (pwgen *PasswordGenerator) GeneratePassword() string
```
PasswordGenerator.GeneratePassword creates a new random password matching the
defined password generator configuration, e. g.:

    NewPasswordGenerator(8, 2, 2).GeneratePassword()

could return a password like Le7§X(8P
