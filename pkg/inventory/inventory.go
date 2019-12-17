package inventory

import (
	"encoding/json"
	models "github.com/rahmadax/Web-Game-V4/pkg/models"
	"io/ioutil"
	"os"
)

func getItems (itemIds []string) ([]models.Item, error) {
	var items []models.Item

	jsonFile, err := os.Open("items.json")
	if err != nil {
		return items, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return items, err
	}

	return items, nil
}