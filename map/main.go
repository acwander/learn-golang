package main

import "fmt"

func main() {
	// 3 ways to declare a map..

	colors := map[string]string{
		"red":   "#ff000",
		"green": "4bf745",
		"white": "#fffff",
	}

	// var colors map[string]string

	// colors := make(map[int]string)

	// // add key value pair
	// colors[10] = "#fffff"

	// // delete a key value pair
	// delete(colors, 10)

	// fmt.Println(colors)
	printMap(colors)
}

// print all key value pairs in a map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
