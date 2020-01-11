package characters

func (character *Character) TakeDamage(change int) bool {
	character.ChangeHp(change)
	if character.Hp >= 0 {
		return false
	}
	return true
}

func (character Character) ChangeHp(change int) {
	character.Hp += change
}

func (character Character) ChangeMp(change int) {
	character.Mp += change
}
