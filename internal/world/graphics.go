package world

// Player graphics — 3x2: L M R columns, two rows
var PlayerGfx = []*Cell{
	{X: 0, Y: 0, Symbol: 'L'},
	{X: 1, Y: 0, Symbol: 'M'},
	{X: 2, Y: 0, Symbol: 'R'},
	{X: 0, Y: 1, Symbol: 'L'},
	{X: 1, Y: 1, Symbol: 'M'},
	{X: 2, Y: 1, Symbol: 'R'},
}

var PlayerGfxBefore = []*Cell{
	{X: 0, Y: 0, Symbol: 'B'},
	{X: 1, Y: 0, Symbol: '4'},
	{X: 2, Y: 0, Symbol: 'F'},
	{X: 0, Y: 1, Symbol: 'B'},
	{X: 1, Y: 1, Symbol: '4'},
	{X: 2, Y: 1, Symbol: 'F'},
}

// Yellow Key (8 wide x 2 terminal rows = 3 pixel rows via half-block chars)
// From objectGfxKey[] = { 3, 0x07, 0xFD, 0xA7 } in Adventure.cpp
// ▄▄▄▄▄█▀█
// ▀ ▀  ▀▀▀
var KeyGfx = []*Cell{
	// Row 0: ▄▄▄▄▄█▀█
	{X: 0, Y: 0, Symbol: '▄'},
	{X: 1, Y: 0, Symbol: '▄'},
	{X: 2, Y: 0, Symbol: '▄'},
	{X: 3, Y: 0, Symbol: '▄'},
	{X: 4, Y: 0, Symbol: '▄'},
	{X: 5, Y: 0, Symbol: '█'},
	{X: 6, Y: 0, Symbol: '▀'},
	{X: 7, Y: 0, Symbol: '█'},
	// Row 1: ▀ ▀  ▀▀▀
	{X: 0, Y: 1, Symbol: '▀'},
	{X: 2, Y: 1, Symbol: '▀'},
	{X: 5, Y: 1, Symbol: '▀'},
	{X: 6, Y: 1, Symbol: '▀'},
	{X: 7, Y: 1, Symbol: '▀'},
}

// Dragon (8 wide x 10 tall terminal rows = 20 pixel rows via half-block chars)
// State 0 from objectGfxDrag[] in Adventure.cpp
//     ▄██▄
// ████▄▄█▀
//     ▀█▀
//    ▄▄█▄
//  ▄██████
// ██▀   ██
// ██   ▄██
// ▀▀████▀▀
// ▄   █▄▄▄
// ▀▀█▄▄▄▄█
var DragonGfx = []*Cell{
	// Row 0: ▄██▄  (cols 4-7)
	{X: 4, Y: 0, Symbol: '▄'},
	{X: 5, Y: 0, Symbol: '█'},
	{X: 6, Y: 0, Symbol: '█'},
	{X: 7, Y: 0, Symbol: '▄'},
	// Row 1: ████▄▄█▀
	{X: 0, Y: 1, Symbol: '█'},
	{X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '█'},
	{X: 3, Y: 1, Symbol: '█'},
	{X: 4, Y: 1, Symbol: '▄'},
	{X: 5, Y: 1, Symbol: '▄'},
	{X: 6, Y: 1, Symbol: '█'},
	{X: 7, Y: 1, Symbol: '▀'},
	// Row 2:     ▀█▀  (cols 4-6)
	{X: 4, Y: 2, Symbol: '▀'},
	{X: 5, Y: 2, Symbol: '█'},
	{X: 6, Y: 2, Symbol: '▀'},
	// Row 3:    ▄▄█▄  (cols 3-6)
	{X: 3, Y: 3, Symbol: '▄'},
	{X: 4, Y: 3, Symbol: '▄'},
	{X: 5, Y: 3, Symbol: '█'},
	{X: 6, Y: 3, Symbol: '▄'},
	// Row 4:  ▄██████  (cols 1-7)
	{X: 1, Y: 4, Symbol: '▄'},
	{X: 2, Y: 4, Symbol: '█'},
	{X: 3, Y: 4, Symbol: '█'},
	{X: 4, Y: 4, Symbol: '█'},
	{X: 5, Y: 4, Symbol: '█'},
	{X: 6, Y: 4, Symbol: '█'},
	{X: 7, Y: 4, Symbol: '█'},
	// Row 5: ██▀   ██
	{X: 0, Y: 5, Symbol: '█'},
	{X: 1, Y: 5, Symbol: '█'},
	{X: 2, Y: 5, Symbol: '▀'},
	{X: 6, Y: 5, Symbol: '█'},
	{X: 7, Y: 5, Symbol: '█'},
	// Row 6: ██   ▄██
	{X: 0, Y: 6, Symbol: '█'},
	{X: 1, Y: 6, Symbol: '█'},
	{X: 5, Y: 6, Symbol: '▄'},
	{X: 6, Y: 6, Symbol: '█'},
	{X: 7, Y: 6, Symbol: '█'},
	// Row 7: ▀▀████▀▀
	{X: 0, Y: 7, Symbol: '▀'},
	{X: 1, Y: 7, Symbol: '▀'},
	{X: 2, Y: 7, Symbol: '█'},
	{X: 3, Y: 7, Symbol: '█'},
	{X: 4, Y: 7, Symbol: '█'},
	{X: 5, Y: 7, Symbol: '█'},
	{X: 6, Y: 7, Symbol: '▀'},
	{X: 7, Y: 7, Symbol: '▀'},
	// Row 8: ▄   █▄▄▄
	{X: 0, Y: 8, Symbol: '▄'},
	{X: 4, Y: 8, Symbol: '█'},
	{X: 5, Y: 8, Symbol: '▄'},
	{X: 6, Y: 8, Symbol: '▄'},
	{X: 7, Y: 8, Symbol: '▄'},
	// Row 9: ▀▀█▄▄▄▄█
	{X: 0, Y: 9, Symbol: '▀'},
	{X: 1, Y: 9, Symbol: '▀'},
	{X: 2, Y: 9, Symbol: '█'},
	{X: 3, Y: 9, Symbol: '▄'},
	{X: 4, Y: 9, Symbol: '▄'},
	{X: 5, Y: 9, Symbol: '▄'},
	{X: 6, Y: 9, Symbol: '▄'},
	{X: 7, Y: 9, Symbol: '█'},
}

