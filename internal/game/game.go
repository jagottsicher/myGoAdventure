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
	RelX         float64
	RelY         float64
	StepX        int
	StepY        int
	Width        int
	Height       int
	Style        tcell.Style
	Shape        []*world.Cell
	Frames       [][]*world.Cell // animation frames; nil = static
	AnimInterval int             // game ticks per frame
	animTick     int
	animFrame    int
}

// Animate advances the animation by one tick. Call once per game update.
func (o *Object) Animate() {
	if len(o.Frames) < 2 {
		return
	}
	o.animTick++
	if o.animTick >= o.AnimInterval {
		o.animTick = 0
		o.animFrame = (o.animFrame + 1) % len(o.Frames)
		o.Shape = o.Frames[o.animFrame]
	}
}

var Player *Object
var YellowKey *Object
var GreenDragon *Object
var Bat *Object
var Portcullis *Object
var Bridge *Object
var Sword *Object
var Chalice *Object
var Magnet *Object
var Dot *Object
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
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	GreenDragon = &Object{
		RelX:         4.0 / 5.0,
		RelY:         0.5,
		Width:        8,
		Height:       10,
		StepX:        0,
		StepY:        0,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x86, 0xd9, 0x22)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30, // ~0.5s at 60 FPS
	}
	AllObjects = append(AllObjects, GreenDragon)
}

func InitBat(w, h int) {
	frames := [][]*world.Cell{world.BatGfx, world.BatGfxOpen}
	Bat = &Object{
		RelX:         1.0 / 5.0,
		RelY:         4.0 / 5.0,
		Width:        8,
		Height:       6,
		StepX:        0,
		StepY:        0,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 20, // ~0.33s at 60 FPS
	}
	AllObjects = append(AllObjects, Bat)
}

func InitPortcullis(w, h int) {
	// Castle template: 40 chars wide, 12 rows tall.
	// Gate opening: cols 18–21 (4 chars) in template row 5.
	doorStartCol := 18 * w / 40
	doorWidth := (4*w + 39) / 40 // ceiling division — never one short
	if doorWidth < 2 {
		doorWidth = 2
	}
	// Height = 1 template row scaled to terminal rows (min 2)
	portHeight := h / 12
	if portHeight < 2 {
		portHeight = 2
	}
	frames := world.MakePortcullisFrames(doorWidth, portHeight)
	// RelY: template row 5 of 12 maps to ~5/12 of terminal height
	Portcullis = &Object{
		RelX:         float64(doorStartCol) / float64(w),
		RelY:         5.0 / 12.0,
		Width:        doorWidth,
		Height:       portHeight,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 45,
	}
	AllObjects = append(AllObjects, Portcullis)
}

func InitBridge(w, h int) {
	Bridge = &Object{
		RelX: 0.1, RelY: 0.15, Width: 8, Height: 12,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x99, 0x00, 0xCC)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.BridgeGfx,
	}
	AllObjects = append(AllObjects, Bridge)
}

func InitSword(w, h int) {
	Sword = &Object{
		RelX: 0.65, RelY: 0.25, Width: 8, Height: 3,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.SwordGfx,
	}
	AllObjects = append(AllObjects, Sword)
}

func InitChalice(w, h int) {
	Chalice = &Object{
		RelX: 0.8, RelY: 0.25, Width: 8, Height: 5,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xAA, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.ChaliceGfx,
	}
	AllObjects = append(AllObjects, Chalice)
}

func InitMagnet(w, h int) {
	Magnet = &Object{
		RelX: 0.75, RelY: 0.65, Width: 8, Height: 4,
		Style: tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.MagnetGfx,
	}
	AllObjects = append(AllObjects, Magnet)
}

func InitDot(w, h int) {
	Dot = &Object{
		RelX: 0.3, RelY: 0.3, Width: 1, Height: 1,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xAA, 0xAA, 0xAA)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.DotGfx,
	}
	AllObjects = append(AllObjects, Dot)
}

func ReinitOnResize(w, h int) {
	if Portcullis != nil {
		doorStartCol := 18 * w / 40
		doorWidth := (4*w + 39) / 40
		if doorWidth < 2 {
			doorWidth = 2
		}
		portHeight := (h + 11) / 12
		if portHeight < 2 {
			portHeight = 2
		}
		frames := world.MakePortcullisFrames(doorWidth, portHeight)
		Portcullis.RelX = float64(doorStartCol) / float64(w)
		Portcullis.Width = doorWidth
		Portcullis.Height = portHeight
		Portcullis.Frames = frames
		Portcullis.Shape = frames[Portcullis.animFrame]
	}
}

func InitPlayer(w, h int) {
	Player = &Object{
		RelX:   float64(w/2) / float64(w),
		RelY:   float64(h/3*2) / float64(h),
		Width:  3,
		Height: 2,
		StepX:  2,
		StepY:  1,
		Style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple),
		Shape:  world.PlayerGfx,
	}
	AllObjects = append(AllObjects, Player)
}
