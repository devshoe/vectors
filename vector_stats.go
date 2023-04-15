package vector

import "math"

func (v Vector[T]) Sum() T {
	var sum T
	for _, value := range v {
		sum += value
	}
	return sum
}

// Mean returns the average value in the vector
func (v Vector[T]) Mean() float64 {
	var sum float64
	for _, value := range v {
		sum += float64(value)
	}
	return sum / float64(v.Length())
}

// Median returns the middle value in the vector
func (v Vector[T]) Median() T {
	sorted := v.Copy().Sort(false)
	if sorted.Length()%2 == 0 {
		return (sorted[sorted.Length()/2-1] + sorted[sorted.Length()/2]) / 2
	}
	return sorted[sorted.Length()/2]
}

// Mode returns the most common value in the vector
func (v Vector[T]) Mode() T {
	var mode T
	var maxCount int
	for _, value := range v {
		count := 0
		for _, v := range v {
			if v == value {
				count++
			}
		}
		if count > maxCount {
			mode = value
			maxCount = count
		}
	}
	return mode
}

// Variance returns the variance of the vector
// ie the average of the squared differences from the Mean
// https://en.wikipedia.org/wiki/Variance
// you take squares of each value, subtract the mean, and then divide by the number of values
func (v Vector[T]) Variance() float64 {
	var sum float64
	var sumOfSquares float64
	for _, value := range v {
		sum += float64(value)
		sumOfSquares += float64(value) * float64(value)
	}
	n := float64(v.Length())
	return (sumOfSquares - (sum*sum)/n) / n
}

// StandardDeviation returns the standard deviation of the vector
// ie the square root of the variance
func (v Vector[T]) StandardDeviation() float64 {
	return math.Sqrt(v.Variance())
}

// Range returns the difference between the largest and smallest values in the vector
func (v Vector[T]) Range() T {
	return v.Max() - v.Min()
}

// InterquartileRange returns the difference between the 75th and 25th percentiles
func (v Vector[T]) InterquartileRange() T {
	sorted := v.Copy().Sort(false)
	return sorted[sorted.Length()*3/4] - sorted[sorted.Length()/4]
}

func (v Vector[T]) ChangePercentAdjacent() Vector[T] {
	var result Vector[T]
	for i, value := range v {
		if i == 0 {
			result = append(result, 0)
		} else {
			result = append(result, (value-v[i-1])/v[i-1])
		}
	}
	return result
}

// Change returns the difference between the first and last values in the vector
func (v Vector[T]) Change() T {
	return v.Last() - v.First()
}

// PercentChange returns the percentage change between the first and last values in the vector
func (v Vector[T]) PercentChange() float64 {
	return float64(v.Change()) / float64(v.First())
}

func (v Vector[T]) MinMax() (min, max T) {
	return v.Min(), v.Max()
}
