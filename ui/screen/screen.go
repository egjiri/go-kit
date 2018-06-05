package screen

import (
	"fmt"
	"strings"

	"github.com/egjiri/go-kit/ui/cursor"
)

var components []*component

// component is a type that keeps track of content to display at the bottom of the screen
type component struct {
	staticContent   string
	dynamicStringer fmt.Stringer
}

// Add creates a new component with input from either a string of a fmt.Stringer() interface
// It then gets displayed immediately and also gets stored in components for later processing
func Add(ss ...interface{}) {
	for _, s := range ss {
		var c component
		if str, ok := s.(string); ok {
			c = component{staticContent: str}
		} else if stringer, ok := s.(fmt.Stringer); ok {
			c = component{dynamicStringer: stringer}
		}
		fmt.Println(s)
		components = append(components, &c)
	}
}

// Refresh updates the content of the components and re-displayes them on the screen
func Refresh() {
	fmt.Println(position() + formattedContent())
}

// Println prints the passed content to the screen before the content of the components
func Println(str string) {
	ensureSpacer()
	fmt.Println(position() + formattedString(str))
	content := formattedContent()
	if content != "" {
		fmt.Println(content)
	}
}

func (c *component) content() string {
	str := string(c.staticContent)
	if c.dynamicStringer != nil {
		str = c.dynamicStringer.String()
	}
	return str
}

func (c *component) heigh() int {
	count := 1
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
	return p + cursor.MoveHorizontal(-1000)
}

func formattedContent() string {
	var str string
	for i, c := range components {
		if i != 0 {
			str += "\n"
		}
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

func ensureSpacer() {
	if len(components) > 0 && components[0].content() != "" {
		// Add a spacer component the first time Println is used
		spacer := component{}
		components = append([]*component{&spacer}, components...)
		fmt.Println()
	}
}
