package game_instance

import (
	"github.com/Rahmadax/Web-Game-V4/pkg/gameplay/characters/player_characters"
	"github.com/Rahmadax/Web-Game-V4/pkg/gameplay/scene_data"
)

type GameInstance struct {
	Players  []player.Player
	scenes []scene_data.s
	GameMode string
}

func NewGame(players []player.Player, gameMode string) GameInstance {
	gameInstance := GameInstance{
		Players:  players,
		GameMode: gameMode,
	}
	return gameInstance
}

func (gameInstance *GameInstance) AddPlayers(players []player.Player) {
	gameInstance.Players = players
}
