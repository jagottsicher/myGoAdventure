package main

import (
	"fmt"
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

			// theRow = "abc"
			// theRow = theRow + " " + fmt.Sprintf("%d", int(percentageX)) + " "
			// emitStr(s, 0, x, roomStyle, theRow)
		}

		fmt.Print(theRow)
	}

	// line by line
	// top and last lines repeat twice, the inner lines 8 times
	// for lines, content := range r.compressedRoomData[0:1] {
	// 	for yScale := 0; yScale < defaultYFactor; yScale++ {
	// 		// column by column
	// 		for columns, char := range content {
	// 			// output each char x times
	// 			for xScale := 0; xScale < defaultXFactor; xScale++ {
	// 				emitStr(s, ((defaultXFactor * columns) + xScale), (lines + yScale), roomStyle, string(char))
	// 			}
	// 		}
	// 	}
	// }

	// for lines, content := range r.compressedRoomData[1:11] {
	// for lines, content := range r.compressedRoomData {

	// 	for yScale := 0; yScale < r.allRoomDimensions.defaultYInnerFactor; yScale++ {
	// 		// column by column
	// 		for columns, char := range content {
	// 			// output each char x times
	// 			for xScale := 0; xScale < defaultXFactor; xScale++ {
	// 				emitStr(s, ((defaultXFactor * columns) + xScale), ((r.allRoomDimensions.defaultYInnerFactor * (lines)) + defaultYFactor + yScale), roomStyle, string(char))
	// 			}
	// 		}
	// 	}
	// }

	// for lines, content := range r.compressedRoomData[11:] {
	// 	for yScale := 0; yScale < defaultYFactor; yScale++ {
	// 		// column by column
	// 		for columns, char := range content {
	// 			// output each char x times
	// 			for xScale := 0; xScale < defaultXFactor; xScale++ {
	// 				emitStr(s, ((defaultXFactor * columns) + xScale), (lines + yScale), roomStyle, string(char))
	// 			}
	// 		}
	// 	}
	// }

	// for _, content := range r.compressedRoomData[11:] {
	// 	for yScale := 0; yScale < defaultYFactor; yScale++ {
	// 		// column by column
	// 		for columns, char := range content {
	// 			// output each char x times
	// 			for xScale := 0; xScale < defaultXFactor; xScale++ {
	// 				// emitStr(s, ((defaultXFactor * columns) + xScale), ((r.allRoomDimensions.dimensions.height) - defaultYFactor + yScale - 4), roomStyle, string(char))
	// 				emitStr(s, ((defaultXFactor * columns) + xScale), (lines + yScale), roomStyle, string(char))
	// 			}
	// 		}
	// 	}
	// }

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

	// build the gates
	if r == &roomYellowCastle {
		// output gates
		displayGate(s, &yellowCastleGate)
	}

	if r == &roomBlackCastle {
		// output gates
		displayGate(s, &blackCastleGate)
	}

	if r == &roomWhiteCastle {
		// output gates
		displayGate(s, &whiteCastleGate)
	}

	// render the bars, if needed
	if r.barLeft == true {
		displayBar(s, &leftBar)
	}
	if r.barRight == true {
		displayBar(s, &rightBar)
	}

	// // fill area black
	// for i := 0; i < 4; i++ {
	// 	emitStr(s, 0, (r.allRoomDimensions.dimensions.height - 4 + i), menuStyle, "                                                                                                                                                                ")
	// }

	// emitStr(s, (r.allRoomDimensions.dimensions.width-len("[ESC] Quit   [F9] Color Mode   [F10] Game Type   [F11] Difficulty   [F12] Reset/Retry"))/2, (r.allRoomDimensions.dimensions.height - 3), menuStyle, "[ESC] Quit   [F9] Color Mode   [F10] Game Type   [F11] Difficulty   [F12] Reset/Retry")

}

func displayGate(s tcell.Screen, g *gate) {
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			emitStr(s, g.pos_x+w, g.pos_y+h, g.objectStyle, string(g.buildRune))
		}
	}
}
func displayBar(s tcell.Screen, b *bar) {
	for h := 0; h < b.height; h++ {
		for w := 0; w < b.width; w++ {
			emitStr(s, b.pos_x+w, h, b.objectStyle, string(b.buildRune))
		}
	}
}
