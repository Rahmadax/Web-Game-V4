package utils

import (
	"encoding/json"
	"github.com/Rahmadax/Web-Game-V4/pkg/models"
	"github.com/Rahmadax/Web-Game-V4/pkg/scene_data"
	"io/ioutil"
	"os"
)

func GetObjectFromJson(route string, i interface{}) error {
	jsonFile, err := os.Open(route)
	if err != nil {
		return err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	switch i.(type) {
	case *map[string]models.Item:
	case *map[string]scene_data.Tile:
	default:
		return nil
	}

	err = json.Unmarshal(byteValue, &i)
	if err != nil {
		return err
	}

	return nil
}
