package main

type UserModel struct {
	ID  string `json:"id"`
	Age int    `json:",string"`
	//age clusters: a<20 (0), 20<a<40 (1), 40<a<60 (2), 60<a (3)
	AgeCluster int
	Actions    []string
}
type ItemModel struct {
	ID           string
	TRecommended int
	TActed       int
	TScored      int //times acted from recommendation
}

var scores []float64
var ranking []ItemModel
