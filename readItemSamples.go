package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

func readDataset(path string, lineSeparation string, valueSeparation string) [][]string {
	var dataset [][]string
	b, err := ioutil.ReadFile(path)
	check(err)
	str := string(b)
	str = strings.Replace(str, "\r", "", -1)
	lines := strings.Split(str, lineSeparation)
	for _, v1 := range lines {
		params := strings.Split(v1, valueSeparation)
		var datasetLine []string
		for _, v2 := range params {
			datasetLine = append(datasetLine, v2)
		}
		dataset = append(dataset, datasetLine)
	}
	return dataset
}
func getItemsFromDataset(dataset [][]string) []ItemModel {
	var items []ItemModel
	for _, v := range dataset {
		var newItem ItemModel
		newItem.ID = v[0]
		items = append(items, newItem)
	}
	return items
}

func datasetToMongodbIfNotExist(c *mgo.Collection, items []ItemModel) {
	fmt.Println(items)
	for _, item := range items {
		saveItem(c, item)
	}
}
