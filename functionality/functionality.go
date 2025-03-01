package functionality

import (
    "fmt"
    "strings"
    // "os/exec"
    "exec/utils"
    // "exec/options"
)

/// When the user provides extra arguments signaling some command
///
/// Example:
///     exec echo this commmand that
///
/// This function appends the file at the end of the query
func AppendingFileToCommandLineArgs(dir string, args []string) error {
    file, err := utils.GetFile(dir)
    if err != nil {
        return err
    }
    cmd := append(args, file)
    return utils.Execute(strings.Join(cmd, " "))
}

/// When the user gives extra arguments in the command line
///
/// Example:
///     exec command arguments here {FILE}
///
/// The point is the `{FILE}` token is present
func FILEWithinCommandLineArgs(dir string, args []string) error {
    file, err := utils.GetFile(dir)
    if err != nil { return err }
    var cmd []string = args
    for i, a := range args {
        if a == "{FILE}" {
            cmd[i] = file
        }
    }
    return utils.Execute(strings.Join(cmd, " "))
}

/// When the user only calls `exec` in their terminal
func NoContextWithinCommandLineArgs(dir string) error {
    file, err := utils.GetFile(dir)
    if err != nil { return err }
    fmt.Println("For file name input, place `{FILE}` as a token")
    cmd := utils.Input("Command: ")
    args := utils.ParseStringAsCL(cmd)
    for i, a := range args {
        if a == "{FILE}" {
            args[i] = file
        }
    }
    return utils.Execute(strings.Join(args, " "))
}

