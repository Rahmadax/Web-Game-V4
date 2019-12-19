package inventory

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Private Helper Functions (lower case)
func getItems(itemIds []string) ([]Item, error) {
	var items = make(map[string]Item)

	jsonFile, err := os.Open(itemJson)
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

func getItem(itemId string) (Item, error) {
	var items = make(map[string]Item)

	jsonFile, err := os.Open(itemJson)
	if err != nil {
		return Item{}, err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return Item{}, err
	}

	return items[itemId], nil
}
