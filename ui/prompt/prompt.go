package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

var reader = bufio.NewReader(os.Stdin)

// Prompt is a type used to prompt the user for input and react to it
type Prompt struct {
	PromptMessage string
	YesCallback   func() error
	NoMessage     string
}

// RequestAndHandleUserInput takes a Prompt type and displays a UI prompting for
// user input and if the response if yes it run the YesCallback ad if no
// it prints the NoMessage
func (prompt *Prompt) RequestAndHandleUserInput() error {
	fmt.Print(color.CyanString(prompt.PromptMessage + " [y/N]: "))
	res, err := reader.ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "cannot read user input")
	}
	fmt.Println()

	switch strings.TrimSuffix(strings.ToLower(res), "\n") {
	case "y", "yes":
		return prompt.YesCallback()
	case "n", "no", "":
		color.Yellow(prompt.NoMessage)
		return nil
	default:
		return prompt.RequestAndHandleUserInput()
	}
}
