package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	encoding.Register()

	// init direction to avoid init error
	initDirections()

	// Player.Init()

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

	// displayPlayer(s)
	// Player.display(s)

	var currentRoom rooms

	currentRoom = roomStartRoomTopEntryRoom

	display(s, &currentRoom)
	// display(s, &TopEntryRoom)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			// Player.display(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				clearScreen()
				fmt.Println("Bye.")
				os.Exit(0)
			} else if ev.Rune() == 'w' || ev.Key() == tcell.KeyUp {
				//Player.Movement(s, 0, -2)
				if currentRoom.up != nil {
					currentRoom = *currentRoom.up
					display(s, &currentRoom)
					s.Sync()
				}
				//Player.display(s)
			} else if ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft {
				//Player.Movement(s, -4, 0)
				//Player.display(s)
				if currentRoom.left != nil {
					currentRoom = *currentRoom.left
					display(s, &currentRoom)
					s.Sync()
				}
			} else if ev.Rune() == 's' || ev.Key() == tcell.KeyDown {
				//Player.Movement(s, 0, 2)
				if currentRoom.down != nil {
					currentRoom = *currentRoom.down
					display(s, &currentRoom)
					s.Sync()
				}
				//Player.display(s)
			} else if ev.Rune() == 'd' || ev.Key() == tcell.KeyRight {
				//Player.Movement(s, 4, 0)
				//Player.display(s)
				if currentRoom.right != nil {
					currentRoom = *currentRoom.right
					display(s, &currentRoom)
					s.Sync()
				}
			}
		}
	}
}
