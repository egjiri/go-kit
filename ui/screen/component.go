package screen

import "fmt"

// component is a type that keeps track of content to display
// at the bottom of the screen
type component struct {
	staticContent   string
	dynamicStringer fmt.Stringer
}

func newComponent(s interface{}) *component {
	switch str := s.(type) {
	case string:
		return &component{staticContent: str}
	case fmt.Stringer:
		return &component{dynamicStringer: str}
	}
	return nil
}

func (c *component) height() int {
	count := 1
	for _, ch := range c.content() {
		if ch == '\n' {
			count++
		}
	}
	return count
}

func (c *component) content() string {
	str := string(c.staticContent)
	if c.dynamicStringer != nil {
		str = c.dynamicStringer.String()
	}
	return str
}
