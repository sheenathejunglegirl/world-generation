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
	width := 40  //random.Int(worldConfig.Map.Min, worldConfig.Map.Max)
	height := 20 //random.Int(worldConfig.Map.Min, worldConfig.Map.Max)
	cells := make([][]Cell, height)
	filters := make([]Filter, height*width*100)
	filterIndex := 0
	for i := range cells {
		cells[i] = make([]Cell, width)
		for j := range cells[i] {
			tree, treeCount := random.BinaryString(worldConfig.Map.TreeCount, .80)
			if treeCount == worldConfig.Map.TreeCount {
				for fi := -5; fi < 5; fi++ {
					for fj := -5; fj < 5; fj++ {
						if i+fi > 0 && i+fi < width && j+fj > 0 && j+fj < height {
							filters[filterIndex] = Filter{
								X:     i + fi,
								Y:     j + fj,
								Count: 1,
							}
							filterIndex++
						}
					}
				}
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
	}

	return cells
}