// Bat State 03 — wings up (compact), decoded from objectGfxBat[] State 03
// 7 pixel rows → 4 terminal rows (3 pairs + 1 lone top-half row)
//
// █      █
// ██    ██
// ▀█▀██▀█▀
//  ▀▀  ▀▀
var BatGfx = []*Cell{
	// Row 0: 0x81+0x81 → █      █
	{X: 0, Y: 0, Symbol: '█'},
	{X: 7, Y: 0, Symbol: '█'},
	// Row 1: 0xC3+0xC3 → ██    ██
	{X: 0, Y: 1, Symbol: '█'},
	{X: 1, Y: 1, Symbol: '█'},
	{X: 6, Y: 1, Symbol: '█'},
	{X: 7, Y: 1, Symbol: '█'},
	// Row 2: 0xFF+0x5A → ▀█▀██▀█▀
	{X: 0, Y: 2, Symbol: '▀'},
	{X: 1, Y: 2, Symbol: '█'},
	{X: 2, Y: 2, Symbol: '▀'},
	{X: 3, Y: 2, Symbol: '█'},
	{X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '▀'},
	{X: 6, Y: 2, Symbol: '█'},
	{X: 7, Y: 2, Symbol: '▀'},
	// Row 3: 0x66 lone → ▀▀  ▀▀ (top-half only)
	{X: 1, Y: 3, Symbol: '▀'},
	{X: 2, Y: 3, Symbol: '▀'},
	{X: 5, Y: 3, Symbol: '▀'},
	{X: 6, Y: 3, Symbol: '▀'},
}

// Bat State FF — wings spread/down, decoded from objectGfxBat[] State FF
// 11 pixel rows → 6 terminal rows (5 pairs + 1 lone top-half row)
//
// ▄      ▀
// ▄      ▀
//  ▄▀██▀▄
// ▄█▀  ▀█▄
// █      █
// ▀      ▀
var BatGfxOpen = []*Cell{
	// Row 0: 0x01+0x80 → ▄      ▀
	{X: 0, Y: 0, Symbol: '▄'},
	{X: 7, Y: 0, Symbol: '▀'},
	// Row 1: 0x01+0x80 → ▄      ▀
	{X: 0, Y: 1, Symbol: '▄'},
	{X: 7, Y: 1, Symbol: '▀'},
	// Row 2: 0x3C+0x5A →  ▄▀██▀▄
	{X: 1, Y: 2, Symbol: '▄'},
	{X: 2, Y: 2, Symbol: '▀'},
	{X: 3, Y: 2, Symbol: '█'},
	{X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '▀'},
	{X: 6, Y: 2, Symbol: '▄'},
	// Row 3: 0x66+0xC3 → ▄█▀  ▀█▄
	{X: 0, Y: 3, Symbol: '▄'},
	{X: 1, Y: 3, Symbol: '█'},
	{X: 2, Y: 3, Symbol: '▀'},
	{X: 5, Y: 3, Symbol: '▀'},
	{X: 6, Y: 3, Symbol: '█'},
	{X: 7, Y: 3, Symbol: '▄'},
	// Row 4: 0x81+0x81 → █      █
	{X: 0, Y: 4, Symbol: '█'},
	{X: 7, Y: 4, Symbol: '█'},
	// Row 5: 0x81 lone → ▀      ▀ (top-half only)
	{X: 0, Y: 5, Symbol: '▀'},
	{X: 7, Y: 5, Symbol: '▀'},
}

// Dragon State 01 — mouth open, decoded from objectGfxDrag[] State 01 in Adventure.cpp
// 22 pixel rows → 11 terminal rows via half-block pairs (bit7=leftmost pixel)
//
// ▀▄
//   ▀▄▄██▄
//     █▄█▀
//   ▄▀▀█▀
// ▄▀  ▄█▄
//   ▄████▄
//  ███████
//  ███████
//   ▀███▀
// ▄▄▄▄█
// █▄▄
var DragonGfxOpen = []*Cell{
	// All Y values shifted -1 so total height matches DragonGfx (10 rows).
	// Row -1 (clipped): pixels 0x80+0x40 → ▀▄  (out of bounds, not drawn)
	{X: 0, Y: -1, Symbol: '▀'},
	{X: 1, Y: -1, Symbol: '▄'},
	// Row 0: pixels 0x26+0x1F → ▀▄▄██▄
	{X: 2, Y: 0, Symbol: '▀'},
	{X: 3, Y: 0, Symbol: '▄'},
	{X: 4, Y: 0, Symbol: '▄'},
	{X: 5, Y: 0, Symbol: '█'},
	{X: 6, Y: 0, Symbol: '█'},
	{X: 7, Y: 0, Symbol: '▄'},
	// Row 1: pixels 0x0B+0x0E → █▄█▀
	{X: 4, Y: 1, Symbol: '█'},
	{X: 5, Y: 1, Symbol: '▄'},
	{X: 6, Y: 1, Symbol: '█'},
	{X: 7, Y: 1, Symbol: '▀'},
	// Row 2: pixels 0x1E+0x24 → ▄▀▀█▀
	{X: 2, Y: 2, Symbol: '▄'},
	{X: 3, Y: 2, Symbol: '▀'},
	{X: 4, Y: 2, Symbol: '▀'},
	{X: 5, Y: 2, Symbol: '█'},
	{X: 6, Y: 2, Symbol: '▀'},
	// Row 3: pixels 0x44+0x8E → ▄▀  ▄█▄
	{X: 0, Y: 3, Symbol: '▄'},
	{X: 1, Y: 3, Symbol: '▀'},
	{X: 4, Y: 3, Symbol: '▄'},
	{X: 5, Y: 3, Symbol: '█'},
	{X: 6, Y: 3, Symbol: '▄'},
	// Rows 4–9: body from DragonGfx (rows 4–9), same Y
	// Row 4:  ▄██████
	{X: 1, Y: 4, Symbol: '▄'},
	{X: 2, Y: 4, Symbol: '█'},
	{X: 3, Y: 4, Symbol: '█'},
	{X: 4, Y: 4, Symbol: '█'},
	{X: 5, Y: 4, Symbol: '█'},
	{X: 6, Y: 4, Symbol: '█'},
	{X: 7, Y: 4, Symbol: '█'},
	// Row 5:  ██▀   ██
	{X: 0, Y: 5, Symbol: '█'},
	{X: 1, Y: 5, Symbol: '█'},
	{X: 2, Y: 5, Symbol: '▀'},
	{X: 6, Y: 5, Symbol: '█'},
	{X: 7, Y: 5, Symbol: '█'},
	// Row 6:  ██   ▄██
	{X: 0, Y: 6, Symbol: '█'},
	{X: 1, Y: 6, Symbol: '█'},
	{X: 5, Y: 6, Symbol: '▄'},
	{X: 6, Y: 6, Symbol: '█'},
	{X: 7, Y: 6, Symbol: '█'},
	// Row 7:  ▀▀████▀▀
	{X: 0, Y: 7, Symbol: '▀'},
	{X: 1, Y: 7, Symbol: '▀'},
	{X: 2, Y: 7, Symbol: '█'},
	{X: 3, Y: 7, Symbol: '█'},
	{X: 4, Y: 7, Symbol: '█'},
	{X: 5, Y: 7, Symbol: '█'},
	{X: 6, Y: 7, Symbol: '▀'},
	{X: 7, Y: 7, Symbol: '▀'},
	// Row 8:  ▄   █▄▄▄
	{X: 0, Y: 8, Symbol: '▄'},
	{X: 4, Y: 8, Symbol: '█'},
	{X: 5, Y: 8, Symbol: '▄'},
	{X: 6, Y: 8, Symbol: '▄'},
	{X: 7, Y: 8, Symbol: '▄'},
	// Row 9:  ▀▀█▄▄▄▄█
	{X: 0, Y: 9, Symbol: '▀'},
	{X: 1, Y: 9, Symbol: '▀'},
	{X: 2, Y: 9, Symbol: '█'},
	{X: 3, Y: 9, Symbol: '▄'},
	{X: 4, Y: 9, Symbol: '▄'},
	{X: 5, Y: 9, Symbol: '▄'},
	{X: 6, Y: 9, Symbol: '▄'},
	{X: 7, Y: 9, Symbol: '█'},
}

