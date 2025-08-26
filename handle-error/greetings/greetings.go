package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name is given return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return a value that embeds the name in a greetings message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}