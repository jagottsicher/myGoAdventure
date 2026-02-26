package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

// InitUserInput starts a goroutine that listens for terminal events and
// returns a channel of key names.  Resize events trigger an immediate
// redraw of the stage geometry.
func InitUserInput() chan string {
	inputChan := make(chan string, 1)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				// Non-blocking send: drop key if the consumer hasn't caught up yet.
				select {
				case inputChan <- ev.Name():
				default:
				}
			case *tcell.EventResize:
				screen.Sync()
				needsRedraw = true
			}
		}
	}()
	return inputChan
}

// ReadInput does a non-blocking read from the input channel.
// Returns an empty string when no key is pending.
func ReadInput(inputChan chan string) string {
	select {
	case key := <-inputChan:
		return key
	default:
		return ""
	}
}

// HandleUserInput processes a key name and updates game state accordingly.
func HandleUserInput(key string) {
	switch key {
	case "Rune[q]", "Ctrl+C":
		screen.Fini()
		fmt.Fprintln(os.Stderr, "Bye.")
		os.Exit(0)

	case "Rune[w]", "Up":
		movePlayer(0, -player.stepY)

	case "Rune[s]", "Down":
		movePlayer(0, +player.stepY)

	case "Rune[a]", "Left":
		movePlayer(-player.stepX, 0)

	case "Rune[d]", "Right":
		movePlayer(+player.stepX, 0)
	}
}
