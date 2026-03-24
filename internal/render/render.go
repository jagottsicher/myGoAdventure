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
	for _, point := range CurrentScreen {
		var style tcell.Style
		if point.Symbol == 'X' || point.Symbol == 'x' {
			wallColor := game.CurrentRoom.Foreground
			if darkMaze && game.Player != nil {
				dx := point.X - playerCX
				dy := point.Y - playerCY
				// aura: 33 wide (±16), 16 high (±8)
				if dx >= -16 && dx <= 16 && dy >= -8 && dy <= 8 {
					wallColor = auraColor
				}
			}
			style = tcell.StyleDefault.Background(wallColor).Foreground(wallColor)
		} else {
			style = tcell.StyleDefault.
				Background(game.CurrentRoom.Background).
				Foreground(game.CurrentRoom.Foreground)
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
			if obj.ZLayer != layer {
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
			style = tcell.StyleDefault.Foreground(objFg).Background(existingBg)
		}
		Screen.SetContent(px, py, point.Symbol, nil, style)
	}
}

func InitGamestate() {
	world.InitDirections()
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

	// Passage barriers — black on black, only visible in their room.
	blackOnBlack := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	game.InitBarrier(&world.RoomTopAccessRight, 1.0/20.0, h, blackOnBlack)    // left of yellow castle: block left side
	game.InitBarrier(&world.RoomCorridorRight, 19.0/20.0, h, blackOnBlack)   // right of yellow castle: block right side
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

	Screen.SetStyle(tcell.StyleDefault.
		Background(world.RoomSplashScreen.Background).
		Foreground(world.RoomSplashScreen.Foreground))

	for i := 0; i < 256; i++ {
		Screen.SetStyle(tcell.StyleDefault.
			Background(world.RoomSplashScreen.Background).
			Foreground(tcell.NewRGBColor(int32(i), int32(i), int32(i))))

		screenWidth, screenHeight := Screen.Size()
		title := "An Adventure - dedicated to Warren Robinett"
		subline := "Press [Q] to quit  [H] for help"
		emitStr(Screen, screenWidth/2-len(title)/2, screenHeight/2-1, tcell.StyleDefault, title)
		emitStr(Screen, screenWidth/2-len(subline)/2, screenHeight/2, tcell.StyleDefault, subline)

		time.Sleep(time.Millisecond * 10)
		Screen.Show()
	}
	time.Sleep(time.Second * 3)
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
	termW, termH := Screen.Size()
	// Check room walls via screen buffer.
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			px, py := screenX+dx, screenY+dy
			if px < 0 || px >= termW || py < 0 || py >= termH {
				continue
			}
			r, _, _, _ := Screen.GetContent(px, py)
			if r == 'X' {
				return true
			}
		}
	}
	// Check barrier and solid objects directly (not via screen buffer).
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
	blank := tcell.StyleDefault.Background(bg).Foreground(bg)
	head := tcell.StyleDefault.Background(bg).Foreground(tcell.ColorBlack)
	body := tcell.StyleDefault.Background(bg).Foreground(tcell.ColorBlack)
	hilite := tcell.StyleDefault.Background(bg).Foreground(tcell.NewRGBColor(0x60, 0x60, 0x60))
	cur := tcell.StyleDefault.Background(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Foreground(tcell.ColorBlack)

	lx := termW/2 - 22
	cy := termH/2 - 10
	boxW := 72 // content max ~68 + 2 padding each side
	boxH := 27 // content 23 lines + 2 padding each side

	// Draw gray box (+2 padding around text).
	for y := cy - 2; y < cy-2+boxH; y++ {
		for x := lx - 2; x < lx-2+boxW; x++ {
			if x >= 0 && x < termW && y >= 0 && y < termH {
				Screen.SetContent(x, y, ' ', nil, blank)
			}
		}
	}

	put := func(y int, s string, st tcell.Style) {
		emitStr(Screen, lx, y, st, s)
	}

	put(cy+0, "An  A D V E N T U R E  —  help", head)
	put(cy+1, "", body)
	put(cy+2, "Objective", hilite)
	put(cy+3, "  Find the enchanted chalice and return it to the golden castle.", body)
	put(cy+4, "", body)
	put(cy+5, "Controls", hilite)
	put(cy+6, "  [W] / [↑]   [A] / [←]   [S] / [↓]   [D] / [→]    move player", body)
	put(cy+7, "  [H]                                toggle this help screen", body)
	put(cy+8, "  [V]                                cycle game variation", body)
	put(cy+9, "  [R]                                reset game", body)
	put(cy+10, "  [Q]                                quit", body)
	put(cy+11, "", body)
	put(cy+12, "Game variations", hilite)
	put(cy+13, "  1  easy    1 castle, 1 slow dragon, chalice in the open,", body)
	put(cy+14, "             bat and bridge not in play", body)
	put(cy+15, "  2  normal  3 castles (golden/white/black), 3 dragons, full maze,", body)
	put(cy+16, "             all objects and keys in play", body)
	put(cy+17, "  3  hard    like variation 2, but dragons are faster", body)
	put(cy+18, "             and significantly more aggressive", body)
	put(cy+19, "", body)
	emitStr(Screen, lx, cy+20, cur, fmt.Sprintf("  Current: Variation %d  ", game.G.GameType))
	put(cy+21, "", body)
	put(cy+22, "  Press [H] to close", hilite)
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
	game.HelpMode = false
	game.GodMode = false
	game.CarriedObject = nil
	game.ResetObjects()
	w, h := Screen.Size()
	game.CurrentRoom = &world.RoomYellowCastle
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
	blackOnBlack := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	game.InitBarrier(&world.RoomTopAccessRight, 1.0/20.0, h, blackOnBlack)
	game.InitBarrier(&world.RoomCorridorRight, 19.0/20.0, h, blackOnBlack)
	game.InitBarrier(&world.RoomSideCorridorOlive, 1.0/20.0, h, blackOnBlack)
	game.InitBarrier(&world.RoomSideCorridorCyan, 19.0/20.0, h, blackOnBlack)
	FillTheScreen()
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
	}
}

