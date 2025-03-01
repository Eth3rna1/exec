package utils

import (
    "fmt"
    "os"
    // "strings"
    "os/exec"
    "runtime"
    // "errors"
    "exec/options"
)

func Input[T any](prompt T) string {
    var response string;
    fmt.Print(prompt);
    fmt.Scanln(&response);
    return response;
}

func ReadDir(path string) ([]string, error) {
    var entries []string
    iterator, err := os.ReadDir(path);
    if err != nil {
        return nil, err
    }
    for _, entry := range iterator {
        entries = append(entries, entry.Name())
    }
    return entries, nil
}

func Execute(command string) error {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        // for windows
        cmd = exec.Command("cmd", "/C", command)
    } else {
        // For Unix-like systems
        cmd = exec.Command("bash", "-c", command)
    }
    if err := cmd.Start(); err != nil {
        return err
    }
    if err := cmd.Wait(); err != nil {
        return err
    }
    return nil
}

func Find[T comparable](arr []T, item T) bool {
    left := 0
    right := len(arr) - 1
    for left <= right {
        if arr[left] == item || arr[right] == item {
            return true
        }
        left++
        right--
    }
    return false
}

func GetFile(path string) (string, error) {
    entries, err := ReadDir(path)
    if err != nil { return "", err }
    var options options.Options = options.OptionsFrom(entries)
    fmt.Println(options.Display())
    for {
        choice := Input("Choice: ")
        result, err := options.Eval(choice)
        if err != nil {
            fmt.Println(err)
            continue
        }
        return result, nil
    }
}

func ParseStringAsCL(str string) []string {
    var splits []string
    var buffer []rune
    var gather bool = false
    for _, char := range str {
        if char == '"' {
            switch gather {
            case true:
                gather = false
            case false:
                gather = true
            }
            continue
        }
        if char == ' ' {
            if gather {
                buffer = append(buffer, char)
            } else {
                splits = append(splits, string(buffer))
                buffer = []rune{}
            }
            continue
        }
        buffer = append(buffer, char)
    }
    if len(buffer) != 0 {
        splits = append(splits, string(buffer))
    }
    return splits
}
