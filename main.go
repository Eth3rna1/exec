package main

import (
	"os"
	"fmt"
	"strings"
	 "exec/utils"
	 "exec/functions"
)

// The token needed to exchange with an entry
const TOKEN string = "{.}"
// The default directory to start executing
var DIR string = func() string {
	res, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res
}()

func main() {
	if utils.DirLen(DIR) == 0 {
		functions.Exec_NoEntriesInCwd(os.Args, TOKEN)
		return
	}
	if len(os.Args) == 1 {
		functions.Exec_OnlyExecInCl(DIR, TOKEN)
		return
	}
	// anonymous any function
	if func(args *[]string) bool {
		for _, el := range *args {
			if strings.Contains(el, TOKEN) {
				return true
			}
		}
		return false
	}(&os.Args) {
		functions.Exec_TokenInCl(DIR, TOKEN, os.Args)
		return
	}
	functions.Exec_AppendEntryToArgs(DIR, TOKEN, os.Args)
	return
}
