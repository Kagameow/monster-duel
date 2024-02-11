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

func readUserInput() (userInput string) {
	userInput, _ = reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return
}

func HandlePlayerInput(allowedInputs []string) (input string) {
	fmt.Print("Enter your command code: ")
	input = readUserInput()
	if !slices.Contains(allowedInputs, input) {
		fmt.Println("Invalid input, please try again.")
		return HandlePlayerInput(allowedInputs)
	}
	return
}
