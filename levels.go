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
