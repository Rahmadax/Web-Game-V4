package player_data

type Player struct {
	Id             string
	AccountName    string
	Title          string
	Characters     []Character
	DeadCharacters []Character
}

func initPlayer(id string, accountName string) Player {
	return Player{
		Id:             id,
		AccountName:    accountName,
		Title:          "",
		Characters:     make([]Character, 0),
		DeadCharacters: make([]Character, 0),
	}
}

func (player *Player) addCharacterToList(character Character) {
	player.Characters = append(player.Characters, character)
}

func (player *Player) characterPermaDeath(character Character) {
	player.Characters = append(player.DeadCharacters, character)
}
