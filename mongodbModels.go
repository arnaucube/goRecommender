package main

type UserModel struct {
	ID      string `json:"id"`
	Age     int    `json:",string"`
	Actions []string
}
type ItemModel struct {
	ID           string
	TRecommended int
	TActed       int
	TScored      int //times acted from recommendation
}

var scores []float64
var ranking []ItemModel
