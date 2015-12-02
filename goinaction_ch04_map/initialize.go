package main

import (
	"fmt"
)

func main() {
	dict := make(map[string]string)
	dict2 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	var dict3 map[string]string

	fmt.Println(dict, dict2, dict3)
}
