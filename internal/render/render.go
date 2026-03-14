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
	for _, obj := range game.AllObjects {
		DrawObject(obj)
	}
}

func DrawObject(obj *game.Object) {
	termW, termH := Screen.Size()
	screenX := int(obj.RelX * float64(termW))
	screenY := int(obj.RelY * float64(termH))
	objFg, _, _ := obj.Style.Decompose()
	for _, point := range obj.Shape {
		px := screenX + point.X
		py := screenY + point.Y
		if px < 0 || px >= termW || py < 0 || py >= termH {
			continue
		}
		style := obj.Style
		if point.Symbol == '▀' || point.Symbol == '▄' {
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
	game.InitGreenDragon(w, h)
	game.InitBat(w, h)
	game.InitBridge(w, h)
	game.InitSword(w, h)
	game.InitChalice(w, h)
	game.InitMagnet(w, h)
	game.InitDot(w, h)
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
