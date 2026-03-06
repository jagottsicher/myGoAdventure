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
var currentRoom *room

func init() {
	// encoding.Register()
	initDirections()
	// uncompressRoomData()
}

func main() {

	initScreen()
	initGamestate()
	inputChan := InitUserInput()

	// Calculate the duration of each frame
	frameDuration := time.Second / game.FPS

	drawStage()

	for !IsGameOver() {
		startTime := time.Now()

		// Handle user input
		HandleUserInput(ReadInput(inputChan))

		// update properties
		//(game, screen, objects, etc.)
		UpdateStates()

		// draw screen (background)
		drawStage()

		// draw objects
		drawAllVisibleobjects()

		// show screen
		screen.Show()

		// Calculate the remaining time until the next frame
		remainingTime := frameDuration - time.Since(startTime)

		// If there is remaining time, sleep for that duration
		if remainingTime > 0 {
			time.Sleep(remainingTime)
		}
	}
}

// UpdateStates updates all game object states for the current tick.
func UpdateStates() {

}

// IsGameOver returns true when the game-ending condition has been met.
func IsGameOver() bool {
	return false
}
