package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	randomFormat()
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hello %v, welcome.",
		"Gooday mate, %v",
		"Top of the morning, %v",
	}

	return formats[rand.Intn(len(formats))]
}
