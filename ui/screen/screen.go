package screen

import (
	"fmt"
	"strings"

	"github.com/egjiri/go-kit/ui/cursor"
)

var components []*Component

// Component is a type that keeps track of content to display at the bottom of the screen
type Component struct {
	staticContent   string
	dynamicStringer fmt.Stringer
}

// NewComponent creates a new component with dynamic content based on the
// resuts of the String() function that gets displayed at the bottom of the screen
func NewComponent(s fmt.Stringer) *Component {
	comp := Component{dynamicStringer: s}
	components = append(components, &comp)
	return &comp
}

// NewStaticComponent creates a new component with static content
// that gets displayed at the bottom of the screen
func NewStaticComponent(s string) *Component {
	comp := Component{staticContent: s}
	components = append(components, &comp)
	return &comp
}

// Show prints all the components to the screen in the order they were created
func Show() {
	fmt.Println(formattedContent())
}

// Refresh updates the content of the components and re-displayes them on the screen
func Refresh() {
	fmt.Println(position() + formattedContent())
}

// Println prints the passed content to the screen before the content of the components
func Println(str string) {
	str = position() + formattedString(str) + "\n" + formattedContent()
	fmt.Println(str)
}

func (c *Component) content() string {
	str := string(c.staticContent)
	if c.dynamicStringer != nil {
		str = c.dynamicStringer.String()
	}
	return "\n" + str
}

func (c *Component) heigh() int {
	count := 0
	for _, ch := range c.content() {
		if ch == '\n' {
			count++
		}
	}
	return count
}

func position() string {
	var p string
	for _, c := range components {
		p += cursor.MoveVertical(c.heigh())
	}
	p += cursor.MoveVertical(1)
	return p + cursor.MoveHorizontal(-1000)
}

func formattedContent() string {
	var str string
	for _, c := range components {
		str += formattedString(c.content())
	}
	return str
}

func formattedString(str string) string {
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
