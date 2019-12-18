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

func (inventory *Inventory) AddItem(itemId string, amount int) bool {
	location, ok := inventory.isItemInInventory(itemId)
	if ok {
		inventory.ItemSlots[location].Amount += amount
		return true
	} else if len(inventory.ItemSlots) < inventory.MaxSize {
			newSlot := ItemSlot{ItemId: itemId, Amount: amount}
			inventory.ItemSlots = append(inventory.ItemSlots, newSlot)
			return true
	}
	return false
}

func (inventory *Inventory) removeItem(itemId string, amount int) {
}

func (inventory *Inventory) isItemInInventory(itemId string) (int, bool) {
	idList := inventory.getInventoryItemIds()
	for i := range idList {
		if idList[i] == itemId {
			return i, true
		}
	}
	return -1, false
}
