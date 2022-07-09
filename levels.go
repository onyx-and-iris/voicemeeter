package voicemeeter

import "math"

// levels
type levels struct {
	iRemote
	k      *kind
	init   int
	offset int
	id     string
}

func (l *levels) convertLevel(i float32) float32 {
	if i > 0 {
		val := 20 * math.Log10(float64(i))
		return float32(roundFloat(float64(val), 1))
	}
	return -200.0
}

var _levelCache *levelCache

// levelCache defines level slices used by the pooler to track updates
type levelCache struct {
	stripMode       int
	stripLevels     []float32
	busLevels       []float32
	stripLevelsBuff []float32
	busLevelsBuff   []float32
	stripComp       []bool
	busComp         []bool
}

// newLevelCache returns a levelCache struct address
func newLevelCache(k *kind) *levelCache {
	stripLevels := make([]float32, (2*k.physIn)+(8*k.virtIn))
	busLevels := make([]float32, 8*k.numBus())
	stripComp := make([]bool, (2*k.physIn)+(8*k.virtIn))
	busComp := make([]bool, 8*k.numBus())
	if _levelCache == nil {
		_levelCache = &levelCache{stripMode: 0, stripLevels: stripLevels, busLevels: busLevels, stripComp: stripComp, busComp: busComp}
	}
	return _levelCache
}
