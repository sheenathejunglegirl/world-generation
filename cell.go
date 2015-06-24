package main

import (
	"github.com/sheenathejunglegirl/world-generation/random"
)

// Cell represents information describing a specific cell of a map
type Cell struct {
	ID       int    `json:"id"`
	Tree     string `json:"tree"`
	Shrub    string `json:"shrub"`
	Rock     string `json:"rock"`
	Treasure bool   `json:"treasure"`
	Enemy    int    `json:"enemy"`
}

func (c *Cell) applyFilter(filter Filter) {
	newChance := .80 + float64(filter.Count)
	c.Tree, _ = random.BinaryString(worldConfig.Map.TreeCount, newChance)
}
