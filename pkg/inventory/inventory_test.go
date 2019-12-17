package inventory

import "fmt"

func tGetItems(){
	itemIds := []string{"1"}
	items, err := getItems(itemIds)

	fmt.Println(items)
	fmt.Println(err)
}
