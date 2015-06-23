package main

import (
	"math/rand"
	"time"
)

var letters = []rune("01")

func randomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func randomBinaryString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(2)]
	}
	return string(b)
}
