package voicemeeter

// levels represents the levels field for a channel
type levels struct {
	iRemote
	k      *kind
	init   int
	offset int
	id     string
}

// returns true if any levels value for a strip/bus have been updated
func (l *levels) IsDirty() bool {
	var vals []bool
	if l.id == "strip" {
		vals = _levelCache.stripComp[l.init : l.init+l.offset]
	} else if l.id == "bus" {
		vals = _levelCache.busComp[l.init : l.init+l.offset]
	}
	return !allTrue(vals, l.offset)
}

var _levelCache *levelCache

// levelCache defines level slices used by the pooler to track updates
type levelCache struct {
	stripMode       int
	stripLevels     []float64
	busLevels       []float64
	stripLevelsBuff []float64
	busLevelsBuff   []float64
	stripComp       []bool
	busComp         []bool
}

// newLevelCache returns a levelCache struct address
func newLevelCache(k *kind) *levelCache {
	stripLevels := make([]float64, (2*k.PhysIn)+(8*k.VirtIn))
	busLevels := make([]float64, 8*k.NumBus())
	stripComp := make([]bool, (2*k.PhysIn)+(8*k.VirtIn))
	busComp := make([]bool, 8*k.NumBus())
	if _levelCache == nil {
		_levelCache = &levelCache{stripMode: 0, stripLevels: stripLevels, busLevels: busLevels, stripComp: stripComp, busComp: busComp}
	}
	return _levelCache
}
