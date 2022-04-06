package pkg

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Name(name string) (string, error) {
	//if no name is given, returns an error
	if name == "" {
		return "", errors.New("no name is given")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//Names returns a map that associates each of the named people with message
func Names(names []string) (map[string]string, error) {
	//a map to associate names with messages
	messages := make(map[string]string)
	//loop through the recieved name from slice
	// use name func to get message for each name from slice
	for _, name := range names {
		message, err := Name(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Whats up %v boss",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
