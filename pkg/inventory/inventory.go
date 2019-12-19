package inventory

// vars and constants
const itemJson = "pkg/inventory/items.json"

// Structures
type Inventory struct {
	MaxSize    int
	Currencies Currencies
	ItemSlots  []ItemSlot
}

type ItemSlot struct {
	ItemId string
	Amount int
}

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Target      string `json:"target"`
	Effect      string `json:"effect"`
	EffectValue int    `json:"effect_value"`
	Value       int    `json:"value"`
	MaxStack    int    `json:"max_stack"`
}

// Exported Functions
// Retrieve a full inventory to be returned to the client.
func (inventory *Inventory) GetInventory() ([]Item, error) {
	idList := inventory.getInventoryItemIds()

	itemList, err := getItems(idList)
	if err != nil {
		return []Item{}, err
	}

	return itemList, nil
}

// Removes empty slots between items. Pushes all items as far left as they can go.
func (inventory *Inventory) CleanInventory() {
	itemSlots := make([]ItemSlot, inventory.MaxSize)

	counter := 0
	for i := 0; i < inventory.MaxSize; i++ {
		if inventory.ItemSlots[i].ItemId != "" {
			itemSlots[counter] = inventory.ItemSlots[i]
			counter += 1
		}
	}

	inventory.ItemSlots = itemSlots
}

// Private Functions
func (inventory *Inventory) addItem(itemId string, amountStillToAdd int) (bool, error) {
	bought := false

	invInfo, inInv := inventory.findItemInInventory(itemId)
	item, err := getItem(itemId)
	if err != nil {
		return false, err
	}
	maxStack := item.MaxStack

	// Fill slots that already have that item
	if inInv {
		bought = true
		for i := range invInfo {
			if amountStillToAdd == 0 {
				break
			}
			space := maxStack - invInfo[i].amount
			if space > 0 {
				if space <= amountStillToAdd {
					inventory.ItemSlots[invInfo[i].location].Amount = maxStack
					amountStillToAdd -= space
				} else if space > amountStillToAdd {
					inventory.ItemSlots[invInfo[i].location].Amount += amountStillToAdd
					amountStillToAdd = 0
				}
			}
		}
	}

	// Fill additional empty slots
	if amountStillToAdd > 0 {
		emptySlotLocations, ok := inventory.findEmptyInventorySlots()
		if ok {
			for i := range emptySlotLocations {
				if amountStillToAdd == 0 {
					return true, nil
				}

				newSlot := ItemSlot{}
				if amountStillToAdd > maxStack {
					newSlot = ItemSlot{ItemId: itemId, Amount: maxStack}
					amountStillToAdd -= maxStack
				} else if amountStillToAdd < maxStack {
					newSlot = ItemSlot{ItemId: itemId, Amount: amountStillToAdd}
					amountStillToAdd -= amountStillToAdd
				}
				inventory.ItemSlots[emptySlotLocations[i]] = newSlot
			}
		}
	}
	return bought, nil
}

func (inventory *Inventory) removeItem(itemId string, position int, amountToRemove int) (int, bool) {
	if len(inventory.ItemSlots) == 0 {
		return 0, false
	}
	thisSlot := inventory.ItemSlots[position]
	if thisSlot.ItemId == itemId {
		if thisSlot.Amount > amountToRemove {
			inventory.ItemSlots[position].Amount -= amountToRemove
			return amountToRemove, true
		} else if thisSlot.Amount == amountToRemove {
			inventory.ItemSlots[position] = ItemSlot{}
			return amountToRemove, true
		} else if thisSlot.Amount < amountToRemove {
			inventory.ItemSlots[position] = ItemSlot{}
			return thisSlot.Amount, true
		}
	}

	return 0, false
}

func (inventory *Inventory) MoveItem(currentSlot int, destinationSlot int) {
	if currentSlot < inventory.MaxSize && destinationSlot < inventory.MaxSize && currentSlot >= 0 && destinationSlot >= 0 {
		slots := inventory.ItemSlots
		if slots[destinationSlot].ItemId == "" {
			slots[destinationSlot] = slots[currentSlot]
			slots[currentSlot] = ItemSlot{}
		} else {
			temp := slots[destinationSlot]
			slots[destinationSlot] = slots[currentSlot]
			slots[currentSlot] = temp
		}
	}
}

// Utility Functions
func (inventory *Inventory) getInventoryItemIds() []string {
	items := inventory.ItemSlots

	idList := make([]string, len(items))
	for i := range items {
		idList[i] = items[i].ItemId
	}
	return idList
}

type invInfo struct {
	location int
	amount   int
}

func (inventory *Inventory) findItemInInventory(itemId string) ([]invInfo, bool) {
	idList := inventory.getInventoryItemIds()

	var invInfoArr []invInfo
	for i := range idList {
		if idList[i] == itemId {
			invInfo := invInfo{
				location: i,
				amount:   inventory.ItemSlots[i].Amount,
			}
			invInfoArr = append(invInfoArr, invInfo)
		}
	}
	if len(invInfoArr) == 0 {
		return invInfoArr, false
	} else {
		return invInfoArr, true
	}
}

func (inventory *Inventory) findEmptyInventorySlots() ([]int, bool) {
	idList := inventory.getInventoryItemIds()

	var emptySlotLocations []int
	for i := range idList {
		if idList[i] == "" {
			emptySlotLocations = append(emptySlotLocations, i)
		}
	}
	if len(emptySlotLocations) == 0 {
		return emptySlotLocations, false
	} else {
		return emptySlotLocations, true
	}
}
