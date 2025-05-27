package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Greet(name string) (string, error) {
	// If no name is given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// retunr a greeting that embeds the name in a message.
	message := fmt.Sprintf(randFormat(), name)
	return message, nil
}

// Hellos returns a map that associates names with their greeting messages.
func Greets(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Greet(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randFormat() string {
	// a slice of message formats.

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// return a randomly selected message format by indexing into the slice.
	return formats[rand.Intn(len(formats))]
}
