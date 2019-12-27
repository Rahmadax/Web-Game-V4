package gameplay

import (
	"github.com/Rahmadax/Web-Game-V4/pkg/inventory"
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
func (character Character) trade() {

}
