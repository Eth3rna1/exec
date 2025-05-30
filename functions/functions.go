package functions

import (
	"os"
	"fmt"
	"strings"
	"exec/utils"
	"path/filepath"
)

// When the token is already predefined in the command line
//     Example: exec example_command.exe {TOKEN} {TOKEN}
func Exec_TokenInCl(dir string, token string, args []string) {
	selection := utils.EntriesDropBox(dir);
	if selection == nil {
		return
	}
	var cmd []string
	relPathSelection := filepath.Join(dir, *selection)
	cmd = append(cmd, args[1:]...)
	utils.SwapToken(&cmd, relPathSelection, token)
	fmt.Println(cmd) // logging
	utils.Execute(cmd)
}

// When only `exec` is called in the command line
//     Example: exec
func Exec_OnlyExecInCl(dir string, token string) {
	selection := utils.EntriesDropBox(dir)
	if selection == nil {
		return
	}
	relPathSelection := filepath.Join(dir, *selection)
	fmt.Printf("\nPlace `%s` as a placeholder to replace with the entry\n", token)
	cmd := utils.ParseQuery(utils.Input("Command>> "))
	utils.SwapToken(&cmd, relPathSelection, token)
	fmt.Println(cmd) // logging
	utils.Execute(cmd)
}

// Regardless of the cl, there are no entries in the
// current working directory to work with
//
// Asks for directory to work with
func Exec_NoEntriesInCwd(args []string, token string) {
	dir := utils.Input("Dir: ")
	stat, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !stat.IsDir() {
		fmt.Printf("`%s` is not a directory.", dir)
		return
	}
	if entries, _ := os.ReadDir(dir); len(entries) == 0 {
		fmt.Printf("`%s` is empty, please select a directory with entries.", dir)
		return
	}
	if len(args) == 1 {
		Exec_OnlyExecInCl(dir, token)
		return
	}
	// my implementation of an `any()` function
	if func(arr *[]string) bool {
		for _, el := range *arr {
			if strings.Contains(el, token) {
				return true
			}
		}
		return false
	}(&args) {
		Exec_TokenInCl(dir, token, args)
		return
	}
	Exec_AppendEntryToArgs(dir, token, args) // else this
}

// When a command given along with calling `exec`, although the token isn't present,
// the program will always assume a token is present, thus, it will open a UI and have
// the client to select and append such entry to the command
//
// Example: exec seek -e -r --use-cache /* Here */     <---
//															 \
//															  |
// The program will have the user select an entry to append  /
func Exec_AppendEntryToArgs(dir string, token string, args []string) {
	var command []string
	selection := utils.EntriesDropBox(dir)
	if selection == nil {
		return
	}
	relPathSelection := filepath.Join(dir, *selection)
	command = append(command, relPathSelection)
	fmt.Println(command) // logging
	utils.Execute(command)
}
