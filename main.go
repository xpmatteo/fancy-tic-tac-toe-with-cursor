package main

import (
	"fmt"
	"strings"
)

func getGreeting(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	name = strings.ToUpper(name[:1]) + name[1:]
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(getGreeting(""))
}
