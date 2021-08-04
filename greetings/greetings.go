package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

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

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	}
	message := fmt.Sprintf(randomString(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomString() string {
	formats := []string{"Hi, %v. Welcome!", "Great to see you, %v!", "How are you, %v!"}

	return formats[rand.Intn(len(formats))]
}
