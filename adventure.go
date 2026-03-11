package main

import (
	"time"

	"development/myGoAdventure/internal/game"
	"development/myGoAdventure/internal/input"
	"development/myGoAdventure/internal/render"
)

func main() {
	render.InitScreen()
	render.InitGamestate()
	input.InitUserInput()

	frameDuration := time.Second / game.G.FPS

	render.DrawStage()

	for !isGameOver() {
		startTime := time.Now()

		updateStates()

		render.DrawStage()
		render.DrawAllVisibleObjects()
		render.Screen.Show()

		remainingTime := frameDuration - time.Since(startTime)
		if remainingTime > 0 {
			time.Sleep(remainingTime)
		}
	}
}

func updateStates() {}

func isGameOver() bool {
	return false
}
