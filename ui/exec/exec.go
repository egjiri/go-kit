package exec

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Heading prints a green heading to the screen followed by ...
func Heading(str string) {
	fmt.Println(HeadingString(str))
}

// HeadingString returns a green heading followed by ...
func HeadingString(str string) string {
	return "\n" + color.GreenString(str+"...")
}

// Command prints a colored command to the screen
func Command(str string) {
	fmt.Println(CommandString(str))
}

// CheckError checks for the existance of an error and if present
// prints it in red and exits with an error code 1
func CheckError(err error) {
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

// CommandString returns a colored command message
func CommandString(str string) string {
	prefix := color.BlueString("==> ")
	boldStr := color.New(color.Bold).Sprint(str)
	return prefix + boldStr
}

// Finished prints a green finished message to the screen
func Finished() {
	fmt.Println(FinishedString())
}

// FinishedString returns a green finished message
func FinishedString() string {
	return color.GreenString("\nFinished!")
}
