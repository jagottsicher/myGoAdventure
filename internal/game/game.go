package game

import (
	"fmt"
	"math"
	"math/rand"
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

// DifficultyLeft/Right mirror the two physical switches on the Atari 2600.
// true = A (harder), false = B (easier).
// DifficultyLeft:  A = shorter roar window after dragon touch (harder to escape).
// DifficultyRight: A = dragons flee from the sword; B = dragons ignore the sword.
var DifficultyLeft = false  // default B
var DifficultyRight = false // default B

func ToggleDifficultyLeft() {
	DifficultyLeft = !DifficultyLeft
}

func ToggleDifficultyRight() {
	DifficultyRight = !DifficultyRight
}

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

// SelOverlay is the selection overlay state for variation and difficulty cycling.
type SelOverlay struct {
	Active bool
	Kind   string // "variation" or "difficulty"
	Value  int    // current preview: variation=1..3, difficulty=0(A)/1(B)
	Ticks  int    // countdown; reaches 0 → apply and close
}

var Overlay SelOverlay

// HandleSelOverlayKey is called when the V or F key is pressed.
// Pressing V while the difficulty overlay is open applies the pending difficulty
// selection first, then opens the variation overlay (and vice versa).
func HandleSelOverlayKey(kind string) {
	if Overlay.Active && Overlay.Kind != kind {
		applySelOverlay()
	}
	if Overlay.Active && Overlay.Kind == kind {
		cycleSelOverlay()
		Overlay.Ticks = 3 * int(G.FPS)
	} else {
		var val int
		if kind == "variation" {
			val = int(G.GameType)
		} else if !DifficultyLeft {
			val = 1 // B
		}
		Overlay = SelOverlay{Active: true, Kind: kind, Value: val, Ticks: 3 * int(G.FPS)}
	}
}

func cycleSelOverlay() {
	switch Overlay.Kind {
	case "variation":
		Overlay.Value = Overlay.Value%3 + 1
	case "difficulty":
		Overlay.Value = 1 - Overlay.Value
	}
}

// NeedFullReset signals adventure.go to call render.InitGamestate on the next tick.
var NeedFullReset bool

func applySelOverlay() {
	if !Overlay.Active {
		return
	}
	switch Overlay.Kind {
	case "variation":
		newType := uint8(Overlay.Value)
		if newType != G.GameType {
			G.GameType = newType
			NeedFullReset = true
		}
	case "difficulty":
		a := Overlay.Value == 0
		DifficultyLeft = a
		DifficultyRight = a
	}
	Overlay.Active = false
}

// ClearForFullReset clears all dynamic game state before a full re-init.
// Called from adventure.go when NeedFullReset is true.
func ClearForFullReset() {
	AllObjects = nil
	CarriedObject = nil
	GameWon = false
	WinOverlayTimer = 0
	HelpMode = false
	ConfirmMode = false
	GodMode = false
	Eaten = false
	NeedFullReset = false
}

// RandomizeObjectsV3 mirrors C++ SetupRoomObjects randomization for game level 2 (variation 3).
// For each object in roomBoundsData, picks a random room in [lower, upper] using the
// same retry-loop strategy as the original.
func RandomizeObjectsV3() {
	if G.GameType != 3 {
		return
	}
	type bound struct {
		obj         **Object
		lower, upper int
	}
	bounds := []bound{
		{&Chalice,      0x13, 0x1A},
		{&RedDragon,    0x01, 0x1D},
		{&YellowDragon, 0x01, 0x1D},
		{&GreenDragon,  0x01, 0x1D},
		{&Sword,        0x01, 0x1D},
		{&Bridge,       0x01, 0x1D},
		{&YellowKey,    0x01, 0x1D},
		{&WhiteKey,     0x01, 0x16},
		{&BlackKey,     0x01, 0x12},
		{&Bat,          0x01, 0x1D},
		{&Magnet,       0x01, 0x1D},
	}
	for _, b := range bounds {
		if *b.obj == nil {
			continue
		}
		for {
			id := rand.Intn(0x1F) // 0x00–0x1E, matches C++ Platform_Random()*0x1f
			if id >= b.lower && id <= b.upper {
				(*b.obj).Room = world.RoomsByID[id]
				break
			}
		}
	}
}

// UpdateSelOverlay decrements the overlay timer and applies the selection on timeout.
// Call once per game tick.
func UpdateSelOverlay() {
	if !Overlay.Active {
		return
	}
	Overlay.Ticks--
	if Overlay.Ticks <= 0 {
		applySelOverlay()
	}
}

// CarriedObject is the object the player is currently carrying (nil = nothing).
var CarriedObject *Object

// carryOffsetX/Y: integer offset from player bounding-box top-left to
// carried object bounding-box top-left, in terminal columns/rows.
// Using integers avoids float rounding jitter on every frame.
var carryOffsetX, carryOffsetY int

// dropCooldown prevents immediately re-picking-up a just-dropped object.
var dropCooldown int

// PlayerMoved is set to true by HandleUserInput whenever the player's position
// changes. Reset to false at the start of each HandleUserInput call.
// TryPickup only fires when this is true — prevents objects drifting (magnet,
// bat drop) from auto-attaching to a stationary player.
var PlayerMoved bool

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// TryPickup checks cells adjacent to the player (cross shape, no corners).
// Snaps the picked-up object flush to the nearest side of the player.
// Only fires when the player moved this frame — objects drifting into the
// pickup zone (via magnet or bat drop) do not auto-attach.
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

		// dx/dy: relative position of object center to player center (used for offset + orientation).
		pCX := pLeft + Player.Width/2
		pCY := pTop + Player.Height/2
		oCX := ox + obj.Width/2
		oCY := oy + obj.Height/2
		dx := oCX - pCX
		dy := oCY - pCY

		var offX, offY int
		if obj == Bridge {
			// Bridge: only grabbable at the two pillar areas (cols 0–1 = left, cols 8–9 = right).
			// - From left/right: player must be outside the bridge touching that pillar.
			// - From above/below: player X must overlap a pillar column, not just the deck (cols 2–7).
			// Offset preserves the exact grab position — no centering.
			leftPillarR := ox + 2
			rightPillarL := ox + 8
			bridgeR := ox + obj.Width

			fromLeft := pRight >= ox && pRight <= leftPillarR &&
				pBottom > oy && pTop < oy+obj.Height
			fromRight := pLeft >= rightPillarL && pLeft <= bridgeR &&
				pBottom > oy && pTop < oy+obj.Height

			atLeftPillar := pRight > ox && pLeft < leftPillarR
			atRightPillar := pRight > rightPillarL && pLeft < bridgeR
			pillarContact := atLeftPillar || atRightPillar
			fromAbove := pBottom >= oy && pBottom <= oy+1 && pillarContact
			fromBelow := pTop >= oy+obj.Height-1 && pTop <= oy+obj.Height && pillarContact

			if !fromLeft && !fromRight && !fromAbove && !fromBelow {
				continue
			}
			offX = ox - pLeft
			offY = oy - pTop
		} else {
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
		}

		if obj == BatCarrying {
			// Player steals from bat — give the player a grace period before bat hunts again.
			BatCarrying = nil
			batFedUpTimer = 0xff
			batHuntDelay = 90 // ~1.5s at 60 FPS
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
	batFedUpTimer  = 0xff // 0xff = hunting; 0..254 = carrying (counts up to 0xff)
	batHuntDelay   int    // frames remaining before bat may hunt again after player steal
	BatCarrying    *Object // object the bat is currently carrying (nil = hunting)
	batDirX        int    // movement direction: -1, 0, or +1
	batDirY        int    // movement direction: -1, 0, or +1
	batTick        int    // tick counter for movement gate
	batCellX       int    // persistent bat position in terminal columns (top-left of bounding box)
	batCellY      int     // persistent bat position in terminal rows
)

// BatDebugState returns a debug string with bat AI state — remove after debugging.
func BatDebugState() string {
	if Bat == nil {
		return "BAT nil"
	}
	carry := "-"
	if BatCarrying != nil {
		carry = "CARRY"
	}
	carried := "-"
	if Bat == CarriedObject {
		carried = "BY_PLAYER"
	}
	return fmt.Sprintf("dx=%+d dy=%+d tick=%d fed=%d carry=%s held=%s cellY=%d relY=%.4f",
		batDirX, batDirY, batTick, batFedUpTimer, carry, carried, batCellY, Bat.RelY)
}

func batPriorityList() []*Object {
	return []*Object{Chalice, Sword, Bridge, YellowKey, WhiteKey, BlackKey, RedDragon, YellowDragon, GreenDragon, Magnet}
}

// UpdateBat drives bat AI: hunting, carrying, room transitions.
// Call once per game tick from adventure.go updateStates().
func UpdateBat(termW, termH int) {
	if Bat == nil {
		return
	}

	// When the player is carrying the bat: sync integer state from player-moved RelX/RelY,
	// update the bat's own carried object, then return.
	if Bat == CarriedObject {
		batCellX = int(Bat.RelX*float64(termW)) - Bat.Width/2
		batCellY = int(Bat.RelY*float64(termH)) - Bat.Height/2
		if BatCarrying != nil {
			cLeft := batCellX + Bat.Width
			BatCarrying.RelX = float64(cLeft+BatCarrying.Width/2) / float64(termW)
			BatCarrying.RelY = float64(batCellY+BatCarrying.Height/2) / float64(termH)
			BatCarrying.Room = Bat.Room
		}
		return
	}

	// While carrying: increment timer toward 0xff, then hunt for a different object.
	if BatCarrying != nil && batFedUpTimer < 0xff {
		batFedUpTimer++
	}

	// Hunting: find highest-priority object in bat's room (different from current carry).
	if batHuntDelay > 0 {
		batHuntDelay--
	}
	if batFedUpTimer >= 0xff && batHuntDelay == 0 {
		// Use persistent integer position — no float64 round-trip.
		const expand = 4
		ebX1 := batCellX - expand
		ebY1 := batCellY - expand
		ebX2 := batCellX + Bat.Width + expand
		ebY2 := batCellY + Bat.Height + expand

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
			bCX := batCellX + Bat.Width/2
			objCX := oLeft + obj.Width/2
			if bCX < objCX {
				batDirX = 1
			} else if bCX > objCX {
				batDirX = -1
			} else {
				batDirX = 0
			}
			bCY := batCellY + Bat.Height/2
			objCY := oTop + obj.Height/2
			if bCY < objCY {
				batDirY = 1
			} else if bCY > objCY {
				batDirY = -1
			} else {
				batDirY = 0
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

	// Ensure bat always has velocity — both axes must not be zero simultaneously.
	if batDirX == 0 && batDirY == 0 {
		batDirY = 1
	}

	// Move 1 cell per axis every 4 frames (15 steps/s at 60 fps).
	// Integer state batCellX/Y is never read back from RelX/RelY — no round-trip rounding.
	batTick = (batTick + 1) % 3
	if batTick == 0 {
		batCellX += batDirX
		batCellY += batDirY
	}

	// Room transitions: cross into adjacent room or bounce off dead-end walls.
	if batCellX < 0 {
		if Bat.Room != nil && Bat.Room.Left != nil {
			Bat.Room = Bat.Room.Left
			batCellX = termW - Bat.Width
		} else {
			batCellX = 0
			batDirX = 1
		}
	} else if batCellX+Bat.Width > termW {
		if Bat.Room != nil && Bat.Room.Right != nil {
			Bat.Room = Bat.Room.Right
			batCellX = 0
		} else {
			batCellX = termW - Bat.Width
			batDirX = -1
		}
	}
	if batCellY < 0 {
		if Bat.Room != nil && Bat.Room.Up != nil {
			Bat.Room = Bat.Room.Up
			batCellY = termH - Bat.Height
		} else {
			batCellY = 0
			batDirY = 1
		}
	} else if batCellY+Bat.Height > termH {
		if Bat.Room != nil && Bat.Room.Down != nil {
			Bat.Room = Bat.Room.Down
			batCellY = 0
		} else {
			batCellY = termH - Bat.Height
			batDirY = -1
		}
	}

	Bat.RelX = float64(batCellX+Bat.Width/2) / float64(termW)
	Bat.RelY = float64(batCellY+Bat.Height/2) / float64(termH)

	// Update carried object position: right side of bat, same top.
	if BatCarrying != nil {
		cLeft := batCellX + Bat.Width
		BatCarrying.RelX = float64(cLeft+BatCarrying.Width/2) / float64(termW)
		BatCarrying.RelY = float64(batCellY+BatCarrying.Height/2) / float64(termH)
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

var playerGodStyle = tcell.StyleDefault.Background(tcell.NewRGBColor(0xFF, 0xAA, 0x00)).Foreground(tcell.NewRGBColor(0xFF, 0xAA, 0x00))

func ToggleGodMode() {
	GodMode = !GodMode
	if Player == nil {
		return
	}
	if GodMode {
		Player.Style = playerGodStyle
	}
	// GodMode off: UpdatePlayerStyle() corrects the color on the next frame.
}

// UpdatePlayerStyle sets the player's style to match the current room color.
// Both fg and bg are identical so the LMR characters render as solid blocks.
// Mirrors C++ behavior: ball is drawn in roomDefs[room].color.
func UpdatePlayerStyle() {
	if Player == nil || CurrentRoom == nil || GodMode {
		return
	}
	c := CurrentRoom.Foreground
	// Dark maze rooms have Foreground == Background (walls invisible).
	// Use the same orange as the torch aura so the player matches the visible walls.
	if c == CurrentRoom.Background {
		c = tcell.NewRGBColor(0xFF, 0x80, 0x00)
	}
	Player.Style = tcell.StyleDefault.Background(c).Foreground(c)
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
	Dead      bool // if true, object is removed from rendering and AI

	// Portcullis state machine (mirrors C++ Portals() logic).
	// PortState is an index into portStatesSeq (0–23).
	// 0 = closed/stopped, 12 = fully open/stopped, >22 → permanent unlock.
	PortState int
	PortTick  int  // frame counter — state only advances every portTicksPerStep frames
	Unlocked  bool // permanently open — gate never becomes solid again

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
var EasterEggBarrier *Object
var AllObjects []*Object
var CurrentRoom *world.Room

// castle portal state — tracked across frames by UpdateCastlePortals
var castlePortalPrevRoom *world.Room

// Win condition state.
var GameWon bool
var WinFlashTimer int   // counts down 255→0: screen flashes
var WinOverlayTimer int // counts down 180→0: "YOU WON!" overlay visible

// Flash color state — hue cycles 0–359, lum cycles 0–200.
// Chalice always uses flash colors; room also flashes during win.
var flashHue int
var flashLum int

// GetFlashColor returns the current cycling color (matches C++ GetFlashColor).
func GetFlashColor() tcell.Color {
	h := float64(flashHue) / (360.0 / 3) // 0–3 range
	var r, g, b float64
	if h < 1 {
		r = h * 255; g = 0; b = (1 - h) * 255
	} else if h < 2 {
		h -= 1; r = (1 - h) * 255; g = h * 255; b = 0
	} else {
		h -= 2; r = 0; g = (1 - h) * 255; b = h * 255
	}
	lum := float64(flashLum)
	if lum > r {
		r = lum
	}
	if lum > g {
		g = lum
	}
	if lum > b {
		b = lum
	}
	return tcell.NewRGBColor(int32(r), int32(g), int32(b))
}

// AdvanceFlashColor increments the flash hue and lum every frame.
func AdvanceFlashColor() {
	flashHue += 2
	if flashHue >= 360 {
		flashHue -= 360
	}
	flashLum += 11
	if flashLum > 200 {
		flashLum = 0
	}
}

// UpdateChaliceColor keeps the chalice style in sync with the flash color.
func UpdateChaliceColor() {
	if Chalice == nil {
		return
	}
	fc := GetFlashColor()
	Chalice.Style = tcell.StyleDefault.Foreground(fc).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd))
}

// CheckWinCondition triggers the win state when the chalice reaches RoomAboveYellowCastle.
func CheckWinCondition() {
	if GameWon || Chalice == nil {
		return
	}
	if Chalice.Room == &world.RoomAboveYellowCastle {
		GameWon = true
		WinFlashTimer = 255
	}
}

// UpdateWinState decrements the flash timer, then starts the overlay timer.
func UpdateWinState() {
	if !GameWon {
		return
	}
	if WinFlashTimer > 0 {
		WinFlashTimer--
		if WinFlashTimer == 0 {
			WinOverlayTimer = 5 * 60 // 5 seconds at 60 fps
		}
	} else if WinOverlayTimer > 0 {
		WinOverlayTimer--
	}
}

func InitYellowKey(w, h int) {
	// C++ V1: room 0x11 (RoomYellowCastle) / V2: room 0x09 (RoomMazeMiddle)
	room := &world.RoomMazeMiddle
	if G.GameType == 1 {
		room = &world.RoomYellowCastle
	}
	YellowKey = &Object{
		RelX: 0.20, RelY: 0.50, Width: 8, Height: 2,
		ZLayer:    1,
		Carryable: true,
		Room:      room,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xF5, 0xCE, 0x42)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.KeyGfx,
	}
	AllObjects = append(AllObjects, YellowKey)
}

func InitWhiteKey(w, h int) {
	// C++ V1: room 0x0E (RoomDeadEndCyan) / V2: room 0x06 (RoomBlueMazeBottom)
	room := &world.RoomBlueMazeBottom
	if G.GameType == 1 {
		room = &world.RoomDeadEndCyan
	}
	WhiteKey = &Object{
		RelX: 0.20, RelY: 0.50, Width: 8, Height: 2,
		ZLayer:    1,
		Carryable: true,
		Room:      room,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xFF, 0xFF)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.KeyGfx,
	}
	AllObjects = append(AllObjects, WhiteKey)
}

func InitBlackKey(w, h int) {
	// C++ V1: room 0x1D (RoomBlackCastleTop) / V2: room 0x19 (RoomRedMazeBottom)
	room := &world.RoomRedMazeBottom
	if G.GameType == 1 {
		room = &world.RoomBlackCastleTop
	}
	BlackKey = &Object{
		RelX: 0.20, RelY: 0.50, Width: 8, Height: 2,
		ZLayer:    1,
		Carryable: true,
		Room:      room,
		Style: tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x00, 0x00, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.KeyGfx,
	}
	AllObjects = append(AllObjects, BlackKey)
}

// dragonState holds persistent integer position and movement state for one dragon.
// Using integers avoids the float64 round-trip truncation bug that froze the bat.
type dragonState struct {
	cellX, cellY int // bounding-box top-left in terminal columns/rows
	dirX, dirY   int // current movement direction: -1, 0, or +1 per axis
	tick         int // tick counter for movement gate
	// State machine mirroring C++ dragon->state:
	// 0 = alive/stalking, 1 = dead, 2 = eaten player, 3 = roaring (timer running)
	State     int
	RoarTimer int
	// Dormant: true until the player first enters the dragon's room.
	// Deviation from C++ original (where dragons always move from game start),
	// but avoids dragons roaming across the entire map before the player finds them.
	Dormant bool
}

// Eaten is true while the player is trapped inside a dragon (state 2).
var Eaten bool

// dragonDiffTable mirrors C++ dragonDiff[]: pairs of (B, A) timer offsets per game level.
// timer = 0xFC - dragonDiffTable[gameLevel*2 + (DifficultyLeft ? 1 : 0)]
// Smaller timer = harder (less time to escape after touching a dragon).
var dragonDiffTable = [6]int{0xD0, 0xE8, 0xF0, 0xF6, 0xF0, 0xF6}

func dragonRoarTimer() int {
	level := int(G.GameType) - 1 // GameType is 1-indexed
	if level < 0 {
		level = 0
	}
	diffIdx := 0
	if DifficultyLeft {
		diffIdx = 1 // A = harder
	}
	return 0xFC - dragonDiffTable[level*2+diffIdx]
}

var (
	GreenDS  dragonState
	YellowDS dragonState
	RedDS    dragonState
)

func initDragonState(ds *dragonState, dragon *Object, w, h int) {
	ds.cellX = int(dragon.RelX*float64(w)) - dragon.Width/2
	ds.cellY = int(dragon.RelY*float64(h)) - dragon.Height/2
	ds.tick = 0
	ds.State = 0
	ds.RoarTimer = 0
	ds.Dormant = true
	Eaten = false
	// V1 (GameType==1): dragons stand still. V2/V3: start moving diagonally.
	if G.GameType == 1 {
		ds.dirX = 0
		ds.dirY = 0
	} else {
		ds.dirX = 1
		ds.dirY = 1
	}
}

func InitGreenDragon(w, h int) {
	// C++ V1: room 0x1D (RoomBlackCastleTop) / V2: room 0x04 (RoomBlueMazeTop)
	room := &world.RoomBlueMazeTop
	if G.GameType == 1 {
		room = &world.RoomBlackCastleTop
	}
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	GreenDragon = &Object{
		RelX: 0.50, RelY: 0.25, Width: 8, Height: 10,
		ZLayer:       1,
		Room:         room,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x86, 0xd9, 0x22)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, GreenDragon)
	initDragonState(&GreenDS, GreenDragon, w, h)
}

func InitYellowDragon(w, h int) {
	// C++ V1: room 0x01 (RoomTopAccessRight) / V2: room 0x19 (RoomRedMazeBottom)
	room := &world.RoomRedMazeBottom
	if G.GameType == 1 {
		room = &world.RoomTopAccessRight
	}
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	YellowDragon = &Object{
		RelX: 0.50, RelY: 0.25, Width: 8, Height: 10,
		ZLayer:       1,
		Room:         room,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xD8, 0x4C)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, YellowDragon)
	initDragonState(&YellowDS, YellowDragon, w, h)
}

