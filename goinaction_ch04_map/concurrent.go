package main

import (
	"fmt"
	"sync"
)

// START OMIT
type SafeMap struct {
	sync.RWMutex
	m map[string]string
}

// END OMIT

var colors SafeMap

func main() {
	colors = SafeMap{m: map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}}

	go add()

	colors.RLock()
	for key, value := range colors.m {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
	colors.RUnlock()
}

func add() {
	colors.Lock()
	colors.m["hoge"] = "#hoge"
	colors.m["fuga"] = "#fuga"
	colors.Unlock()
}
