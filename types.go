package mta

type SingleTypeArray interface {
	Type() int64
	New(size int64) SingleTypeArray
	Size() int64
	Insert(elements interface{}, size, pos int64) error
	Delete(start, end int64) error
	Copy(source SingleTypeArray, srcStart, srcEnd, destStart int64) (extendedSize int64, err error)
	Move(source SingleTypeArray, srcStart, srcEnd, destStart int64) error
}

type blockType struct {
	start int64
	array SingleTypeArray
}

type MultiTypeArray struct {
	registeredTypes map[int64]SingleTypeArray
	blocks          []blockType
}
