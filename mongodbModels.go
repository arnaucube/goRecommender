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
}

var scores []float64
var ranking []ItemModel
