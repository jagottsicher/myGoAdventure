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

func DrawStage() {
	if game.CurrentRoom == nil {
		return
	}
	for _, point := range CurrentScreen {
		var style tcell.Style
		if point.Symbol == 'X' || point.Symbol == 'x' {
			style = tcell.StyleDefault.
				Background(game.CurrentRoom.Foreground).
				Foreground(game.CurrentRoom.Foreground)
		} else {
			style = tcell.StyleDefault.
				Background(game.CurrentRoom.Background).
				Foreground(game.CurrentRoom.Foreground)
		}
		Screen.SetContent(point.X, point.Y, point.Symbol, nil, style)
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
	game.InitPortcullis(w, h)
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
		subline := "Press [Q] to exit."
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
	// Check barrier objects directly (not via screen buffer).
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
			if cell.Symbol != 'X' {
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
