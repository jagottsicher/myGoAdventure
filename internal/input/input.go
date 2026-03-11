package input

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"development/myGoAdventure/internal/game"
	"development/myGoAdventure/internal/render"
	"github.com/gdamore/tcell/v2"
)

const keyReleaseTimeout = 150 * time.Millisecond

var (
	pressedKeys = map[string]bool{}
	keyTimers   = map[string]*time.Timer{}
	keyMu       sync.Mutex
)

func InitUserInput() {
	go func() {
		for {
			switch ev := render.Screen.PollEvent().(type) {
			case *tcell.EventKey:
				key := ev.Name()
				keyMu.Lock()
				pressedKeys[key] = true
				if t, ok := keyTimers[key]; ok {
					t.Reset(keyReleaseTimeout)
				} else {
					k := key
					keyTimers[k] = time.AfterFunc(keyReleaseTimeout, func() {
						keyMu.Lock()
						delete(pressedKeys, k)
						delete(keyTimers, k)
						keyMu.Unlock()
					})
				}
				keyMu.Unlock()
				HandleUserInput()
			case *tcell.EventResize:
				render.CurrentScreen = nil
				render.FillTheScreen()
				render.Screen.Sync()
			}
		}
	}()
}

func HandleUserInput() {
	keyMu.Lock()
	keys := make(map[string]bool, len(pressedKeys))
	for k, v := range pressedKeys {
		keys[k] = v
	}
	keyMu.Unlock()

	if keys["Rune[q]"] {
		render.Screen.Fini()
		clearScreen()
		fmt.Println("Bye.")
		os.Exit(0)
	}

	w, h := render.Screen.Size()
	roomChanged := false

	if keys["Rune[w]"] || keys["Up"] {
		game.Player.RelY -= float64(game.Player.StepY) / float64(h)
		if game.Player.RelY < 0 {
			if game.CurrentRoom != nil && game.CurrentRoom.Up != nil {
				game.CurrentRoom = game.CurrentRoom.Up
				game.Player.RelY = 1.0 - float64(game.Player.Height)/float64(h)
				roomChanged = true
			} else {
				game.Player.RelY = 0
			}
		}
	}

	if keys["Rune[s]"] || keys["Down"] {
		game.Player.RelY += float64(game.Player.StepY) / float64(h)
		if game.Player.RelY >= 1.0 {
			if game.CurrentRoom != nil && game.CurrentRoom.Down != nil {
				game.CurrentRoom = game.CurrentRoom.Down
				game.Player.RelY = 0
				roomChanged = true
			} else {
				game.Player.RelY = float64(h-game.Player.Height) / float64(h)
			}
		}
	}

	if keys["Rune[a]"] || keys["Left"] {
		newRelX := game.Player.RelX - float64(game.Player.StepX)/float64(w)
		newScreenX := int(newRelX * float64(w))
		if newScreenX < 0 {
			if game.CurrentRoom != nil && game.CurrentRoom.Left != nil {
				game.CurrentRoom = game.CurrentRoom.Left
				game.Player.RelX = float64(w-game.Player.Width) / float64(w)
				roomChanged = true
			} else {
				game.Player.RelX = 0
			}
		} else {
			game.Player.RelX = float64(newScreenX) / float64(w)
		}
	}

	if keys["Rune[d]"] || keys["Right"] {
		newRelX := game.Player.RelX + float64(game.Player.StepX)/float64(w)
		newScreenX := int(newRelX * float64(w))
		if newScreenX >= w {
			if game.CurrentRoom != nil && game.CurrentRoom.Right != nil {
				game.CurrentRoom = game.CurrentRoom.Right
				game.Player.RelX = 0
				roomChanged = true
			} else {
				game.Player.RelX = float64(w-game.Player.Width) / float64(w)
			}
		} else {
			game.Player.RelX = float64(newScreenX) / float64(w)
		}
	}

	if roomChanged {
		render.FillTheScreen()
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