// Dragon State 02 — dead, decoded from objectGfxDrag[] State 02 in Adventure.cpp
// 17 pixel rows → 8 terminal rows (7 pairs + 1 lone top-half row)
// dragonStates[] = {0,2,0,1}: game state 1 (dead) → graphic frame index 2 → this sprite.
//
//     ██
//     ██▄
//  ▄▄██▄██
// █▀  ▀▀▀
// ██████▄
// ▀██████
//  ▀█▀▀
//  █▀ ▀▀█
//  ▀▀▀▀▀▀
var DragonGfxDead = []*Cell{
	// Row 0: 0x0C+0x0C →     ██
	{X: 4, Y: 0, Symbol: '█'},
	{X: 5, Y: 0, Symbol: '█'},
	// Row 1: 0x0C+0x0E →     ██▄
	{X: 4, Y: 1, Symbol: '█'},
	{X: 5, Y: 1, Symbol: '█'},
	{X: 6, Y: 1, Symbol: '▄'},
	// Row 2: 0x1B+0x7F →  ▄▄██▄██
	{X: 1, Y: 2, Symbol: '▄'},
	{X: 2, Y: 2, Symbol: '▄'},
	{X: 3, Y: 2, Symbol: '█'},
	{X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '▄'},
	{X: 6, Y: 2, Symbol: '█'},
	{X: 7, Y: 2, Symbol: '█'},
	// Row 3: 0xCE+0x80 → █▀  ▀▀▀
	{X: 0, Y: 3, Symbol: '█'},
	{X: 1, Y: 3, Symbol: '▀'},
	{X: 4, Y: 3, Symbol: '▀'},
	{X: 5, Y: 3, Symbol: '▀'},
	{X: 6, Y: 3, Symbol: '▀'},
	// Row 4: 0xFC+0xFE → ██████▄
	{X: 0, Y: 4, Symbol: '█'},
	{X: 1, Y: 4, Symbol: '█'},
	{X: 2, Y: 4, Symbol: '█'},
	{X: 3, Y: 4, Symbol: '█'},
	{X: 4, Y: 4, Symbol: '█'},
	{X: 5, Y: 4, Symbol: '█'},
	{X: 6, Y: 4, Symbol: '▄'},
	// Row 5: 0xFE+0x7E → ▀██████
	{X: 0, Y: 5, Symbol: '▀'},
	{X: 1, Y: 5, Symbol: '█'},
	{X: 2, Y: 5, Symbol: '█'},
	{X: 3, Y: 5, Symbol: '█'},
	{X: 4, Y: 5, Symbol: '█'},
	{X: 5, Y: 5, Symbol: '█'},
	{X: 6, Y: 5, Symbol: '█'},
	// Row 6: 0x78+0x20 →  ▀█▀▀
	{X: 1, Y: 6, Symbol: '▀'},
	{X: 2, Y: 6, Symbol: '█'},
	{X: 3, Y: 6, Symbol: '▀'},
	{X: 4, Y: 6, Symbol: '▀'},
	// Row 7: 0x6E+0x42 →  █▀ ▀▀█
	{X: 1, Y: 7, Symbol: '█'},
	{X: 2, Y: 7, Symbol: '▀'},
	{X: 4, Y: 7, Symbol: '▀'},
	{X: 5, Y: 7, Symbol: '▀'},
	{X: 6, Y: 7, Symbol: '█'},
	// Row 8: 0x7E lone →  ▀▀▀▀▀▀  (top half only)
	{X: 1, Y: 8, Symbol: '▀'},
	{X: 2, Y: 8, Symbol: '▀'},
	{X: 3, Y: 8, Symbol: '▀'},
	{X: 4, Y: 8, Symbol: '▀'},
	{X: 5, Y: 8, Symbol: '▀'},
	{X: 6, Y: 8, Symbol: '▀'},
}

// Bridge — objectGfxBridge, 24 pixel rows → 12 terminal rows
// ██      ██  (top caps ×2)
//  █      █   (pillars ×8)
// ██      ██  (bottom caps ×2)
var BridgeGfx = []*Cell{
	// Rows 0-1: top caps → ██      ██
	{X: 0, Y: 0, Symbol: '█'}, {X: 1, Y: 0, Symbol: '█'}, {X: 8, Y: 0, Symbol: '█'}, {X: 9, Y: 0, Symbol: '█'},
	{X: 0, Y: 1, Symbol: '█'}, {X: 1, Y: 1, Symbol: '█'}, {X: 8, Y: 1, Symbol: '█'}, {X: 9, Y: 1, Symbol: '█'},
	// Rows 2-9: pillars →  █      █
	{X: 1, Y: 2, Symbol: '█'}, {X: 8, Y: 2, Symbol: '█'},
	{X: 1, Y: 3, Symbol: '█'}, {X: 8, Y: 3, Symbol: '█'},
	{X: 1, Y: 4, Symbol: '█'}, {X: 8, Y: 4, Symbol: '█'},
	{X: 1, Y: 5, Symbol: '█'}, {X: 8, Y: 5, Symbol: '█'},
	{X: 1, Y: 6, Symbol: '█'}, {X: 8, Y: 6, Symbol: '█'},
	{X: 1, Y: 7, Symbol: '█'}, {X: 8, Y: 7, Symbol: '█'},
	{X: 1, Y: 8, Symbol: '█'}, {X: 8, Y: 8, Symbol: '█'},
	{X: 1, Y: 9, Symbol: '█'}, {X: 8, Y: 9, Symbol: '█'},
	// Rows 10-11: bottom caps → ██      ██
	{X: 0, Y: 10, Symbol: '█'}, {X: 1, Y: 10, Symbol: '█'}, {X: 8, Y: 10, Symbol: '█'}, {X: 9, Y: 10, Symbol: '█'},
	{X: 0, Y: 11, Symbol: '█'}, {X: 1, Y: 11, Symbol: '█'}, {X: 8, Y: 11, Symbol: '█'}, {X: 9, Y: 11, Symbol: '█'},
}

