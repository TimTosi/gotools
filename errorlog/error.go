package errorlog

import (
	"fmt"
	"runtime"
)

// -----------------------------------------------------------------------------

// Error returns an error formated according to GNU's error message standard.
func Error(err error) error {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		return fmt.Errorf("%s:%d: %s", file, line, err)
	}
	return nil
}
