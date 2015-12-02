package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["Red"] = "#da1337"

	value := colors["Red"]
	// Did this key exist?
	if value != "" {
		fmt.Println(value)
	}
}
