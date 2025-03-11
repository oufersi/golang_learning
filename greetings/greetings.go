package greetings

import (
	"fmt"
	"math/rand"
)

func Hello(name string) string {
	message := fmt.Sprintf(randFormat(), name)
	return message
}

func randFormat() string {
	formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

	return formats[rand.Intn(len(formats))]
}