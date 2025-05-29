package utils

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"exec/options"
)

func Input(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.TrimSpace(response)
}

func ParseQuery(query string) []string {
	var fragments []string
	var buffer string
	var quoted bool = false
	var runes []rune = []rune(query)
	for i, char := range runes {
		switch (char) {
		case '"':
			if i != 0 && runes[i - 1] == '\\' {
				buffer = buffer[:len(buffer) - 1]
				buffer += string(char)
			} else {
				quoted = !quoted
			}
		case ' ':
			if !quoted {
				if buffer == "" {
					continue
				}
				fragments = append(fragments, buffer)
				buffer = ""
			} else {
				buffer += string(char)
			}
		default:
			buffer += string(char)
		}
	}
	if len(buffer) != 0 {
		fragments = append(fragments, buffer)
	}
	return fragments
}

func EntriesDropBox(dir string) *string {
	var entries []string
	{
		// populating the entries var with entry names
		ent, _ := os.ReadDir(dir) // assuming that dir exists
		for _, e := range ent {
			entries = append(entries, e.Name())
		}
	}
	// assuming that the directory is not empty
	o, _ := options.NewOptions(entries)
	fmt.Println(o.Display())
	choice := Input("\nChoice: ")
	return o.Evaluate(choice)
}

func SwapToken(args *[]string, entry string, token string) {
	for i := 0; i < len(*args); i++ {
		(*args)[i] = strings.Replace((*args)[i], token, entry, -1)
	}
}
