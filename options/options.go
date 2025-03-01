/*
    A structure that helps with the UI
*/
package options

import (
    "strings"
    "strconv"
    "errors"
)

type Options struct {
    hashmap map[string]string
}

func OptionsFrom(values []string) Options {
    hashmap := make(map[string]string);
    for i, val := range values {
        hashmap[strconv.Itoa(i + 1)] = val;
    }
    return Options { hashmap };
}

func (o *Options) Display() string {
    var buffer []string;
    for i := 1; i <= len(o.hashmap); i++ {
        str_index := strconv.Itoa(i);
        buffer = append(buffer, str_index + ".) " + o.hashmap[str_index]);
    }
    return strings.Join(buffer, "\n");
}

func (o *Options) Eval(choice string) (string, error) {
    for key, value := range o.hashmap {
        if key == choice || choice == value {
            return value, nil;
        }
    }
    return "", errors.New("Invalid Choice")
}
