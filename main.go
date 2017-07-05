package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("starting")

	//mongodb start
	readMongodbConfig("./mongodbConfig.json")
	c := connectMongodb()
	fmt.Println(c)
	color.Green("mongodb connected")

	//http server start
	readServerConfig("./serverConfig.json")
	color.Green("server running")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, router))

}
