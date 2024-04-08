package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

var someContent rune

func DrawState() {
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

	for _, obj := range allObjects {
		drawObject(obj)
	}

}

func drawObject(obj *object) {

	for _, point := range obj.shape {
		screen.SetContent(obj.posX+point.x, obj.posY+point.y, point.symbol, nil, obj.style)
	}
	// screen.Sync()

}

func initGamestate() {
	initDirections()
	uncompressRooms()
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

	fillTheScreen()

	screen.Show()
}

func fillTheScreen() {
	// width, height := screen.Size()

	// fill screen with random content
	// for y := 0; y < height; y++ {
	// 	for x := 0; x < width; x++ {
	// 		someContent = rune(rand.Intn(92) + 32)
	// 		screen.SetContent(x, y, someContent, nil, tcell.StyleDefault.
	// 			Background(tcell.ColorBlack).
	// 			Foreground(tcell.ColorYellow))
	// 		aPoint := point{x, y, someContent}
	// 		currentScreen = append(currentScreen, &aPoint)
	// 	}
	// }

	// fill screen with yellow castle content

	// currentScreen = nil
	// for _, cell := range roomYellowCastle.uncompressedRoomData {
	// 	screen.SetContent(cell.x, cell.y, cell.symbol, nil, tcell.StyleDefault.
	// 		Background(tcell.ColorBlack).
	// 		Foreground(tcell.ColorYellow))
	// 	currentScreen = append(currentScreen, cell)
	// }

	currentScreen = nil
	// y := 0
	// for _, cell := range *roomYellowCastle.compressedRoomData {
	// 	emitStr(screen, 0, y, tcell.StyleDefault.
	// 		Background(tcell.ColorBlack).
	// 		Foreground(tcell.ColorYellow), cell)
	// 	y += 1
	// 	// currentScreen = append(currentScreen, cell)
	// }

	for _, cell := range roomYellowCastle.uncompressedRoomData {
		emitStr(screen, cell.x, cell.y, tcell.StyleDefault.
			Background(tcell.ColorBlack).
			Foreground(tcell.ColorYellow), string(cell.symbol))

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
