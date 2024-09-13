package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/owlsome-official/zlogtime"
	"github.com/rs/zerolog"
)

var (
	timeTracker           zlogtime.ZLogTime = zlogtime.New()
	timeTrackerWithConfig zlogtime.ZLogTime = zlogtime.New(zlogtime.Config{LogLevel: zerolog.DebugLevel.String()})
)

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Get("/get_data_from_db", Handler) // GET http://localhost:8000/get_data_from_db

	fmt.Println("Listening on http://localhost:8000")
	fmt.Println("Try to send a request :D")
	go app.Listen(":8000")

	customApp := fiber.New(fiber.Config{DisableStartupMessage: true})
	customApp.Get("/one", HandlerBySecond(1))            // GET http://localhost:8000/one
	customApp.Get("/onepointfive", HandlerBySecond(1.5)) // GET http://localhost:8000/onepointfive
	customApp.Get("/three", HandlerBySecond(3))          // GET http://localhost:8000/three
	customApp.Listen(":8001")
}

func Handler(c *fiber.Ctx) error {
	StopFor(0.02)
	return c.SendString("It's took 20ms!")
}

func StopFor(sec float64) {
	defer timeTracker.TimeTrack("Stop for "+fmt.Sprintf("%v", sec)+"s", time.Now())
	time.Sleep(time.Duration(sec*1000) * time.Millisecond)
}

// ===========================
// CUSTOM APP WITH DEBUG LEVEL
// ===========================
func HandlerBySecond(sec float64) fiber.Handler {
	return func(c *fiber.Ctx) error {
		DebugStopFor(sec)
		return c.SendString("Watch your app logs!")
	}
}

func DebugStopFor(sec float64) {
	defer timeTrackerWithConfig.TimeTrack("Stop for "+fmt.Sprintf("%v", sec)+"s", time.Now())
	time.Sleep(time.Duration(sec*1000) * time.Millisecond)
}

// ===========================
