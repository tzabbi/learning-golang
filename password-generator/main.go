package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!_?-#"
	numberSet      = "0123456789"
)

func main() {

	var passwordLength int
	argsWithProg := os.Args[1:2]

	fmt.Println("This is the only password generator you'll ever need!")
	if len(argsWithProg) == 0 {
		passwordLength = checkUserInput()
	} else if len(argsWithProg) == 1 {
		passwordLength, _ = strconv.Atoi(argsWithProg[0])
	} else {
		fmt.Println("Just enter one value!")
	}
	for passwordLength == 0 {
		fmt.Println("Your value is not a valid number!")
		passwordLength = checkUserInput()
	}
	fmt.Println(generatePassword(passwordLength))
}

func generatePassword(passwordLength int) (password string) {

	// TODO: generate always another string not the same
	charSet := lowerCharSet + upperCharSet + specialCharSet + numberSet
	characterCount := len(charSet)

	for i := 0; i < passwordLength; i++ {
		randomPosition := rand.Intn(characterCount)
		randomChar := string(charSet[randomPosition])
		password = password + randomChar
	}
	return
}

func checkUserInput() (userInputhLength int) {
	var userInput string
	fmt.Print("Please enter the length of your password: ")
	// TODO: check single input
	fmt.Scan(&userInput)
	// replace all whitespaces, because if a multiple imputs crashes the input
	userInput = strings.ReplaceAll(userInput, " ", "")
	if passwordLength, err := strconv.Atoi(userInput); err == nil {
		userInputhLength = passwordLength
	}
	return
}
