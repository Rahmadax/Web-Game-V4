package utils

import (
	"encoding/json"
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

	switch i.(type){}

	err = json.Unmarshal(byteValue, &i)
	if err != nil {
		return err
	}

	return nil
}
