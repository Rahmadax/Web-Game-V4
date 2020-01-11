package fishing

// Any new attributes need to have a capital letter and a json tag that matches a value in the fish json.
// You can't sort by name in the json due to weird issues.
// It makes more sense to sort by a string value of an integer "0", "1"...

type Fish struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
