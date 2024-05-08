package conf

import (
	"flag"
	"nowhere-home/internal/logger"

	"go.uber.org/zap"
)

var (
	DEV = true
)

const (
	GAME_TITLE   = "Nowhere Home"
	GAME_VERSION = "0.0.1"
	WINDOW_W     = 640
	WINDOW_H     = 480
	GAME_W       = 640
	GAME_H       = 480
)

func init() {
	flag.BoolVar(&DEV, "dev", false, "developer mode")
	flag.Parse()
}

func LogConf() {
	logger.Log().Info(
		"Game Config",
		zap.Bool("dev", DEV),
		zap.String("title", GAME_TITLE),
		zap.String("version", GAME_VERSION),
		zap.Int("window width", WINDOW_W),
		zap.Int("window height", WINDOW_H),
		zap.Int("game width", GAME_W),
		zap.Int("game height", GAME_H),
	)
}
