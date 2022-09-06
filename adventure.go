package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/encoding"
)

// This program just prints "Hello, World!".  Press ESC to exit.
func main() {
	encoding.Register()

	// ball = player{
	// 	// playerPosition: position{
	// 	// 	x: 80,
	// 	// y: 40,
	// 	// }
	// 	// size:           spriteSize{
	// 	// 	width: 4,
	// 	// 	height:2,
	// 	// }
	// 	color: tcell.Style{
	// 		Default.Foreground(tcell.ColorYellow.TrueColor()).Background(tcell.ColorGreen),
	// 	},
	// }

	Player.init()

	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	// displayPlayer(s)
	Player.display(s)

	for {
		switch ev := s.PollEvent().(type) {
		case *tcell.EventResize:
			s.Sync()
			Player.display(s)
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				s.Fini()
				clearScreen()
				fmt.Println("Bye.")
				os.Exit(0)
			} else if ev.Rune() == 'w' || ev.Key() == tcell.KeyUp {
				_, h := s.Size()
				if Player.pos_y-1 < 0 {
					Player.pos_y += h
				}
				Player.pos_y -= 1
				Player.display(s)
			} else if ev.Rune() == 'a' || ev.Key() == tcell.KeyLeft {
				w, _ := s.Size()
				if Player.pos_x-2 < 0 {
					Player.pos_x += w
				}
				Player.pos_x -= 2
				Player.display(s)
			} else if ev.Rune() == 's' || ev.Key() == tcell.KeyDown {
				_, h := s.Size()
				if Player.pos_y > h-Player.height {
					Player.pos_y -= h
				}
				Player.pos_y += 1
				Player.display(s)
			} else if ev.Rune() == 'd' || ev.Key() == tcell.KeyRight {
				w, _ := s.Size()
				if Player.pos_x > w-Player.width {
					Player.pos_x -= w
				}
				Player.pos_x += 2
				Player.display(s)
			}
		}
	}
}

// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/gdamore/tcell/v2"
// )

// func main() {

// 	InitColors()
// 	InitContent()
// 	InitPlayer()

// 	currentScreen, err := tcell.NewScreen()
// 	if err != nil {
// 		log.Fatalf("%+v", err)
// 	}
// 	if err := currentScreen.Init(); err != nil {
// 		log.Fatalf("%+v", err)
// 	}
// 	currentScreen.SetStyle(bgrStyle)
// 	// s.EnableMouse()
// 	// s.EnablePaste()
// 	currentScreen.Clear()

// 	// Event loop
// 	// ox, oy := -1, -1
// 	quit := func() {
// 		currentScreen.Fini()
// 		os.Exit(0)
// 	}

// 	switchScreen(currentScreen, castles, YellowDrawStyle)
// 	// getUnderPlayer(currentScreen, playerX, playerY)
// 	drawPlayer(currentScreen, playerY, playerX)

// 	for {
// 		currentScreen.Show()

// 		// Poll event
// 		ev := currentScreen.PollEvent()

// 		// Process event
// 		switch ev := ev.(type) {
// 		case *tcell.EventResize:
// 			currentScreen.Sync()
// 		case *tcell.EventKey:
// 			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
// 				quit()
// 			} else if ev.Key() == tcell.KeyCtrlL {
// 				currentScreen.Sync()
// 			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
// 				currentScreen.Clear()
// 			} else if ev.Rune() == 'n' {
// 				switchScreen(currentScreen, castles, BlackDrawStyle)
// 			} else if ev.Rune() == 'p' {
// 				switchScreen(currentScreen, bottomOpen, DarkGreenDrawStyle)
// 			} else if ev.Rune() == 'w' {
// 				playerY -= 2
// 			} else if ev.Rune() == 'a' {
// 				playerX -= 4
// 			} else if ev.Rune() == 's' {
// 				playerY += 2
// 			} else if ev.Rune() == 'd' {
// 				playerX += 4
// 			}
// 			erasePlayer(currentScreen, playerX, playerY)
// 			// getUnderPlayer(currentScreen, playerX, playerY)
// 			drawPlayer(currentScreen, playerX, playerY)

// 			// case *tcell.EventMouse:
// 			// 	x, y := ev.Position()
// 			// 	button := ev.Buttons()
// 			// 	// Only process button events, not wheel events
// 			// 	button &= tcell.ButtonMask(0xff)
// 			// 	if button != tcell.ButtonNone && ox < 0 {
// 			// 		ox, oy = x, y
// 			// 	}
// 			// 	switch ev.Buttons() {
// 			// 	case tcell.ButtonNone:
// 			// 		if ox >= 0 {
// 			// 			label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
// 			// 			drawBox(s, ox, oy, x, y, boxStyle, label)
// 			// 			ox, oy = -1, -1
// 			// 		}
// 			// 	}
// 		}
// 	}
// }

// func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
// 	if y2 < y1 {
// 		y1, y2 = y2, y1
// 	}
// 	if x2 < x1 {
// 		x1, x2 = x2, x1
// 	}

// 	// Fill background
// 	for row := y1; row <= y2; row++ {
// 		for col := x1; col <= x2; col++ {
// 			s.SetContent(col, row, ' ', nil, style)
// 		}
// 	}

// 	// Draw borders
// 	// for col := x1; col <= x2; col++ {
// 	// 	s.SetContent(col, y1, tcell.RuneHLine, nil, style)
// 	// 	s.SetContent(col, y2, tcell.RuneHLine, nil, style)
// 	// }
// 	// for row := y1 + 1; row < y2; row++ {
// 	// 	s.SetContent(x1, row, tcell.RuneVLine, nil, style)
// 	// 	s.SetContent(x2, row, tcell.RuneVLine, nil, style)
// 	// }

// 	// Only draw corners if necessary
// 	// if y1 != y2 && x1 != x2 {
// 	// 	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
// 	// 	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
// 	// 	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
// 	// 	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
// 	// }

// 	// drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
// }

// // func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
// // 	row := y1
// // 	col := x1
// // 	for _, r := range []rune(text) {
// // 		s.SetContent(col, row, r, nil, style)
// // 		col++
// // 		if col >= x2 {
// // 			row++
// // 			col = x1
// // 		}
// // 		if row > y2 {
// // 			break
// // 		}
// // 	}
// // }
