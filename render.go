package main

import (
	"math/rand"

	"github.com/gdamore/tcell"
)

var someContent rune

func render(screen tcell.Screen) {

	width, height := screen.Size()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			someContent = rune(rand.Intn(92) + 32)
			screen.SetContent(x, y, someContent, nil, tcell.StyleDefault.
				Background(tcell.ColorBlack).
				Foreground(tcell.ColorYellow))
		}
	}

	player.display(screen)

	screen.Show()

	// set the room colors
	// roomStyle := tcell.StyleDefault.
	// 	Background(tcell.ColorDarkGray).
	// 	Foreground(tcell.ColorDarkGray)
	// screen.SetStyle(roomStyle)

	// display room

	// remember what under player

	// emit player
	// player.display(screen)
	// 	screen.Sync()

	// // Any maze

	// // defaultXFactor = int(math.Round(float64(stageWidth) / 40))
	// // defaultYFactor = 2

	// stageYFactor = int(math.Floor(float64(stageHeight) / 12))

	// var percentageY float64
	// var rowValue int
	// MaxRowValue := len(r.compressedRoomData)
	// var percentageX float64
	// var columnValue int
	// MaxColumnValue := 40

	// var theRow string
	// var theSpot rune

	// for y := 0; y < stageHeight; y++ {
	// 	percentageY = float64((y * 100) / (stageHeight - 1))
	// 	rowValue = int(12 * int(percentageY) / 100)
	// 	if rowValue == MaxRowValue {
	// 		rowValue = MaxRowValue - 1
	// 	}
	// 	theRow = r.compressedRoomData[rowValue]

	// 	for x := 0; x < stageWidth; x++ {
	// 		percentageX = float64((x * 100) / (stageWidth - 1))
	// 		columnValue = int(40 * int(percentageX) / 100)
	// 		if columnValue == MaxColumnValue {
	// 			columnValue = MaxColumnValue - 1
	// 		}

	// 		theSpot = rune(theRow[columnValue])

	// 		s.SetContent(x, y, theSpot, nil, roomStyle)
	// 	}
	// }

	// if r == &roomMazeEntry || r == &roomMazeMiddle || r == &roomMazeSide {
	// 	// print the player surrounding
	// 	roomUncoveredStyle := tcell.StyleDefault.
	// 		Background(tcell.ColorOrange).
	// 		Foreground(tcell.ColorDarkGray)

	// 	const surroundingDimX = 36
	// 	const surroundingDimY = 18

	// 	var surroundingContent [surroundingDimY][surroundingDimX]rune

	// 	for surY := 0; surY < surroundingDimY; surY++ {
	// 		for surX := 0; surX < surroundingDimX; surX++ {
	// 			surroundingContent[surY][surX], _, _, _ = s.GetContent(player.pos_x-16+surX, player.pos_y-8+surY)
	// 		}
	// 	}

	// 	for surY := 0; surY < surroundingDimY; surY++ {
	// 		for surX := 0; surX < surroundingDimX; surX++ {
	// 			emitStr(s, player.pos_x-16+surX, player.pos_y-8+surY, roomUncoveredStyle, string(surroundingContent[surY][surX]))
	// 		}
	// 	}
	// }
}
