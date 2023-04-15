package vector

import "math/rand"

// New creates a new vector from a list of values
func New[T Mathable](values ...T) Vector[T] {
	return Vector[T](values)
}

// NewOfLength creates a new vector of length `length` with each element set to 0
// if `value` is provided, each element will be set to first `value`
func NewOfLength[T Mathable](length int, value ...T) Vector[T] {
	var result Vector[T]
	for i := 0; i < length; i++ {
		v := T(0)
		if len(value) > 0 {
			v = value[0]
		}
		result = append(result, v)
	}
	return result
}

// NewRange creates a new vector starting from `from` and ending at `to`.
func NewRange[T Mathable](from, to, stepsize T) Vector[T] {
	var result Vector[T]
	for i := from; i <= to; i += stepsize {
		result = append(result, i)
	}
	return result
}

// NewRandomInt creates a new vector of type `int` with random values between `min` and `max` (min included, max excluded)
func NewRandomInt(min, max, count int) Vector[int] {
	pmax, pmin, add := max, min, 0
	if max < min {
		pmax, pmin = min, max
	} else if max == min {
		return NewOfLength(count, min)
	}

	// set range to [0, max-min]
	add = pmin
	pmax -= pmin
	pmin = 0

	var values Vector[int]
	for i := 0; i < count; i++ {
		values = append(values, rand.Intn(pmax)+add)
	}
	return values
}

// NewRandomFloat64 creates a new vector of type `float64` with random values between `min` and `max` (min included, max excluded)
func NewRandomFloat64(min, max float64, count int) Vector[float64] {
	pmax, pmin, add := max, min, 0.0
	if max < min {
		pmax, pmin = min, max
	} else if max == min {
		return NewOfLength(count, min)
	}

	// set range to [0, max-min]
	add = pmin
	pmax -= pmin
	pmin = 0

	var values Vector[float64]
	for i := 0; i < count; i++ {
		values = append(values, rand.Float64()*(pmax-pmin)+add)
	}
	return values
}
