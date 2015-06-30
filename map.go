package main

import (
	"time"
)

// Map represents information describing a randomly generated map
type Map struct {
	ID            int       `json:"id"`
	X             int       `json:"x"`
	Y             int       `json:"y"`
	Starting      Starting  `json:"starting"`
	GeneratedTime time.Time `json:"generated_time"`
	Cells         [][]Cell  `json:"cells"`
}

type Starting struct {
	X int `json:"x"`
	Y int `json:"y"`
}
