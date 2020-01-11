package inventory

import (
	"github.com/Rahmadax/Web-Game-V4/pkg/jsons"
	"github.com/Rahmadax/Web-Game-V4/pkg/utils"
)

func GetItem(id string) (Item, error) {
	itemStruct := map[string]Item{}
	err := utils.GetObjectFromJson(jsons.ItemJson, itemStruct)
	if err != nil {
		return Item{}, err
	}

	return itemStruct[id], nil
}

func GetItems(idList []string) ([]Item, error) {
	itemStruct := map[string]Item{}

	err := utils.GetObjectFromJson(jsons.ItemJson, &itemStruct)
	if err != nil {
		return []Item{}, err
	}

	returnArr := make([]Item, len(idList))
	for i := 0; i < len(idList); i++ {
		returnArr[i] = itemStruct[idList[i]]
	}

	return returnArr, nil
}
