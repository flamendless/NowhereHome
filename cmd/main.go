package main

import (
	"errors"
	"remembering-home/src/assets"
	"remembering-home/src/common"
	"remembering-home/src/conf"
	"remembering-home/src/game"
	"remembering-home/src/logger"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	logger.InitLog()
	conf.LogConf()

	logger.Log().Info("Setting up resources loader...")
	loader := assets.NewAssetsLoader()

	logger.Log().Info("Launching game...")
	sceneManager := game.NewSceneManager()
	inputSystem := game.NewInputSystem()
	gameState := game.NewGame(loader, sceneManager, inputSystem)

	// sceneManager.GoTo(game.NewDummyScene(gameState))
	// sceneManager.GoTo(game.NewSplashScene(gameState))
	sceneManager.GoTo(game.NewMainMenuScene(gameState))

	if err := ebiten.RunGame(gameState); err != nil {
		if errors.Is(err, common.ERR_QUIT) {
			logger.Log().Info("Successfully exited the game")
			os.Exit(0)
			return
		}
		logger.Log().Fatal(err.Error())
	}
}
