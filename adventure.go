package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

// global size of terminal windows
// var stageWidth, stageHeight, fd int

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	encoding.Register()

	// init handler to get determinate screensize
	// fd = int(os.Stdin.Fd())

	// get terminal dimensions for the first time
	// stageWidth, stageHeight, _ = terminal.GetSize(fd)

	// init all screens
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	// s.SetStyle(tcell.StyleDefault.
	// 	Background(tcell.ColorDarkGray).
	// 	Foreground(tcell.ColorYellow))

	player.init()

	render(s)

	// initStages(rooms)

	// init a player for the first time
	// player.init()

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			// reinit all screens in new sizes
			render(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				clearScreen()
				fmt.Println("Bye.")
				os.Exit(0)
			} else if ev.Rune() == 'w' || ev.Key() == tcell.KeyUp {
				// if checkPlayerWallCollision(s, 1) != true {
				player.pos_y -= 1
				// }
			} else if ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft {
				// if checkPlayerWallCollision(s, 4) != true {
				player.pos_x -= 2
				// }
			} else if ev.Rune() == 's' || ev.Key() == tcell.KeyDown {
				// if checkPlayerWallCollision(s, 3) != true {
				player.pos_y += 1
				// }
			} else if ev.Rune() == 'd' || ev.Key() == tcell.KeyRight {
				// if checkPlayerWallCollision(s, 2) != true {
				player.pos_x += 2
				// }
			}
		}
		for i := 0; i < player.width; i++ {
			s.SetContent(player.pos_x+i, player.pos_y, player.contentX[i], nil, player.contentStyle[i])
		}
		player.display(s)
		s.Show()
	}
}
