package main

import (
	"encoding/json"
	"github.com/ojrac/opensimplex-go"
	"github.com/sheenathejunglegirl/world-generation/random"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MapSearch gets a stubbed map
func MapSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	x, err := strconv.Atoi(vars["x"])
	if err != nil {
		handleInternalServerError(w, err)
		return
	}
	y, err := strconv.Atoi(vars["y"])
	if err != nil {
		handleInternalServerError(w, err)
		return
	}
	log.Printf("x=%s y=%s", x, y)

	json, err := json.Marshal(getStubbedMap(x, y))
	if err != nil {
		handleInternalServerError(w, err)
		return
	}

	writeJsonResponse(w, json)
}

func getStubbedMap(x int, y int) Map {
	stubbedCells := getStubbedCells()
	startingX, startingY := getStartPoints(stubbedCells)
	printMap(stubbedCells, startingX, startingY)
	starting := Starting{
		X: startingX,
		Y: startingY,
	}
	return Map{
		ID:            1,
		X:             x,
		Y:             y,
		Starting:      starting,
		Cells:         stubbedCells,
		GeneratedTime: time.Now(),
	}
}

func getStartPoints(cells [][]Cell) (int, int) {
	for true {
		randomX := random.Int(0, len(cells))
		randomY := random.Int(0, len(cells[0]))
		if cells[randomX][randomY].Water == "000000000" {
			return randomX, randomY
		}
	}
	return 0, 0
}

func getStubbedCells() [][]Cell {
	width := random.Int(worldConfig.Map.Min, worldConfig.Map.Max)
	height := random.Int(worldConfig.Map.Min, worldConfig.Map.Max)
	cells := make([][]Cell, height)
	treeSimplex := opensimplex.NewOpenSimplexWithSeed(time.Now().UTC().UnixNano())
	waterSimplex := opensimplex.NewOpenSimplexWithSeed(time.Now().UTC().UnixNano())
	scale := .02
	for i := range cells {
		cells[i] = make([]Cell, width)
		for j := range cells[i] {
			treeFreq := treeSimplex.Eval2(float64(i)*scale, float64(j)*scale)
			waterFreq := waterSimplex.Eval2(float64(i)*scale, float64(j)*scale)
			rock, _ := random.BinaryString(worldConfig.Map.RockCount, .10)
			shrub, _ := random.BinaryString(worldConfig.Map.ShrubCount, .50)

			cells[i][j] = Cell{
				ID:       1,
				Rock:     rock,
				Treasure: false,
				Enemy:    10,
				Shrub:    shrub,
			}

			cells[i][j].generateTree(treeFreq)
			cells[i][j].generateWater(waterFreq)
		}
	}

	cells = smoothWater(cells)

	return cells
}

func smoothWater(cells [][]Cell) [][]Cell {
	water := "000000001"
	noWater := "000000000"
	above, below, left, right := water, water, water, water
	for i := range cells {
		for j := range cells[i] {
			if water == cells[i][j].Water {
				if i-1 > 0 {
					above = cells[i-1][j].Water
				}
				if i+1 < len(cells)-1 {
					below = cells[i+1][j].Water
				}
				if j-1 > 0 {
					left = cells[i][j-1].Water
				}
				if j+1 < len(cells[i])-1 {
					right = cells[i][j+1].Water
				}
				if above == noWater && left == noWater {
					cells[i][j].Water = "100000000"
				} else if below == noWater && left == noWater {
					cells[i][j].Water = "001000000"
				} else if left == noWater {
					cells[i][j].Water = "010000000"
				} else if above == noWater && right == noWater {
					cells[i][j].Water = "000100000"
				} else if below == noWater && right == noWater {
					cells[i][j].Water = "000010000"
				} else if above == noWater {
					cells[i][j].Water = "000000100"
				} else if below == noWater {
					cells[i][j].Water = "000000010"
				}
			}
		}
	}
	return cells
}

func printMap(cells [][]Cell, startingX int, startingY int) {
	for i := range cells {
		row := ""
		for j := range cells[i] {
			tree := cells[i][j].Tree
			water := cells[i][j].Water

			switch {
			case startingX == i && startingY == j:
				row += "X"
			case water == "000000001":
				row += "="
			case water != "000000000":
				row += "~"
			case tree == "000":
				row += " "
			case tree == "100" || tree == "010" || tree == "001":
				row += "'"
			case tree == "110" || tree == "101" || tree == "011":
				row += "\""
			case tree == "111":
				row += "#"
			}
		}
		log.Println(row)
	}
}
