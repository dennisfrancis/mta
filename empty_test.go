package mta

import (
	"testing"
)

var emptyArrayGenerator = &emptyTypeArray{}

func Test_Type(t *testing.T) {
	emptyArray := emptyArrayGenerator
	if emptyArray.Type() != emptyTypeCode {
		t.Error("Incorrect code returned by Type() : expected = ", emptyTypeCode, " got = ", emptyArray.Type())
	}
}

func Test_Size(t *testing.T) {
	size := int64(100000)
	emptyArray := emptyArrayGenerator.New(size)
	if emptyArray.Size() != size {
		t.Error("Incorrect size array returned by New() : expected = ", size, " got = ", emptyArray.Size())
	}
}

func Test_Insert(t *testing.T) {
	size := int64(1000)
	emptyArray := emptyArrayGenerator.New(size)

	type testCase struct {
		size   int64
		pos    int64
		errMsg string
	}

	cases := []testCase{
		{-100, 10, "Expected error about negative size, got err = nil"},
		{100, size, "Expected error about out of bounds pos argument, got err = nil"},
		{100, size + 1, "Expected error about out of bounds pos argument, got err = nil"},
	}

	for _, cas := range cases {
		err := emptyArray.Insert(struct{}{}, cas.size, cas.pos)
		if err == nil {
			t.Fatal(cas.errMsg)
		}

		if emptyArray.Size() != size {
			t.Fatal("Incorrect size after failed Insert() : expected = ", size, " got = ", emptyArray.Size())
		}
	}

	insertSize := int64(300)
	arr := make([]struct{}, insertSize)
	err := emptyArray.Insert(arr, insertSize, size-1)
	if err != nil {
		t.Error("Expected err = nil, but got err = ", err)
	}
	if emptyArray.Size() != size+insertSize {
		t.Error("Incorrect size after Insert() : expected = ", size+insertSize, " got = ", emptyArray.Size())
	}
}

func Test_Delete(t *testing.T) {
	size := int64(1000)
	emptyArray := emptyArrayGenerator.New(size)

	type testCase struct {
		start  int64
		end    int64
		errMsg string
	}

	cases := []testCase{
		{-100, 10, "Expected error about negative index"},
		{10, -100, "Expected error about negative index"},
		{size, 10, "Expected error about out of bounds index"},
		{10, size, "Expected error about out of bounds index"},
		{101, 100, "Expected error about end < start"},
	}

	for _, cas := range cases {
		err := emptyArray.Delete(cas.start, cas.end)
		if err == nil {
			t.Fatal(cas.errMsg)
		}
		if emptyArray.Size() != size {
			t.Fatal("Incorrect size after failed Delete() : expected = ", size, " got = ", emptyArray.Size())
		}
	}

	start, end := int64(863), size-1
	delSize := end - start + 1
	err := emptyArray.Delete(start, end)
	if err != nil {
		t.Error("Expected nil error, but got err = ", err)
	}
	if emptyArray.Size() != size-delSize {
		t.Error("Incorrect size after Delete() : expected = ", size-delSize, " got = ", emptyArray.Size())
	}
}