// Sword — objectGfxSword, 5 pixel rows → 3 terminal rows
//  ▄▀
// ▀█▀▀▀▀▀▀
//   ▀
var SwordGfx = []*Cell{
	// Row 1: guard tip (shifted down 1 to center in Height=4 box)
	{X: 1, Y: 1, Symbol: '▄'}, {X: 2, Y: 1, Symbol: '▀'},
	// Row 2: blade (center row)
	{X: 0, Y: 2, Symbol: '▀'}, {X: 1, Y: 2, Symbol: '█'},
	{X: 2, Y: 2, Symbol: '▀'}, {X: 3, Y: 2, Symbol: '▀'},
	{X: 4, Y: 2, Symbol: '▀'}, {X: 5, Y: 2, Symbol: '▀'},
	{X: 6, Y: 2, Symbol: '▀'}, {X: 7, Y: 2, Symbol: '▀'},
	// Row 3: pommel
	{X: 2, Y: 3, Symbol: '▀'},
}

// SwordGfxLeft — tip pointing left (horizontal mirror of SwordGfx)
//      ▀▄
// ▀▀▀▀▀▀█▀
//      ▀
var SwordGfxLeft = []*Cell{
	// Row 1: guard tip (shifted down 1)
	{X: 5, Y: 1, Symbol: '▀'}, {X: 6, Y: 1, Symbol: '▄'},
	// Row 2: blade (center row)
	{X: 0, Y: 2, Symbol: '▀'}, {X: 1, Y: 2, Symbol: '▀'},
	{X: 2, Y: 2, Symbol: '▀'}, {X: 3, Y: 2, Symbol: '▀'},
	{X: 4, Y: 2, Symbol: '▀'}, {X: 5, Y: 2, Symbol: '▀'},
	{X: 6, Y: 2, Symbol: '█'}, {X: 7, Y: 2, Symbol: '▀'},
	// Row 3: pommel
	{X: 5, Y: 3, Symbol: '▀'},
}

// SwordGfxUp — tip pointing up (90° CCW rotation of SwordGfx, centered in 8-wide box)
// Pixel grid 8px tall × 5px wide, packed into 4 terminal rows × 8 cols (X offset +2)
//     █
//     █
//   ▄ █ ▄
//   ▀█▀
var SwordGfxUp = []*Cell{
	// Row 0: tip (blade end)
	{X: 4, Y: 0, Symbol: '█'},
	// Row 1: blade
	{X: 4, Y: 1, Symbol: '█'},
	// Row 2: guard crosspiece
	{X: 2, Y: 2, Symbol: '▄'}, {X: 4, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '▄'},
	// Row 3: handle/pommel
	{X: 3, Y: 3, Symbol: '▀'}, {X: 4, Y: 3, Symbol: '█'}, {X: 5, Y: 3, Symbol: '▀'},
}

// SwordGfxDown — tip pointing down (vertical flip of SwordGfxUp, ▄↔▀)
//   ▄█▄
//   ▀ █ ▀
//     █
//     █
var SwordGfxDown = []*Cell{
	// Row 0: handle/pommel (flipped from row 3 of Up)
	{X: 3, Y: 0, Symbol: '▄'}, {X: 4, Y: 0, Symbol: '█'}, {X: 5, Y: 0, Symbol: '▄'},
	// Row 1: guard crosspiece (flipped from row 2 of Up)
	{X: 2, Y: 1, Symbol: '▀'}, {X: 4, Y: 1, Symbol: '█'}, {X: 6, Y: 1, Symbol: '▀'},
	// Row 2: blade
	{X: 4, Y: 2, Symbol: '█'},
	// Row 3: tip (blade end)
	{X: 4, Y: 3, Symbol: '█'},
}

// Chalice — objectGfxChallise, 9 pixel rows → 5 terminal rows
// █      █
// ▀█▄▄▄▄█▀
//  ▀████▀
//    ██
//  ▀▀▀▀▀▀
var ChaliceGfx = []*Cell{
	// Row 0: 0x81+0x81 → █      █
	{X: 0, Y: 0, Symbol: '█'}, {X: 7, Y: 0, Symbol: '█'},
	// Row 1: 0xC3+0x7E → ▀█▄▄▄▄█▀
	{X: 0, Y: 1, Symbol: '▀'}, {X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '▄'}, {X: 3, Y: 1, Symbol: '▄'},
	{X: 4, Y: 1, Symbol: '▄'}, {X: 5, Y: 1, Symbol: '▄'},
	{X: 6, Y: 1, Symbol: '█'}, {X: 7, Y: 1, Symbol: '▀'},
	// Row 2: 0x7E+0x3C →  ▀████▀
	{X: 1, Y: 2, Symbol: '▀'}, {X: 2, Y: 2, Symbol: '█'},
	{X: 3, Y: 2, Symbol: '█'}, {X: 4, Y: 2, Symbol: '█'},
	{X: 5, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '▀'},
	// Row 3: 0x18+0x18 →    ██
	{X: 3, Y: 3, Symbol: '█'}, {X: 4, Y: 3, Symbol: '█'},
	// Row 4: 0x7E lone →  ▀▀▀▀▀▀ (top-half only)
	{X: 1, Y: 4, Symbol: '▀'}, {X: 2, Y: 4, Symbol: '▀'},
	{X: 3, Y: 4, Symbol: '▀'}, {X: 4, Y: 4, Symbol: '▀'},
	{X: 5, Y: 4, Symbol: '▀'}, {X: 6, Y: 4, Symbol: '▀'},
}

