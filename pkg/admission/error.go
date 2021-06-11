package admission

import "fmt"

var (
	ErrUnmarshal error = fmt.Errorf("raw bytes invalid")
	ErrInvalidParam error = fmt.Errorf("tapp param invalid")
	ErrNoEnoughResource error = fmt.Errorf("resource not enough")
)