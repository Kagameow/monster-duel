package logger

import (
	"fmt"
	"os"
)

var FILE *os.File

func init() {
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	FILE = file
}

func WriteStringToLog(logItem string) {
	fmt.Fprint(FILE, logItem)
}
