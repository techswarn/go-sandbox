package main

import (
	"math/rand"
	// "time"
)

func returnRandomText() string {
	rand.Seed(36)
	names := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
	}
	return names[rand.Intn(len(names))]
}