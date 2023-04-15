package vector

// Mathable is any type that supports arithmetic operations
// This type constraints the possible types that can be used in a vector
type Mathable interface {
	int | float64 | float32 | int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 | uint
}

type Splitter interface {
	Split()
}
