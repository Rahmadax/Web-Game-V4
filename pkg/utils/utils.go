package utils

import (
	"encoding/json"
	"github.com/Rahmadax/Web-Game-V4/pkg/jsons"
	"github.com/Rahmadax/Web-Game-V4/pkg/models"
	"io/ioutil"
	"os"
)

func GetObjectFromJson(route string) (interface{}, error) {
	var inter interface{}

	switch route {
	case jsons.ItemJson:
		inter = make(map[string]models.Item)
	}

	jsonFile, err := os.Open(route)
	if err != nil {
		return nil, nil
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &inter)
	if err != nil {
		return nil, err
	}

	return inter, nil
}