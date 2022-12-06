package exception

import "fmt"

var (
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrNotFound       = fmt.Errorf("not found")
	ErrConflict       = fmt.Errorf("conflict")
	ErrTimeout        = fmt.Errorf("timeout")
	ErrCancel         = fmt.Errorf("cancel")
)
