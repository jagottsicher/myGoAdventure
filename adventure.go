package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

var currentScreen []*cell

// global vars
var screen tcell.Screen
var player *object
var allObjects []*object

func main() {
	// encoding.Register()
	// initDirections()
	// uncompressRoomData()

	initScreen()
	initGamestate()
	inputChan := InitUserInput()

	// Calculate the duration of each frame
	frameDuration := time.Second / game.FPS

	DrawState()

	for !IsGameOver() {
		startTime := time.Now()

		HandleUserInput(ReadInput(inputChan))
		UpdateState()
		DrawState()
		screen.Show()

		// Calculate the remaining time until the next frame
		remainingTime := frameDuration - time.Since(startTime)

		// If there is remaining time, sleep for that duration
		if remainingTime > 0 {
			time.Sleep(remainingTime)
		}
	}
}

func UpdateState() {

}

func IsGameOver() bool {
	return false
}
