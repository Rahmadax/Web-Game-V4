package game_instance

import (
	"github.com/rahmadax/Web-Game-V4/pkg/player_data"
)

type GameInstance struct {
	Players []player_data.Player
	GameMode string
}

func NewGame(players []player_data.Player, gameMode string) GameInstance {
	gameInstance := GameInstance{
		Players:  players,
		GameMode: gameMode,
	}
	return gameInstance
}

func (gameInstance *GameInstance) AddPlayers(players []player_data.Player) {
	gameInstance.Players = players
}