func InitRedDragon(w, h int) {
	// C++ V1: room 0x0E (RoomDeadEndCyan) / V2: room 0x14 (RoomBlackMaze2)
	room := &world.RoomBlackMaze2
	if G.GameType == 1 {
		room = &world.RoomDeadEndCyan
	}
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	RedDragon = &Object{
		RelX: 0.50, RelY: 0.25, Width: 8, Height: 10,
		ZLayer:       1,
		Room:         room,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFA, 0x52, 0x55)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 30,
	}
	AllObjects = append(AllObjects, RedDragon)
	initDragonState(&RedDS, RedDragon, w, h)
}

// greenDragonMatrix returns the flee/seek pairs for the green dragon.
// Mirrors C++ greenDragonMatrix[]: (SWORD,GREENDRAGON),(GREENDRAGON,BALL),
// (GREENDRAGON,CHALISE),(GREENDRAGON,BRIDGE),(GREENDRAGON,MAGNET),(GREENDRAGON,BLACKKEY).
// self-ref entries (GreenDragon) are returned as nil — callers skip nil flee/seek.
func greenDragonMatrix() [][2]*Object {
	return [][2]*Object{
		{Sword, nil},        // flee SWORD; no seek (self-ref → nil)
		{nil, Player},       // no flee (self-ref → nil); seek BALL (player)
		{nil, Chalice},      // seek CHALICE
		{nil, Bridge},       // seek BRIDGE
		{nil, Magnet},       // seek MAGNET
		{nil, BlackKey},     // seek BLACKKEY
	}
}

