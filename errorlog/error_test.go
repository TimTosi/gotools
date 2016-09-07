package errorlog

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/facebookgo/ensure"
)

// -----------------------------------------------------------------------------

func helperFuncGood() error { return Error(fmt.Errorf("This is an error")) }

func helperFuncNil() error { return Error(nil) }

// -----------------------------------------------------------------------------

func TestError_good(t *testing.T) {
	ensure.Err(
		t,
		helperFuncGood(),
		regexp.MustCompile(".*errorlog/error_test.go:13.*"),
	)
}

func TestError_nil(t *testing.T) { ensure.DeepEqual(t, nil, helperFuncNil()) }
