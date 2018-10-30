package mta

import (
	"fmt"
)

const (
	EmptyTypeCode int64 = -1
)

type EmptyTypeArray struct {
	size int64
}

func (self *EmptyTypeArray) Type() int64 {
	return -1
}

func (self *EmptyTypeArray) New(size int64) *EmptyTypeArray {
	return &EmptyTypeArray{size: size}
}

func (self *EmptyTypeArray) Size() int64 {
	return self.size
}

// Not really going to be used, but still has to satisfy SingleTypeArray interface
func (self *EmptyTypeArray) Insert(elements interface{}, size, pos int64) error {
	if size < 0 {
		return fmt.Errorf("size < 0 passed")
	}
	if pos >= self.size {
		return fmt.Errorf("pos is out of bounds")
	}
	self.size += size
	return nil
}

func (self *EmptyTypeArray) Delete(start, end int64) error {
	err1, err2 := validateIndex(start, self.size), validateIndex(end, self.size)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if end < start {
		return fmt.Errorf("end < start passed")
	}

	size := end - start + 1
	self.size -= size
	return nil
}

func (self *EmptyTypeArray) Copy(source SingleTypeArray, srcStart, srcEnd, destStart int64) (extendedSize int64, err error) {
	extendedSize = 0
	err = nil
	if source.Type() != EmptyTypeCode {
		err = fmt.Errorf("array type mismatch(%d != %d)", source.Type(), EmptyTypeCode)
		return
	}

	srcSize := source.Size()
	err = validateBeginEndIndices(srcStart, srcEnd, srcSize)
	if err != nil {
		return
	}

	// Overwrite/extend at destination
	copySize := srcEnd - srcStart + 1
	destEnd := destStart + copySize - 1
	if destEnd >= self.size {
		extendedSize = destEnd - self.size + 1
		self.size += extendedSize
	}

	return
}

func (self *EmptyTypeArray) Move(source SingleTypeArray, srcStart, srcEnd, destStart int64) (err error) {
	err = nil
	if source.Type() != EmptyTypeCode {
		err = fmt.Errorf("array type mismatch(%d != %d)", source.Type(), EmptyTypeCode)
		return
	}

	srcSize := source.Size()
	err = validateBeginEndIndices(srcStart, srcEnd, srcSize)
	if err != nil {
		return
	}

	// Delete from source
	source.Delete(srcStart, srcEnd)
	// Insert to destination
	copySize := srcEnd - srcStart + 1
	self.size += copySize
	return
}
