package main

import (
	"fmt"
	"os"

	"development/myGoAdventure/graphics"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	encoding.Register()

	// init direction to avoid init error
	graphics.InitDirections(&graphics.TopEntryRoom)

	// graphics.Player.Init()

	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	// displaygraphics.Player(s)
	// graphics.Player.Display(s)

	var currentRoom graphics.Rooms

	currentRoom = graphics.YellowCastle

	graphics.Display(s, &currentRoom)
	// graphics.Display(s, &graphics.TopEntryRoom)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			// graphics.Player.Display(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				graphics.ClearScreen()
				fmt.Println("Bye.")
				os.Exit(0)
			} else if ev.Rune() == 'w' || ev.Key() == tcell.KeyUp {
				graphics.Player.Movement(s, 0, -2)
				currentRoom = *currentRoom.Up
				graphics.Display(s, &currentRoom)
				s.Sync()
				//graphics.Player.Display(s)
			} else if ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft {
				graphics.Player.Movement(s, -4, 0)
				//graphics.Player.Display(s)
			} else if ev.Rune() == 's' || ev.Key() == tcell.KeyDown {
				graphics.Player.Movement(s, 0, 2)
				currentRoom = *currentRoom.Down
				graphics.Display(s, &currentRoom)
				s.Sync()
				//graphics.Player.Display(s)
			} else if ev.Rune() == 'd' || ev.Key() == tcell.KeyRight {
				graphics.Player.Movement(s, 4, 0)
				//graphics.Player.Display(s)
			}
		}
	}
}
