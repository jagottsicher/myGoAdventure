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

var GodMode bool

var HelpMode bool
var helpPreviousRoom *world.Room

func ResetObjects() {
	AllObjects = nil
}

var ConfirmMode bool
var ConfirmAction string // "quit" or "reset"

func StartConfirm(action string) {
	ConfirmMode = true
	ConfirmAction = action
}

func CancelConfirm() {
	ConfirmMode = false
	ConfirmAction = ""
}

func CycleVariation() {
	G.GameType = G.GameType%3 + 1
}

// CarriedObject is the object the player is currently carrying (nil = nothing).
var CarriedObject *Object

// carryOffsetX/Y: integer offset from player bounding-box top-left to
// carried object bounding-box top-left, in terminal columns/rows.
// Using integers avoids float rounding jitter on every frame.
var carryOffsetX, carryOffsetY int

// dropCooldown prevents immediately re-picking-up a just-dropped object.
var dropCooldown int

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// TryPickup checks cells adjacent to the player (cross shape, no corners).
// Snaps the picked-up object flush to the nearest side of the player.
func TryPickup(termW, termH int) {
	if Player == nil {
		return
	}
	if dropCooldown > 0 {
		dropCooldown--
		return
	}
	pLeft := int(Player.RelX*float64(termW)) - Player.Width/2
	pTop := int(Player.RelY*float64(termH)) - Player.Height/2
	pRight := pLeft + Player.Width
	pBottom := pTop + Player.Height

	hx1, hx2 := pLeft-1, pRight+1
	hy1, hy2 := pTop, pBottom
	vx1, vx2 := pLeft, pRight
	vy1, vy2 := pTop-1, pBottom+1

	overlaps := func(ox1, oy1, ox2, oy2 int) bool {
		hOvr := ox2 > hx1 && ox1 < hx2 && oy2 > hy1 && oy1 < hy2
		vOvr := ox2 > vx1 && ox1 < vx2 && oy2 > vy1 && oy1 < vy2
		return hOvr || vOvr
	}

	for _, obj := range AllObjects {
		if !obj.Carryable || obj == CarriedObject {
			continue
		}
		if obj.Room != nil && obj.Room != CurrentRoom {
			continue
		}
		ox := int(obj.RelX*float64(termW)) - obj.Width/2
		oy := int(obj.RelY*float64(termH)) - obj.Height/2
		if !overlaps(ox, oy, ox+obj.Width, oy+obj.Height) {
			continue
		}

		// Determine which side the object is on relative to the player center.
		pCX := pLeft + Player.Width/2
		pCY := pTop + Player.Height/2
		oCX := ox + obj.Width/2
		oCY := oy + obj.Height/2
		dx := oCX - pCX
		dy := oCY - pCY

		var offX, offY int
		if absInt(dx) >= absInt(dy) {
			// Left or right of player.
			if dx < 0 {
				offX = -obj.Width
			} else {
				offX = Player.Width
			}
			offY = (Player.Height - obj.Height) / 2
		} else {
			// Above or below player — center object on player.
			offX = Player.Width/2 - obj.Width/2
			if dy < 0 {
				offY = -obj.Height
			} else {
				offY = Player.Height
			}
		}

		if obj == BatCarrying {
			// Player steals from bat — bat immediately hunts for something new.
			BatCarrying = nil
			batFedUpTimer = 0xff
		}
		CarriedObject = obj
		carryOffsetX = offX
		carryOffsetY = offY

		// Sword: orient tip away from player on pickup.
		if obj == Sword && len(obj.Frames) >= 4 {
			var frame int
			if absInt(dx) >= absInt(dy) {
				if dx < 0 {
					frame = 1 // left
				} else {
					frame = 0 // right
				}
			} else {
				if dy < 0 {
					frame = 2 // up
				} else {
					frame = 3 // down
				}
			}
			obj.animFrame = frame
			obj.Shape = obj.Frames[frame]
		}

		// Magnet: field lines point away from player; always centered on player.
		if obj == Magnet && obj.SubFrameCount > 0 {
			var orient int
			if absInt(dx) >= absInt(dy) {
				if dx < 0 {
					orient = 3 // left
				} else {
					orient = 1 // right
				}
				// Center vertically on player.
				carryOffsetY = Player.Height/2 - obj.Height/2
			} else {
				if dy < 0 {
					orient = 2 // up
				} else {
					orient = 0 // down
				}
				// Center horizontally on player.
				carryOffsetX = Player.Width/2 - obj.Width/2
			}
			obj.OrientationFrame = orient
			obj.Shape = obj.Frames[orient*obj.SubFrameCount+obj.subFrame]
		}
		return
	}
}

