package regex

import (
	"fmt"
	"regexp"
	"strings"
)

// Replace provides a way to change the content of a string based on regex matches
func Replace(content *[]byte, regexQuery string, replaceMap map[int]string) error {
	rx, err := regexp.Compile(regexQuery)
	if err != nil {
		return err
	}
	strContent := string(*content)
	matches := rx.FindStringSubmatch(strContent)
	for i, value := range replaceMap {
		if len(matches) <= i {
			return fmt.Errorf("Out of bounds")
		}
		updatedSection := strings.Replace(matches[0], matches[i], value, 1)
		*content = []byte(strings.Replace(strContent, matches[0], updatedSection, 1))
	}
	return nil
}
