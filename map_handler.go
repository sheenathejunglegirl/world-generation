package main

import (
	"encoding/json"
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
	width := randomInt(worldConfig.Map.Min, worldConfig.Map.Max)
	height := randomInt(worldConfig.Map.Min, worldConfig.Map.Max)
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
		for j := range cells[i] {
			cells[i][j] = Cell{
				ID:       1,
				Tree:     randomBinaryString(worldConfig.Map.TreeCount),
				Rock:     randomBinaryString(worldConfig.Map.RockCount),
				Treasure: false,
				Enemy:    10,
				Shrub:    randomBinaryString(worldConfig.Map.ShrubCount),
			}
		}
	}
	return cells
}
