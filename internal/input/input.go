package input

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"development/myGoAdventure/internal/game"
	"development/myGoAdventure/internal/render"
	"github.com/gdamore/tcell/v2"
)

func InitUserInput() chan string {
	inputChan := make(chan string)
	go func() {
		for {
			switch ev := render.Screen.PollEvent().(type) {
			case *tcell.EventKey:
				inputChan <- ev.Name()
			case *tcell.EventResize:
				render.CurrentScreen = nil
				render.FillTheScreen()
				render.Screen.Sync()
			}
		}
	}()
	return inputChan
}

func ReadInput(inputChan chan string) string {
	var key string
	select {
	case key = <-inputChan:
	default:
		key = ""
	}
	return key
}

func HandleUserInput(key string) {
	if key == "Rune[q]" {
		render.Screen.Fini()
		clearScreen()
		fmt.Println("Bye.")
		os.Exit(0)
	} else if key == "Rune[w]" || key == "Up" {
		_, h := render.Screen.Size()
		game.Player.RelY -= float64(game.Player.StepY) / float64(h)
		if game.Player.RelY < 0 {
			if game.CurrentRoom != nil && game.CurrentRoom.Up != nil {
				game.CurrentRoom = game.CurrentRoom.Up
				game.Player.RelY = 1.0 - float64(game.Player.Height)/float64(h)
				render.FillTheScreen()
			} else {
				game.Player.RelY = 0
			}
		}
	} else if key == "Rune[s]" || key == "Down" {
		_, h := render.Screen.Size()
		game.Player.RelY += float64(game.Player.StepY) / float64(h)
		if game.Player.RelY >= 1.0 {
			if game.CurrentRoom != nil && game.CurrentRoom.Down != nil {
				game.CurrentRoom = game.CurrentRoom.Down
				game.Player.RelY = 0
				render.FillTheScreen()
			} else {
				game.Player.RelY = float64(h-game.Player.Height) / float64(h)
			}
		}
	} else if key == "Rune[a]" || key == "Left" {
		w, _ := render.Screen.Size()
		newRelX := game.Player.RelX - float64(game.Player.StepX)/float64(w)
		newScreenX := int(newRelX * float64(w))
		if newScreenX < 0 {
			if game.CurrentRoom != nil && game.CurrentRoom.Left != nil {
				game.CurrentRoom = game.CurrentRoom.Left
				game.Player.RelX = float64(w-game.Player.Width) / float64(w)
				render.FillTheScreen()
			} else {
				game.Player.RelX = 0
			}
		} else {
			game.Player.RelX = float64(newScreenX) / float64(w)
		}
	} else if key == "Rune[d]" || key == "Right" {
		w, _ := render.Screen.Size()
		newRelX := game.Player.RelX + float64(game.Player.StepX)/float64(w)
		newScreenX := int(newRelX * float64(w))
		if newScreenX >= w {
			if game.CurrentRoom != nil && game.CurrentRoom.Right != nil {
				game.CurrentRoom = game.CurrentRoom.Right
				game.Player.RelX = 0
				render.FillTheScreen()
			} else {
				game.Player.RelX = float64(w-game.Player.Width) / float64(w)
			}
		} else {
			game.Player.RelX = float64(newScreenX) / float64(w)
		}
	}
}

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
