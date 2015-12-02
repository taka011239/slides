package main

import "fmt"

func main() {
	var nilMap map[string]int

	// ok
	fmt.Println(nilMap["hoge"])

	// panic!!
	nilMap["hoge"] = 1
}
