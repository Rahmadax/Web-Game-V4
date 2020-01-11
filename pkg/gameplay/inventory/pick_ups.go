package inventory

func (inventory *Inventory) PickUpItem(itemId string, amount int) (bool, error) {
	ok, err := inventory.addItem(itemId, amount)
	if err != nil {
		return false, err
	}
	if ok {
		return true, nil
	}

	return false, nil
}
