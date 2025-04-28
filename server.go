package main

import (
	"Kadane/core"
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

var world core.World
var running bool = true

// Prints a usage message to os.Stderr (standard error).
func usage() {
	fmt.Fprintf(os.Stderr, "usage: server [options]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

// Flag Variables (addr)
var (
	addr = flag.String("addr", "localhost:8080", "address to serve")
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	// Parse flags.
	flag.Usage = usage
	flag.Parse()

	// Parse and validate arguments (none).
	args := flag.Args()
	if len(args) != 0 {
		usage()
	}

	world = CreateGameWorld()
	go gameLoop()

	// Serve static files from the "frontend/build" directory
	fs := http.FileServer(http.Dir(filepath.Join(".", "client", "build")))
	http.Handle("/", fs)

	// Register handlers.
	// All requests not otherwise mapped with go to greet.
	// /version is mapped specifically to version.
	http.HandleFunc("/api/game", runGame)
	http.HandleFunc("/version", version)

	log.Printf("serving http://%s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func version(w http.ResponseWriter, r *http.Request) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		http.Error(w, "no build information available", 500)
		return
	}

	fmt.Fprintf(w, "<!DOCTYPE html>\n<pre>\n")
	fmt.Fprintf(w, "%s\n", html.EscapeString(info.String()))
}

func runGame(w http.ResponseWriter, r *http.Request) {
	formatted := fmt.Sprintf("Currently there are %d entities in the game world.", world.GetNumEntities())
	message := Message{Text: formatted}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// Game loop function
func gameLoop() {
	const fps = 60
	frameDuration := time.Second / fps

	var elapsedTime time.Duration
	for running {
		startTime := time.Now()

		// Game update logic
		world.DoTick(float64(elapsedTime.Milliseconds()))

		// Control frame rate
		elapsedTime = time.Since(startTime)
		sleepDuration := frameDuration - elapsedTime
		if sleepDuration > 0 {
			time.Sleep(sleepDuration)
		}
	}
}
