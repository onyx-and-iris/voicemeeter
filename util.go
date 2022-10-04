package voicemeeter

import "math"

// allTrue accepts a boolean slice and evaluates if all elements are True
func allTrue(s []bool, sz int) bool {
	for i := 0; i < sz; i++ {
		if !s[i] {
			return false
		}
	}
	return true
}

// update copies the contents of one float slice into another
func update(s1 []float64, s2 []float64, sz int) {
	for i := 0; i < sz; i++ {
		s1[i] = s2[i]
	}
}

// roundFloat rounds a float value to a given precision
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// convertLevel performs the necessary math for a channel level
func convertLevel(i float64) float64 {
	if i > 0 {
		val := 20 * math.Log10(i)
		return roundFloat(val, 1)
	}
	return -200.0
}
