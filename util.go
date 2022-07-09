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

func update(s1 []float32, s2 []float32, sz int) {
	for i := 0; i < sz; i++ {
		s1[i] = s2[i]
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