// Magnet — objectGfxMagnet, 8 pixel rows → 4 terminal rows
//  ▄████▄
// ██▀  ▀██
// ██    ██
// ██    ██
var MagnetGfx = []*Cell{
	// Row 0: 0x3C+0x7E →  ▄████▄
	{X: 1, Y: 0, Symbol: '▄'}, {X: 2, Y: 0, Symbol: '█'},
	{X: 3, Y: 0, Symbol: '█'}, {X: 4, Y: 0, Symbol: '█'},
	{X: 5, Y: 0, Symbol: '█'}, {X: 6, Y: 0, Symbol: '▄'},
	// Row 1: 0xE7+0xC3 → ██▀  ▀██
	{X: 0, Y: 1, Symbol: '█'}, {X: 1, Y: 1, Symbol: '█'},
	{X: 2, Y: 1, Symbol: '▀'}, {X: 5, Y: 1, Symbol: '▀'},
	{X: 6, Y: 1, Symbol: '█'}, {X: 7, Y: 1, Symbol: '█'},
	// Rows 2-3: 0xC3+0xC3 → ██    ██
	{X: 0, Y: 2, Symbol: '█'}, {X: 1, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '█'}, {X: 7, Y: 2, Symbol: '█'},
	{X: 0, Y: 3, Symbol: '█'}, {X: 1, Y: 3, Symbol: '█'}, {X: 6, Y: 3, Symbol: '█'}, {X: 7, Y: 3, Symbol: '█'},
}

