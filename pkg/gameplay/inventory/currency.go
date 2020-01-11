package inventory

import (
	"math"
)

type Currencies struct {
	Gold int
}

func (inventory *Inventory) SellItem(itemId string, position int, amount int) (bool, int, error) {
	numRemoved, ok := inventory.removeItem(itemId, position, amount)
	if ok {
		item, err := GetItem(itemId)
		if err != nil {
			return false, 0, err
		}
		inventory.Currencies.Gold += item.Value * numRemoved
		return true, numRemoved, nil
	}

	return false, 0, nil
}

func (inventory *Inventory) BuyItem(itemId string, amount int) (bool, int, error) {
	item, err := GetItem(itemId)
	if err != nil {
		return false, 0, err
	}
	if item.Value*amount > inventory.Currencies.Gold {
		amount = int(math.Floor(float64(inventory.Currencies.Gold / item.Value)))
	}

	ok, _ := inventory.addItem(itemId, amount)
	if ok {
		inventory.Currencies.Gold -= amount * item.Value
		return true, amount, nil
	}

	return false, 0, nil
}
