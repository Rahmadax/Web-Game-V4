package scene_data

import "github.com/Rahmadax/Web-Game-V4/pkg/constants"

type Coords struct {
	x int
	y int
}

// These algorithms return an X,Y coordinate that the user will move to.
func (tileMap TileMap) PathFind(start Coords, end Coords)  {
	tileMap.aStar(start, end)
}

func (tileMap TileMap) aStar(start Coords, end Coords) {
	m := tileMap.Map
	sTile := m[start.x][start.y]

	nodes := make([]Coords, 0)


}

type heuristic interface {
	calculate(node Coords, end Coords)
}

func manhattanHeuristic(node Coords, end Coords) Coords {
	return Coords{
		x: node.x - end.x,
		y: node.y - end.y,
	}
}

func (tileMap TileMap) validNode (node Coords, direction string) bool {
	switch direction {
	case constants.n
	}


	if tileMap[]
}