// MakeMagnetFrames returns 16 frames for the magnet (Width=12, Height=8).
//
// Layout: Frames[orientation*4 + fieldPhase]
//   orientation 0=Down, 1=Right, 2=Up, 3=Left
//   fieldPhase  0=no arc, 1=inner, 2=medium, 3=outer
//
// Orientation and field-line phase are driven by independent timers (SubFrameCount=4).
// Arcs use box-drawing chars so they pick up the background color of whatever they overlap.
//
// Down  (poles ↓, arch ↑):  arcs below  with ╰─╯  (arch top, poles+arcs bottom)
// Right (poles →, arch ←):  arcs right  with ─╮╯  (arch left, poles+arcs right)
// Up    (poles ↑, arch ↓):  arcs above  with ╭─╮  (arch bottom, poles+arcs top)
// Left  (poles ←, arch →):  arcs left   with ╭╰─  (arch right, poles+arcs left)
func MakeMagnetFrames() [][]*Cell {
	cp := func(src []*Cell) []*Cell {
		dst := make([]*Cell, len(src))
		copy(dst, src)
		return dst
	}
	join := func(a, b []*Cell) []*Cell { return append(cp(a), b...) }

	// ── DOWN (poles down, arch up) ────────────────────────────────────────────
	// Body at Y=0–3, shifted X+2. Poles: X=2–3 left, X=8–9 right.
	dBody := []*Cell{
		{X: 3, Y: 0, Symbol: '▄'}, {X: 4, Y: 0, Symbol: '█'}, {X: 5, Y: 0, Symbol: '█'},
		{X: 6, Y: 0, Symbol: '█'}, {X: 7, Y: 0, Symbol: '█'}, {X: 8, Y: 0, Symbol: '▄'},
		{X: 2, Y: 1, Symbol: '█'}, {X: 3, Y: 1, Symbol: '█'}, {X: 4, Y: 1, Symbol: '▀'},
		{X: 7, Y: 1, Symbol: '▀'}, {X: 8, Y: 1, Symbol: '█'}, {X: 9, Y: 1, Symbol: '█'},
		{X: 2, Y: 2, Symbol: '█'}, {X: 3, Y: 2, Symbol: '█'}, {X: 8, Y: 2, Symbol: '█'}, {X: 9, Y: 2, Symbol: '█'},
		{X: 2, Y: 3, Symbol: '█'}, {X: 3, Y: 3, Symbol: '█'}, {X: 8, Y: 3, Symbol: '█'}, {X: 9, Y: 3, Symbol: '█'},
	}
	d0 := cp(dBody)
	d1 := join(dBody, []*Cell{ // inner: tight 4-wide arc in center, 1 row
		{X: 4, Y: 4, Symbol: '╰'}, {X: 5, Y: 4, Symbol: '─'},
		{X: 6, Y: 4, Symbol: '─'}, {X: 7, Y: 4, Symbol: '╯'},
	})
	d2 := join(dBody, []*Cell{ // medium: │ at poles, 6-wide arc
		{X: 3, Y: 4, Symbol: '│'}, {X: 8, Y: 4, Symbol: '│'},
		{X: 3, Y: 5, Symbol: '╰'}, {X: 4, Y: 5, Symbol: '─'}, {X: 5, Y: 5, Symbol: '─'},
		{X: 6, Y: 5, Symbol: '─'}, {X: 7, Y: 5, Symbol: '─'}, {X: 8, Y: 5, Symbol: '╯'},
	})
	d3 := join(dBody, []*Cell{ // outer: 2×│ at poles, 10-wide arc
		{X: 1, Y: 4, Symbol: '│'}, {X: 10, Y: 4, Symbol: '│'},
		{X: 1, Y: 5, Symbol: '│'}, {X: 10, Y: 5, Symbol: '│'},
		{X: 1, Y: 6, Symbol: '╰'}, {X: 2, Y: 6, Symbol: '─'}, {X: 3, Y: 6, Symbol: '─'},
		{X: 4, Y: 6, Symbol: '─'}, {X: 5, Y: 6, Symbol: '─'}, {X: 6, Y: 6, Symbol: '─'},
		{X: 7, Y: 6, Symbol: '─'}, {X: 8, Y: 6, Symbol: '─'}, {X: 9, Y: 6, Symbol: '─'},
		{X: 10, Y: 6, Symbol: '╯'},
	})

	// ── RIGHT (arch left, poles right, field lines expand rightward) ──────────
	// Exact horizontal mirror of original rBody (arch was right, now left).
	// back=X=0, corner=X=1 (▄/█/█/▀), transition=X=2 (█/▀/▄/█), rails=X=2–7. X=8–11 free.
	rBody := []*Cell{
		{X: 1, Y: 2, Symbol: '▄'}, {X: 2, Y: 2, Symbol: '█'}, {X: 3, Y: 2, Symbol: '█'},
		{X: 4, Y: 2, Symbol: '█'}, {X: 5, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '█'}, {X: 7, Y: 2, Symbol: '█'},
		{X: 0, Y: 3, Symbol: '█'}, {X: 1, Y: 3, Symbol: '█'}, {X: 2, Y: 3, Symbol: '▀'},
		{X: 0, Y: 4, Symbol: '█'}, {X: 1, Y: 4, Symbol: '█'}, {X: 2, Y: 4, Symbol: '▄'},
		{X: 1, Y: 5, Symbol: '▀'}, {X: 2, Y: 5, Symbol: '█'}, {X: 3, Y: 5, Symbol: '█'},
		{X: 4, Y: 5, Symbol: '█'}, {X: 5, Y: 5, Symbol: '█'}, {X: 6, Y: 5, Symbol: '█'}, {X: 7, Y: 5, Symbol: '█'},
	}
	r0 := cp(rBody)
	r1 := join(rBody, []*Cell{ // inner: tight arc just past pole tips at X=7
		{X: 8, Y: 3, Symbol: '─'}, {X: 9, Y: 3, Symbol: '╮'},
		{X: 8, Y: 4, Symbol: '─'}, {X: 9, Y: 4, Symbol: '╯'},
	})
	r2 := join(rBody, []*Cell{ // medium: arc spanning pole rows + 1 above/below
		{X: 8, Y: 2, Symbol: '─'}, {X: 9, Y: 2, Symbol: '─'}, {X: 10, Y: 2, Symbol: '╮'},
		{X: 10, Y: 3, Symbol: '│'},
		{X: 10, Y: 4, Symbol: '│'},
		{X: 8, Y: 5, Symbol: '─'}, {X: 9, Y: 5, Symbol: '─'}, {X: 10, Y: 5, Symbol: '╯'},
	})
	r3 := join(rBody, []*Cell{ // outer: arc spanning 2 rows beyond pole tips
		{X: 8, Y: 1, Symbol: '─'}, {X: 9, Y: 1, Symbol: '─'}, {X: 10, Y: 1, Symbol: '─'}, {X: 11, Y: 1, Symbol: '╮'},
		{X: 11, Y: 2, Symbol: '│'},
		{X: 11, Y: 3, Symbol: '│'},
		{X: 11, Y: 4, Symbol: '│'},
		{X: 11, Y: 5, Symbol: '│'},
		{X: 8, Y: 6, Symbol: '─'}, {X: 9, Y: 6, Symbol: '─'}, {X: 10, Y: 6, Symbol: '─'}, {X: 11, Y: 6, Symbol: '╯'},
	})

	// ── UP (poles up, arch down) ──────────────────────────────────────────────
	// Vertical flip of Down body (Y→7-Y, ▄↔▀). Body at Y=4–7. Poles: X=2–3,8–9 at Y=4–5.
	uBody := []*Cell{
		{X: 2, Y: 4, Symbol: '█'}, {X: 3, Y: 4, Symbol: '█'}, {X: 8, Y: 4, Symbol: '█'}, {X: 9, Y: 4, Symbol: '█'},
		{X: 2, Y: 5, Symbol: '█'}, {X: 3, Y: 5, Symbol: '█'}, {X: 8, Y: 5, Symbol: '█'}, {X: 9, Y: 5, Symbol: '█'},
		{X: 2, Y: 6, Symbol: '█'}, {X: 3, Y: 6, Symbol: '█'}, {X: 4, Y: 6, Symbol: '▄'},
		{X: 7, Y: 6, Symbol: '▄'}, {X: 8, Y: 6, Symbol: '█'}, {X: 9, Y: 6, Symbol: '█'},
		{X: 3, Y: 7, Symbol: '▀'}, {X: 4, Y: 7, Symbol: '█'}, {X: 5, Y: 7, Symbol: '█'},
		{X: 6, Y: 7, Symbol: '█'}, {X: 7, Y: 7, Symbol: '█'}, {X: 8, Y: 7, Symbol: '▀'},
	}
	u0 := cp(uBody)
	u1 := join(uBody, []*Cell{ // inner: tight 4-wide arc in center, 1 row
		{X: 4, Y: 3, Symbol: '╭'}, {X: 5, Y: 3, Symbol: '─'},
		{X: 6, Y: 3, Symbol: '─'}, {X: 7, Y: 3, Symbol: '╮'},
	})
	u2 := join(uBody, []*Cell{ // medium: │ at poles, 6-wide arc
		{X: 3, Y: 3, Symbol: '│'}, {X: 8, Y: 3, Symbol: '│'},
		{X: 3, Y: 2, Symbol: '╭'}, {X: 4, Y: 2, Symbol: '─'}, {X: 5, Y: 2, Symbol: '─'},
		{X: 6, Y: 2, Symbol: '─'}, {X: 7, Y: 2, Symbol: '─'}, {X: 8, Y: 2, Symbol: '╮'},
	})
	u3 := join(uBody, []*Cell{ // outer: 2×│ at poles, 10-wide arc
		{X: 1, Y: 3, Symbol: '│'}, {X: 10, Y: 3, Symbol: '│'},
		{X: 1, Y: 2, Symbol: '│'}, {X: 10, Y: 2, Symbol: '│'},
		{X: 1, Y: 1, Symbol: '╭'}, {X: 2, Y: 1, Symbol: '─'}, {X: 3, Y: 1, Symbol: '─'},
		{X: 4, Y: 1, Symbol: '─'}, {X: 5, Y: 1, Symbol: '─'}, {X: 6, Y: 1, Symbol: '─'},
		{X: 7, Y: 1, Symbol: '─'}, {X: 8, Y: 1, Symbol: '─'}, {X: 9, Y: 1, Symbol: '─'},
		{X: 10, Y: 1, Symbol: '╮'},
	})

	// ── LEFT (arch right, poles left, field lines expand leftward) ───────────
	// Exact horizontal mirror of original lBody (arch was left, now right).
	// rails=X=4–9, transition=X=9 (█/▀/▄/█), corner=X=10 (▄/█/█/▀), back=X=11. X=0–3 free.
	lBody := []*Cell{
		{X: 4, Y: 2, Symbol: '█'}, {X: 5, Y: 2, Symbol: '█'}, {X: 6, Y: 2, Symbol: '█'},
		{X: 7, Y: 2, Symbol: '█'}, {X: 8, Y: 2, Symbol: '█'}, {X: 9, Y: 2, Symbol: '█'}, {X: 10, Y: 2, Symbol: '▄'},
		{X: 9, Y: 3, Symbol: '▀'}, {X: 10, Y: 3, Symbol: '█'}, {X: 11, Y: 3, Symbol: '█'},
		{X: 9, Y: 4, Symbol: '▄'}, {X: 10, Y: 4, Symbol: '█'}, {X: 11, Y: 4, Symbol: '█'},
		{X: 4, Y: 5, Symbol: '█'}, {X: 5, Y: 5, Symbol: '█'}, {X: 6, Y: 5, Symbol: '█'},
		{X: 7, Y: 5, Symbol: '█'}, {X: 8, Y: 5, Symbol: '█'}, {X: 9, Y: 5, Symbol: '█'}, {X: 10, Y: 5, Symbol: '▀'},
	}
	l0 := cp(lBody)
	l1 := join(lBody, []*Cell{ // inner: arc ends at X=3 (just before rail start at X=4)
		{X: 2, Y: 3, Symbol: '╭'}, {X: 3, Y: 3, Symbol: '─'},
		{X: 2, Y: 4, Symbol: '╰'}, {X: 3, Y: 4, Symbol: '─'},
	})
	l2 := join(lBody, []*Cell{ // medium
		{X: 1, Y: 2, Symbol: '╭'}, {X: 2, Y: 2, Symbol: '─'}, {X: 3, Y: 2, Symbol: '─'},
		{X: 1, Y: 3, Symbol: '│'},
		{X: 1, Y: 4, Symbol: '│'},
		{X: 1, Y: 5, Symbol: '╰'}, {X: 2, Y: 5, Symbol: '─'}, {X: 3, Y: 5, Symbol: '─'},
	})
	l3 := join(lBody, []*Cell{ // outer: arc spanning 2 rows beyond pole tips
		{X: 0, Y: 1, Symbol: '╭'}, {X: 1, Y: 1, Symbol: '─'}, {X: 2, Y: 1, Symbol: '─'}, {X: 3, Y: 1, Symbol: '─'},
		{X: 0, Y: 2, Symbol: '│'},
		{X: 0, Y: 3, Symbol: '│'},
		{X: 0, Y: 4, Symbol: '│'},
		{X: 0, Y: 5, Symbol: '│'},
		{X: 0, Y: 6, Symbol: '╰'}, {X: 1, Y: 6, Symbol: '─'}, {X: 2, Y: 6, Symbol: '─'}, {X: 3, Y: 6, Symbol: '─'},
	})

	// Frames[orientation*4 + fieldPhase]  orientation: 0=Down 1=Right 2=Up 3=Left
	// Right slot: arch LEFT (X=0–2), poles+arcs expand RIGHT.
	// Left  slot: arch RIGHT (X=9–11), poles+arcs expand LEFT.
	return [][]*Cell{d0, d1, d2, d3, r0, r1, r2, r3, u0, u1, u2, u3, l0, l1, l2, l3}
}

