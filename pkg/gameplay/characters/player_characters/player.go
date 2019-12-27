package gameplay

import "github.com/Rahmadax/Web-Game-V4/pkg/gameplay/characters"

type Player struct {
	Id             string
	AccountName    string
	Title          string
	Characters     []characters.Character
	DeadCharacters []characters.Character
}

func initPlayer(id string, accountName string) Player {
	return Player{
		Id:             id,
		AccountName:    accountName,
		Title:          "",
		Characters:     make([]characters.Character, 0),
		DeadCharacters: make([]characters.Character, 0),
	}
}

func (player *Player) addCharacterToList(character characters.Character) {
	player.Characters = append(player.Characters, character)
}

func (player *Player) characterPermaDeath(character characters.Character) {
	player.Characters = append(player.DeadCharacters, character)
}
