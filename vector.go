package vector

import (
	"sort"
)

// Vector is a slice of any type that supports arithmetic operations
type Vector[T Mathable] []T

// Length returns the length of the vector
func (v Vector[T]) Length() int {
	return len(v)
}

// IsEmpty returns true if the vector is empty or uninitialized
func (v Vector[T]) IsEmpty() bool {
	return v.Length() == 0
}

// Contains returns true if the vector contains the given `value`
func (v Vector[T]) Contains(value T) bool {
	for _, v := range v {
		if v == value {
			return true
		}
	}
	return false
}

// First returns the first element of the vector or a zero value if the vector is empty
func (v Vector[T]) First() T {
	if len(v) == 0 {
		return *new(T)
	}
	return v[0]
}

// Last returns the last element of the vector or a zero value if the vector is empty
func (v Vector[T]) Last() T {
	if len(v) == 0 {
		return *new(T)
	}
	return v[len(v)-1]
}

// Min returns the minimum value in the vector
func (v Vector[T]) Min() T {
	min, _ := v.MinWithIndex()
	return min
}

// Max returns the maximum value in the vector
func (v Vector[T]) Max() T {
	max, _ := v.MaxWithIndex()
	return max
}

// Copy returns a deep copy of the vector
func (v Vector[T]) Copy() Vector[T] {
	var copy Vector[T]
	return append(copy, v...)
}

// Sort sorts the vector in ascending order. If `descending` is true, the vector is sorted in descending order
// original vector is modified inplace
func (v Vector[T]) Sort(descending ...bool) Vector[T] {
	sort.Slice(v, func(i, j int) bool {
		if len(descending) > 0 || descending[0] {
			return v[i] > v[j]
		}
		return v[i] < v[j]
	})
	return v
}

// Reverse reverses the vector
// original vector is modified inplace
func (v Vector[T]) Reverse() Vector[T] {
	for i := 0; i < len(v)/2; i++ {
		j := len(v) - i - 1
		v[i], v[j] = v[j], v[i]
	}
	return v
}

// CountOf returns the number of times `value` appears in the vector
func (v Vector[T]) CountOf(value T) int {
	var count int
	for _, v := range v {
		if v == value {
			count++
		}
	}
	return count
}

// IndexOf returns the index of the first occurrence of `value` in the vector or -1 if not found
func (v Vector[T]) IndexOf(value T) int {
	for i, v := range v {
		if v == value {
			return i
		}
	}
	return -1
}

// Unique returns a new vector with unique elements
func (v Vector[T]) Unique() Vector[T] {
	var result Vector[T]
	for _, value := range v {
		if !result.Contains(value) {
			result = append(result, value)
		}
	}
	return result
}

// MinWithIndex returns the minimum value and its index in the vector
func (v Vector[T]) MinWithIndex() (min T, idx int) {
	min = v.First()
	for i, value := range v {
		if value < min {
			min = value
			idx = i
		}
	}
	return
}

// MaxWithIndex returns the maximum value and its index in the vector
func (v Vector[T]) MaxWithIndex() (max T, idx int) {
	max = v.First()
	for i, value := range v {
		if value > max {
			max = value
			idx = i
		}
	}
	return
}
