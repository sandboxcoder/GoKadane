package main

import (
	"Kadane/core"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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

	router := gin.Default()
	router.GET("/api/game", gameApiHandler)
	router.Static("/static", "./client/build")
	// serve index.html for unknown routes (SPA fallback)
	router.NoRoute(func(c *gin.Context) {
		c.File("./client/build/index.html")
	})

	router.Run("localhost:8080")
}

func gameApiHandler(c *gin.Context) {
	formatted := fmt.Sprintf("Currently there are %d entities in the game world.", world.GetNumEntities())
	message := Message{Text: formatted}
	c.IndentedJSON(http.StatusOK, message)
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
