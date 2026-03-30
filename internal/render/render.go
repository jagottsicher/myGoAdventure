package render

import (
	"fmt"
	"os"
	"time"

	"development/myGoAdventure/internal/game"
	"development/myGoAdventure/internal/world"
	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

var Screen tcell.Screen
var CurrentScreen []*world.Cell

var auraColor = tcell.NewRGBColor(0xFF, 0x80, 0x00) // orange torch glow

// eeFlashHue is a slow independent hue for the Easter Egg room text color.
var eeFlashHue int
var eeFlashTick int

func DrawStage() {
	if game.CurrentRoom == nil {
		return
	}
	// Torch aura only in dark maze rooms (Foreground == Background → walls invisible).
	darkMaze := game.CurrentRoom.Foreground == game.CurrentRoom.Background
	termW, termH := Screen.Size()
	var playerCX, playerCY int
	if game.Player != nil {
		playerCX = int(game.Player.RelX * float64(termW))
		playerCY = int(game.Player.RelY * float64(termH))
	}
	winFlash := game.GameWon && game.WinFlashTimer > 0
	for _, point := range CurrentScreen {
		var style tcell.Style
		if point.Symbol == 'X' || point.Symbol == 'x' {
			wallColor := game.CurrentRoom.Foreground
			if winFlash {
				wallColor = game.GetFlashColor()
			} else if darkMaze && game.Player != nil {
				dx := point.X - playerCX
				dy := point.Y - playerCY
				// aura: 33 wide (±16), 16 high (±8)
				if dx >= -16 && dx <= 16 && dy >= -8 && dy <= 8 {
					wallColor = auraColor
				}
			}
			style = tcell.StyleDefault.Background(wallColor).Foreground(wallColor)
		} else {
			bg := game.CurrentRoom.Background
			fg := game.CurrentRoom.Foreground
			if winFlash {
				fc := game.GetFlashColor()
				bg = fc
				fg = fc
			}
			style = tcell.StyleDefault.Background(bg).Foreground(fg)
		}
		Screen.SetContent(point.X, point.Y, point.Symbol, nil, style)
	}
}

// DrawDebugBat draws bat state on the last row — temporary debug aid, remove after bat is confirmed working.
func DrawDebugBat() {
	termW, termH := Screen.Size()
	if game.Bat == nil {
		return
	}
	msg := game.BatDebugState()
	style := tcell.StyleDefault.Foreground(tcell.ColorYellow).Background(tcell.ColorDarkBlue)
	for i, r := range msg {
		if i >= termW {
			break
		}
		Screen.SetContent(i, termH-1, r, nil, style)
	}
}

func DrawAllVisibleObjects() {
	for layer := 0; layer <= 2; layer++ {
		for _, obj := range game.AllObjects {
			if obj.Dead || obj.ZLayer != layer {
				continue
			}
			if obj.Room != nil && obj.Room != game.CurrentRoom {
				continue
			}
			DrawObject(obj)
		}
	}
}

func DrawObject(obj *game.Object) {
	termW, termH := Screen.Size()
	ox, oy := obj.Width/2, obj.Height/2
	if obj.BodyOffsets != nil {
		off := obj.BodyOffsets[obj.OrientationFrame%len(obj.BodyOffsets)]
		ox, oy = off[0], off[1]
	}
	screenX := int(obj.RelX*float64(termW)) - ox
	screenY := int(obj.RelY*float64(termH)) - oy
	objFg, _, _ := obj.Style.Decompose()
	for _, point := range obj.Shape {
		px := screenX + point.X
		if obj.Flipped {
			px = screenX + (obj.Width - 1 - point.X)
		}
		py := screenY + point.Y
		if px < 0 || px >= termW || py < 0 || py >= termH {
			continue
		}
		style := obj.Style
		s := point.Symbol
		isHalfBlock := s >= 0x2580 && s <= 0x259F && s != '█'
		isBoxDrawing := s >= 0x2500 && s <= 0x257F
		if isHalfBlock || isBoxDrawing {
			_, _, existingStyle, _ := Screen.GetContent(px, py)
			_, existingBg, _ := existingStyle.Decompose()
			if obj == game.Dot && s == '▀' {
				// Use wall color as background only when the dot's own cell is a wall cell.
				// Mirror the aura logic from DrawStage for dark maze rooms.
				existingBg = game.CurrentRoom.Background
				darkMaze := game.CurrentRoom.Foreground == game.CurrentRoom.Background
				playerCX := int(game.Player.RelX * float64(termW))
				playerCY := int(game.Player.RelY * float64(termH))
				for _, cell := range CurrentScreen {
					if cell.X == px && cell.Y == py && (cell.Symbol == 'X' || cell.Symbol == 'x') {
						wallColor := game.CurrentRoom.Foreground
						if darkMaze {
							ddx := px - playerCX
							ddy := py - playerCY
							if ddx >= -16 && ddx <= 16 && ddy >= -8 && ddy <= 8 {
								wallColor = auraColor
							}
						}
						existingBg = wallColor
						break
					}
				}
			}
			style = tcell.StyleDefault.Foreground(objFg).Background(existingBg)
		}
		Screen.SetContent(px, py, point.Symbol, nil, style)
	}
}

func InitGamestate() {
	world.InitDirections(int(game.G.GameType))
	game.CurrentRoom = &world.RoomYellowCastle
	w, h := Screen.Size()
	game.InitPlayer(w, h)
	game.InitYellowKey(w, h)
	game.InitWhiteKey(w, h)
	game.InitBlackKey(w, h)
	game.InitGreenDragon(w, h)
	game.InitYellowDragon(w, h)
	game.InitRedDragon(w, h)
	game.InitBat(w, h)
	game.InitPortcullises(w, h)
	game.InitBridge(w, h)
	game.InitSword(w, h)
	game.InitChalice(w, h)
	game.InitMagnet(w, h)
	game.InitDot(w, h)

	// Variation 3: randomize object rooms per C++ roomBoundsData.
	game.RandomizeObjectsV3()

	// Passage barriers — black on black, only visible in their room.
	blackOnBlack := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	game.InitBarrier(&world.RoomTopAccessRight, 1.0/20.0, h, blackOnBlack)    // left of yellow castle: block left side
	game.EasterEggBarrier = game.InitBarrier(&world.RoomCorridorRight, 19.0/20.0, h, blackOnBlack) // Easter Egg: passable with Dot
	game.InitBarrier(&world.RoomSideCorridorOlive, 1.0/20.0, h, blackOnBlack)  // below white castle: block left side
	game.InitBarrier(&world.RoomSideCorridorCyan, 19.0/20.0, h, blackOnBlack) // cyan room next to it: block right side

	FillTheScreen()
}

func InitScreen() {
	var err error
	Screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := Screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	bg := world.RoomSplashScreen.Background
	Screen.SetStyle(tcell.StyleDefault.Background(bg).Foreground(bg))
	Screen.Clear()

	screenWidth, screenHeight := Screen.Size()
	drawSplashOnScreen(screenWidth, screenHeight, bg)
	Screen.Show()
	time.Sleep(time.Second * 5)
}

func FillTheScreen() {
	if game.CurrentRoom == nil {
		return
	}
	termW, termH := Screen.Size()
	template := *game.CurrentRoom.RoomData
	templateH := len(template)
	if templateH == 0 {
		return
	}
	templateW := len([]rune(template[0]))
	if templateW == 0 {
		return
	}

	CurrentScreen = nil
	for ty := 0; ty < termH; ty++ {
		srcY := (2*ty + 1) * templateH / (2 * termH)
		row := []rune(template[srcY])
		for tx := 0; tx < termW; tx++ {
			srcX := (2*tx + 1) * templateW / (2 * termW)
			var ch rune = ' '
			if srcX < len(row) {
				ch = row[srcX]
			}
			CurrentScreen = append(CurrentScreen, &world.Cell{X: tx, Y: ty, Symbol: ch})
		}
	}
}

// WouldCollideWall checks whether a bounding box at (screenX, screenY) with the given
// width/height overlaps any 'X' cell — either from the room template on screen or from
// a barrier object belonging to the current room (checked directly, not via screen buffer
// to avoid race conditions with the render loop).
func WouldCollideWall(screenX, screenY, width, height int) bool {
	// Build set of positions covered by the bridge — the entire bounding box is passable.
	// Applies whether the bridge is carried or dropped, as long as it is in the current room.
	bridged := map[[2]int]bool{}
	if b := game.Bridge; b != nil && b.Room == game.CurrentRoom {
		termW, termH := Screen.Size()
		box := int(b.RelX*float64(termW)) - b.Width/2
		boy := int(b.RelY*float64(termH)) - b.Height/2
		for y := 0; y < b.Height; y++ {
			for x := 0; x < b.Width; x++ {
				bridged[[2]int{box + x, boy + y}] = true
			}
		}
	}

	// Check room walls via CurrentScreen — authoritative wall list, never corrupted
	// by object rendering (unlike Screen.GetContent which objects overwrite).
	for _, cell := range CurrentScreen {
		if cell.Symbol != 'X' {
			continue
		}
		if bridged[[2]int{cell.X, cell.Y}] {
			continue // wall covered by bridge deck — passable
		}
		if cell.X >= screenX && cell.X < screenX+width &&
			cell.Y >= screenY && cell.Y < screenY+height {
			return true
		}
	}
	// Check barrier and solid objects directly (not via screen buffer).
	termW, termH := Screen.Size()
	for _, obj := range game.AllObjects {
		if obj.Room == nil || obj.Room != game.CurrentRoom {
			continue
		}
		if len(obj.Shape) == 0 {
			continue
		}
		ox := int(obj.RelX*float64(termW)) - obj.Width/2
		oy := int(obj.RelY*float64(termH)) - obj.Height/2
		for _, cell := range obj.Shape {
			// Barriers: only 'X' cells block.
			// Solid objects (e.g. closed portcullis): all cells block.
			if !obj.Solid && cell.Symbol != 'X' {
				continue
			}
			px, py := ox+cell.X, oy+cell.Y
			if px >= screenX && px < screenX+width && py >= screenY && py < screenY+height {
				return true
			}
		}
	}
	return false
}

func DrawHelp() {
	termW, termH := Screen.Size()
	bg := tcell.NewRGBColor(0xcd, 0xcd, 0xcd)
	gold := tcell.NewRGBColor(0xFF, 0xD8, 0x4C)
	head := tcell.StyleDefault.Background(bg).Foreground(tcell.ColorBlack).Bold(true)
	body := tcell.StyleDefault.Background(bg).Foreground(tcell.ColorBlack)
	hilite := tcell.StyleDefault.Background(bg).Foreground(tcell.NewRGBColor(0x40, 0x40, 0xC0))
	cur := tcell.StyleDefault.Background(gold).Foreground(tcell.ColorBlack)

	lx := termW/2 - 22
	cy := termH/2 - 11

	put := func(y int, s string, st tcell.Style) {
		emitStr(Screen, lx, y, st, s)
	}

	diffStr := func(on bool) string {
		if on {
			return "A"
		}
		return "B"
	}

	put(cy+0, "An  A D V E N T U R E  —  help", head)
	put(cy+1, "", body)
	put(cy+2, "Objective", hilite)
	put(cy+3, "  Find the enchanted chalice and return it to the golden castle.", body)
	put(cy+4, "", body)
	put(cy+5, "Controls", hilite)
	put(cy+6, "  [W] / [↑]   [A] / [←]   [S] / [↓]   [D] / [→]    move player", body)
	put(cy+7, "  [Space]                            drop carried object", body)
	put(cy+8, "  [H]                                toggle this help screen", body)
	put(cy+9, "  [V]                                select game variation  (1 / 2 / 3)", body)
	put(cy+10, "  [N]                                select difficulty  (A / B)", body)
	put(cy+11, "  [R]                                reset game", body)
	put(cy+12, "  [Q]                                quit", body)
	put(cy+13, "", body)
	put(cy+14, "Game variations  [V]", hilite)
	put(cy+15, "  1  easy    1 castle, 1 slow dragon, chalice in the open,", body)
	put(cy+16, "             bat and bridge not in play", body)
	put(cy+17, "  2  normal  3 castles (golden/white/black), 3 dragons, full maze,", body)
	put(cy+18, "             all objects and keys in play", body)
	put(cy+19, "  3  hard    like variation 2, but dragons are faster", body)
	put(cy+20, "             and significantly more aggressive", body)
	put(cy+21, "", body)
	put(cy+22, "Difficulty  [N]  (A = harder / B = easier)", hilite)
	put(cy+23, "  A  dragons flee from sword, shorter roar window after touch", body)
	put(cy+24, "  B  dragons ignore the sword, longer roar window", body)
	put(cy+25, "", body)
	emitStr(Screen, lx, cy+26, cur, fmt.Sprintf("  Variation %d   Difficulty: %s  ",
		game.G.GameType, diffStr(game.DifficultyLeft)))
	put(cy+27, "", body)
	put(cy+28, "  Press [H] to close", hilite)
}

func DrawWinOverlay() {
	if !game.GameWon || game.WinOverlayTimer <= 0 {
		return
	}
	termW, termH := Screen.Size()
	fc := game.GetFlashColor()
	bg := tcell.StyleDefault.Background(fc).Foreground(fc)
	text := tcell.StyleDefault.Background(fc).Foreground(tcell.ColorBlack)

	msg := "  YOU WON!  "
	sub := "  Press [R] to play again  "
	boxW := len(sub) + 4
	boxH := 5
	bx := termW/2 - boxW/2
	by := termH/2 - boxH/2

	for y := by; y < by+boxH; y++ {
		for x := bx; x < bx+boxW; x++ {
			if x >= 0 && x < termW && y >= 0 && y < termH {
				Screen.SetContent(x, y, ' ', nil, bg)
			}
		}
	}
	emitStr(Screen, bx+(boxW-len(msg))/2, by+1, text, msg)
	emitStr(Screen, bx+(boxW-len(sub))/2, by+3, text, sub)
}

func DrawConfirm() {
	termW, termH := Screen.Size()
	bg := tcell.NewRGBColor(0xaa, 0xaa, 0xaa)
	blank := tcell.StyleDefault.Background(bg).Foreground(bg)
	body := tcell.StyleDefault.Background(bg).Foreground(tcell.ColorBlack)
	yes := tcell.StyleDefault.Background(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Foreground(tcell.ColorBlack)

	var line1, line2 string
	if game.ConfirmAction == "quit" {
		line1 = "  Quit the game?"
		line2 = "  Are you sure?  [Y] yes   any other key: no"
	} else {
		line1 = "  Reset the game?"
		line2 = "  Are you sure?  [Y] yes   any other key: no"
	}
	boxW := len(line2) + 8
	boxH := 6
	bx := termW/2 - boxW/2
	by := termH/2 - boxH/2

	for y := by; y < by+boxH; y++ {
		for x := bx; x < bx+boxW; x++ {
			if x >= 0 && x < termW && y >= 0 && y < termH {
				Screen.SetContent(x, y, ' ', nil, blank)
			}
		}
	}
	emitStr(Screen, bx+2, by+1, body, line1)
	emitStr(Screen, bx+2, by+3, body, "  Are you sure?  ")
	emitStr(Screen, bx+2+17, by+3, yes, "[Y] yes")
	emitStr(Screen, bx+2+25, by+3, body, "   any other key: no")
}

func ResetGame() {
	w, h := Screen.Size()
	game.SoftReset(w, h)
	FillTheScreen()
}

// DrawSelOverlay renders the selection overlay for variation or difficulty cycling.
// DrawEaten renders the "eaten by dragon" overlay. Player can only reset from here.
func DrawEaten() {
	if !game.Eaten {
		return
	}
	termW, termH := Screen.Size()
	bg := tcell.StyleDefault.Background(tcell.NewRGBColor(0x66, 0x00, 0x00)).Foreground(tcell.NewRGBColor(0xFF, 0xAA, 0xAA))
	hi := tcell.StyleDefault.Background(tcell.NewRGBColor(0x66, 0x00, 0x00)).Foreground(tcell.NewRGBColor(0xFF, 0xFF, 0xFF))

	lines := []string{
		"  You have been eaten by a dragon!  ",
		"",
		"  Press [R] to reset the game.  ",
	}
	boxW := 40
	boxH := len(lines) + 4
	bx := termW/2 - boxW/2
	by := termH/2 - boxH/2

	for y := by; y < by+boxH; y++ {
		for x := bx; x < bx+boxW; x++ {
			if x >= 0 && x < termW && y >= 0 && y < termH {
				Screen.SetContent(x, y, ' ', nil, bg)
			}
		}
	}
	for i, line := range lines {
		st := bg
		if i == 0 {
			st = hi
		}
		emitStr(Screen, bx+(boxW-len(line))/2, by+2+i, st, line)
	}
}

func DrawSelOverlay() {
	o := game.Overlay
	if !o.Active {
		return
	}
	termW, termH := Screen.Size()

	bg := tcell.StyleDefault.Background(tcell.NewRGBColor(0x33, 0x33, 0x33)).Foreground(tcell.NewRGBColor(0xcc, 0xcc, 0xcc))
	sel := tcell.StyleDefault.Background(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Foreground(tcell.ColorBlack)

	var title string
	var labels []string
	var cur int

	if o.Kind == "variation" {
		title = " Game Variation "
		labels = []string{" 1 ", " 2 ", " 3 "}
		cur = o.Value - 1 // 1-indexed → 0-indexed
	} else {
		title = "  Difficulty  "
		labels = []string{" A ", " B "}
		cur = o.Value
	}

	itemsRow := ""
	for i, l := range labels {
		if i > 0 {
			itemsRow += "  "
		}
		itemsRow += l
	}

	boxW := len(title) + 4
	if len(itemsRow)+4 > boxW {
		boxW = len(itemsRow) + 4
	}
	boxH := 6 // +1 row for progress bar
	bx := termW/2 - boxW/2
	by := termH/2 - boxH/2

	// Fill background.
	for y := by; y < by+boxH; y++ {
		for x := bx; x < bx+boxW; x++ {
			if x >= 0 && x < termW && y >= 0 && y < termH {
				Screen.SetContent(x, y, ' ', nil, bg)
			}
		}
	}

	// Title row.
	emitStr(Screen, bx+(boxW-len(title))/2, by+1, bg, title)

	// Items row: draw each item individually to apply selection highlight.
	ix := bx + (boxW-len(itemsRow))/2
	col := ix
	for i, l := range labels {
		st := bg
		if i == cur {
			st = sel
		}
		emitStr(Screen, col, by+3, st, l)
		col += len(l)
		if i < len(labels)-1 {
			emitStr(Screen, col, by+3, bg, "  ")
			col += 2
		}
	}

	// Progress bar — smooth 8-level half-block countdown.
	// Full at overlay open, empties as timer runs out.
	barW := boxW - 2
	barX := bx + 1
	barY := by + 4
	filled := tcell.StyleDefault.Background(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Foreground(tcell.NewRGBColor(0x33, 0x33, 0x33))
	empty := bg

	ratio := 1.0
	if o.MaxTicks > 0 {
		ratio = float64(o.Ticks) / float64(o.MaxTicks)
	}
	eighths := int(ratio * float64(barW*8))
	fullCells := eighths / 8
	partial := eighths % 8
	partialChars := []rune{' ', '▏', '▎', '▍', '▌', '▋', '▊', '▉'}

	for i := 0; i < barW; i++ {
		x := barX + i
		if i < fullCells {
			Screen.SetContent(x, barY, '█', nil, filled)
		} else if i == fullCells && partial > 0 {
			Screen.SetContent(x, barY, partialChars[partial], nil, filled)
		} else {
			Screen.SetContent(x, barY, ' ', nil, empty)
		}
	}
}

func GetWidth() int {
	screenWidth, _ := Screen.Size()
	return screenWidth
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

// DrawSpecialRooms renders UI overlays for rooms with special content.
// Must be called after DrawAllVisibleObjects, only when not in HelpMode.
func DrawSpecialRooms() {
	switch game.CurrentRoom {
	case &world.RoomNumberRoom:
		drawNameRoom()
	case &world.RoomSplashScreen:
		drawEasterEggRoom()
	}
}

func drawNameRoom() {
	termW, termH := Screen.Size()
	bg := game.CurrentRoom.Background

	gold := tcell.NewRGBColor(0xFF, 0xD8, 0x4C)
	goBlue := tcell.NewRGBColor(0x00, 0xAD, 0xD8)
	white := tcell.NewRGBColor(0xFF, 0xFF, 0xFF)
	orange := tcell.NewRGBColor(0xFF, 0x80, 0x00)
	dim := tcell.NewRGBColor(0x88, 0x88, 0x88)
	boxBg := tcell.NewRGBColor(0x18, 0x00, 0x30)

	inner := []struct {
		text  string
		color tcell.Color
	}{
		{"", gold},
		{"   An  A D V E N T U R E   ", gold},
		{"", gold},
		{"  ═══════════════════════  ", orange},
		{"", gold},
		{"   ported to Go", goBlue},
		{"   by Jens Schendel", white},
		{"", gold},
		{"   dedicated to", dim},
		{"   Warren Robinett", orange},
		{"", gold},
	}

	boxW := 0
	for _, l := range inner {
		if len(l.text) > boxW {
			boxW = len(l.text)
		}
	}
	boxW += 4
	boxH := len(inner) + 4

	bx := termW/2 - boxW/2
	by := termH/2 - boxH/2

	// Fill box background
	fill := tcell.StyleDefault.Background(boxBg).Foreground(boxBg)
	for y := by; y < by+boxH; y++ {
		for x := bx; x < bx+boxW; x++ {
			if x >= 0 && x < termW && y >= 0 && y < termH {
				Screen.SetContent(x, y, ' ', nil, fill)
			}
		}
	}

	// Box border (gold, box-drawing chars)
	bord := tcell.StyleDefault.Background(boxBg).Foreground(gold)
	emitStr(Screen, bx, by, bord, "╔"+repeatStr("═", boxW-2)+"╗")
	emitStr(Screen, bx, by+boxH-1, bord, "╚"+repeatStr("═", boxW-2)+"╝")
	for y := by + 1; y < by+boxH-1; y++ {
		Screen.SetContent(bx, y, '║', nil, bord)
		Screen.SetContent(bx+boxW-1, y, '║', nil, bord)
	}

	// Content lines
	for i, l := range inner {
		st := tcell.StyleDefault.Background(boxBg).Foreground(l.color)
		emitStr(Screen, bx+2, by+2+i, st, l.text)
	}

	_ = bg // room bg visible through castle graphic behind the box
}

// adventureTitleGlyphs holds the 7×5 half-block glyphs for "An ADVENTURE".
var adventureTitleGlyphs = map[rune][5]string{
	'A': {"  ███  ", " █   █ ", "███████", "█     █", "█     █"},
	'n': {"       ", "       ", " ████▄ ", " █   █ ", " █   █ "},
	'D': {"█████▄ ", "█     █", "█     █", "█     █", "█████▀ "},
	'V': {"█     █", "█     █", " █   █ ", "  █ █  ", "   █   "},
	'E': {"███████", "█      ", "█████  ", "█      ", "███████"},
	'N': {"█▄    █", "█ █   █", "█  █  █", "█   █ █", "█    ▀█"},
	'T': {"███████", "   █   ", "   █   ", "   █   ", "   █   "},
	'U': {"█     █", "█     █", "█     █", "█     █", " █████ "},
	'R': {"██████ ", "█     █", "██████ ", "█  █   ", "█   ██ "},
	' ': {"       ", "       ", "       ", "       ", "       "},
}

// drawAdventureTitle renders the "An ADVENTURE" half-block title centered at centerY.
// Returns the Y coordinate of the row immediately below the title (centerY + 5).
func drawAdventureTitle(termW, termH int, bg tcell.Color, centerY int) int {
	gold := tcell.NewRGBColor(0xFF, 0xD8, 0x4C)
	dim := tcell.NewRGBColor(0xAA, 0xAA, 0xAA)
	rainbow := []tcell.Color{
		tcell.NewRGBColor(0xFF, 0x44, 0x44),
		tcell.NewRGBColor(0xFF, 0x88, 0x00),
		tcell.NewRGBColor(0xFF, 0xD8, 0x4C),
		tcell.NewRGBColor(0x88, 0xFF, 0x44),
		tcell.NewRGBColor(0x00, 0xFF, 0xAA),
		tcell.NewRGBColor(0x44, 0xCC, 0xFF),
		tcell.NewRGBColor(0xAA, 0x44, 0xFF),
		tcell.NewRGBColor(0xFF, 0x44, 0xCC),
		tcell.NewRGBColor(0xFF, 0x44, 0x44),
	}
	title := []rune("An ADVENTURE")
	titleColors := []tcell.Color{
		gold, gold, dim,
		rainbow[0], rainbow[1], rainbow[2], rainbow[3],
		rainbow[4], rainbow[5], rainbow[6], rainbow[7], rainbow[8],
	}
	const letterW = 7
	const gap = 1
	titleW := len(title)*letterW + (len(title)-1)*gap
	startX := termW/2 - titleW/2
	for i, r := range title {
		g, ok := adventureTitleGlyphs[r]
		if !ok {
			continue
		}
		cx := startX + i*(letterW+gap)
		st := tcell.StyleDefault.Background(bg).Foreground(titleColors[i])
		for row := 0; row < 5; row++ {
			for col, ch := range []rune(g[row]) {
				sx, sy := cx+col, centerY+row
				if sx >= 0 && sx < termW && sy >= 0 && sy < termH {
					Screen.SetContent(sx, sy, ch, nil, st)
				}
			}
		}
	}
	return centerY + 5
}

// splashTitleGlyphs holds 8×8 block-pixel glyphs for the startup splash screen title.
var splashTitleGlyphs = map[rune][8]string{
	'A': {
		"   ██   ",
		"  █  █  ",
		" █    █ ",
		"█      █",
		"████████",
		"█      █",
		"█      █",
		"█      █",
	},
	'n': {
		"        ",
		"        ",
		"        ",
		" ██████ ",
		"██     █",
		"█      █",
		"█      █",
		"█      █",
	},
	'D': {
		"███████ ",
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		"███████ ",
	},
	'V': {
		"█      █",
		"█      █",
		" █    █ ",
		" █    █ ",
		"  █  █  ",
		"  █  █  ",
		"   ██   ",
		"   ██   ",
	},
	'E': {
		"████████",
		"█       ",
		"█       ",
		"██████  ",
		"█       ",
		"█       ",
		"█       ",
		"████████",
	},
	'N': {
		"█      █",
		"██     █",
		"█ █    █",
		"█  █   █",
		"█   █  █",
		"█    █ █",
		"█     ██",
		"█      █",
	},
	'T': {
		"████████",
		"   █    ",
		"   █    ",
		"   █    ",
		"   █    ",
		"   █    ",
		"   █    ",
		"   █    ",
	},
	'U': {
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		"█      █",
		" ██████ ",
	},
	'R': {
		"███████ ",
		"█      █",
		"█      █",
		"███████ ",
		"█   █   ",
		"█    █  ",
		"█     █ ",
		"█      █",
	},
	' ': {
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
	},
}

// drawSplashTitle renders the "An ADVENTURE" title using 8×8 glyphs, centered at centerY.
// Returns the Y coordinate immediately below the title (centerY + 8).
func drawSplashTitle(termW, termH int, bg tcell.Color, centerY int) int {
	gold := tcell.NewRGBColor(0xFF, 0xD8, 0x4C)
	dim := tcell.NewRGBColor(0xAA, 0xAA, 0xAA)
	rainbow := []tcell.Color{
		tcell.NewRGBColor(0xFF, 0x44, 0x44),
		tcell.NewRGBColor(0xFF, 0x88, 0x00),
		tcell.NewRGBColor(0xFF, 0xD8, 0x4C),
		tcell.NewRGBColor(0x88, 0xFF, 0x44),
		tcell.NewRGBColor(0x00, 0xFF, 0xAA),
		tcell.NewRGBColor(0x44, 0xCC, 0xFF),
		tcell.NewRGBColor(0xAA, 0x44, 0xFF),
		tcell.NewRGBColor(0xFF, 0x44, 0xCC),
		tcell.NewRGBColor(0xFF, 0x44, 0x44),
	}
	title := []rune("An ADVENTURE")
	titleColors := []tcell.Color{
		gold, gold, dim,
		rainbow[0], rainbow[1], rainbow[2], rainbow[3],
		rainbow[4], rainbow[5], rainbow[6], rainbow[7], rainbow[8],
	}
	const letterW = 8
	const gap = 1
	titleW := len(title)*letterW + (len(title)-1)*gap
	startX := termW/2 - titleW/2
	for i, r := range title {
		g, ok := splashTitleGlyphs[r]
		if !ok {
			continue
		}
		cx := startX + i*(letterW+gap)
		st := tcell.StyleDefault.Background(bg).Foreground(titleColors[i])
		for row := 0; row < 8; row++ {
			for col, ch := range []rune(g[row]) {
				sx, sy := cx+col, centerY+row
				if sx >= 0 && sx < termW && sy >= 0 && sy < termH {
					Screen.SetContent(sx, sy, ch, nil, st)
				}
			}
		}
	}
	return centerY + 8
}

func drawSplashOnScreen(termW, termH int, bg tcell.Color) {
	gold := tcell.NewRGBColor(0xFF, 0xD8, 0x4C)
	dim := tcell.NewRGBColor(0xAA, 0xAA, 0xAA)

	titleY := termH/2 - 4 // center 9-row title: rows termH/2-4 … termH/2+4
	below := drawSplashTitle(termW, termH, bg, titleY)

	subtitle := "ported to Go by Jens Schendel, dedicated to Warren Robinett"
	subY := below + 1
	emitStr(Screen, termW/2-len(subtitle)/2, subY,
		tcell.StyleDefault.Background(bg).Foreground(dim), subtitle)

	hintY := subY + 1
	hintParts := []struct {
		text  string
		color tcell.Color
	}{
		{"Press ", dim},
		{"[H]", gold},
		{" for help", dim},
	}
	hintLen := 0
	for _, p := range hintParts {
		hintLen += len(p.text)
	}
	hx := termW/2 - hintLen/2
	for _, p := range hintParts {
		emitStr(Screen, hx, hintY, tcell.StyleDefault.Background(bg).Foreground(p.color), p.text)
		hx += len(p.text)
	}
}

// drawEasterEggRoom renders the Easter Egg room (C++ 0x1E):
// "An ADVENTURE" in the splash-screen pixel font, then credit text in cycling flash color.
func drawEasterEggRoom() {
	termW, termH := Screen.Size()
	bg := game.CurrentRoom.Background

	// Advance the local slow hue: step 1 every 4 frames → ~6× slower than chalice.
	eeFlashTick++
	if eeFlashTick >= 4 {
		eeFlashTick = 0
		eeFlashHue++
		if eeFlashHue >= 360 {
			eeFlashHue = 0
		}
	}
	// Compute color from eeFlashHue (same formula as GetFlashColor but own counter).
	h := float64(eeFlashHue) / (360.0 / 3)
	var cr, cg, cb float64
	if h < 1 {
		cr = h * 255; cg = 0; cb = (1 - h) * 255
	} else if h < 2 {
		h -= 1; cr = (1 - h) * 255; cg = h * 255; cb = 0
	} else {
		h -= 2; cr = 0; cg = (1 - h) * 255; cb = h * 255
	}
	fc := tcell.NewRGBColor(int32(cr), int32(cg), int32(cb))

	// Title: 12 glyphs × 7 + 11 gaps = 95 cols wide.
	// Text wrapped to 70 cols yields ~4 lines.
	fullText := "ported with \u2665 to Go by Jens Schendel with intense use of the awesome tcell package provided by Garrett D\u2019Amore, and with heartfelt thanks for countless hours spent in front of my Atari 2600, dedicated to Warren Robinett. The ADVENTURE goes on and on!"
	lines := wordWrap(fullText, 70)

	totalRows := 5 + 1 + len(lines) // title height + gap + text lines
	startY := termH/2 - totalRows/2
	below := drawAdventureTitle(termW, termH, bg, startY)

	st := tcell.StyleDefault.Background(bg).Foreground(fc)
	for i, l := range lines {
		lw := len([]rune(l))
		emitStr(Screen, termW/2-lw/2, below+1+i, st, l)
	}
}

// wordWrap breaks s into lines of at most maxW runes, splitting at word boundaries.
func wordWrap(s string, maxW int) []string {
	words := []string{}
	cur := ""
	for _, r := range s {
		if r == ' ' {
			if cur != "" {
				words = append(words, cur)
				cur = ""
			}
		} else {
			cur += string(r)
		}
	}
	if cur != "" {
		words = append(words, cur)
	}

	var lines []string
	line := ""
	for _, w := range words {
		if line == "" {
			line = w
		} else if len([]rune(line))+1+len([]rune(w)) <= maxW {
			line += " " + w
		} else {
			lines = append(lines, line)
			line = w
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

func repeatStr(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
