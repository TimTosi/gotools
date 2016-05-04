package errorlog

import (
	"fmt"
	"runtime"
)

// -----------------------------------------------------------------------------

// Error returns an error formated according to GNU's error message standard.
func Error(err error) error {
	_, file, line, _ := runtime.Caller(0)
	return fmt.Errorf("%s:%d: %s", file, line, err)
}
