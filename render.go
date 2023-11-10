package main

import (
	"math"

	"github.com/gdamore/tcell"
)

func display(s tcell.Screen, r *rooms) {

	// set the room colors
	roomStyle := r.roomStyle
	s.SetStyle(roomStyle)

	// Any maze

	// defaultXFactor = int(math.Round(float64(stageWidth) / 40))
	// defaultYFactor = 2

	stageYFactor = int(math.Floor(float64(stageHeight) / 12))

	var percentageY float64
	var rowValue int
	MaxRowValue := len(r.compressedRoomData)
	var percentageX float64
	var columnValue int
	MaxColumnValue := 40

	var theRow string
	var theSpot rune

	for y := 0; y < stageHeight; y++ {
		percentageY = float64((y * 100) / (stageHeight - 1))
		rowValue = int(12 * int(percentageY) / 100)
		if rowValue == MaxRowValue {
			rowValue = MaxRowValue - 1
		}
		theRow = r.compressedRoomData[rowValue]

		for x := 0; x < stageWidth; x++ {
			percentageX = float64((x * 100) / (stageWidth - 1))
			columnValue = int(40 * int(percentageX) / 100)
			if columnValue == MaxColumnValue {
				columnValue = MaxColumnValue - 1
			}

			theSpot = rune(theRow[columnValue])

			s.SetContent(x, y, theSpot, nil, roomStyle)
		}

		//fmt.Print(theRow)
	}

	if r == &roomMazeEntry || r == &roomMazeMiddle || r == &roomMazeSide {
		// print the player surrounding
		roomUncoveredStyle := tcell.StyleDefault.
			Background(tcell.ColorOrange).
			Foreground(tcell.ColorDarkGray)

		const surroundingDimX = 36
		const surroundingDimY = 18

		var surroundingContent [surroundingDimY][surroundingDimX]rune

		for surY := 0; surY < surroundingDimY; surY++ {
			for surX := 0; surX < surroundingDimX; surX++ {
				surroundingContent[surY][surX], _, _, _ = s.GetContent(player.pos_x-16+surX, player.pos_y-8+surY)
			}
		}

		for surY := 0; surY < surroundingDimY; surY++ {
			for surX := 0; surX < surroundingDimX; surX++ {
				emitStr(s, player.pos_x-16+surX, player.pos_y-8+surY, roomUncoveredStyle, string(surroundingContent[surY][surX]))
			}
		}
	}

	// emitStr(s, (r.allRoomDimensions.dimensions.width-len("[ESC] Quit   [F9] Color Mode   [F10] Game Type   [F11] Difficulty   [F12] Reset/Retry"))/2, (r.allRoomDimensions.dimensions.height - 3), menuStyle, "[ESC] Quit   [F9] Color Mode   [F10] Game Type   [F11] Difficulty   [F12] Reset/Retry")

}
