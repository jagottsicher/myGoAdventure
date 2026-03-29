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

const (
	// keyInitialTimeout covers the OS key-repeat initial delay (~500–660 ms on Linux).
	// Keeps the key "pressed" until the first OS repeat event arrives.
	keyInitialTimeout = 750 * time.Millisecond
	// keyRepeatTimeout is used once OS repeats are flowing (~30 ms interval).
	// Short enough that releasing the key stops movement within ~80 ms.
	keyRepeatTimeout = 80 * time.Millisecond
)

var (
	pressedKeys  = map[string]bool{}
	keyTimers    = map[string]*time.Timer{}
	keyConsumed  = map[string]bool{} // true after first step; cleared on OS repeat
	keyMu        sync.Mutex
	playerMoveTick int
)

func InitUserInput() {
	go func() {
		for {
			switch ev := render.Screen.PollEvent().(type) {
			case *tcell.EventKey:
				handleKey(ev)
			case *tcell.EventResize:
				render.CurrentScreen = nil
				render.FillTheScreen()
				w, h := render.Screen.Size()
				game.ReinitOnResize(w, h)
				render.Screen.Sync()
			}
		}
	}()
}

func handleKey(ev *tcell.EventKey) {
	// Confirmation dialog takes priority.
	if game.ConfirmMode {
		if ev.Key() == tcell.KeyRune && (ev.Rune() == 'y' || ev.Rune() == 'Y') {
			action := game.ConfirmAction
			game.CancelConfirm()
			if action == "quit" {
				render.Screen.Fini()
				clearScreen()
				fmt.Println("Bye.")
				os.Exit(0)
			} else if action == "reset" {
				render.ResetGame()
			}
		} else {
			game.CancelConfirm()
		}
		return
	}

	// Any key closes the help screen.
	if game.HelpMode {
		game.ToggleHelp()
		render.FillTheScreen()
		return
	}

	// One-shot special keys.
	switch {
	case ev.Key() == tcell.KeyCtrlG:
		game.ToggleGodMode()
		return
	case ev.Key() == tcell.KeyRune && (ev.Rune() == 'h' || ev.Rune() == 'H'):
		game.ToggleHelp()
		render.FillTheScreen()
		return
	case ev.Key() == tcell.KeyRune && (ev.Rune() == 'v' || ev.Rune() == 'V'):
		game.HandleSelOverlayKey("variation")
		return
	case ev.Key() == tcell.KeyRune && (ev.Rune() == 'n' || ev.Rune() == 'N'):
		game.HandleSelOverlayKey("difficulty")
		return
	case ev.Key() == tcell.KeyRune && (ev.Rune() == 'r' || ev.Rune() == 'R'):
		game.StartConfirm("reset")
		return
	case ev.Key() == tcell.KeyRune && ev.Rune() == ' ':
		if game.CarriedObject != nil {
			game.DropCarried()
		} else {
			w, h := render.Screen.Size()
			game.TryPickup(w, h)
		}
		return
	case ev.Key() == tcell.KeyRune && (ev.Rune() == 'q' || ev.Rune() == 'Q'):
		game.StartConfirm("quit")
		return
	}

	if game.HelpMode {
		return
	}

	// Movement keys via held-key system.
	key := ev.Name()
	keyMu.Lock()
	pressedKeys[key] = true
	if t, ok := keyTimers[key]; ok {
		// Repeat event: OS repeat is flowing — allow next step and reset short timeout.
		keyConsumed[key] = false
		t.Reset(keyRepeatTimeout)
	} else {
		// First press: use long timeout to bridge the OS initial repeat delay.
		k := key
		keyTimers[k] = time.AfterFunc(keyInitialTimeout, func() {
			keyMu.Lock()
			delete(pressedKeys, k)
			delete(keyTimers, k)
			delete(keyConsumed, k)
			keyMu.Unlock()
		})
	}
	keyMu.Unlock()
}

