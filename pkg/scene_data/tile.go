package scene_data

type Tile struct {
	Id     int
	Cost   int // Cost to cross when calculating A*
	Action func()
}