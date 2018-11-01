package mta

import (
	"fmt"
)

const (
	emptyTypeCode string = "emptyType"
)

type emptyTypeArray struct {
	size int64
}

func (self *emptyTypeArray) Type() string {
	return emptyTypeCode
}

func (self *emptyTypeArray) New(size int64) SingleTypeArray {
	return &emptyTypeArray{size: size}
}

func (self *emptyTypeArray) Size() int64 {
	return self.size
}

// Not really going to be used, but still has to satisfy SingleTypeArray interface
func (self *emptyTypeArray) GetSlice(start, end int64) (elements interface{}, err error) {

	err = validateBeginEndIndices(start, end, self.size)
	elements = nil

	if err != nil {
		return
	}

	elements = struct{}{}
	return
}

// Not really going to be used, but still has to satisfy SingleTypeArray interface
func (self *emptyTypeArray) Insert(elements interface{}, size, pos int64) error {

	if size < 0 {
		return fmt.Errorf("size < 0 passed")
	}

	err := validateIndex(pos, self.size)
	if err != nil {
		return err
	}

	self.size += size
	return nil
}

func (self *emptyTypeArray) Delete(start, end int64) error {

	err := validateBeginEndIndices(start, end, self.size)
	if err != nil {
		return err
	}

	size := end - start + 1
	self.size -= size
	return nil
}

func (self *emptyTypeArray) Replace(elements interface{}, size, pos int64) error {

	if size < 0 {
		return fmt.Errorf("size < 0 passed")
	}
	err := validateIndex(pos, self.size)
	if err != nil {
		return err
	}

	endIdx := pos + size - 1
	err = validateIndex(endIdx, self.size)
	if err != nil {
		return err
	}

	// This function is really a no-op for emptyTypeArray
	return nil
}
