package regex

import (
	"fmt"
	"regexp"
	"strings"
)

// Replace provides a way to change the content of a string based on regex matches
func Replace(content string, regexQuery string, replaceMap map[int]string) (string, error) {
	rx, err := regexp.Compile(regexQuery)
	if err != nil {
		return content, err
	}
	matches := rx.FindStringSubmatch(content)
	for i, value := range replaceMap {
		if len(matches) <= i {
			return content, fmt.Errorf("Out of bounds")
		}
		updatedSection := strings.Replace(matches[0], matches[i], value, 1)
		content = strings.Replace(content, matches[0], updatedSection, 1)
	}
	return content, nil
}
