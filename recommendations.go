package main

import (
	"fmt"

	"github.com/fatih/color"
)

func getRecommendations(userid string, nrec int) []ItemModel {
	user, err := getUserById(userid)
	check(err)

	items, err := getItemsNotActed(user.Actions)
	check(err)

	//select nrec items from the items array

	color.Blue("recommended items: ")
	fmt.Println(items)
	return items
}