// yellowDragonMatrix returns the flee/seek pairs for the yellow dragon.
// Mirrors C++ yellowDragonMatrix[]: (SWORD,YELLOWDRAGON),(YELLOWKEY,YELLOWDRAGON),
// (YELLOWDRAGON,BALL),(YELLOWDRAGON,CHALISE).
func yellowDragonMatrix() [][2]*Object {
	return [][2]*Object{
		{Sword, nil},        // flee SWORD; no seek (self-ref → nil)
		{YellowKey, nil},    // flee YELLOWKEY; no seek (self-ref → nil)
		{nil, Player},       // seek BALL (player)
		{nil, Chalice},      // seek CHALICE
	}
}

// redDragonMatrix returns the flee/seek pairs for the red dragon.
// Mirrors C++ redDragonMatrix[]: (SWORD,REDDRAGON),(REDDRAGON,BALL),
// (REDDRAGON,CHALISE),(REDDRAGON,WHITEKEY).
func redDragonMatrix() [][2]*Object {
	return [][2]*Object{
		{Sword, nil},        // flee SWORD; no seek (self-ref → nil)
		{nil, Player},       // seek BALL (player)
		{nil, Chalice},      // seek CHALICE
		{nil, WhiteKey},     // seek WHITEKEY
	}
}

