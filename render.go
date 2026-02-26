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
	template := *roomYellowCastle.compressedRoomData
	templateH := float64(len(template))
	templateW := float64(len([]rune(template[0])))

	screenX := int(obj.posX / templateW * float64(termW))
	screenY := int(obj.posY / templateH * float64(termH))

	for _, point := range obj.shape {
		screen.SetContent(screenX+point.x, screenY+point.y, point.symbol, nil, obj.style)
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
	template := *roomYellowCastle.compressedRoomData
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
		srcY := ty * templateH / termH
		row := []rune(template[srcY])
		for tx := 0; tx < termW; tx++ {
			srcX := tx * templateW / termW
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
