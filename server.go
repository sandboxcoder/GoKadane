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
	// All requests will need the static prefix to avoid filepath collision.
	// See: https://stackoverflow.com/questions/36357791/gin-router-path-segment-conflicts-with-existing-wildcard
	router.Static("/static", "./client/build")
	// serve index.html for unknown routes (SPA fallback)
	router.NoRoute(func(c *gin.Context) {
		c.File("./client/build/index.html")
	})

	router.Run("localhost:8080")
}

func gameApiHandler(c *gin.Context) {
	entCollection := world.GetCollection()
	list := make([]core.EntityInfo, 0, 5)
	entityList := *entCollection.GetEntities()
	for _, value := range entityList {
		entityInfo := core.CreateInfo(&value)
		list = append(list, entityInfo)
	}
	c.IndentedJSON(http.StatusOK, list)
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
