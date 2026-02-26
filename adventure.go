package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

// screen is the global tcell screen used by all render and input functions.
var screen tcell.Screen

// newScreen creates and initialises a new tcell screen.
func newScreen() (tcell.Screen, error) {
	s, err := tcell.NewScreen()
	if err != nil {
		return nil, fmt.Errorf("tcell.NewScreen: %w", err)
	}
	if err := s.Init(); err != nil {
		return nil, fmt.Errorf("screen.Init: %w", err)
	}
	return s, nil
}

func init() {
	initRooms()
}

func main() {
	initScreen()
	defer screen.Fini()

	initPlayer()

	inputChan := InitUserInput()

	frameDuration := time.Second / game.fps

	// Initial draw
	screen.Clear()
	drawStage()
	drawAllVisibleObjects()
	screen.Show()

	for !IsGameOver() {
		frameStart := time.Now()

		HandleUserInput(ReadInput(inputChan))
		UpdateStates()

		if needsRedraw {
			screen.Clear()
			needsRedraw = false
		}

		drawStage()
		drawAllVisibleObjects()
		drawStatusBar()
		screen.Show()

		if remaining := frameDuration - time.Since(frameStart); remaining > 0 {
			time.Sleep(remaining)
		}
	}
}

// UpdateStates updates all game object states for the current tick.
// Dragons, bat, magnet etc. will be driven from here in future iterations.
func UpdateStates() {
}

// clearTerminal clears the terminal using an ANSI escape sequence.
// Only used outside of the tcell screen (e.g. on exit).
func clearTerminal() {
	fmt.Fprint(os.Stdout, "\033[2J\033[H")
}
