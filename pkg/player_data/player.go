package player

import uuid "github.com/satori/go.uuid"

type Player struct {
	Id             uuid.UUID
	AccountName    string
	Title          string
	Characters     []Character
	DeadCharacters []Character
}

func InitPlayer(accountName string) Player {
	id := uuid.NewV4()

	return Player{
		Id:             id,
		AccountName:    accountName,
		Title:          "",
		Characters:     make([]Character, 0),
		DeadCharacters: make([]Character, 0),
	}
}

func (player *Player) NewCharacter(name string, race string, class string) {
	id := len(player.Characters) + len(player.DeadCharacters) + 1
	char := Character{
		Id:        id,
		Name:      name,
		Race:      race,
		Class:     class,
		Hp:        10,
		Mp:        10,
		Inventory: generateBasicInventory(),
	}
	player.Characters = append(player.Characters, char)
}

func (player *Player) addCharacterToList(character Character) {
	player.Characters = append(player.Characters, character)
}

func (player *Player) characterPermaDeath(character Character) {
	player.Characters = append(player.DeadCharacters, character)
}
