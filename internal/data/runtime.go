package data

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = fmt.Errorf("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	return []byte(strconv.Quote(jsonValue)), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquoted, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquoted, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	value, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(value)
	return nil
}
