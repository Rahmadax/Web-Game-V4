package pkg

type Player struct {
	name  string
	class string
	hp    int
	mp    int
	gold  int
}

func (player Player) New (name string, class string, hp int, mp int, gold int) {
	player.name = name
	player.class = class
	player.hp = hp
	player.mp = mp
	player.gold = gold
}

func (player Player) ChangeHp (change int) {
	player.hp += change
}

func (player Player) ChangeMp (change int) {
	player.mp += change
}

func (player Player) ChangeGold (change int) {
	player.gold += change
}
