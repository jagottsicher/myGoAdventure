package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

type playBall struct {
	width  int
	height int
	pos_x  int
	pos_y  int
	style  tcell.Style
}

func (player *playBall) Init() {
	player.width = 4
	player.height = 2
	player.pos_x = 80
	player.pos_y = 40
	// player.style = tcell.StyleDefault.Foreground(tcell.ColorDarkBlue.TrueColor()).Background(tcell.ColorYellow)
}

func (player playBall) Display(s tcell.Screen) {
	// s.Clear()
	emitStr(s, player.pos_x, player.pos_y, player.style, "PPPP")
	emitStr(s, player.pos_x, player.pos_y+1, player.style, "PPPP")
	emitStr(s, 80, 24, player.style, fmt.Sprintf("%d/%d", player.pos_x, player.pos_y))
	s.Show()
}

func (player *playBall) Movement(s tcell.Screen, deltaX, deltaY int) {
	w, h := s.Size()
	if player.pos_x+deltaX >= w {
		player.pos_x -= w
	}
	if player.pos_y+deltaY >= h {
		player.pos_y -= h
	}
	if player.pos_x+deltaX < 0 {
		player.pos_x += w
	}
	if player.pos_y+deltaY < 0 {
		player.pos_y += h
	}
	player.pos_x += deltaX
	player.pos_y += deltaY
}

var player playBall

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

// func displayPlayer(s tcell.Screen) {
// 	_, h := s.Size()

// 	// if positionX > 0 {
// 	// 	positionX = (w/2 - 2) + playerX
// 	// } else {
// 	// 	positionX = w - playerX
// 	// }

// 	positionY := (h / 2) + playerY

// 	s.Clear()
// 	style := tcell.StyleDefault.Foreground(tcell.ColorDarkBlue.TrueColor()).Background(tcell.ColorYellow)

// 	emitStr(s, positionX, positionY, style, "PPPP")
// 	emitStr(s, positionX, positionY+1, style, fmt.Sprintf("%d", positionX))
// 	// emitStr(s, w/2-9, h/2+1, tcell.StyleDefault, "Press ESC to exit.")
// 	s.Show()
//}
