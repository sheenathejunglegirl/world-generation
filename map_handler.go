package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// MapSearch gets a stubbed map
func MapSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Printf("id=%s", id)

	json, err := json.Marshal(getStubbedMap())
	if err != nil {
		handleInternalServerError(w, err)
		return
	}

	writeJsonResponse(w, json)
}

func getStubbedMap() Map {
	return Map{
		ID:            1,
		X:             120,
		Y:             40,
		Cells:         getStubbedCells(),
		GeneratedTime: time.Now(),
	}
}

func getStubbedCells() [][]Cell {
	cells := make([][]Cell, 10) /* type declaration */
	for i := range cells {
		cells[i] = make([]Cell, 10) /* again the type? */
		for j := range cells[i] {
			cells[i][j] = Cell{
				ID:       1,
				Tree:     "101",
				Rock:     "01",
				Treasure: false,
				Enemy:    10,
				Shrub:    "00",
			}
		}
	}
	return cells
}