// moveDragon updates one dragon's direction from its matrix, then moves it.
// tickPeriod controls speed: 4 = medium (green/yellow, same as bat), 3 = fast (red).
func moveDragon(dragon *Object, ds *dragonState, matrix [][2]*Object, tickPeriod, termW, termH int) {
	if dragon == nil || ds.State == 1 {
		return // dead: nothing
	}

	// State 2: eaten — freeze player inside dragon, don't move dragon.
	if ds.State == 2 {
		if Player != nil {
			Player.RelX = dragon.RelX
			Player.RelY = dragon.RelY
			CurrentRoom = dragon.Room
		}
		return
	}

	// State 3: roaring — count down, then eat or release.
	if ds.State == 3 {
		ds.RoarTimer--
		if ds.RoarTimer <= 0 {
			if dragon.Room == CurrentRoom && Player != nil &&
				CollisionCheckObjects(dragon, Player, termW, termH) {
				ds.State = 2
				Eaten = true
			} else {
				ds.State = 0
			}
		}
		return // no movement during roar
	}

	// Dormant: wait until the player first enters this dragon's room, then wake permanently.
	if ds.Dormant {
		if dragon.Room == CurrentRoom {
			ds.Dormant = false
		} else {
			return
		}
	}

	dCX := ds.cellX + dragon.Width/2
	dCY := ds.cellY + dragon.Height/2

	// Difficulty Right B: skip the first matrix pair (sword flee).
	// Mirrors C++ `matrix+2` when gameDifficultyRight == DIFFICULTY_B.
	activeMatrix := matrix
	if !DifficultyRight && len(activeMatrix) > 1 {
		activeMatrix = activeMatrix[1:]
	}

	// Determine direction from matrix: flee takes priority over seek within each pair.
	// dirSet is true when the matrix fired — prevents the adjacent-room override below.
	dirSet := false
	for _, pair := range activeMatrix {
		flee, seek := pair[0], pair[1]

		// Flee check (nil = self-ref, skip).
		if flee != nil {
			inRoom := (flee == Player && dragon.Room == CurrentRoom) ||
				(flee != Player && flee.Room == dragon.Room)
			if inRoom {
				fx := int(flee.RelX*float64(termW))
				fy := int(flee.RelY*float64(termH))
				if dCX < fx {
					ds.dirX = -1
				} else if dCX > fx {
					ds.dirX = 1
				} else {
					ds.dirX = 0
				}
				if dCY < fy {
					ds.dirY = -1
				} else if dCY > fy {
					ds.dirY = 1
				} else {
					ds.dirY = 0
				}
				dirSet = true
				break
			}
		}

		// Seek check (nil = self-ref, skip).
		// Only reached when flee is nil (self-ref) — matches C++ else-branch logic.
		if flee == nil && seek != nil {
			inRoom := (seek == Player && dragon.Room == CurrentRoom) ||
				(seek != Player && seek.Room == dragon.Room)
			if inRoom {
				sx := int(seek.RelX*float64(termW))
				sy := int(seek.RelY*float64(termH))
				if dCX < sx {
					ds.dirX = 1
				} else if dCX > sx {
					ds.dirX = -1
				} else {
					ds.dirX = 0
				}
				if dCY < sy {
					ds.dirY = 1
				} else if dCY > sy {
					ds.dirY = -1
				} else {
					ds.dirY = 0
				}
				dirSet = true
				break
			}
		}
	}

	// Adjacent-room following: when the matrix didn't find any target in the dragon's
	// current room, steer toward the exit that leads to the player's room (1-step lookahead).
	// This is the fix for "dragon never follows across screen boundaries" — without it,
	// the dragon just keeps its old direction which may point at a dead-end wall.
	if !dirSet && dragon.Room != nil {
		switch CurrentRoom {
		case dragon.Room.Up:
			ds.dirY = -1
		case dragon.Room.Down:
			ds.dirY = 1
		case dragon.Room.Left:
			ds.dirX = -1
		case dragon.Room.Right:
			ds.dirX = 1
		}
	}

	// Ensure dragon always has velocity.
	if ds.dirX == 0 && ds.dirY == 0 {
		ds.dirY = 1
	}

	// Move 1 cell per axis every tickPeriod frames.
	ds.tick = (ds.tick + 1) % tickPeriod
	if ds.tick == 0 {
		ds.cellX += ds.dirX
		ds.cellY += ds.dirY
	}

	// Room transitions: cross into adjacent room or bounce off dead-end walls.
	if ds.cellX < 0 {
		if dragon.Room != nil && dragon.Room.Left != nil {
			dragon.Room = dragon.Room.Left
			ds.cellX = termW - dragon.Width
		} else {
			ds.cellX = 0
			ds.dirX = 1
		}
	} else if ds.cellX+dragon.Width > termW {
		if dragon.Room != nil && dragon.Room.Right != nil {
			dragon.Room = dragon.Room.Right
			ds.cellX = 0
		} else {
			ds.cellX = termW - dragon.Width
			ds.dirX = -1
		}
	}
	if ds.cellY < 0 {
		if dragon.Room != nil && dragon.Room.Up != nil {
			dragon.Room = dragon.Room.Up
			ds.cellY = termH - dragon.Height
		} else {
			ds.cellY = 0
			ds.dirY = 1
		}
	} else if ds.cellY+dragon.Height > termH {
		if dragon.Room != nil && dragon.Room.Down != nil {
			dragon.Room = dragon.Room.Down
			ds.cellY = 0
		} else {
			ds.cellY = termH - dragon.Height
			ds.dirY = -1
		}
	}

	dragon.RelX = float64(ds.cellX+dragon.Width/2) / float64(termW)
	dragon.RelY = float64(ds.cellY+dragon.Height/2) / float64(termH)
}

