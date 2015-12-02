package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["Red"] = "#da1337"

	value, exists := colors["Red"]
	// Did this key exist?
	if exists {
		fmt.Println(value)
	}
}
