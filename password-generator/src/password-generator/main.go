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
