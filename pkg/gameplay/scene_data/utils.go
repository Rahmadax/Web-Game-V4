package scene_data

import (
"github.com/Rahmadax/Web-Game-V4/pkg/jsons"
"github.com/Rahmadax/Web-Game-V4/pkg/utils"
)

func GetTile(id string) (Tile, error) {
	tileStruct := map[string]Tile{}
	err := utils.GetObjectFromJson(jsons.TileJson, tileStruct)
	if err != nil {
		return Tile{}, err
	}

	return tileStruct[id], nil
}

func GetTiles(idList []string) ([]Tile, error) {
	tileStruct := map[string]Tile{}

	err := utils.GetObjectFromJson(jsons.TileJson, &tileStruct)
	if err != nil {
		return []Tile{}, err
	}

	returnArr := make([]Tile, len(idList))
	for i := 0; i < len(idList); i++ {
		returnArr[i] = tileStruct[idList[i]]
	}

	return returnArr, nil
}