// killDragonIfSwordHits checks whether the sword overlaps the dragon and kills it.
// Room check is the cheap early exit — matches C++ CollisionCheckObjectObject order.
// Optimization: skips entirely when sword and dragon are in different rooms (O(1)).
func killDragonIfSwordHits(dragon *Object, ds *dragonState, termW, termH int) {
	// Only state 0 (alive/stalking) can be killed — matches C++ check inside state==0 block.
	if dragon == nil || ds.State != 0 {
		return
	}
	if Sword == nil || Sword.Room != dragon.Room {
		return
	}
	if CollisionCheckObjects(dragon, Sword, termW, termH) {
		ds.State = 1
		ds.dirX = 0
		ds.dirY = 0
		// Switch to dead graphic and freeze animation (dragonStates[1]=frame2=dead).
		dragon.Shape = world.DragonGfxDead
		dragon.Frames = dragon.Frames[:1] // Animate() skips when len(Frames)<=1
	}
}

// touchDragonIfPlayerHits checks whether the player overlaps the dragon and triggers
// the roar state (state 3). Only fires in state 0 (stalking) — mirrors C++ check.
func touchDragonIfPlayerHits(dragon *Object, ds *dragonState, termW, termH int) {
	// Already eaten by another dragon — don't trigger again.
	if Eaten {
		return
	}
	if dragon == nil || ds.State != 0 || Player == nil {
		return
	}
	if dragon.Room != CurrentRoom {
		return
	}
	if CollisionCheckObjects(dragon, Player, termW, termH) {
		ds.Dormant = false // wake up on contact even if not yet awake
		ds.State = 3
		ds.RoarTimer = dragonRoarTimer()
		// Dragon snaps to player position (C++ behavior).
		dragon.RelX = Player.RelX
		dragon.RelY = Player.RelY
		ds.cellX = int(dragon.RelX*float64(termW)) - dragon.Width/2
		ds.cellY = int(dragon.RelY*float64(termH)) - dragon.Height/2
		ds.dirX = 0
		ds.dirY = 0
	}
}

