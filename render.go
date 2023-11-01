package main

import "github.com/gdamore/tcell"

func display(s tcell.Screen, r *rooms) {

	// set the room colors
	roomStyle := r.roomStyle
	s.SetStyle(roomStyle)

	// Any maze

	// line by line
	// top and last lines repeat twice, the inner lines 8 times
	for lines, content := range r.compressedRoomData[0:1] {
		for yScale := 0; yScale < r.allRoomDimensions.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.allRoomDimensions.defaultXFactor; xScale++ {
					emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), (lines + yScale), roomStyle, string(char))
				}
			}
		}
	}

	// var barsStyle = tcell.StyleDefault.
	// 	Background(tcell.ColorDarkGray).
	// 	Foreground(tcell.ColorBlack)

	for lines, content := range r.compressedRoomData[1:11] {
		for yScale := 0; yScale < r.allRoomDimensions.defaultYInnerFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.allRoomDimensions.defaultXFactor; xScale++ {
					if string(char) != "+" {
						emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), ((r.allRoomDimensions.defaultYInnerFactor * (lines)) + r.allRoomDimensions.defaultYFactor + yScale), roomStyle, string(char))
					} else {
						// emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), ((r.allRoomDimensions.defaultYInnerFactor * (lines)) + r.allRoomDimensions.defaultYFactor + yScale), barsStyle, string(char))
					}
				}
			}
		}
	}

	for _, content := range r.compressedRoomData[11:] {
		for yScale := 0; yScale < r.allRoomDimensions.defaultYFactor; yScale++ {
			// column by column
			for columns, char := range content {
				// output each char x times
				for xScale := 0; xScale < r.allRoomDimensions.defaultXFactor; xScale++ {
					emitStr(s, ((r.allRoomDimensions.defaultXFactor * columns) + xScale), ((r.allRoomDimensions.dimensions.height) - r.allRoomDimensions.defaultYFactor + yScale - 4), roomStyle, string(char))
				}
			}
		}
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

	if r == &roomYellowCastle {
		// output gates
		displayObject(s, &yellowCastleGate)
	}

	if r == &roomBlackCastle {
		// output gates
		displayObject(s, &blackCastleGate)
	}

	if r == &roomWhiteCastle {
		// output gates
		displayObject(s, &whiteCastleGate)
	}

	// fill area black
	for i := 0; i < 4; i++ {
		emitStr(s, 0, (r.allRoomDimensions.dimensions.height - 4 + i), menuStyle, "                                                                                                                                                                ")
	}

	emitStr(s, (r.allRoomDimensions.dimensions.width-len("[ESC] Quit   [F9] Color Mode   [F10] Game Type   [F11] Difficulty   [F12] Reset/Retry"))/2, (r.allRoomDimensions.dimensions.height - 3), menuStyle, "[ESC] Quit   [F9] Color Mode   [F10] Game Type   [F11] Difficulty   [F12] Reset/Retry")

}

func displayObject(s tcell.Screen, g *gate) {
	for h := 0; h < g.height; h++ {
		for w := 0; w < g.width; w++ {
			emitStr(s, g.pos_x+w, g.pos_y+h, g.objectStyle, string(g.buildRune))
		}
	}
}
