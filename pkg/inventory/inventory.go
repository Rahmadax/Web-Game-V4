package inventory

import (
	"encoding/json"
	models "github.com/rahmadax/Web-Game-V4/pkg/models"
	inventory "github.com/rahmadax/Web-Game-V4/pkg/inventory"
)

func getItems (itemIds []string) ([]models.Item, error) {
	var items []models.Item

	err := json.Unmarshal(items.json, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}