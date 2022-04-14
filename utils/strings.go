package utils

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

// StringInput interactively takes the user's textual input
func StringInput() (string, error) {
	var input string
	reader := bufio.NewReader(os.Stdin)

	for {
		var err error
		input, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}

		if input != "" {
			break
		}
	}

	return strings.TrimSpace(input), nil
}

// CleanFileName returns a cleaned-up version of the given name that:
// - is lower case
// - ends with `.desktop`
// - `-` is the only allowed special character
// - `-` is the word separator
func CleanFileName(fileName string) (string, error) {
	if fileName == "" {
		return "", errors.New("empty file name")
	}

	// replace special characters with a space
	re, err := regexp.Compile(`[^a-zA-Z0-9]`)
	if err != nil {
		return "", err
	}
	cleanName := re.ReplaceAllString(fileName, " ")

	// remove multiple spaces
	re, err = regexp.Compile(`[ ]{2,}`)
	if err != nil {
		return "", err
	}
	cleanName = re.ReplaceAllString(cleanName, " ")

	// final cleanup
	cleanName = strings.TrimSpace(cleanName)
	cleanName = strings.ReplaceAll(cleanName, " ", "-")

	return strings.ToLower(cleanName) + ".desktop", nil
}
