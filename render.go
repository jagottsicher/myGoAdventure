package main

import (
	"fmt"
	"os"

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

// drawStatusBar renders a single-line HUD at the very bottom of the terminal
// showing the current room and player coordinates.
func drawStatusBar() {
	termW, termH := screen.Size()
	if termH < 2 {
		return
	}
	name := roomNames[player.roomID]
	if name == "" {
		name = "???"
	}
	status := fmt.Sprintf(" Room 0x%02X: %-28s  pos (%.1f, %.1f)  WASD/Arrows to move  Q to quit ",
		player.roomID, name, player.posX, player.posY)
	if len(status) > termW {
		status = status[:termW]
	}
	style := tcell.StyleDefault.Background(tcell.ColorNavy).Foreground(tcell.ColorWhite)
	emitStr(screen, 0, termH-1, style, status)
	// Pad to full width
	for i := len([]rune(status)); i < termW; i++ {
		screen.SetContent(i, termH-1, ' ', nil, style)
	}
}

// roomNames maps room IDs to human-readable names.
var roomNames = map[int]string{
	0x00: "Number Room",
	0x01: "Top Access (S)",
	0x02: "Top Access (N)",
	0x03: "Left of Name Room",
	0x04: "Top of Blue Maze",
	0x05: "Blue Maze #1",
	0x06: "Bottom of Blue Maze",
	0x07: "Center of Blue Maze",
	0x08: "Blue Maze Entry",
	0x09: "Maze Middle",
	0x0A: "Maze Entry",
	0x0B: "Maze Side",
	0x0C: "Side Corridor (S)",
	0x0D: "Side Corridor (N)",
	0x0E: "Top Entry Room",
	0x0F: "White Castle",
	0x10: "Black Castle",
	0x11: "Yellow Castle",
	0x12: "Yellow Castle Entry",
	0x13: "Black Maze #1",
	0x14: "Black Maze #2",
	0x15: "Black Maze #3",
	0x16: "Black Maze Entry",
	0x17: "Red Maze #1",
	0x18: "Top of Red Maze",
	0x19: "Bottom of Red Maze",
	0x1A: "White Castle Entry",
	0x1B: "Black Castle Entry",
	0x1C: "Purple Room",
	0x1D: "Red Top Entry",
	0x1E: "Name Room (Easter Egg)",
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

// initScreen initialises the tcell screen, shows the splash screen, and
// blocks until the player presses a key (or ESC/Q to quit).
func initScreen() {
	var err error
	screen, err = newScreen()
	if err != nil {
		panic(err)
	}
	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
	screen.Clear()
	drawSplash()

	// Wait for any keypress before starting the game.
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Name() == "Rune[q]" {
				screen.Fini()
				os.Exit(0)
			}
			return // any other key starts the game
		case *tcell.EventResize:
			screen.Sync()
			drawSplash()
		}
	}
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
