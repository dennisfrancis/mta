package mta

import (
	"fmt"
)

func validateIndex(idx, size int64) error {
	if size < 0 {
		return fmt.Errorf("size < 0")
	}
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
	errStart := validateIndex(beg, size)
	if errStart != nil {
		err = errStart
		return
	}
	errEnd := validateIndex(end, size)
	if errEnd != nil {
		err = errEnd
		return
	}
	if end < beg {
		err = fmt.Errorf("end < beg")
	}
	return
}
