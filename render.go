package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func drawStage() {
	// if isGamePaused {
	// 	return
	// }

	// screen.Clear()

	// PrintString(0, 0, debugLog)
	//

	for _, point := range currentScreen {
		screen.SetContent(point.x, point.y, point.symbol, nil, tcell.StyleDefault.
			Background(roomYellowCastle.background).
			Foreground(roomYellowCastle.foreground))
	}
}

func drawAllVisibleobjects() {
	// TODO: later only the ones which are visible
	for _, obj := range allObjects {
		drawObject(obj)
	}
}

func drawObject(obj *object) {
	termW, termH := screen.Size()
	screenX := int(obj.relX * float64(termW))
	screenY := int(obj.relY * float64(termH))
	for _, point := range obj.shape {
		px := screenX + point.x
		py := screenY + point.y
		if px < 0 || px >= termW || py < 0 || py >= termH {
			continue
		}
		screen.SetContent(px, py, point.symbol, nil, obj.style)
	}
}

func initGamestate() {
	initDirections()
	// uncompressRooms()
	initPlayer()
}

func initScreen() {
	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	screen.SetStyle(tcell.StyleDefault.
		Background(roomSplashScreen.background).
		Foreground(roomSplashScreen.foreground))

	for i := 0; i < 256; i++ {

		screen.SetStyle(tcell.StyleDefault.
			Background(roomSplashScreen.background).
			Foreground(tcell.NewRGBColor(int32(i), int32(i), int32(i))))

		screenWidth, screenHeight := screen.Size()
		title := "An Adventure - dedicated to Warren Robinett"
		subline := "Press [Q] to exit."
		emitStr(screen, screenWidth/2-len(title)/2, screenHeight/2-1, tcell.StyleDefault, title)
		emitStr(screen, screenWidth/2-len(subline)/2, screenHeight/2, tcell.StyleDefault, subline)

		time.Sleep(time.Millisecond * 10)
		screen.Show()
	}
	time.Sleep(time.Second * 3)

	fillTheScreen()
}

func fillTheScreen() {
	termW, termH := screen.Size()
	template := *roomYellowCastle.roomData
	templateH := len(template)
	if templateH == 0 {
		return
	}
	templateW := len([]rune(template[0]))
	if templateW == 0 {
		return
	}

	currentScreen = nil
	for ty := 0; ty < termH; ty++ {
		srcY := (2*ty + 1) * templateH / (2 * termH)
		row := []rune(template[srcY])
		for tx := 0; tx < termW; tx++ {
			srcX := (2*tx + 1) * templateW / (2 * termW)
			var ch rune = ' '
			if srcX < len(row) {
				ch = row[srcX]
			}
			currentScreen = append(currentScreen, &cell{x: tx, y: ty, symbol: ch})
		}
	}
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}
