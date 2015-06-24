package random

import (
	"math/rand"
	"time"
)

func BinaryString(n int, chanceOfAnyOnes float64) (string, int) {
	individualChance := chanceOfAnyOnes / float64(n)
	count := 0

	randstr := make([]byte, n) // Random string to return
	for i := 0; i < n; i++ {
		randNum := rand.Float64()
		if randNum < individualChance {
			randstr[i] = '1'
			count++
		} else {
			randstr[i] = '0'
		}
	}
	return string(randstr), count
}

func Int(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
