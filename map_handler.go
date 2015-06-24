package main

import (
	"encoding/json"
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
	cells := make([][]Cell, width)
	for i := range cells {
		log.Println(i)
		cells[i] = make([]Cell, height)
		for j := range cells[i] {
			log.Println(j)
			tree := random.BinaryString(worldConfig.Map.TreeCount, .80)
			rock := random.BinaryString(worldConfig.Map.RockCount, .10)
			shrub := random.BinaryString(worldConfig.Map.ShrubCount, .50)

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
	return cells
}