// UpdateDragons drives movement AI for all three dragons.
// Call once per game tick from adventure.go updateStates().
func UpdateDragons(termW, termH int) {
	touchDragonIfPlayerHits(GreenDragon, &GreenDS, termW, termH)
	touchDragonIfPlayerHits(YellowDragon, &YellowDS, termW, termH)
	touchDragonIfPlayerHits(RedDragon, &RedDS, termW, termH)
	killDragonIfSwordHits(GreenDragon, &GreenDS, termW, termH)
	killDragonIfSwordHits(YellowDragon, &YellowDS, termW, termH)
	killDragonIfSwordHits(RedDragon, &RedDS, termW, termH)
	moveDragon(GreenDragon, &GreenDS, greenDragonMatrix(), 4, termW, termH)  // slow
	moveDragon(YellowDragon, &YellowDS, yellowDragonMatrix(), 3, termW, termH) // medium
	moveDragon(RedDragon, &RedDS, redDragonMatrix(), 2, termW, termH)          // fast
}

func InitBat(w, h int) {
	// C++ V1: room 0x1A (RoomWhiteCastleEntry), no movement
	//         V2: room 0x02 (RoomBelowYellowCastle), movementY=-3
	batRoom := &world.RoomBelowYellowCastle
	if G.GameType == 1 {
		batRoom = &world.RoomWhiteCastleEntry
	}
	frames := [][]*world.Cell{world.BatGfx, world.BatGfxOpen}
	Bat = &Object{
		RelX: 0.20, RelY: 0.25, Width: 8, Height: 6,
		ZLayer:       1,
		Carryable:    true,
		Room:         batRoom,
		Style:        tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 12, // ~5 Hz wing flap — visually comfortable in terminal
	}
	AllObjects = append(AllObjects, Bat)
	// Reset bat AI state on every init/reset.
	// C++ V2 table: OBJECT_BAT, room=0x02, movementX=0, movementY=-3
	// movementY=-3 in Atari coords = moving downward (low y = roomDown in Atari).
	// In terminal rows (y increases downward) that maps to batDirY=+1.
	batFedUpTimer = 0xff
	batHuntDelay = 0
	BatCarrying = nil
	batDirX = 0
	if G.GameType == 1 {
		batDirY = 0 // V1: bat starts stationary
	} else {
		batDirY = 1 // V2/V3: bat starts moving immediately
	}
	batTick = 0
	batCellX = int(0.20*float64(w)) - Bat.Width/2
	batCellY = int(0.25*float64(h)) - Bat.Height/2
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
	bg := tcell.NewRGBColor(0xcd, 0xcd, 0xcd)
	makePort := func(room *world.Room) *Object {
		style := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(bg)
		o := &Object{
			RelX: relX, RelY: relY,
			Width: doorWidth, Height: portHeight,
			ZLayer:       2,
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
	// C++ V1: room 0x04 (RoomBlueMazeTop) / V2: room 0x0B (RoomMazeSide)
	room := &world.RoomMazeSide
	if G.GameType == 1 {
		room = &world.RoomBlueMazeTop
	}
	Bridge = &Object{
		RelX: 0.40, RelY: 0.50, Width: 10, Height: 12,
		ZLayer: 0,
		Carryable: true,
		Room:      room,
		Style:     tcell.StyleDefault.Foreground(tcell.NewRGBColor(0x99, 0x00, 0xCC)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:  world.BridgeGfx,
	}
	AllObjects = append(AllObjects, Bridge)
}

func InitSword(w, h int) {
	// C++ V1: room 0x12 (RoomAboveYellowCastle) / V2: room 0x11 (RoomYellowCastle)
	room := &world.RoomYellowCastle
	if G.GameType == 1 {
		room = &world.RoomAboveYellowCastle
	}
	frames := [][]*world.Cell{world.SwordGfx, world.SwordGfxLeft, world.SwordGfxUp, world.SwordGfxDown}
	Sword = &Object{
		RelX: 0.20, RelY: 0.25, Width: 8, Height: 4,
		ZLayer:       1,
		Carryable:    true,
		Room:         room,
		Style:        tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xF5, 0xCE, 0x42)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:        frames[0],
		Frames:       frames,
		AnimInterval: 60,
		Paused:       true,
	}
	AllObjects = append(AllObjects, Sword)
}

func InitChalice(w, h int) {
	// C++ V1: room 0x1C (RoomOtherPurpleRoom) / V2: room 0x14 (RoomBlackMaze2)
	room := &world.RoomBlackMaze2
	if G.GameType == 1 {
		room = &world.RoomOtherPurpleRoom
	}
	Chalice = &Object{
		RelX: 0.30, RelY: 0.25, Width: 8, Height: 5,
		ZLayer:    1,
		Carryable: true,
		Room:      room,
		Style:     tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xFF, 0xAA, 0x00)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:     world.ChaliceGfx,
	}
	AllObjects = append(AllObjects, Chalice)
}

func InitMagnet(w, h int) {
	// C++ V1: room 0x1B (RoomBlackCastleEntry) / V2: room 0x0E (RoomDeadEndCyan)
	room := &world.RoomDeadEndCyan
	if G.GameType == 1 {
		room = &world.RoomBlackCastleEntry
	}
	frames := world.MakeMagnetFrames()
	Magnet = &Object{
		RelX: 0.80, RelY: 0.25, Width: 12, Height: 8,
		Carryable:        true,
		Room:             room,
		Style:            tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape:            frames[0],
		Frames:           frames,
		SubFrameCount:    4,
		SubFrameInterval: 15,
		AnimInterval:     0, // 0 = no orientation rotation; field lines still animate
		BodyOffsets:      [][2]int{{6, 2}, {4, 4}, {6, 6}, {8, 4}},
	}
	AllObjects = append(AllObjects, Magnet)
	magnetTick = 0
}

// Magnet attraction state.
var magnetTick int

func magnetPriorityList() []*Object {
	return []*Object{YellowKey, WhiteKey, BlackKey, Sword, Bridge, Chalice}
}

// magnetFieldStopOffsets gives the (dx, dy) stop displacement from the magnet anchor
// per orientation (0=Down, 1=Right, 2=Up, 3=Left), in terminal cells.
// Attracted objects stop at the outer field-line arc, not on the magnet body itself.
var magnetFieldStopOffsets = [4][2]int{
	{0, +4},  // Down:  outer arc bottom-centre at local Y=6, anchor at Y=2 → +4
	{+7, 0},  // Right: outer arc right-centre at local X=11, anchor at X=4 → +7
	{0, -5},  // Up:    outer arc top-centre at local Y=1, anchor at Y=6 → -5
	{-8, 0},  // Left:  outer arc left-centre at local X=0, anchor at X=8 → -8
}

// UpdateMagnet attracts the highest-priority eligible object in the magnet's room
// toward the outer field-line arc at ±1 cell/frame.
// Only the player's CarriedObject is excluded (matches C++ original).
func UpdateMagnet(termW, termH int) {
	if Magnet == nil {
		return
	}

	// Compute stop position: magnet anchor displaced to outer arc tip for current orientation.
	anchorX := int(math.Round(Magnet.RelX * float64(termW)))
	anchorY := int(math.Round(Magnet.RelY * float64(termH)))
	orient := Magnet.OrientationFrame % 4
	targetX := anchorX + magnetFieldStopOffsets[orient][0]
	targetY := anchorY + magnetFieldStopOffsets[orient][1]

	// Tick gate runs unconditionally so rhythm is independent of object eligibility.
	magnetTick = (magnetTick + 1) % 4

	for _, obj := range magnetPriorityList() {
		if obj == nil || obj == CarriedObject || obj == BatCarrying {
			continue
		}
		if obj.Room != Magnet.Room {
			continue
		}

		// Use math.Round for center coords to prevent the same float64 truncation
		// rounding error that froze the bat: int(n/termW * termW) can give n-1.
		oCX := int(math.Round(obj.RelX * float64(termW)))
		oCY := int(math.Round(obj.RelY * float64(termH)))

		// Move object 1 cell per axis every 4 frames toward the target.
		if magnetTick == 0 {
			oLeft := oCX - obj.Width/2
			oTop := oCY - obj.Height/2
			if oCX < targetX {
				oLeft++
			} else if oCX > targetX {
				oLeft--
			}
			if oCY < targetY {
				oTop++
			} else if oCY > targetY {
				oTop--
			}
			obj.RelX = float64(oLeft+obj.Width/2) / float64(termW)
			obj.RelY = float64(oTop+obj.Height/2) / float64(termH)
		}
		break // only the highest-priority object per frame
	}
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
		RelX: 0.40, RelY: 0.917, Width: 1, Height: 1,
		Carryable: true,
		Room:      &world.RoomBlackMaze3,
		Style:     tcell.StyleDefault.Foreground(tcell.NewRGBColor(0xAA, 0xAA, 0xAA)).Background(tcell.NewRGBColor(0xcd, 0xcd, 0xcd)),
		Shape: world.DotGfx,
	}
	AllObjects = append(AllObjects, Dot)
}

