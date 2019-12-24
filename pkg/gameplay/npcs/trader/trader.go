package trader

import (
	"github.com/Rahmadax/Web-Game-V4/pkg/inventory"
	player "github.com/Rahmadax/Web-Game-V4/pkg/player_data"
)

type Trader struct {
	Hp              int
	TraderInventory []inventory.ItemSlot
}

func (trader *Trader) Interaction() {
	trader.trade()
}

func (trader *Trader) generateInventory() {

}

// Private Functions
func (character *player.Character) trade() {

}
