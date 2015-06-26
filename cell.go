package main

import (
	"github.com/sheenathejunglegirl/world-generation/random"
)

// Cell represents information describing a specific cell of a map
type Cell struct {
	ID       int    `json:"id"`
	Tree     string `json:"tree"`
	Water    string `json:"water"`
	Shrub    string `json:"shrub"`
	Rock     string `json:"rock"`
	Treasure bool   `json:"treasure"`
	Enemy    int    `json:"enemy"`
}

func (c *Cell) generateTree(frequency float64) {
	chanceOfTree := 1.0 + frequency
	if frequency > 0 {
		chanceOfTree = chanceOfTree * (3 * frequency)
	} else {
		chanceOfTree = chanceOfTree / 2
	}
	c.Tree, _ = random.BinaryString(worldConfig.Map.TreeCount, chanceOfTree)
}

func (c *Cell) generateWater(frequency float64) {
	if frequency > .5 {
		c.Water = "000000001"
	} else {
		c.Water = "000000000"
	}
}
