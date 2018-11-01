package stas

//go:generate go run generate/generate.go -type float64
//go run generate/generate.go -type Formula -pkgpath github.com/dennisfrancis/sample

import (
	"fmt"
	"github.com/dennisfrancis/mta"
	//"typepackage"
)

type ElementType struct{}

type ElementTypeArray struct {
	arr []ElementType
}

func (self *ElementTypeArray) Type() string {
	return "[]ElementType"
}

func (self *ElementTypeArray) New(size int64) mta.SingleTypeArray {
	ret := &ElementTypeArray{}
	ret.arr = make([]ElementType, size)
	return ret
}

func (self *ElementTypeArray) Size() int64 {
	return len(self.arr)
}

func (self *ElementTypeArray) GetSlice(start, end int64) (elements interface{}, err error) {

	err = validateBeginEndIndices(start, end, self.size)
	elements = nil

	if err != nil {
		return
	}

	elements = self.arr[start : end+1]
	return
}

func (self *ElementTypeArray) Append(elements interface{}, size int64) error {

	if size < 0 {
		return fmt.Errorf("size < 0 passed")
	}

	srcElements, ok := elements.([]ElementType)
	if !ok {
		return fmt.Errorf("elements is not of type []ElementType")
	}

	if len(srcElements) < size {
		return fmt.Errorf("elements does not have size = %d elements", size)
	}

	self.arr = append(self.arr, srcElements[:size]...)
	return nil
}

func (self *ElementTypeArray) Insert(elements interface{}, size, pos int64) error {

	if size <= 0 {
		return fmt.Errorf("size <= 0")
	}

	if err := validateIndex(pos, len(self.arr)); err != nil {
		return err
	}

	newElems, ok := elements.([]ElementType)
	if !ok {
		return fmt.Errorf("elements is not of type []ElementType")
	}

	first := self.arr[:pos]
	last := self.arr[pos:]

	self.arr = append(first, newElems...)
	if pos < size-1 {
		self.arr = append(self.arr, last...)
	}

	return nil
}

func (self *ElementTypeArray) Delete(start, end int64) error {

	err := validateBeginEndIndices(start, end, self.size)
	if err != nil {
		return err
	}

	if end == (len(self.arr) - 1) {
		self.arr = self.arr[:start]
	} else {
		self.arr = append(self.arr[:start], self.arr[end+1:]...)
	}

	return nil
}

func (self *ElementTypeArray) Replace(elements interface{}, size, pos int64) error {

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

	srcElements, ok := elements.([]ElementType)
	if !ok {
		return fmt.Errorf("elements is not of type []ElementType")
	}

	for srcIdx, element := range srcElements {
		self.arr[pos+srcIdx] = element
	}

	return nil
}
