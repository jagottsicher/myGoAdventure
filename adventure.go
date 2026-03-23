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
		if game.HelpMode {
			render.DrawHelp()
		} else {
			render.DrawAllVisibleObjects()
			render.DrawSpecialRooms()
		}
		if game.ConfirmMode {
			render.DrawConfirm()
		}
		render.Screen.Show()

		remainingTime := frameDuration - time.Since(startTime)
		if remainingTime > 0 {
			time.Sleep(remainingTime)
		}
	}
}

func updateStates() {
	game.GreenDragon.Animate()
	game.YellowDragon.Animate()
	game.RedDragon.Animate()
	game.Bat.Animate()
	game.UpdatePortcullis(game.PortcullisYellow, game.YellowKey)
	game.UpdatePortcullis(game.PortcullisWhite, game.WhiteKey)
	game.UpdatePortcullis(game.PortcullisBlack, game.BlackKey)
	game.PortcullisYellow.Animate()
	game.PortcullisWhite.Animate()
	game.PortcullisBlack.Animate()
	game.Sword.Animate()
	game.Magnet.Animate()
	termW, termH := render.Screen.Size()
	game.TryPickup(termW, termH)
	game.UpdateCarriedObject(termW, termH)
	game.UpdateBat(termW, termH)
	game.UpdateMagnet(termW, termH)
	facePlayer(game.GreenDragon, termW)
	facePlayer(game.YellowDragon, termW)
	facePlayer(game.RedDragon, termW)
}

func facePlayer(dragon *game.Object, termW int) {
	dragonCenterX := dragon.RelX + float64(dragon.Width)/(2*float64(termW))
	playerCenterX := game.Player.RelX + float64(game.Player.Width)/(2*float64(termW))
	dragon.Flipped = dragonCenterX < playerCenterX
}

func isGameOver() bool {
	return false
}
