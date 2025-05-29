package main

import (
	"fmt"
	_ "exec/utils"
	_ "exec/options"
	_ "exec/functions"
)

func PrintArray(arr []string) {
	var buffer string
	buffer += "["
	for i, a := range arr {
		buffer += "\""
		buffer += a
		buffer += "\""
		if i != len(arr) - 1 {
			buffer += ", "
		}
	}
	buffer += "]"
	fmt.Println(buffer)
}

func main() {}
