package main

import (
	"time"
)

// Map represents information describing a randomly generated map
type Map struct {
	ID            int       `json:"id"`
	X             int       `json:"x"`
	Y             int       `json:"y"`
	GeneratedTime time.Time `json:"generated_time"`
	Cells         [][]Cell  `json:"cells"`
}
