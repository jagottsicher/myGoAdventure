# myGoAdventure

[![CI](https://github.com/jagottsicher/myGoAdventure/actions/workflows/ci.yml/badge.svg)](https://github.com/jagottsicher/myGoAdventure/actions/workflows/ci.yml)

A terminal remake of Warren Robinett's [Adventure](https://en.wikipedia.org/wiki/Adventure_(Atari_2600)) (Atari 2600, 1980) — written in Go, rendered entirely in the terminal using Unicode block characters.

No graphics. No sprites. Just `▀`, `▄`, `█` and a healthy respect for the original.

---

## Features

- All 3 game variations (1: easy, 2: standard, 3: random placement)
- All rooms, dragons, objects and castle mechanics from the original
- Dragon AI with state machine: roaming, roaring, eating, fleeing
- Magnet, portcullis, bat, bridge — all behave as in the original
- Adaptive layout: rooms scale to terminal size
- God mode (`G`) with debug overlay
- Easter egg room (enter your name)

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
| Arrow keys | Move |
| Space | Pick up / drop object |
| `1` / `2` / `3` | Select game variation |
| `H` | Help screen |
| `Q` / `Esc` | Quit |

## Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| [github.com/gdamore/tcell/v2](https://github.com/gdamore/tcell) | v2.7.4 | Terminal rendering, input, color |
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
- All 31 rooms with correct connections for variations 1 and 2
- Variation 3 random placement (mirrors original C++ bounds table)
- Full dragon behavior (roam, roar, eat, flee from sword)
- All object interactions: keys, castles, portcullis, bridge, bat, magnet, sword, chalice
- Win condition and win overlay
- Help screen and Easter egg room

Known issues / not yet implemented:
- Variation 3 connections not fully verified against original
- Number room (game select screen) is functional but visually minimal
- Audio: none (terminal limitation)

## References

- Original 6502 assembly: [Greg Troutman's disassembly](http://www.atariage.com/forums/topic/33233-adventure-disassembly/)
- C++ port used as primary logic reference: [AdventureRevisited](https://github.com/fenix/AdventureRevisited) by Mike Sutton
- tcell library: [github.com/gdamore/tcell](https://github.com/gdamore/tcell)
- Room map reference: Maurice Molyneaux's annotated Adventure map

---

Dedicated to Warren Robinett, whose hidden room started it all.
