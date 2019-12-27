package scene_data

type Coords struct {
	X int
	Y int
}

// These algorithms return an X,Y coordinate that the user will move to.
func (tileMap TileMap) PathFind(start Coords, end Coords) {
	//tileMap.aStar(start, end)
}

//func (tileMap TileMap) aStar(start Coords, end Coords) {
//	m := tileMap.tiles
//	sTile := m[start.x][start.y]
//
//	nodes := make([]Coords, 0)
//
//
//}

func manhattanHeuristic(node Coords, end Coords) Coords {
	return Coords{
		X: node.X - end.X,
		Y: node.Y - end.Y,
	}
}

func (tileMap TileMap) CheckIfWall(node Coords) bool {
	if tileMap.Tiles[node.X][node.Y].Cost == -1 {
		return true
	} else {
		return false
	}
}

func (tileMap TileMap) CheckValidIndex(node Coords, direction string) bool {
	switch direction {
	case "north":
		if node.Y-1 <= -1 {
			return false
		}
		return true

	case "south":
		if node.Y+1 >= len(tileMap.Tiles) {
			return false
		}
		return true

	case "west":
		if node.X-1 <= -1 {
			return false
		}
		return true

	case "east":
		if node.X+1 >= len(tileMap.Tiles[node.Y]) {
			return false
		}
		return true
	}

	return false
}

func GenTileMap() TileMap {
	tm := TileMap{}

	tiles := make([][]Tile, 0)

	tile := Tile{}

	for row := 0; row < 6; row++ {
		tileArr := make([]Tile, row+1)
		tiles = append(tiles, tileArr)
		for col := 0; col < row+1; col++ {
			tiles[row][col] = tile
		}
	}
	tm.Tiles = tiles

	return tm
}
