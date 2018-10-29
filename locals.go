package mta

import (
	"fmt"
)

func validateIndex(idx, size int64) error {
	if idx < 0 {
		return fmt.Errorf("index is negative")
	}
	if idx >= size {
		return fmt.Errorf("index is out of bounds")
	}
	return nil
}

func validateBeginEndIndices(beg, end, size int64) (err error) {
	err = nil
	errStart, errEnd := validateIndex(beg, size), validateIndex(end, size)
	if errStart != nil {
		err = errStart
		return
	}
	if errEnd != nil {
		err = errEnd
		return
	}
	return
}
