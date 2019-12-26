package inventory

import (
	"github.com/Rahmadax/Web-Game-V4/pkg/utils"
)

func GetItem(id string) (Item, error) {
	items, err := utils.GetObjectFromJson(ItemJson)
	if err != nil {
		return Item{}, err
	}

	return items.(map[string]Item)[id], nil
}

func GetItems(idList []string) ([]Item, error) {
	items, err := utils.GetObjectFromJson(ItemJson)
	if err != nil {
		return []Item{}, err
	}
	it := items.(map[string]interface{})

	returnArr := make([]Item, len(idList))
	for i := 0; i < len(idList); i++ {
		returnArr[i] = it[idList[i]].(Item)
	}

	return returnArr, nil
}

