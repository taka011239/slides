package main

import "fmt"

var colors map[string]string

func main() {
	colors = map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// START OMIT
	go add()

	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
}

func add() {
	colors["hoge"] = "#hoge"
	colors["fuga"] = "#fuga"
}

// END OMIT
