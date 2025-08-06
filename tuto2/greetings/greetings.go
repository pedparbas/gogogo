package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	//If no name was given, return an error with a message

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())
	return message, nil
}

// randomFormat returns one oj a set of greetings messages. The returned message is
// selected at random.

func randomFormat() string {
	//A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v Well met!",
	}

	//Return a randomly selected message format by specifying a random index for the
	//slice formats.
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people with a greetings message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with the messages
	messages := make(map[string]string)
	//Loop through the received slice of names, calling the Hello function to get a
	//message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}
