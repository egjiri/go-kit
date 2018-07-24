package print

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Heading prints a green heading to the screen followed by ...
func Heading(str string) {
	fmt.Println("\n" + color.GreenString(str+"..."))
}

// Command prints a colored command to the screen
func Command(str string) {
	prefix := color.BlueString("==> ")
	boldStr := color.New(color.Bold).Sprint(str)
	fmt.Println(prefix + boldStr)
}

// Finished prints a green finished message to the screen
func Finished() {
	color.Green("\nFinished!")
}

// Fatal acceps a string or an error type and prints a red message
// to the screen and exits with an error code 1
func Fatal(str interface{}) {
	switch msg := str.(type) {
	case error:
		color.Red(msg.Error())
	case string:
		color.Red(msg)
	default:
		color.Red("Invalid argument to Fatal")
	}
	os.Exit(1)
}

// CheckError checks for the existance of an error and if present
// prints it in red and exits with an error code 1
func CheckError(err error) {
	if err != nil {
		Fatal(err)
	}
}
