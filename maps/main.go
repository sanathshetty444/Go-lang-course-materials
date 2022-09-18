package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#fff021",
		"blue": "#fff021",
	}

	colors2 := make(map[string]string)
	colors2["blue"] = "ffrree"

	delete(colors2, "blue")

	fmt.Print(colors2["blue"])

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
