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
				// currentScreen = nil
				// initScreen()
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
		// fmt.Println(fmt.Sprintf("%b", convertToBinary("XXXXXXXXXXX X X X      X X X XXXXXXXXXXX")))
		fmt.Println(roomYellowCastle.uncompressedRoomData)
		os.Exit(0)
	} else if key == "Rune[w]" || key == "Up" {
		player.posY -= player.stepY
		_, screenHeight := screen.Size()
		if player.posY < 0 {
			player.posY = screenHeight - 1
		}
	} else if key == "Rune[s]" || key == "Down" {
		player.posY += player.stepY
		_, screenHeight := screen.Size()
		if player.posY > screenHeight-1 {
			player.posY = 0
		}
	} else if key == "Rune[a]" || key == "Left" {
		player.posX -= player.stepX
		screenWidth, _ := screen.Size()
		if player.posX < 0 {
			player.posX = screenWidth - 2
		}
	} else if key == "Rune[d]" || key == "Right" {
		player.posX += player.stepX
		screenWidth, _ := screen.Size()
		if player.posX > screenWidth-1 {
			player.posX = 0
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
