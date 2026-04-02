# myGoAdventure

[![CI](https://github.com/jagottsicher/myGoAdventure/actions/workflows/ci.yml/badge.svg)](https://github.com/jagottsicher/myGoAdventure/actions/workflows/ci.yml)

A terminal remake of Warren Robinett's [Adventure](https://en.wikipedia.org/wiki/Adventure_(Atari_2600)) (Atari 2600, 1980) — written in Go, rendered entirely in the terminal using Unicode block characters.

No graphics. No sprites. Just `▀`, `▄`, `█` and a healthy respect for the original.

## An ADVENTURE 

ported with ♥ to Go by Jens Schendel with intense use of the awesome tcell package provided by Garrett D'Amore, and with heartfelt thanks for countless hours spent in front of my Atari 2600, dedicated to Warren Robinett. The ADVENTURE goes on and on!

---

## Features

- All 3 game variations (1: easy, 2: standard, 3: random placement)
- All rooms, dragons, objects and castle mechanics from the original
- Dragon AI with state machine: roaming, roaring, eating, fleeing
- Magnet, portcullis, bat, bridge — all behave as in the original
- Adaptive layout: rooms scale to terminal size
- Easter egg room with credit text

## Requirements

- Go 1.20+
- A terminal with Unicode and 256-color support (most modern terminals qualify)
- Recommended: at least 80×24 — larger is better

## Build & Run

```sh
go run .
```

```sh
go build -o adventure && ./adventure
```

## Controls

| Key | Action |
|-----|--------|
| Arrow keys / `W` `A` `S` `D` | Move |
| `Space` | Pick up / drop object |
| `V` | Cycle game variation (1 / 2 / 3) |
| `N` | Cycle difficulty (A / B) |
| `R` | Reset game |
| `H` | Toggle help screen |
| `Q` | Quit |

## Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| [github.com/gdamore/tcell/v3](https://github.com/gdamore/tcell) | v3.1.2 | Terminal rendering, input, color |
| [github.com/mattn/go-runewidth](https://github.com/mattn/go-runewidth) | v0.0.15 | Unicode character width handling |

Indirect dependencies pulled in by tcell: `go-colorful`, `encoding`, `uniseg`, `golang.org/x/sys`, `golang.org/x/term`, `golang.org/x/text`.

## Project Structure

```
myGoAdventure/
├── main.go              # Entry point, game loop
├── adventure.go         # Top-level update/render orchestration
├── go.mod / go.sum
├── internal/
│   ├── game/            # Game state, objects, dragons, player, input handling
│   ├── render/          # tcell rendering, coordinate mapping, overlays
│   ├── input/           # Key bindings
│   └── world/           # Room definitions, level data, graphics data
└── assets/              # Embedded data files (if any)
```

## Current State

Working:
- All 31 rooms implemented; connections verified for variation 1
- Variation 2/3 connections mostly correct — see known issues below
- Variation 3 random object placement (mirrors original C++ bounds table)
- Full dragon behavior (roam, roar, eat, flee from sword)
- V2/V3 dragons start moving immediately from game start (matches original)
- All object interactions: keys, castles, portcullis, bridge, bat, magnet, sword, chalice
- Win condition and win overlay
- Help screen and Easter egg room

Known issues / not yet implemented:
- 5 room connections wrong in V2/V3: BlackMazeEntry.Down, BlackCastleTop.Up, BlackCastleEntry.Up/Right/Left
- Easter egg dot mechanic not implemented (dot exists but carrying it has no effect yet)
- Audio: none (terminal limitation)

## References

- Original 6502 assembly: [Greg Troutman's disassembly](http://www.atariage.com/forums/topic/33233-adventure-disassembly/)
- C++ port used as primary logic reference: [AdventureRevisited](https://github.com/fenix/AdventureRevisited) by Mike Sutton
- tcell library: [github.com/gdamore/tcell](https://github.com/gdamore/tcell)
- Room map reference: Maurice Molyneaux's annotated Adventure map

---

Dedicated to Warren Robinett, whose hidden room started it all.