// UpdateEasterEggBarrier mirrors C++ ROOMFLAG_RIGHTTHINWALL logic:
// the right barrier of RoomCorridorRight is passable only when the Dot is in that room.
func UpdateEasterEggBarrier() {
	if EasterEggBarrier == nil {
		return
	}
	if Dot != nil && Dot.Room == &world.RoomCorridorRight {
		EasterEggBarrier.Room = nil // disable — no 'X' cells drawn, wall passable
	} else {
		EasterEggBarrier.Room = &world.RoomCorridorRight // restore barrier
	}
}

// portStatesSeq mirrors C++ portStates[]: maps PortState index (0–23) to graphic state (0–6).
// 0 = fully closed, 6 = fully open.
// Sequence: opens from state 1→12, holds at 12, closes from 13→22, permanent unlock at >22.
var portStatesSeq = [24]int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 5, 5, 4, 4, 3, 3, 2, 2, 1, 1}

// portTicksPerStep: frames between each portcullis state increment.
// 12 steps × 30 frames / 60 fps ≈ 6 seconds to fully open.
const portTicksPerStep = 30

// CollisionCheckObjects returns true if the bounding boxes of a and b overlap.
func CollisionCheckObjects(a, b *Object, termW, termH int) bool {
	ax := int(a.RelX*float64(termW)) - a.Width/2
	ay := int(a.RelY*float64(termH)) - a.Height/2
	bx := int(b.RelX*float64(termW)) - b.Width/2
	by := int(b.RelY*float64(termH)) - b.Height/2
	return ax < bx+b.Width && ax+a.Width > bx && ay < by+b.Height && ay+a.Height > by
}

// UpdatePortcullis drives the portcullis open animation and permanent unlock.
// Simplified behaviour (not the original open+close cycle):
//   - Trigger: player in castle room carrying key, gate closed (PortState=0), adjacent to gate.
//   - Gate opens over 11 steps (PortState 1→11), throttled by portTicksPerStep.
//   - At PortState=12: gate permanently unlocked — stays open forever.
func UpdatePortcullis(port, key *Object, termW, termH int) {
	if port == nil || key == nil {
		return
	}

	if port.Unlocked {
		port.Solid = false
		return
	}

	// Trigger: gate must be fully closed (PortState=0).
	// Key counts as "held" if the player carries it directly, or if the player
	// carries the bat which is itself carrying the key.
	playerHasKey := CarriedObject == key || (CarriedObject == Bat && BatCarrying == key)
	if CurrentRoom == port.Room && playerHasKey && port.PortState == 0 {
		portAx := int(port.RelX*float64(termW)) - port.Width/2
		portAy := int(port.RelY*float64(termH)) - port.Height/2
		plAx := int(Player.RelX*float64(termW)) - Player.Width/2
		plAy := int(Player.RelY*float64(termH)) - Player.Height/2
		touching := plAx < portAx+port.Width && plAx+Player.Width > portAx &&
			plAy <= portAy+port.Height && plAy+Player.Height >= portAy
		if touching {
			port.PortState++
		}
	}

	// Auto-increment while opening — throttled by portTicksPerStep.
	if port.PortState != 0 {
		port.PortTick++
		if port.PortTick >= portTicksPerStep {
			port.PortTick = 0
			port.PortState++
		}
	}

	// Permanently unlock when fully open (PortState reaches 12).
	if port.PortState >= 12 {
		port.PortState = 0
		port.Unlocked = true
		port.Solid = false
		if len(port.Frames) > 0 {
			port.Shape = port.Frames[len(port.Frames)-1]
		}
		return
	}

	// Map PortState (0–11) to opening frame.
	gfxState := portStatesSeq[port.PortState]
	numFrames := len(port.Frames)
	if numFrames > 1 {
		frame := gfxState * (numFrames - 1) / 6
		port.Shape = port.Frames[frame]
	}
	port.Solid = true // solid throughout opening animation
}

