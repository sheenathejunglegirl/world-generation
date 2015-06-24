package random

import (
	"math/rand"
	"time"
)

func BinaryString(n int, chanceOfAnyOnes float64) string {
	individualChance := chanceOfAnyOnes / float64(n)

	randstr := make([]byte, n) // Random string to return
	for i := 0; i < n; i++ {
		randNum := rand.Float64()
		if randNum < individualChance {
			randstr[i] = '1'
		} else {
			randstr[i] = '0'
		}
	}
	return string(randstr)
}

func Int(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
