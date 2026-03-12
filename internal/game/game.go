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
var YellowKey *Object
var GreenDragon *Object
var AllObjects []*Object
var CurrentRoom *world.Room

func InitYellowKey(w, h int) {
	YellowKey = &Object{
		RelX:   1.0 / 5.0,
		RelY:   0.5,
		Width:  8,
		Height: 2,
		StepX:  0,
		StepY:  0,
		Style:  tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.KeyGfx,
	}
	AllObjects = append(AllObjects, YellowKey)
}

func InitGreenDragon(w, h int) {
	GreenDragon = &Object{
		RelX:   4.0 / 5.0,
		RelY:   0.5,
		Width:  8,
		Height: 10,
		StepX:  0,
		StepY:  0,
		Style:  tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x86, 0xd9, 0x22)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.DragonGfx,
	}
	AllObjects = append(AllObjects, GreenDragon)
}

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
