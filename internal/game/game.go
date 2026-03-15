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
	Flipped      bool // mirror sprite horizontally
	ZLayer       int  // draw order: 0=default, 1=dragon, 2=bridge
	Style        tcell.Style
	Shape        []*world.Cell
	Frames       [][]*world.Cell // animation frames; nil = static
	AnimInterval int             // game ticks per orientation change

	// 2D animation: SubFrameCount > 0 enables independent field-line + orientation timers.
	// Frames layout: Frames[orientationFrame * SubFrameCount + subFrame]
	// orientationFrame cycles via AnimInterval; subFrame cycles via SubFrameInterval.
	SubFrameCount    int // number of field-line phases per orientation (0 = flat animation)
	SubFrameInterval int // game ticks per field-line phase step
	OrientationFrame int // current orientation index (settable from outside)

	// BodyOffsets: per-orientation anchor point within the bounding box.
	// screenX = RelX*termW - BodyOffsets[orientation][0]
	// screenY = RelY*termH - BodyOffsets[orientation][1]
	// nil = fall back to Width/2, Height/2 (center of bounding box)
	BodyOffsets [][2]int

	animTick  int
	animFrame int // used as orientationFrame in flat mode
	subFrame  int
	subTick   int
}

// Animate advances the animation by one tick. Call once per game update.
func (o *Object) Animate() {
	if len(o.Frames) < 2 {
		return
	}
	if o.SubFrameCount > 0 {
		// 2D mode: field-line phase and orientation run independently.
		o.subTick++
		if o.subTick >= o.SubFrameInterval {
			o.subTick = 0
			o.subFrame = (o.subFrame + 1) % o.SubFrameCount
		}
		o.animTick++
		if o.animTick >= o.AnimInterval {
			o.animTick = 0
			numOrientations := len(o.Frames) / o.SubFrameCount
			o.OrientationFrame = (o.OrientationFrame + 1) % numOrientations
		}
		o.Shape = o.Frames[o.OrientationFrame*o.SubFrameCount+o.subFrame]
		return
	}
	// Flat mode.
	o.animTick++
	if o.animTick >= o.AnimInterval {
		o.animTick = 0
		o.animFrame = (o.animFrame + 1) % len(o.Frames)
		o.Shape = o.Frames[o.animFrame]
	}
}

var Player *Object
var YellowKey *Object
var WhiteKey *Object
var BlackKey *Object
var GreenDragon *Object
var YellowDragon *Object
var RedDragon *Object
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
		RelX:   1.0/5.0 + 4.0/float64(w),
		RelY:   0.5 + 1.0/float64(h),
		Width:  8,
		Height: 2,
		StepX:  0,
		StepY:  0,
		Style:  tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.KeyGfx,
	}
	AllObjects = append(AllObjects, YellowKey)
}

func InitWhiteKey(w, h int) {
	WhiteKey = &Object{
		RelX:   0.15 + 4.0/float64(w),
		RelY:   0.25 + 1.0/float64(h),
		Width:  8,
		Height: 2,
		Style:  tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xFF, 0xFF)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.KeyGfx,
	}
	AllObjects = append(AllObjects, WhiteKey)
}

func InitBlackKey(w, h int) {
	BlackKey = &Object{
		RelX:   0.15 + 4.0/float64(w),
		RelY:   0.65 + 1.0/float64(h),
		Width:  8,
		Height: 2,
		Style:  tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x00, 0x00, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.KeyGfx,
	}
	AllObjects = append(AllObjects, BlackKey)
}

func InitGreenDragon(w, h int) {
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	GreenDragon = &Object{
		RelX:         4.0/5.0 + 4.0/float64(w),
		RelY:         0.5 + 5.0/float64(h),
		Width:        8,
		Height:       10,
		StepX:        0,
		StepY:        0,
		ZLayer:       1,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x86, 0xd9, 0x22)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30, // ~0.5s at 60 FPS
	}
	AllObjects = append(AllObjects, GreenDragon)
}

