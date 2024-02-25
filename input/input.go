package input

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"slices"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readUserInput() (userInput string, err error) {
	userInput, err = reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return
}

func HandlePlayerInput(allowedInputs []string) string {
	fmt.Print("Enter your command code: ")
	input, err := readUserInput()

	if err != nil {
		fmt.Println("An error occurred while reading your input. Please try again.")
		return HandlePlayerInput(allowedInputs)
	}
	if !slices.Contains(allowedInputs, input) {
		fmt.Println("Invalid input, please try again.")
		return HandlePlayerInput(allowedInputs)
	}

	return input
}
