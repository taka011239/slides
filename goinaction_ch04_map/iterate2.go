package main

import (
	"fmt"
	"sort"
)

func main() {
	// START OMIT
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	var keys []string
	for k := range colors {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", colors[k])
	}
	// END OMIT
}