// MakeBarrierGfx returns a vertical column of 'X' with the given height.
// Placed as an Object, these act as invisible passage blockers:
// collision detection checks for 'X' on screen, so they block movement automatically.
func MakeBarrierGfx(height int) []*Cell {
	cells := make([]*Cell, height)
	for i := 0; i < height; i++ {
		cells[i] = &Cell{X: 0, Y: i, Symbol: 'X'}
	}
	return cells
}

// Dot — objectGfxDot, 1 pixel row → 1 terminal row (top-half only)
var DotGfx = []*Cell{
	{X: 0, Y: 0, Symbol: '▀'},
}

// MakePortcullisGfx generates a portcullis grating of '╋' characters with the given height.
func MakePortcullisGfx(width, height int) []*Cell {
	var cells []*Cell
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cells = append(cells, &Cell{X: x, Y: y, Symbol: '╋'})
		}
	}
	return cells
}

// MakePortcullisFrames returns 4 animation frames: fully closed → 3 opening steps.
// Each frame the gate rises: heights are H, 2/3·H, 1/3·H, 1 row.
func MakePortcullisFrames(width, height int) [][]*Cell {
	h2 := height * 2 / 3
	if h2 < 1 {
		h2 = 1
	}
	h3 := height / 3
	if h3 < 1 {
		h3 = 1
	}
	return [][]*Cell{
		MakePortcullisGfx(width, height), // closed
		MakePortcullisGfx(width, h2),     // 1st step
		MakePortcullisGfx(width, h3),     // 2nd step
		MakePortcullisGfx(width, 1),      // open — 1 row at top
	}
}

// Castle
var RoomGfxCastle = &[]string{
	"XXXXXXXXXXX X X X      X X X XXXXXXXXXXX",
	"X         X X X X      X X X X         X",
	"X         XXXXXXX      XXXXXXX         X",
	"X         XXXXXXXXXXXXXXXXXXXX         X",
	"X           XXXXXXXXXXXXXXXX           X",
	"X           XXXXXX    XXXXXX           X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Left of Name room: solid top wall, open sides, opening at bottom
var RoomCorridorRightGfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Room below yellow castle: opening at top, open sides, solid bottom
var RoomBelowYellowCastleGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Room above yellow castle
var RoomAboveYellowCastleGfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"X                                      X",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze Entry: complex maze walls, opening at bottom center
var RoomBlueMazeEntryGfx = &[]string{
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"XXXX  XX  XX  XXXXXXXXXXXX  XX  XX  XXXX",
	"      XX  XX                XX  XX      ",
	"      XX  XX                XX  XX      ",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"XXXXXXXX  XXXXXXXXXXXXXXXXXXXX  XXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze #1 (Room 5 in C++)
var RoomBlueMaze1Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXXXXXXXX  XXXXXXXXXXXXXXXX  XXXXXXXXXX",
	"XXXX              XXXX              XXXX",
	"XXXX              XXXX              XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"XXXX  XXXXXXXXXX  XXXX  XXXXXXXXXX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"      XX      XX  XXXX  XX      XX      ",
	"XXXXXXXX  XX  XX  XXXX  XX  XX  XXXXXXXX",
}

// Red Maze #1 (Room 0x17 in C++)
var RoomRedMaze1Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"              XX        XX              ",
	"              XX        XX              ",
	"XXXXXXXXXXXX  XX        XX  XXXXXXXXXXXX",
	"XXXXXXXXXXXX  XX        XX  XXXXXXXXXXXX",
	"XXXX      XX  XX  XXXX  XX  XX      XXXX",
	"XXXX      XX  XX  XXXX  XX  XX      XXXX",
	"XXXX  XX  XXXXXX  XXXX  XXXXXX  XX  XXXX",
}

// Bottom of Red Maze (Room 0x19 in C++)
var RoomRedMazeBottomGfx = &[]string{
	"XXXX  XX  XXXXXX  XXXX  XXXXXX  XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"      XX  XX                XX  XX      ",
	"      XX  XX                XX  XX      ",
	"XXXXXXXXXXXX                XXXXXXXXXXXX",
	"XXXXXXXXXXXX                XXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top of Red Maze (Room 0x18 in C++)
var RoomRedMazeTopGfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XXXX                  ",
	"                  XXXX                  ",
	"XXXXXXXXXXXXXXXX  XXXX  XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX  XXXX  XXXXXXXXXXXXXXXX",
	"              XX  XXXX  XX              ",
	"              XX  XXXX  XX              ",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XXXXXXXXXXXXXXXXXXXX  XX  XXXX",
	"XXXX  XX  XX                XX  XX  XXXX",
	"XXXX  XX  XX                XX  XX  XXXX",
	"XXXX  XXXXXX  XX        XX  XXXXXX  XXXX",
}

