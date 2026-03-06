package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"
)

// InitUserInput starts a goroutine that listens for terminal events and returns a channel of key names.
func InitUserInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			case *tcell.EventResize:
				currentScreen = nil
				fillTheScreen()
				screen.Sync()
			}
		}
	}()

	return inputChan
}

// ReadInput performs a non-blocking read from the input channel, returning an empty string if no key is pending.
func ReadInput(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}

	return key
}

// HandleUserInput processes a key name and updates game state accordingly.
func HandleUserInput(key string) {

	if key == "Rune[q]" {
		screen.Fini()
		clearScreen()
		fmt.Println("Bye.")
		os.Exit(0)
	} else if key == "Rune[w]" || key == "Up" {
		_, h := screen.Size()
		player.relY -= float64(player.stepY) / float64(h)
		if player.relY < 0 {
			player.relY = float64(h-player.height) / float64(h)
		}
	} else if key == "Rune[s]" || key == "Down" {
		_, h := screen.Size()
		player.relY += float64(player.stepY) / float64(h)
		if player.relY >= 1.0 {
			player.relY = 0
		}
	} else if key == "Rune[a]" || key == "Left" {
		w, _ := screen.Size()
		newRelX := player.relX - float64(player.stepX)/float64(w)
		newScreenX := int(newRelX * float64(w))
		if newScreenX < 0 {
			player.relX = float64(w-player.width) / float64(w)
		} else {
			player.relX = float64(newScreenX) / float64(w)
		}
	} else if key == "Rune[d]" || key == "Right" {
		w, _ := screen.Size()
		newRelX := player.relX + float64(player.stepX)/float64(w)
		newScreenX := int(newRelX * float64(w))
		if newScreenX >= w {
			player.relX = 0
		} else {
			player.relX = float64(newScreenX) / float64(w)
		}
	}
}

// clear the screen depending on your OS
func clearScreen() {
	if runtime.GOOS != "windows" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
