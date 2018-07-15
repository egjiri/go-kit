package screen

import (
	"fmt"
	"strings"

	"github.com/egjiri/go-kit/ui/cursor"
)

// Printer keeps track of the components to be displayed on the screen
type Printer struct {
	components []*component
}

// Write allows the Printer type to implement the io.Writer interface
func (pr *Printer) Write(p []byte) (n int, err error) {
	s := strings.TrimSuffix(string(p), "\n") // Strip the added Line Feed in the end
	Println(s)
	return len(s), nil
}

func (pr *Printer) position() string {
	var p string
	for _, c := range pr.components {
		p += cursor.MoveVertical(c.height())
	}
	return p + cursor.MoveHorizontal(-1000)
}

func (pr *Printer) formattedContent() string {
	var str string
	for i, c := range pr.components {
		if i != 0 {
			str += "\n"
		}
		str += pr.formattedString(c.content())
	}
	return str
}

func (pr *Printer) formattedString(str string) string {
	var newStr string
	parts := strings.Split(str, "\n")
	for i, part := range parts {
		newStr += part + cursor.ClearLineForward()
		if i != len(parts)-1 {
			newStr += "\n"
		}
	}
	return newStr
}

func (pr *Printer) ensureSpacer() {
	if len(pr.components) > 0 && pr.components[0].content() != "" {
		// Add a spacer component the first time Println is used
		spacer := component{}
		pr.components = append([]*component{&spacer}, pr.components...)
		fmt.Println()
	}
}