func HandleUserInput() {
	game.PlayerMoved = false // reset every frame
	if game.Eaten {
		return // player trapped inside dragon — no movement
	}
	// Throttle to every 2nd frame (~30 moves/sec at 60 fps) to match
	// the original speed before the game-loop-driven input change.
	playerMoveTick++
	if playerMoveTick%2 != 0 {
		return
	}

	// Only pass keys that haven't been consumed yet; mark them consumed immediately.
	// A key is un-consumed again only when an OS repeat event arrives (held key).
	// Tap = 1 step. Hold = 1 step, brief pause, then one step per OS repeat.
	keyMu.Lock()
	keys := make(map[string]bool, len(pressedKeys))
	for k := range pressedKeys {
		if !keyConsumed[k] {
			keys[k] = true
			keyConsumed[k] = true
		}
	}
	keyMu.Unlock()

	w, h := render.Screen.Size()
	roomChanged := false

	playerScreenXY := func() (int, int) {
		return int(game.Player.RelX*float64(w)) - game.Player.Width/2,
			int(game.Player.RelY*float64(h)) - game.Player.Height/2
	}

	canMove := func(candX, candY int) bool {
		if game.GodMode {
			return true
		}
		return !render.WouldCollideWall(candX, candY, game.Player.Width, game.Player.Height)
	}

	if keys["Rune[w]"] || keys["Up"] {
		newRelY := game.Player.RelY - float64(game.Player.StepY)/float64(h)
		candX, _ := playerScreenXY()
		candY := int(newRelY*float64(h)) - game.Player.Height/2
		if canMove(candX, candY) {
			game.Player.RelY = newRelY
			game.PlayerMoved = true
			if candY < 0 {
				if game.CurrentRoom != nil && game.CurrentRoom.Up != nil {
					game.CurrentRoom = game.CurrentRoom.Up
					game.Player.RelY = 1.0 - float64(game.Player.Height/2)/float64(h)
					roomChanged = true
				} else {
					game.Player.RelY = float64(game.Player.Height/2) / float64(h)
				}
			}
		}
	}

	if keys["Rune[s]"] || keys["Down"] {
		newRelY := game.Player.RelY + float64(game.Player.StepY)/float64(h)
		candX, _ := playerScreenXY()
		candY := int(newRelY*float64(h)) - game.Player.Height/2
		if canMove(candX, candY) {
			game.Player.RelY = newRelY
			game.PlayerMoved = true
			if candY+game.Player.Height > h {
				if game.CurrentRoom != nil && game.CurrentRoom.Down != nil {
					game.CurrentRoom = game.CurrentRoom.Down
					game.Player.RelY = float64(game.Player.Height/2) / float64(h)
					roomChanged = true
				} else {
					game.Player.RelY = float64(h-game.Player.Height/2) / float64(h)
				}
			}
		}
	}

	if keys["Rune[a]"] || keys["Left"] {
		anchorX := int(game.Player.RelX * float64(w))
		_, candY := playerScreenXY()
		newAnchorX := anchorX - game.Player.StepX
		candX := newAnchorX - game.Player.Width/2
		if !canMove(candX, candY) {
			newAnchorX = anchorX - 1
			candX = newAnchorX - game.Player.Width/2
		}
		if canMove(candX, candY) {
			game.Player.RelX = float64(newAnchorX) / float64(w)
			game.PlayerMoved = true
			if candX < 0 {
				if game.CurrentRoom != nil && game.CurrentRoom.Left != nil {
					game.CurrentRoom = game.CurrentRoom.Left
					game.Player.RelX = float64(w-game.Player.Width+game.Player.Width/2) / float64(w)
					roomChanged = true
				} else {
					game.Player.RelX = float64(game.Player.Width/2) / float64(w)
				}
			}
		}
	}

	if keys["Rune[d]"] || keys["Right"] {
		anchorX := int(game.Player.RelX * float64(w))
		_, candY := playerScreenXY()
		newAnchorX := anchorX + game.Player.StepX
		candX := newAnchorX - game.Player.Width/2
		if !canMove(candX, candY) {
			newAnchorX = anchorX + 1
			candX = newAnchorX - game.Player.Width/2
		}
		if canMove(candX, candY) {
			game.Player.RelX = float64(newAnchorX) / float64(w)
			game.PlayerMoved = true
			if candX+game.Player.Width > w {
				if game.CurrentRoom != nil && game.CurrentRoom.Right != nil {
					game.CurrentRoom = game.CurrentRoom.Right
					game.Player.RelX = float64(game.Player.Width/2) / float64(w)
					roomChanged = true
				} else {
					game.Player.RelX = float64(w-game.Player.Width+game.Player.Width/2) / float64(w)
				}
			}
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
		_ = cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}
