package models

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Group       string `json:"group"`
	Target      string `json:"target"`
	Effect      string `json:"effect"`
	EffectValue int    `json:"effect_value"`
	Value       int    `json:"value"`
	MaxStack    int    `json:"max_stack"`
}
