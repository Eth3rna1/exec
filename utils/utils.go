package utils

import (
	"os"
	"fmt"
	"bufio"
	"runtime"
	"strings"
	"os/exec"
	"exec/options"
)

func Execute(command []string) {
	var args []string
	if runtime.GOOS == "windows" {
		args = append(args, "cmd", "/C")
	}
	args = append(args, command...)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
		if len(ent) == 0 {
			return nil
		}
		for _, e := range ent {
			entries = append(entries, e.Name())
		}
	}
	o, _ := options.NewOptions(entries)
	fmt.Println(o.Display())
	fmt.Println("\nSelect an entry by its index or value.")
	choice := Input("\nEntry: ")
	return o.Evaluate(choice)
}

func SwapToken(args *[]string, entry string, token string) {
	for i := 0; i < len(*args); i++ {
		(*args)[i] = strings.Replace((*args)[i], token, entry, -1)
	}
}

func DirLen(dir string) int {
	stat, err := os.Stat(dir)
	if err != nil {
		return 0
	}
	if !stat.IsDir() {
		return 0
	}
	entries, _ := os.ReadDir(dir)
	return len(entries)
}
