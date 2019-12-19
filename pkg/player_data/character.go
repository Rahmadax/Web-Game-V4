package player_data

import "github.com/Rahmadax/Web-Game-V4/pkg/inventory"

type Character struct {
	Id        int
	Name      string
	Class     string
	Hp        int
	Mp        int
	inventory inventory.Inventory
}

func (character *Character) New(id int, name string, class string, hp int, mp int) {
	character.Id = id
	character.Name = name
	character.Class = class
	character.Hp = hp
	character.Mp = mp
}

func (character Character) ChangeHp(change int) {
	character.Hp += change
}

func (character Character) ChangeMp(change int) {
	character.Mp += change
}