func InitYellowDragon(w, h int) {
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	YellowDragon = &Object{
		RelX:         0.15 + 4.0/float64(w),
		RelY:         0.4 + 5.0/float64(h),
		Width:        8,
		Height:       10,
		ZLayer:       1,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, YellowDragon)
}

func InitRedDragon(w, h int) {
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	RedDragon = &Object{
		RelX:         0.75 + 4.0/float64(w),
		RelY:         0.65 + 5.0/float64(h),
		Width:        8,
		Height:       10,
		ZLayer:       1,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFA, 0x52, 0x55)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, RedDragon)
}

func InitBat(w, h int) {
	frames := [][]*world.Cell{world.BatGfx, world.BatGfxOpen}
	Bat = &Object{
		RelX:         1.0/5.0 + 4.0/float64(w),
		RelY:         4.0/5.0 + 3.0/float64(h),
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
		RelX:         float64(doorStartCol+doorWidth/2) / float64(w),
		RelY:         5.0/12.0 + float64(portHeight/2)/float64(h),
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
		RelX: 0.1 + 5.0/float64(w), RelY: 0.15 + 6.0/float64(h), Width: 10, Height: 12,
		ZLayer: 2,
		Style:  tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x99, 0x00, 0xCC)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.BridgeGfx,
	}
	AllObjects = append(AllObjects, Bridge)
}

func InitSword(w, h int) {
	frames := [][]*world.Cell{world.SwordGfx, world.SwordGfxLeft, world.SwordGfxUp, world.SwordGfxDown}
	Sword = &Object{
		RelX: 0.65 + 4.0/float64(w), RelY: 0.25 + 2.0/float64(h), Width: 8, Height: 4,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 60,
	}
	AllObjects = append(AllObjects, Sword)
}

func InitChalice(w, h int) {
	Chalice = &Object{
		RelX: 0.8 + 4.0/float64(w), RelY: 0.25 + 2.0/float64(h), Width: 8, Height: 5,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xAA, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.ChaliceGfx,
	}
	AllObjects = append(AllObjects, Chalice)
}

func InitMagnet(w, h int) {
	frames := world.MakeMagnetFrames()
	Magnet = &Object{
		RelX:             0.75 - 6.0/float64(w), RelY: 0.65 + 2.0/float64(h), Width: 12, Height: 8,
		Style:            tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:            frames[0],
		Frames:           frames,
		SubFrameCount:    4,  // 4 field-line phases per orientation
		SubFrameInterval: 15, // ~0.5s per field-line step
		AnimInterval:     90, // ~3s per orientation (for testing all 4 directions)
		// BodyOffsets: per-orientation body center within the 12x8 bounding box.
		// Down(0): body X=2–9, Y=0–3 → center=(6,2)
		// Right(1): body X=0–7, Y=2–5 → center=(4,4) (arch left, poles right)
		// Up(2):   body X=2–9, Y=4–7 → center=(6,6)
		// Left(3):  body X=4–11, Y=2–5 → center=(8,4) (arch right, poles left)
		BodyOffsets: [][2]int{{6, 2}, {4, 4}, {6, 6}, {8, 4}},
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
		Portcullis.RelX = float64(doorStartCol+doorWidth/2) / float64(w)
		Portcullis.Width = doorWidth
		Portcullis.Height = portHeight
		Portcullis.Frames = frames
		Portcullis.Shape = frames[Portcullis.animFrame]
	}
}

func InitPlayer(w, h int) {
	Player = &Object{
		RelX:   float64(w/2+1) / float64(w),
		RelY:   float64(h/3*2+1) / float64(h),
		Width:  3,
		Height: 2,
		StepX:  2,
		StepY:  1,
		Style:  tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple),
		Shape:  world.PlayerGfx,
	}
	AllObjects = append(AllObjects, Player)
}
