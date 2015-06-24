package main

import (
	"encoding/json"
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
	"github.com/sheenathejunglegirl/world-generation/random"
	"log"
	"math/rand"
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
	scale := rand.Float64() / 20
	for i := range cells {
		cells[i] = make([]Cell, width)
		row := ""
		for j := range cells[i] {
			forest := simplexnoise.Noise2(float64(i)*scale, float64(j)*scale)
			chanceOfTree := 1.0 + forest
			if forest > 0 {
				chanceOfTree = chanceOfTree * (3 * forest)
			} else {
				chanceOfTree = chanceOfTree / 2
			}
			tree, treeCount := random.BinaryString(worldConfig.Map.TreeCount, chanceOfTree)

			switch {
			case treeCount == 0:
				row += " "
			case treeCount == 1:
				row += "'"
			case treeCount == 2:
				row += "\""
			case treeCount == 3:
				row += "\""
			}
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
		log.Println(row)
	}
	return cells
}