// Half-block digit designs (6 chars wide, 4 rows tall).
// Each row encodes two pixel rows via ▀ ▄ █ and space.
var digitRows = map[int][4]string{
	1: {" ▄██  ", "  ██  ", "  ██  ", " ▀▀▀▀▀"},
	2: {"▄▀▀▀▀▄", " ▄▄▄▄▀", "▄▀    ", "▀▀▀▀▀▀"},
	3: {"▄▀▀▀▀▄", " ▄▄▄▄▀", "▄    █", " ▀▀▀▀ "},
}

var (
	colorEasy   = tcell.NewRGBColor(0xFF, 0xD8, 0x4C) // gold
	colorNormal = tcell.NewRGBColor(0x88, 0xCC, 0xFF) // sky blue
	colorHard   = tcell.NewRGBColor(0xFF, 0x55, 0x44) // red

	dimEasy   = tcell.NewRGBColor(0x66, 0x55, 0x1C)
	dimNormal = tcell.NewRGBColor(0x33, 0x55, 0x77)
	dimHard   = tcell.NewRGBColor(0x66, 0x22, 0x1C)
)

func drawNumberRoom() {
	termW, termH := Screen.Size()
	bg := game.CurrentRoom.Background

	cols := []struct {
		n     int
		label string
		on    tcell.Color
		off   tcell.Color
	}{
		{1, "easy", colorEasy, dimEasy},
		{2, "normal", colorNormal, dimNormal},
		{3, "hard", colorHard, dimHard},
	}

	const digitW = 6
	const gap = 10
	totalW := 3*digitW + 2*gap
	startX := termW/2 - totalW/2
	startY := termH/2 - 5

	title := "select  variation"
	ts := tcell.StyleDefault.Background(bg).Foreground(tcell.NewRGBColor(0x88, 0x88, 0x88))
	emitStr(Screen, termW/2-len(title)/2, startY-2, ts, title)

	for i, col := range cols {
		x := startX + i*(digitW+gap)
		selected := uint8(i+1) == game.G.GameType
		fg := col.off
		if selected {
			fg = col.on
		}
		ds := tcell.StyleDefault.Background(bg).Foreground(fg)
		for row, line := range digitRows[col.n] {
			emitStr(Screen, x, startY+row, ds, line)
		}
		// label row
		ls := tcell.StyleDefault.Background(bg).Foreground(fg)
		label := col.label
		if selected {
			label = "▶ " + label
		} else {
			label = "  " + label
		}
		emitStr(Screen, x-1, startY+5, ls, label)
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

func repeatStr(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
