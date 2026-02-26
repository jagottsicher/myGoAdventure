package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// needsRedraw is set to true by the resize handler so the main loop knows to
// rebuild the stage geometry on the next frame.
var needsRedraw bool

// gridToScreen maps a room grid coordinate to terminal cell coordinates.
// Room grid: gridCols × gridRows (40 × 7).
// Result is clamped to the current terminal dimensions.
func gridToScreen(gx, gy float64, termW, termH int) (int, int) {
	sx := int(gx / float64(gridCols) * float64(termW))
	sy := int(gy / float64(gridRows) * float64(termH))
	return sx, sy
}

// drawStage renders the walls of the current player's room.
// It iterates over the room's decoded wall map and paints each cell.
func drawStage() {
	termW, termH := screen.Size()
	r := roomByID(player.roomID)
	if r == nil {
		return
	}

	bg := tcell.ColorBlack
	fg := r.color

	wallStyle := tcell.StyleDefault.Background(bg).Foreground(fg)
	emptyStyle := tcell.StyleDefault.Background(bg).Foreground(bg)

	// Calculate cell dimensions for the current terminal size.
	cellW := termW / gridCols // terminal columns per grid column (may be 0 for tiny terminals)
	cellH := termH / gridRows // terminal rows per grid row
	if cellW < 1 {
		cellW = 1
	}
	if cellH < 1 {
		cellH = 1
	}

	// Paint each grid cell individually.
	for row := 0; row < gridRows; row++ {
		for col := 0; col < gridCols; col++ {
			// Top-left terminal coordinate for this grid cell.
			startX, startY := gridToScreen(float64(col), float64(row), termW, termH)

			// Determine how many terminal cells this grid cell occupies.
			endX, endY := gridToScreen(float64(col+1), float64(row+1), termW, termH)

			var style tcell.Style
			var ch rune
			if r.walls[row][col] {
				style = wallStyle
				ch = '█'
			} else {
				style = emptyStyle
				ch = ' '
			}

			for ty := startY; ty < endY && ty < termH; ty++ {
				for tx := startX; tx < endX && tx < termW; tx++ {
					screen.SetContent(tx, ty, ch, nil, style)
				}
			}
		}
	}

	// Fill any remaining terminal area not covered by the grid with background.
	// (Happens when termW or termH is not evenly divisible by gridCols/gridRows.)
	_ = cellW
	_ = cellH
}

// drawAllVisibleObjects renders every object in allObjects whose roomID
// matches the player's current room.
func drawAllVisibleObjects() {
	for _, obj := range allObjects {
		if obj.roomID == player.roomID {
			drawObject(obj)
		}
	}
}

// drawObject paints a single object at its current grid position,
// scaled to the current terminal size.
func drawObject(obj *object) {
	termW, termH := screen.Size()
	sx, sy := gridToScreen(obj.posX, obj.posY, termW, termH)

	for _, pt := range obj.shape {
		tx := sx + pt.x
		ty := sy + pt.y
		if tx >= 0 && tx < termW && ty >= 0 && ty < termH {
			screen.SetContent(tx, ty, pt.symbol, nil, obj.style)
		}
	}
}

// drawSplash renders the animated startup/title screen and waits for the user
// to press a key or for a brief timeout.
func drawSplash() {
	termW, termH := screen.Size()

	title := "An Adventure – dedicated to Warren Robinett"
	sub := "Use [WASD] or arrow keys to move. Press [Q] to quit."

	titleX := termW/2 - len(title)/2
	subX := termW/2 - len(sub)/2
	titleY := termH/2 - 1
	subY := termH / 2

	screen.Clear()
	emitStr(screen, titleX, titleY, tcell.StyleDefault.Bold(true), title)
	emitStr(screen, subX, subY, tcell.StyleDefault, sub)
	screen.Show()
}

// initScreen initialises the tcell screen, displays the splash screen, and
// returns when the user is ready to play.
func initScreen() {
	var err error
	screen, err = newScreen()
	if err != nil {
		panic(err)
	}
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	screen.Clear()
	drawSplash()
}

// emitStr writes a string at (x, y) using the given style, handling
// multi-column Unicode characters via go-runewidth.
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
