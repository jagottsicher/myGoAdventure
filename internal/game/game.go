package game

import (
	"time"

	"development/myGoAdventure/internal/world"
	"github.com/gdamore/tcell/v2"
)

type Game struct {
	GameType uint8
	FPS      time.Duration
}

var G = Game{
	GameType: 2,
	FPS:      60,
}

type Object struct {
	RelX   float64
	RelY   float64
	StepX  int
	StepY  int
	Width  int
	Height int
	Style  tcell.Style
	Shape  []*world.Cell
}

var Player *Object
var AllObjects []*Object
var CurrentRoom *world.Room

func InitPlayer(w, h int) {
	Player = &Object{
		RelX:   float64(w/2) / float64(w),
		RelY:   float64(h/3*2) / float64(h),
		Width:  2,
		Height: 1,
		StepX:  2,
		StepY:  1,
		Style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple),
		Shape:  world.PlayerGfx,
	}
	AllObjects = append(AllObjects, Player)
}
