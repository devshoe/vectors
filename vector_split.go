package vector

// SplitIntoWindows splits a vector into windows of size `windowSize`
// If the vector is not evenly divisible by `windowSize`, the first window will be padded with `padwith`
func (v Vector[T]) SplitIntoWindows(windowSize int, padwith *T) Vectors[T] {
	if windowSize <= 0 {
		panic("windowSize must be greater than 0")
	}

	if padwith != nil {
		padwithValue := *padwith
		countToPad := windowSize - (v.Length() % windowSize)
		padding := []T{}
		for i := 0; i < countToPad; i++ {
			padding = append(padding, padwithValue)
		}
		v = append(padding, v...)
	}

	var result Vectors[T]
	for i := 0; i < v.Length(); i += windowSize {
		result = append(result, v[i:min(i+windowSize, v.Length())])
	}

	return result
}

func (v Vector[T]) SplitIntoRollingWindows(windowSize int, padwith *T) Vectors[T] {
	if windowSize <= 0 {
		panic("windowSize must be greater than 0")
	}

	var result Vectors[T]
	for i := 1; i <= v.Length(); i++ {
		window := v[max(0, i-windowSize):i]
		if padwith != nil {
			padwithValue := *padwith
			countToPad := windowSize - window.Length()
			padding := Vector[T]{}
			for i := 0; i < countToPad; i++ {
				padding = append(padding, padwithValue)
			}
			window = append(padding, window...)
		}
		result = append(result, window)
	}

	return result
}
