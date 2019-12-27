package scene_data

type Tile struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Cost   int    `json:"cost"`
	Sprite string `json:"sprite"`
	Action string `json:"action"`
}