// White Castle Entry — red (Room 0x1A in C++)
var RoomWhiteCastleEntryGfx = &[]string{
	"XXXX  XXXXXX  XX        XX  XXXXXX  XXXX",
	"XXXX          XX        XX          XXXX",
	"XXXX          XX        XX          XXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX                        XX  XXXX",
	"XXXX  XX                        XX  XXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Top Entry Room — solid bottom (Room 0x0E cyan in C++)
var RoomTopEntryRoomGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Top Entry Room — opening at bottom (Room 0x1D red in C++, above Black Castle)
var RoomBlackCastleTopGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Blue Maze Bottom (Room 6 in C++)
var RoomBlueMazeBottomGfx = &[]string{
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
	"      XX      XX        XX      XX      ",
	"      XX      XX        XX      XX      ",
	"XXXX  XXXXXXXXXX        XXXXXXXXXX  XXXX",
	"XXXX  XXXXXXXXXX        XXXXXXXXXX  XXXX",
	"XXXX                                XXXX",
	"XXXX                                XXXX",
	"XXXXXXXX                        XXXXXXXX",
	"XXXXXXXX                        XXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Blue Maze Top (Room 4 in C++)
var RoomBlueMazeTopGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"        XX    XX        XX    XX        ",
	"        XX    XX        XX    XX        ",
	"XXXX    XX    XXXX    XXXX    XX    XXXX",
	"XXXX    XX    XXXX    XXXX    XX    XXXX",
	"XXXX    XX                    XX    XXXX",
	"XXXX    XX                    XX    XXXX",
	"XXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXXXX    XXXXXXXXXXXXXXXXXX",
	"      XX        XX    XX        XX      ",
	"      XX        XX    XX        XX      ",
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
}

// Blue Maze Center (Room 7 in C++)
var RoomBlueMazeCenterGfx = &[]string{
	"XXXX  XX  XXXXXXXX    XXXXXXXX  XX  XXXX",
	"      XX      XXXX    XXXX      XX      ",
	"      XX      XXXX    XXXX      XX      ",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"XXXXXXXXXXXX  XXXX    XXXX  XXXXXXXXXXXX",
	"          XX  XXXX    XXXX  XX          ",
	"          XX  XXXX    XXXX  XX          ",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"XXXX  XX  XX  XXXX    XXXX  XX  XX  XXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXXXXXX  XX  XX        XX  XX  XXXXXXXX",
}

// Black Maze #1 (Room 0x13 in C++, ROOMFLAG_NONE: right = reverse(left))
var RoomBlackMaze1Gfx = &[]string{
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"            XX            XX            ",
	"            XX            XX            ",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXX            XXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX",
	"XX    XXXXXXXXXXXXXXXXXXXXXXXXXXXX    XX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
}

// Black Maze #2 (Room 0x14 in C++, ROOMFLAG_MIRROR: right = same as left)
var RoomBlackMaze2Gfx = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"                  XX                  XX",
	"                  XX                  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"XXXXXXXXXXXXXXXX  XXXXXXXXXXXXXXXXXX  XX",
	"              XX                  XX    ",
	"              XX                  XX    ",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"XXXX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXX",
	"        XXXX      XX        XXXX      XX",
	"        XXXX      XX        XXXX      XX",
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
}

// Black Maze #3 (Room 0x15 in C++, ROOMFLAG_MIRROR: right = same as left)
var RoomBlackMaze3Gfx = &[]string{
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XX                  XX                  ",
	"XX                  XX                  ",
	"XX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXX",
	"XX    XXXXXXXXXXXXXXXX    XXXXXXXXXXXXXX",
	"      XX                  XX            ",
	"      XX                  XX            ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
	"XX          XX      XX          XX      ",
	"XX          XX      XX          XX      ",
	"XXXXXXXX    XXXXXXXXXXXXXXXX    XXXXXXXX",
}

// Black Maze Entry (Room 0x16 in C++, ROOMFLAG_NONE: right = reverse(left))
var RoomBlackMazeEntryGfx = &[]string{
	"XX  XX  XXXX  XX  XXXX  XX  XXXX  XX  XX",
	"    XX        XX  XXXX  XX        XX    ",
	"    XX        XX  XXXX  XX        XX    ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Easter Egg Room (Room 0x00): opening at top, no side walls, solid bottom
var RoomEasterEggGfx = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// Number Room (Rooms 0x12, 0x1C in C++)
var RoomGfxNumberRoom = &[]string{
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XX                                    XX",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Side Corridor (Rooms 0x0C, 0x0D in C++)
var RoomGfxSideCorridor = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"                                        ",
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
}

// Maze Entry (Room 0x0A in C++)
var RoomGfxMazeEntry = &[]string{
	"XXXXXXXXXXXXXXXX        XXXXXXXXXXXXXXXX",
	"      XX                        XX      ",
	"      XX                        XX      ",
	"XXXX  XX    XXXXXXXXXXXXXXXX    XX  XXXX",
	"XXXX  XX    XXXXXXXXXXXXXXXX    XX  XXXX",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"XXXXXXXX  XX      XXXX      XX  XXXXXXXX",
	"          XX      XXXX      XX          ",
	"          XX      XXXX      XX          ",
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
}

// Maze Middle (Room 0x09 in C++)
var RoomGfxMazeMiddle = &[]string{
	"XXXXXXXXXXXX  XX  XXXX  XX  XXXXXXXXXXXX",
	"              XX  XXXX  XX              ",
	"              XX  XXXX  XX              ",
	"XXXX      XXXXXX  XXXX  XXXXXX      XXXX",
	"XXXX      XXXXXX  XXXX  XXXXXX      XXXX",
	"          XX                XX          ",
	"          XX                XX          ",
	"XXXXXXXX  XX  XXXXXXXXXXXX  XX  XXXXXXXX",
	"XXXXXXXX  XX  XXXXXXXXXXXX  XX  XXXXXXXX",
	"      XX  XX  XX        XX  XX  XX      ",
	"      XX  XX  XX        XX  XX  XX      ",
	"XXXX  XX  XX  XX  XXXX  XX  XX  XX  XXXX",
}

// Maze Side (Room 0x0B in C++)
var RoomGfxMazeSide = &[]string{
	"XXXX  XX  XX  XX  XXXX  XX  XX  XX  XXXX",
	"      XX      XX  XXXX  XX      XX      ",
	"      XX      XX  XXXX  XX      XX      ",
	"      XXXXXXXXXX  XXXX  XXXXXXXXXX      ",
	"      XXXXXXXXXX  XXXX  XXXXXXXXXX      ",
	"                  XXXX                  ",
	"                  XXXX                  ",
	"      XXXXXXXX    XXXX    XXXXXXXX      ",
	"      XXXXXXXX    XXXX    XXXXXXXX      ",
	"      XX          XXXX          XX      ",
	"      XX          XXXX          XX      ",
	"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
}

// ConvertToBinary converts a row of room graphics to a bitmask for collision detection.
func ConvertToBinary(data string) int64 {
	binary := int64(0)
	for _, char := range data {
		binary <<= 1
		if char == 'X' {
			binary |= 1
		}
	}
	return binary
}
