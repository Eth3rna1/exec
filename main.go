/*
A script for productivity, making it easier to work with longer object names
*/
package main

import (
    // "fmt"
    "os"
    // "exec/options"
    // "errors"
    "exec/utils"
    "exec/functionality"
)

const DEFAULT_DIR string = "."

func main() {
    var err error
    switch len(os.Args) {
    case 0, 1:
        err = functionality.NoContextWithinCommandLineArgs(DEFAULT_DIR)
    default:
        if utils.Find(os.Args, "{FILE}") {
            err = functionality.FILEWithinCommandLineArgs(DEFAULT_DIR, os.Args[1:])
        } else {
            err = functionality.AppendingFileToCommandLineArgs(DEFAULT_DIR, os.Args[1:])
        }
    }
    if err != nil {
        println(err)
        os.Exit(1)
    }
}
