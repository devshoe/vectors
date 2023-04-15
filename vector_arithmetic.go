package vector

// ScalarAdd adds a `value` to each element of the vector
func (v Vector[T]) ScalarAdd(value T) Vector[T] {
	for i := range v {
		v[i] += value
	}
	return v
}

// ScalarSub subtracts a `value` from each element of the vector
func (v Vector[T]) ScalarSub(value T) Vector[T] {
	for i := range v {
		v[i] -= value
	}
	return v
}

// ScalarMul multiplies each element of the vector by a `value`
func (v Vector[T]) ScalarMul(value T) Vector[T] {
	for i := range v {
		v[i] *= value
	}
	return v
}

// ScalarDiv divides each element of the vector by a `value`
func (v Vector[T]) ScalarDiv(value T) Vector[T] {
	for i := range v {
		v[i] /= value
	}
	return v
}
