package main

import (
	"fmt"

	"github.com/fatih/color"
)

func getRecommendations(userid string, nrec int) {

	fmt.Println(userid)
	fmt.Println(nrec)

	user, err := getUserById(userid)
	check(err)
	color.Blue("user: ")
	fmt.Println(user)

	items, err := getAllItems()
	check(err)
	color.Blue("all items: ")
	fmt.Println(items)
}