// UpdateCarriedObject moves the carried object with the player every frame.
// All arithmetic stays in integer screen coordinates to avoid rounding jitter.
func UpdateCarriedObject(termW, termH int) {
	if CarriedObject == nil || Player == nil {
		return
	}
	pLeft := int(Player.RelX*float64(termW)) - Player.Width/2
	pTop := int(Player.RelY*float64(termH)) - Player.Height/2
	oLeft := pLeft + carryOffsetX
	oTop := pTop + carryOffsetY
	CarriedObject.RelX = float64(oLeft+CarriedObject.Width/2) / float64(termW)
	CarriedObject.RelY = float64(oTop+CarriedObject.Height/2) / float64(termH)
	CarriedObject.Room = CurrentRoom
}

// DropCarried drops the carried object at its current position.
func DropCarried() {
	if CarriedObject != nil {
		CarriedObject = nil
		dropCooldown = 20 // ~0.33s at 60 FPS — prevents immediate re-pickup
	}
}

// Bat AI state.
var (
	batFedUpTimer = 0xff  // 0xff = hunting; 0..254 = carrying (counts up to 0xff)
	BatCarrying   *Object // object the bat is currently carrying (nil = hunting)
	batMovX       int     // current X velocity in terminal cols (set while hunting)
	batMovY       int     // current Y velocity in terminal rows
)

func batPriorityList() []*Object {
	return []*Object{Chalice, Sword, Bridge, YellowKey, WhiteKey, BlackKey, RedDragon, YellowDragon, GreenDragon, Magnet}
}

// UpdateBat drives bat AI: hunting, carrying, room transitions.
// Call once per game tick from adventure.go updateStates().
func UpdateBat(termW, termH int) {
	if Bat == nil || Bat == CarriedObject {
		return
	}

	// While carrying: increment timer toward 0xff, then hunt for a different object.
	if BatCarrying != nil && batFedUpTimer < 0xff {
		batFedUpTimer++
	}

	// Hunting: find highest-priority object in bat's room (different from current carry).
	if batFedUpTimer >= 0xff {
		batLeft := int(Bat.RelX*float64(termW)) - Bat.Width/2
		batTop := int(Bat.RelY*float64(termH)) - Bat.Height/2

		// Expand bat extents by 4 terminal cells for proximity detection (≈7 Atari px).
		const expand = 4
		ebX1 := batLeft - expand
		ebY1 := batTop - expand
		ebX2 := batLeft + Bat.Width + expand
		ebY2 := batTop + Bat.Height + expand

		for _, obj := range batPriorityList() {
			if obj == nil || obj == BatCarrying {
				continue
			}
			if obj.Room != Bat.Room {
				continue
			}

			oLeft := int(obj.RelX*float64(termW)) - obj.Width/2
			oTop := int(obj.RelY*float64(termH)) - obj.Height/2

			// Steer toward target center.
			if batLeft+Bat.Width/2 < oLeft+obj.Width/2 {
				batMovX = 2
			} else if batLeft+Bat.Width/2 > oLeft+obj.Width/2 {
				batMovX = -2
			} else {
				batMovX = 0
			}
			if batTop+Bat.Height/2 < oTop+obj.Height/2 {
				batMovY = 1
			} else if batTop+Bat.Height/2 > oTop+obj.Height/2 {
				batMovY = -1
			} else {
				batMovY = 0
			}

			// Pick up if expanded bat extents overlap target.
			if ebX1 < oLeft+obj.Width && ebX2 > oLeft && ebY1 < oTop+obj.Height && ebY2 > oTop {
				if obj == CarriedObject {
					// Steal from player.
					CarriedObject = nil
					dropCooldown = 20
				}
				BatCarrying = obj
				batFedUpTimer = 0
			}
			break // only highest-priority target per frame
		}
	}

	// Apply current velocity (bat always moves, even while carrying).
	batLeft := int(Bat.RelX*float64(termW)) - Bat.Width/2
	batTop := int(Bat.RelY*float64(termH)) - Bat.Height/2
	batLeft += batMovX
	batTop += batMovY

	// Room transitions.
	if batLeft < 0 {
		if Bat.Room != nil && Bat.Room.Left != nil {
			Bat.Room = Bat.Room.Left
			batLeft = termW - Bat.Width
		} else {
			batLeft = 0
			batMovX = 0
		}
	} else if batLeft+Bat.Width > termW {
		if Bat.Room != nil && Bat.Room.Right != nil {
			Bat.Room = Bat.Room.Right
			batLeft = 0
		} else {
			batLeft = termW - Bat.Width
			batMovX = 0
		}
	}
	if batTop < 0 {
		if Bat.Room != nil && Bat.Room.Up != nil {
			Bat.Room = Bat.Room.Up
			batTop = termH - Bat.Height
		} else {
			batTop = 0
			batMovY = 0
		}
	} else if batTop+Bat.Height > termH {
		if Bat.Room != nil && Bat.Room.Down != nil {
			Bat.Room = Bat.Room.Down
			batTop = 0
		} else {
			batTop = termH - Bat.Height
			batMovY = 0
		}
	}

	Bat.RelX = float64(batLeft+Bat.Width/2) / float64(termW)
	Bat.RelY = float64(batTop+Bat.Height/2) / float64(termH)

	// Update carried object position: right side of bat, same top.
	if BatCarrying != nil {
		cLeft := batLeft + Bat.Width
		BatCarrying.RelX = float64(cLeft+BatCarrying.Width/2) / float64(termW)
		BatCarrying.RelY = float64(batTop+BatCarrying.Height/2) / float64(termH)
		BatCarrying.Room = Bat.Room
	}
}

