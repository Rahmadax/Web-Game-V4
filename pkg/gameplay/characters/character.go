package characters

import inventory "github.com/Rahmadax/Web-Game-V4/pkg/gameplay/inventory"

type Character struct {
	Id        int
	Name      string
	Race      string
	Class     string
	Hp        int
	Mp        int
	Inventory inventory.Inventory
}

func (character *Character) New(id int, name string) {
	character.Id = id
	character.Name = name
	character.Class = "human"
	character.Hp = 10
	character.Mp = 10
	character.Inventory = generateBasicInventory()
}

func generateBasicInventory() inventory.Inventory {
	return inventory.Inventory{
		MaxSize:    12,
		Currencies: inventory.Currencies{Gold: 20},
		ItemSlots:  make([]inventory.ItemSlot, 12),
	}
}


