package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// init is called automaticaly at program startup
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Functions that start with lower case are like 'private'
// Other modules can't call them
func randomFormat() string {
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

// Functions that start with upper case are like 'public'
// Other modules can call them
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	// range returns the index and a copy of the item's value
	// We don't need the index so we use the blank identifier _
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}