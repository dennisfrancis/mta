package mta

import (
	"fmt"
)

func NewMultiTypeArray(initSize int64, types []SingleTypeArray) (*MultiTypeArray, error) {
	ret := &MultiTypeArray{}
	var err error = nil
	for _, arr := range types {
		err = ret.AddType(arr)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (self *MultiTypeArray) AddType(arr SingleTypeArray) error {
	typ := arr.Type()
	if _, present := self.registeredTypes[typ]; present {
		return fmt.Errorf("SingleTypeArray of Type = %s already registered", typ)
	}
	self.registeredTypes[typ] = arr
	return nil
}

func (self *MultiTypeArray) Insert(elements interface{}, typeCode string, size, pos int64) error {
	return nil
}

func (self *MultiTypeArray) Delete(start, end int64) error {
	return nil
}

func (self *MultiTypeArray) Copy(srcStart, srcEnd, destStart int64) error {
	return nil
}

func (self *MultiTypeArray) Move(srcStart, srcEnd, destStart int64) error {
	return nil
}
