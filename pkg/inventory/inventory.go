package inventory

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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
	Name   string `json:"name"`
	Group  string `json:"group"`
	Effect string `json:"effect"`
	Target string `json:"target"`
	Count  int    `json:"count"`
	Value  int    `json:"value"`
}

type Currencies struct {
	Gold int
}

func (inventory *Inventory) getInventoryItemIds() []string {
	items := inventory.ItemSlots

	idList := make([]string, len(items))
	for i := range items {
		idList[i] = items[i].ItemId
	}

	return idList
}

func (inventory *Inventory) getInventory() []Item {
	idList := inventory.getInventoryItemIds()

	itemList, err := GetItems(idList)
	if err != nil {

	}
	return itemList

}

func GetItems(itemIds []string) ([]Item, error) {
	var items = make(map[string]Item)

	jsonFile, err := os.Open("Web-Game-V4/pkg/inventory/items.json")
	if err != nil {
		return []Item{}, err
	}



	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return []Item{}, err
	}

	var outputArray = make([]Item, len(itemIds))
	for i, val := range itemIds {
		outputArray[i] = items[val]
	}

	return outputArray, nil
}

func (inventory *Inventory) AddItem(itemId string, amountToAdd int) bool {
	location, _, ok := inventory.findItemInInventory(itemId)
	if ok {
		inventory.ItemSlots[location].Amount += amountToAdd
		return true
	} else if len(inventory.ItemSlots) < inventory.MaxSize {
			newSlot := ItemSlot{ItemId: itemId, Amount: amountToAdd}
			inventory.ItemSlots = append(inventory.ItemSlots, newSlot)
			return true
	}

	return false
}

func (inventory *Inventory) RemoveItem(itemId string, amountToRemove int) bool {
	if len(inventory.ItemSlots) == 0 {
		return false
	}

	location, num, ok := inventory.findItemInInventory(itemId)
	if ok  {
		if num > amountToRemove {
			inventory.ItemSlots[location].Amount -= amountToRemove
		} else if num == amountToRemove {
			inventory.ItemSlots[location] = ItemSlot{}
		}
	}
	return false
}

// Removes empty slots between items. Pushes all items as left as they can go.
func (inventory *Inventory) CleanInventory() {
	inv := Inventory{
		MaxSize: inventory.MaxSize,
		Currencies:inventory.Currencies,
		ItemSlots: make([]ItemSlot, inventory.MaxSize),
	}

	counter := 0
	for i := 1; i <= inventory.MaxSize; i++ {
		if inventory.ItemSlots[i].ItemId != "" {
			inv.ItemSlots[counter] = inventory.ItemSlots[i]
			counter += 1
		}
	}
}

func (inventory *Inventory) findItemInInventory(itemId string) (int, int, bool) {
	idList := inventory.getInventoryItemIds()
	for i := range idList {
		if idList[i] == itemId {
			return i, inventory.ItemSlots[i].Amount, true
		}
	}
	return -1, -1, false
}
