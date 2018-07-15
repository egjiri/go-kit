package screen

import (
	"fmt"
)

var pr = new(Printer)

// Writer aliases the default pr *Printer to be used by other packages importing
// this package and use it as an io.Writer
var Writer = pr

// Add creates a new component with input from either a string of a fmt.Stringer() interface
// It then gets displayed immediately and also gets stored in components for later processing
func Add(ss ...interface{}) {
	for _, s := range ss {
		if c := newComponent(s); c != nil {
			fmt.Println(s)
			pr.components = append(pr.components, c)
		}
	}
}

// Refresh updates the content of the components and re-displayes them on the screen
func Refresh() {
	fmt.Println(pr.position() + pr.formattedContent())
}

// Println prints the passed content to the screen before the content of the components
func Println(str string) {
	pr.ensureSpacer()
	fmt.Println(pr.position() + pr.formattedString(str))
	if content := pr.formattedContent(); content != "" {
		fmt.Println(content)
	}
}
