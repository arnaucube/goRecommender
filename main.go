package main

import (
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/fatih/color"
)

var userCollection *mgo.Collection
var itemCollection *mgo.Collection

func main() {
	fmt.Println("starting")

	//mongodb start
	readMongodbConfig("./mongodbConfig.json")
	session, err := getSession()
	check(err)
	userCollection = getCollection(session, "users")
	itemCollection = getCollection(session, "items")
	color.Green("mongodb connected")

	//read items dataset
	itemsDataset := readDataset("./itemSamples.data", "\n", ",")
	items := getItemsFromDataset(itemsDataset)
	datasetToMongodbIfNotExist(itemCollection, items)

	//http server start
	readServerConfig("./serverConfig.json")
	color.Green("server running")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, router))

}
