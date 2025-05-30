package options

import (
	"errors"
	"strings"
	"strconv"
)

// A structure that deals with UI selection
type Options struct {
	values map[string]string
}

// Initializer function
func NewOptions(vals []string) (*Options, error) {
	if len(vals) == 0 {
		return nil, errors.New("given values for `options` cannot be empty")
	}
	var options Options = Options { values: make(map[string]string) }
	for i, el := range vals {
		options.values[strconv.Itoa(i + 1)] = el
	}
	return &options, nil
}

// Returns a UI displaying all options provided
func (o *Options) Display() string {
	var buffer_lines []string
	for i := 1; i < len(o.values) + 1; i++ {
		key := strconv.Itoa(i)
		val, _ := o.values[key]
		buffer := key + ".) " + val
		buffer_lines = append(buffer_lines, buffer)
	}
	return strings.Join(buffer_lines, "\n")
}

// Handles user input and returns the selected choice
func (o *Options) Evaluate(choice string) *string {
	for i := 1; i < len(o.values) + 1; i++ {
		key := strconv.Itoa(i)
		val, _ := o.values[key]
		if choice == key || choice == val {
			return &val
		}
	}
	return nil
}