// UpdateCastlePortals manages portcullis portal teleports for all three castles.
// Returns true when a room change happened (caller should call render.FillTheScreen).
//
//   - Player entering castle room from entry room (going Down) spawns below the portcullis
//     instead of at the top of the screen.
//   - Player in castle room crossing upward through open portcullis → teleport to entry room bottom.
func UpdateCastlePortals(termW, termH int) bool {
	roomChanged := false
	defer func() {
		castlePortalPrevRoom = CurrentRoom
	}()

	type castlePair struct {
		castleRoom *world.Room
		entryRoom  *world.Room
		port       *Object
	}
	pairs := []castlePair{
		{&world.RoomYellowCastle, &world.RoomAboveYellowCastle, PortcullisYellow},
		{&world.RoomWhiteCastle, &world.RoomWhiteCastleEntry, PortcullisWhite},
		{&world.RoomBlackCastle, &world.RoomBlackCastleTop, PortcullisBlack},
	}

	for _, pair := range pairs {
		port := pair.port
		if port == nil {
			continue
		}

		// Use the same portAy formula as the renderer.
		portAy := int(port.RelY*float64(termH)) - port.Height/2

		// Case 1: Player just entered castle room from entry room (going Down).
		// Reposition to just below the portcullis TOP edge (portAy+1).
		// No room change — player is already in castleRoom; just fix RelY.
		if CurrentRoom == pair.castleRoom && castlePortalPrevRoom == pair.entryRoom {
			Player.RelY = float64(portAy+1+Player.Height/2) / float64(termH)
		}

		// Case 2: Portcullis is open and player's top has reached portAy
		// → teleport to entry room bottom.
		// Simple positional check: no crossing-detection timing issues.
		if CurrentRoom == pair.castleRoom && !port.Solid {
			currTop := int(Player.RelY*float64(termH)) - Player.Height/2
			if currTop <= portAy {
				CurrentRoom = pair.entryRoom
				Player.RelY = 1.0 - float64(Player.Height/2)/float64(termH)
				roomChanged = true
			}
		}
	}
	return roomChanged
}

func ReinitOnResize(w, h int) {
	// Re-sync bat integer position for the new terminal size.
	if Bat != nil {
		batCellX = int(Bat.RelX*float64(w)) - Bat.Width/2
		batCellY = int(Bat.RelY*float64(h)) - Bat.Height/2
	}
	// Re-sync dragon integer positions for the new terminal size.
	if GreenDragon != nil {
		GreenDS.cellX = int(GreenDragon.RelX*float64(w)) - GreenDragon.Width/2
		GreenDS.cellY = int(GreenDragon.RelY*float64(h)) - GreenDragon.Height/2
	}
	if YellowDragon != nil {
		YellowDS.cellX = int(YellowDragon.RelX*float64(w)) - YellowDragon.Width/2
		YellowDS.cellY = int(YellowDragon.RelY*float64(h)) - YellowDragon.Height/2
	}
	if RedDragon != nil {
		RedDS.cellX = int(RedDragon.RelX*float64(w)) - RedDragon.Width/2
		RedDS.cellY = int(RedDragon.RelY*float64(h)) - RedDragon.Height/2
	}

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
		var frame int
		if p.Unlocked {
			frame = len(frames) - 1
		} else {
			gfxState := portStatesSeq[p.PortState]
			frame = gfxState * (len(frames) - 1) / 6
		}
		p.Shape = frames[frame]
	}
}

func InitPlayer(w, h int) {
	Player = &Object{
		RelX:    float64(w/2+1) / float64(w),
		RelY:    float64(h/3*2+1) / float64(h),
		Width:   3,
		Height:  2,
		ZLayer:  1,
		StepX:   2,
		StepY:   1,
		Style:   tcell.StyleDefault, // updated each frame by UpdatePlayerStyle()
		Shape:   world.PlayerGfx,
	}
	AllObjects = append(AllObjects, Player)
}

// SoftReset mirrors the original Atari "game reset": player returns to the
// Yellow Castle at the start position; all objects, dragons, and bat keep
// their current positions and rooms. Dragon states are cleared (dead dragons
// become alive again, eaten state is cleared). This matches the C++ original.
func SoftReset(w, h int) {
	// Player back to Yellow Castle start.
	CurrentRoom = &world.RoomYellowCastle
	if Player != nil {
		Player.RelX = float64(w/2+1) / float64(w)
		Player.RelY = float64(h/3*2+1) / float64(h)
	}
	// Drop carried object in place.
	CarriedObject = nil
	dropCooldown = 0
	// Clear eaten state.
	Eaten = false
	// Reset dragon states but keep positions/rooms.
	// Also restore graphics if a dragon was killed (State==1 truncated Frames).
	resetDragonForSoftReset(GreenDragon, &GreenDS)
	resetDragonForSoftReset(YellowDragon, &YellowDS)
	resetDragonForSoftReset(RedDragon, &RedDS)
	// Clear UI state.
	HelpMode = false
	GodMode = false
	CancelConfirm()
	Overlay.Active = false
	// Clear win state.
	GameWon = false
	WinFlashTimer = 0
	WinOverlayTimer = 0
}

func resetDragonForSoftReset(dragon *Object, ds *dragonState) {
	if dragon == nil {
		return
	}
	ds.State = 0
	ds.RoarTimer = 0
	// Restore animation frames if dragon was killed (frames were truncated).
	frames := [][]*world.Cell{world.DragonGfx, world.DragonGfxOpen}
	dragon.Shape = frames[0]
	dragon.Frames = frames
	dragon.animFrame = 0
	dragon.animTick = 0
}
