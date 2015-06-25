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
	return Map{
		ID:            1,
		X:             x,
		Y:             y,
		Cells:         getStubbedCells(),
		GeneratedTime: time.Now(),
	}
}

func getStubbedCells() [][]Cell {
	width := random.Int(worldConfig.Map.Min, worldConfig.Map.Max)
	height := random.Int(worldConfig.Map.Min, worldConfig.Map.Max)
	cells := make([][]Cell, height)
	simplex := opensimplex.NewOpenSimplexWithSeed(time.Now().UTC().UnixNano())
	scale := .02
	for i := range cells {
		cells[i] = make([]Cell, width)
		for j := range cells[i] {
			forest := simplex.Eval2(float64(i)*scale, float64(j)*scale)
			chanceOfTree := 1.0 + forest
			if forest > 0 {
				chanceOfTree = chanceOfTree * (3 * forest)
			} else {
				chanceOfTree = chanceOfTree / 2
			}
			tree, _ := random.BinaryString(worldConfig.Map.TreeCount, chanceOfTree)
			rock, _ := random.BinaryString(worldConfig.Map.RockCount, .10)
			shrub, _ := random.BinaryString(worldConfig.Map.ShrubCount, .50)

			cells[i][j] = Cell{
				ID:       1,
				Tree:     tree,
				Rock:     rock,
				Treasure: false,
				Enemy:    10,
				Shrub:    shrub,
			}
		}
	}

	printMap(cells)
	return cells
}

func printMap(cells [][]Cell) {
	for i := range cells {
		row := ""
		for j := range cells[i] {
			tree := cells[i][j].Tree

			switch {
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
