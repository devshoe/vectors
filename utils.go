package vector

// IsInArray checks if value is in array
func IsInArray[T comparable](value T, array ...T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// ToAnyArray converts a slice of any type to a slice of interface{}
func ToAnyArray[T any](array []T) []any {
	var result []any
	for _, v := range array {
		result = append(result, v)
	}
	return result
}

func min(values ...int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func max(values ...int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}