func ToggleHelp() {
	HelpMode = !HelpMode
	if HelpMode {
		helpPreviousRoom = CurrentRoom
		CurrentRoom = &world.RoomSplashScreen
	} else {
		CurrentRoom = helpPreviousRoom
	}
}

var (
	playerNormalStyle = tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorPurple)
	playerGodStyle    = tcell.StyleDefault.Background(tcell.NewRGBColor(0xFF, 0xAA, 0x00)).Foreground(tcell.ColorPurple)
)

func ToggleGodMode() {
	GodMode = !GodMode
	if Player == nil {
		return
	}
	if GodMode {
		Player.Style = playerGodStyle
	} else {
		Player.Style = playerNormalStyle
	}
}

type Object struct {
	RelX         float64
	RelY         float64
	StepX        int
	StepY        int
	Width        int
	Height       int
	Flipped      bool         // mirror sprite horizontally
	ZLayer       int          // draw order: 0=default, 1=dragon, 2=bridge
	Room         *world.Room // if set, only rendered when CurrentRoom == Room
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

	Paused    bool // if true, Animate() does nothing (animation frozen)
	Solid     bool // if true, WouldCollideWall treats all cells as blocking
	Carryable bool // if true, player can pick this up

	animTick  int
	animFrame int // used as orientationFrame in flat mode
	subFrame  int
	subTick   int
}

// Animate advances the animation by one tick. Call once per game update.
func (o *Object) Animate() {
	if o.Paused || len(o.Frames) < 2 {
		return
	}
	if o.SubFrameCount > 0 {
		// 2D mode: field-line phase and orientation run independently.
		o.subTick++
		if o.subTick >= o.SubFrameInterval {
			o.subTick = 0
			o.subFrame = (o.subFrame + 1) % o.SubFrameCount
		}
		if o.AnimInterval > 0 {
			o.animTick++
			if o.animTick >= o.AnimInterval {
				o.animTick = 0
				numOrientations := len(o.Frames) / o.SubFrameCount
				o.OrientationFrame = (o.OrientationFrame + 1) % numOrientations
			}
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
var PortcullisYellow *Object // C++ PORT1 — castle screen 0x11 (RoomYellowCastle)
var PortcullisWhite *Object  // C++ PORT2 — castle screen 0x0F (RoomWhiteCastle)
var PortcullisBlack *Object  // C++ PORT3 — castle screen 0x10 (RoomBlackCastle)
var Bridge *Object
var Sword *Object
var Chalice *Object
var Magnet *Object
var Dot *Object
var AllObjects []*Object
var CurrentRoom *world.Room

func InitYellowKey(w, h int) {
	// C++ V2: room 0x09 (RoomMazeMiddle), X=0x20, Y=0x40
	YellowKey = &Object{
		RelX: 0.20, RelY: 0.50, Width: 8, Height: 2,
		Carryable: true,
		Room:      &world.RoomMazeMiddle,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xF5, 0xCE, 0x42)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.KeyGfx,
	}
	AllObjects = append(AllObjects, YellowKey)
}

func InitWhiteKey(w, h int) {
	// C++ V2: room 0x06 (RoomBlueMazeBottom), X=0x20, Y=0x40
	WhiteKey = &Object{
		RelX: 0.20, RelY: 0.50, Width: 8, Height: 2,
		Carryable: true,
		Room:      &world.RoomBlueMazeBottom,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xFF, 0xFF)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.KeyGfx,
	}
	AllObjects = append(AllObjects, WhiteKey)
}

func InitBlackKey(w, h int) {
	// C++ V2: room 0x19 (RoomRedMazeBottom), X=0x20, Y=0x40
	BlackKey = &Object{
		RelX: 0.20, RelY: 0.50, Width: 8, Height: 2,
		Carryable: true,
		Room:      &world.RoomRedMazeBottom,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x00, 0x00, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.KeyGfx,
	}
	AllObjects = append(AllObjects, BlackKey)
}

func InitGreenDragon(w, h int) {
	// C++ V2: room 0x04 (RoomBlueMazeTop), X=0x50, Y=0x20
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	GreenDragon = &Object{
		RelX: 0.50, RelY: 0.25, Width: 8, Height: 10,
		ZLayer:       1,
		Room:         &world.RoomBlueMazeTop,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x86, 0xd9, 0x22)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, GreenDragon)
}

