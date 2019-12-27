package inventory

import (
	"github.com/Rahmadax/Web-Game-V4/pkg/jsons"
	"github.com/Rahmadax/Web-Game-V4/pkg/models"
	"github.com/Rahmadax/Web-Game-V4/pkg/utils"
)

func GetItem(id string) (models.Item, error) {
	itemStruct := map[string]models.Item{}
	err := utils.GetObjectFromJson(jsons.ItemJson, itemStruct)
	if err != nil {
		return models.Item{}, err
	}

	return itemStruct[id], nil
}

func GetItems(idList []string) ([]models.Item, error) {
	itemStruct := map[string]models.Item{}

	err := utils.GetObjectFromJson(jsons.ItemJson, &itemStruct)
	if err != nil {
		return []models.Item{}, err
	}

	returnArr := make([]models.Item, len(idList))
	for i := 0; i < len(idList); i++ {
		returnArr[i] = itemStruct[idList[i]]
	}

	return returnArr, nil
}
