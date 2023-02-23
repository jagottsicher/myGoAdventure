package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

// room needed to be globally available
var currentRoom *rooms = nil

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	encoding.Register()

	// init the room colors
	// roomColorInit()

	// init direction to avoid init error
	initDirections()

	// player init
	player.init()

	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	// init the startroom
	currentRoom = &roomStartRoomTopEntryRoom

	// show the room for the first time
	display(s, currentRoom)

	// displayPlayer
	player.display(s, currentRoom)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			display(s, currentRoom)
			player.display(s, currentRoom)
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				clearScreen()
				fmt.Println("Bye.")
				os.Exit(0)
			} else if ev.Rune() == 'w' || ev.Key() == tcell.KeyUp {
				if checkPlayerWallCollision(s, 1) != true {
					player.movement(s, 0, -2)
				}
			} else if ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft {
				if checkPlayerWallCollision(s, 4) != true {
					player.movement(s, -2, 0)
				}
			} else if ev.Rune() == 's' || ev.Key() == tcell.KeyDown {
				if checkPlayerWallCollision(s, 3) != true {
					player.movement(s, 0, 2)
				}
			} else if ev.Rune() == 'd' || ev.Key() == tcell.KeyRight {
				if checkPlayerWallCollision(s, 2) != true {
					player.movement(s, 2, 0)
				}
			}
		}

		display(s, currentRoom)
		player.display(s, currentRoom)
		spot, _, _, _ := s.GetContent(player.pos_x-1, player.pos_y-1)
		s.SetContent(5, 47, spot, nil, menuStyle)
		s.Sync()
	}
}