func InitYellowDragon(w, h int) {
	// C++ V2: room 0x19 (RoomRedMazeBottom), X=0x50, Y=0x20
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	YellowDragon = &Object{
		RelX: 0.50, RelY: 0.25, Width: 8, Height: 10,
		ZLayer:       1,
		Room:         &world.RoomRedMazeBottom,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, YellowDragon)
}

func InitRedDragon(w, h int) {
	// C++ V2: room 0x14 (RoomBlackMaze2), X=0x50, Y=0x20
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	RedDragon = &Object{
		RelX: 0.50, RelY: 0.25, Width: 8, Height: 10,
		ZLayer:       1,
		Room:         &world.RoomBlackMaze2,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFA, 0x52, 0x55)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, RedDragon)
}

func InitBat(w, h int) {
	// C++ V2: room 0x02 (RoomBelowYellowCastle), X=0x20, Y=0x20
	frames := [][]*world.Cell{world.BatGfx, world.BatGfxOpen}
	Bat = &Object{
		RelX: 0.20, RelY: 0.25, Width: 8, Height: 6,
		Carryable:    true,
		Room:         &world.RoomBelowYellowCastle,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 1, // toggle every frame — matches original
	}
	AllObjects = append(AllObjects, Bat)
	// Reset bat AI state on every init/reset.
	batFedUpTimer = 0xff
	BatCarrying = nil
	batMovX = 0
	batMovY = 0
}

func InitPortcullises(w, h int) {
	// Castle template: 40 chars wide, 12 rows tall.
	// Gate opening: cols 18–21 (4 chars) in template row 5.
	doorStartCol := 18 * w / 40
	doorWidth := (4*w + 39) / 40 // ceiling division — never one short
	if doorWidth < 2 {
		doorWidth = 2
	}
	portHeight := h / 12
	if portHeight < 2 {
		portHeight = 2
	}
	frames := world.MakePortcullisFrames(doorWidth, portHeight)
	relX := float64(doorStartCol+doorWidth/2) / float64(w)
	relY := 5.0/12.0 + float64(portHeight/2)/float64(h)
	style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd))

	makePort := func(room *world.Room) *Object {
		o := &Object{
			RelX: relX, RelY: relY,
			Width: doorWidth, Height: portHeight,
			Style:        style,
			Shape:        frames[0],
			Frames:       frames,
			AnimInterval: 45,
			Room:         room,
			Paused:       true,
			Solid:        true,
		}
		AllObjects = append(AllObjects, o)
		return o
	}

	PortcullisYellow = makePort(&world.RoomYellowCastle)
	PortcullisWhite = makePort(&world.RoomWhiteCastle)
	PortcullisBlack = makePort(&world.RoomBlackCastle)
}

func InitBridge(w, h int) {
	// C++ V2: room 0x0B (RoomMazeSide), X=0x40, Y=0x40
	Bridge = &Object{
		RelX: 0.40, RelY: 0.50, Width: 10, Height: 12,
		ZLayer: 2,
		Carryable: true,
		Room:      &world.RoomMazeSide,
		Style:     tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x99, 0x00, 0xCC)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.BridgeGfx,
	}
	AllObjects = append(AllObjects, Bridge)
}

