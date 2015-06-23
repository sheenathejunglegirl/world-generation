package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var port = flag.Int("port", 8080, "web service port")
var worldConfigFile = flag.String("world-config-file", "world.json", "world configuration file")
var worldConfig WorldConfiguration

func main() {
	var err error
	worldConfig, err = getWorldConfiguration(*worldConfigFile)

	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	r := NewLoggedRouter()
	err = http.ListenAndServe(":"+strconv.Itoa(*port), r)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal("The server returned without error.")
}
