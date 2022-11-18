package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!_?-#"
	numberSet      = "0123456789"
)

func main() {

	var passwordLength int

	for passwordLength == 0 {
		fmt.Println("Your value is not a valid number!")
		if passwordLength == 0 {
			passwordLength = checkUserInput()
		}
	}
	fmt.Println(generatePassword(passwordLength))
}

func generatePassword(passwordLength int) (password string) {

	// TODO: generate always another string not the same
	charSet := lowerCharSet + upperCharSet + specialCharSet + numberSet
	characterCount := len(charSet)

	rand.Seed(time.Now().UTC().UnixNano())

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