func InitSword(w, h int) {
	// C++ V2: room 0x11 (RoomYellowCastle), X=0x20, Y=0x20
	frames := [][]*world.Cell{world.SwordGfx, world.SwordGfxLeft, world.SwordGfxUp, world.SwordGfxDown}
	Sword = &Object{
		RelX: 0.20, RelY: 0.25, Width: 8, Height: 4,
		Carryable:    true,
		Room:         &world.RoomYellowCastle,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xF5, 0xCE, 0x42)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 60,
		Paused:       true,
	}
	AllObjects = append(AllObjects, Sword)
}

func InitChalice(w, h int) {
	// C++ V2: room 0x14 (RoomBlackMaze2), X=0x30, Y=0x20
	Chalice = &Object{
		RelX: 0.30, RelY: 0.25, Width: 8, Height: 5,
		Carryable: true,
		Room:      &world.RoomBlackMaze2,
		Style:     tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xAA, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:     world.ChaliceGfx,
	}
	AllObjects = append(AllObjects, Chalice)
}

func InitMagnet(w, h int) {
	// C++ V2: room 0x0E (RoomDeadEndCyan), X=0x80, Y=0x20
	frames := world.MakeMagnetFrames()
	Magnet = &Object{
		RelX: 0.80, RelY: 0.25, Width: 12, Height: 8,
		Carryable:        true,
		Room:             &world.RoomDeadEndCyan,
		Style:            tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:            frames[0],
		Frames:           frames,
		SubFrameCount:    4,
		SubFrameInterval: 15,
		AnimInterval:     0, // 0 = no orientation rotation; field lines still animate
		BodyOffsets:      [][2]int{{6, 2}, {4, 4}, {6, 6}, {8, 4}},
	}
	AllObjects = append(AllObjects, Magnet)
}

// InitBarrier creates a 1-wide vertical 'X' barrier at the given relative position.
// relX/relY is the center of the bounding box (Width=1, so no horizontal offset).
// Style is typically black-on-black so the barrier is invisible but still blocks movement.
// Because DrawObject renders 'X' onto the screen, WouldCollideWall detects them automatically.
func InitBarrier(room *world.Room, relX float64, h int, style tcell.Style) *Object {
	height := h
	b := &Object{
		RelX:   relX,
		RelY:   0.5,
		Width:  1,
		Height: height,
		Room:   room,
		Style:  style,
		Shape:  world.MakeBarrierGfx(height),
	}
	AllObjects = append(AllObjects, b)
	return b
}

func InitDot(w, h int) {
	// C++ V2: room 0x15 (RoomBlackMaze3), X=0x45, Y=0x12
	Dot = &Object{
		RelX: 0.43, RelY: 0.14, Width: 1, Height: 1,
		Carryable: true,
		Room:      &world.RoomBlackMaze3,
		Style:     tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xAA, 0xAA, 0xAA)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.DotGfx,
	}
	AllObjects = append(AllObjects, Dot)
}

// UpdatePortcullis drives the open/close state machine for one portcullis.
// key is the matching key object (YellowKey, WhiteKey, BlackKey).
// Opens when key is in the same room; stays open once fully open.
func UpdatePortcullis(port *Object, key *Object) {
	if port == nil || key == nil {
		return
	}
	fullyOpen := port.animFrame == len(port.Frames)-1

	if fullyOpen {
		// Already open — keep paused, not solid.
		port.Paused = true
		port.Solid = false
		return
	}

	keyPresent := key.Room == port.Room
	if keyPresent {
		// Key is here — animate open.
		port.Paused = false
		port.Solid = true // still solid until fully open
	} else {
		// Key gone before fully open — freeze.
		port.Paused = true
		port.Solid = true
	}
}

func ReinitOnResize(w, h int) {
	ports := []*Object{PortcullisYellow, PortcullisWhite, PortcullisBlack}
	if ports[0] == nil {
		return
	}
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
	relX := float64(doorStartCol+doorWidth/2) / float64(w)
	relY := 5.0/12.0 + float64(portHeight/2)/float64(h)
	for _, p := range ports {
		if p == nil {
			continue
		}
		p.RelX = relX
		p.RelY = relY
		p.Width = doorWidth
		p.Height = portHeight
		p.Frames = frames
		p.Shape = frames[p.animFrame]
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
		Style:  playerNormalStyle,
		Shape:  world.PlayerGfx,
	}
	AllObjects = append(AllObjects, Player)
}
