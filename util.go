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
func update(s1 []float32, s2 []float32, sz int) {
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
func convertLevel(i float32) float32 {
	if i > 0 {
		val := 20 * math.Log10(float64(i))
		return float32(roundFloat(float64(val), 1))
	}
	return -200.0
}
