package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {

	InitColor()
	// defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	// bgrStyle := tcell.StyleDefault.Foreground(tcell.ColorLightGray).Background(tcell.ColorLightGray)
	// OrangeDrawStyle := tcell.StyleDefault.Foreground(tcell.ColorOrange).Background(tcell.ColorOrange)
	// YellowDrawStyle := tcell.StyleDefault.Foreground(tcell.Color226).Background(tcell.Color226)

	// fmt.Printf("%T", YellowDrawStyle)
	// fmt.Scanln()

	testStage := []string{
		"XXXXXXXXXXXXXXXX   XXX   XXX   XXX            XXX   XXX   XXX   XXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXX   XXX   XXX   XXX            XXX   XXX   XXX   XXXXXXXXXXXXXXXX",
		"XX           XXX   XXX   XXX   XXX            XXX   XXX   XXX   XXX           XX",
		"XX           XXX   XXX   XXX   XXX            XXX   XXX   XXX   XXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXX            XXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
		"XX           XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX           XX",
		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                  XXXXXXXXXXXXXXX   ####   XXXXXXXXXXXXXXX                  XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XX                                                                            XX",
		"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"XXXXXXXXXXXXXXXXXXXXXXXXXXXX                        XXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	}

	testStage2 := []string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                                                                ",
	"                                                                                ",
	"                                                                                ",
	"                                                                                ",
	"                                                                                ",
	"                                                                                ",
	"                                                                                ",
	"                                                                                ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"                              XX                XX                              ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXX    XX                XX    XXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX            XX    XX    XXXXXXXX    XX    XX            XXXXXXXXXXXX",
	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
	"XXXXXXXXXXXX    XXXX    XXXXXXXX    XXXXXXXX    XXXXXXXX    XXXX    XXXXXXXXXXXX",
}

	//  Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(bgrStyle)
	// s.EnableMouse()
	// s.EnablePaste()
	s.Clear()

	// fmt.Println(testStage[1])

	// Draw initial boxes
	// drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	// drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")
	//drawBox(s, 0, 0, 1, 0, dotStyle)

	var tempString string

	for i, _ := range testStage {
		tempString = testStage[i]
		characters := []rune(tempString)
		for j := 0; j < len(tempString); j++ {
			if string(characters[j]) == "X" {
				col := j * 2
				s.SetContent(col, i, 'X', nil, YellowDrawStyle)
				s.SetContent(col+1, i, 'X', nil, YellowDrawStyle)
			}
		}
	}

	s2, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s2.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s2.SetStyle(bgrStyle)
	// s.EnableMouse()
	// s2.EnablePaste()
	s2.Clear()

	for i, _ := range testStage2 {
		tempString = testStage2[i]
		characters := []rune(tempString)
		for j := 0; j < len(tempString); j++ {
			if string(characters[j]) == "X" {
				col := j * 2
				s2.SetContent(col, i, 'X', nil, OrangeDrawStyle)
				s2.SetContent(col+1, i, 'X', nil, OrangeDrawStyle)
			}
		}
	}

	// Event loop
	// ox, oy := -1, -1
	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	nextScreen := false

	for {
		// Update screen
		if nextScreen == false {
			s.Show()
		} else {
			s2.Show()
		}

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			} else if ev.Rune() == 'd' {
				s.Clear()
				drawBox(s, 5, 0, 9, 1, OrangeDrawStyle)
			} else if ev.Rune() == 'n' {
				nextScreen = true
			}
			// case *tcell.EventMouse:
			// 	x, y := ev.Position()
			// 	button := ev.Buttons()
			// 	// Only process button events, not wheel events
			// 	button &= tcell.ButtonMask(0xff)

			// 	if button != tcell.ButtonNone && ox < 0 {
			// 		ox, oy = x, y
			// 	}
			// 	switch ev.Buttons() {
			// 	case tcell.ButtonNone:
			// 		if ox >= 0 {
			// 			label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
			// 			drawBox(s, ox, oy, x, y, boxStyle, label)
			// 			ox, oy = -1, -1
			// 		}
			// 	}
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	// for col := x1; col <= x2; col++ {
	// 	s.SetContent(col, y1, tcell.RuneHLine, nil, style)
	// 	s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	// }
	// for row := y1 + 1; row < y2; row++ {
	// 	s.SetContent(x1, row, tcell.RuneVLine, nil, style)
	// 	s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	// }

	// Only draw corners if necessary
	// if y1 != y2 && x1 != x2 {
	// 	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	// 	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	// 	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	// 	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	// }

	// drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

// func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
// 	row := y1
// 	col := x1
// 	for _, r := range []rune(text) {
// 		s.SetContent(col, row, r, nil, style)
// 		col++
// 		if col >= x2 {
// 			row++
// 			col = x1
// 		}
// 		if row > y2 {
// 			break
// 		}
// 	}
// }
