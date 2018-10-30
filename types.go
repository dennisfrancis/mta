package mta

type SingleTypeArray interface {

	// Returns type code of the underlying type
	Type() int64

	// Returns newly created SingleTypeArray of the same type with "size" number of elements
	New(size int64) SingleTypeArray

	// Returns the number of elements in the array
	Size() int64

	// Returns a slice [start, end] of the bare array. Both indices are inclusive.
	GetSlice(start, end int64) (elements interface{}, err error)

	// Inserts the slice "elements" of length "size" at "pos" position in the array.
	// The underlying type of "elements" must match that of the array.
	Insert(elements interface{}, size, pos int64) error

	// Deletes the elements in the array from index "start" to the index "end", both inclusive.
	Delete(start, end int64) error

	// Replaces the content of the array that starts from "pos" to "pos + size - 1" (both inclusive)
	// with the first "size" elements of "elements".
	Replace(elements interface{}, size, pos int64) error
}

type blockType struct {
	start int64
	array SingleTypeArray
}

type MultiTypeArray struct {
	registeredTypes map[int64]SingleTypeArray
	blocks          []blockType
}
