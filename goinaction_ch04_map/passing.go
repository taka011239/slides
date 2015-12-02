package main

import "fmt"

func main() {
	// START OMIT
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}

	removeColor(colors, "Coral")

	for key, value := range colors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

// END OMIT
