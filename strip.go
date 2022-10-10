package voicemeeter

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// iStrip defines the interface strip types must satisfy
type iStrip interface {
	String() string
	GetMute() bool
	SetMute(val bool)
	GetMono() bool
	SetMono(val bool)
	GetSolo() bool
	SetSolo(val bool)
	GetLimit() int
	SetLimit(val int)
	GetLabel() string
	SetLabel(val string)
	GetGain() float64
	SetGain(val float64)
	GetMc() bool
	SetMc(val bool)
	GetComp() float64
	SetComp(val float64)
	GetGate() float64
	SetGate(val float64)
	GetAudibility() float64
	SetAudibility(val float64)
	GainLayer() []gainLayer
	Levels() *levels
	FadeTo(target float64, time_ int)
	FadeBy(change float64, time_ int)
	AppGain(name string, gain float64)
	AppMute(name string, val bool)
	iOutputs
}

// strip represents a strip channel
type strip struct {
	iRemote
	outputs
	gainLayer []gainLayer
	levels
}

// GetMute returns the value of the Mute parameter
func (s *strip) GetMute() bool {
	return s.getter_bool("Mute")
}

// SetMute sets the value of the Mute parameter
func (s *strip) SetMute(val bool) {
	s.setter_bool("Mute", val)
}

// GetMono returns the value of the Mono parameter
func (s *strip) GetMono() bool {
	return s.getter_bool("Mono")
}

// SetMono sets the value of the Mono parameter
func (s *strip) SetMono(val bool) {
	s.setter_bool("Mono", val)
}

// GetSolo returns the value of the Solo parameter
func (s *strip) GetSolo() bool {
	return s.getter_bool("Solo")
}

// SetSolo sets the value of the Solo parameter
func (s *strip) SetSolo(val bool) {
	s.setter_bool("Solo", val)
}

// GetLimit returns the value of the Limit parameter
func (s *strip) GetLimit() int {
	return s.getter_int("Limit")
}

// SetLimit sets the value of the Limit parameter
func (s *strip) SetLimit(val int) {
	s.setter_int("Limit", val)
}

// GetLabel returns the value of the Label parameter
func (s *strip) GetLabel() string {
	return s.getter_string("Label")
}

// SetLabel sets the value of the Label parameter
func (s *strip) SetLabel(val string) {
	s.setter_string("Label", val)
}

// GetGain returns the value of the Gain parameter
func (s *strip) GetGain() float64 {
	return s.getter_float("Gain")
}

// SetGain sets the value of the Gain parameter
func (s *strip) SetGain(val float64) {
	s.setter_float("Gain", val)
}

// GainLayer returns the gainlayer field
func (s *strip) GainLayer() []gainLayer {
	return s.gainLayer
}

// Levels returns the levels field
func (s *strip) Levels() *levels {
	return &s.levels
}

// FadeTo sets the value of gain to target over at time interval of time_
func (s *strip) FadeTo(target float64, time_ int) {
	s.setter_string("FadeTo", fmt.Sprintf("(\"%f\", %d)", target, time_))
	time.Sleep(time.Millisecond)
}

// FadeBy adjusts the value of gain by change over a time interval of time_
func (s *strip) FadeBy(change float64, time_ int) {
	s.setter_string("FadeBy", fmt.Sprintf("(\"%f\", %d)", change, time_))
	time.Sleep(time.Millisecond)
}

// physicalStrip represents a single physical strip
type physicalStrip struct {
	strip
}

// newPhysicalStrip returns a physicalStrip type cast to an iStrip
func newPhysicalStrip(i int, k *kind) iStrip {
	o := newOutputs(fmt.Sprintf("strip[%d]", i), i)
	gl := make([]gainLayer, 8)
	for j := 0; j < 8; j++ {
		gl[j] = newGainLayer(i, j)
	}
	l := newStripLevels(i, k)
	ps := physicalStrip{strip{iRemote{fmt.Sprintf("strip[%d]", i), i}, o, gl, l}}
	return iStrip(&ps)
}

// String implements fmt.stringer interface
func (p *physicalStrip) String() string {
	return fmt.Sprintf("PhysicalStrip%d", p.index)
}

// GetComp returns the value of the Comp parameter
func (p *physicalStrip) GetComp() float64 {
	return p.getter_float("Comp")
}

// SetComp sets the value of the Comp parameter
func (p *physicalStrip) SetComp(val float64) {
	p.setter_float("Comp", val)
}

// GetGate returns the value of the Gate parameter
func (p *physicalStrip) GetGate() float64 {
	return p.getter_float("Gate")
}

// SetGate sets the value of the Gate parameter
func (p *physicalStrip) SetGate(val float64) {
	p.setter_float("Gate", val)
}

// GetAudibility returns the value of the Audibility parameter
func (p *physicalStrip) GetAudibility() float64 {
	return p.getter_float("Audibility")
}

// SetAudibility sets the value of the Audibility parameter
func (p *physicalStrip) SetAudibility(val float64) {
	p.setter_float("Audibility", val)
}

// GetMc logs a warning reason invalid parameter
// it always returns zero value
func (p *physicalStrip) GetMc() bool {
	log.Warn("invalid parameter MC for physicalStrip")
	return false
}

// SetMc logs a warning reason invalid parameter
func (p *physicalStrip) SetMc(val bool) {
	log.Warn("invalid parameter MC for physicalStrip")
}

// virtualStrip represents a single virtual strip
type virtualStrip struct {
	strip
}

// newVirtualStrip returns a virtualStrip type cast to an iStrip
func newVirtualStrip(i int, k *kind) iStrip {
	o := newOutputs(fmt.Sprintf("strip[%d]", i), i)
	gl := make([]gainLayer, 8)
	for j := 0; j < 8; j++ {
		gl[j] = newGainLayer(i, j)
	}
	l := newStripLevels(i, k)
	vs := virtualStrip{strip{iRemote{fmt.Sprintf("strip[%d]", i), i}, o, gl, l}}
	return iStrip(&vs)
}

// String implements fmt.stringer interface
func (v *virtualStrip) String() string {
	return fmt.Sprintf("VirtualStrip%d", v.index)
}

// GetMc returns the value of the MC parameter
func (v *virtualStrip) GetMc() bool {
	return v.getter_bool("MC")
}

// SetMc sets the value of the MC parameter
func (v *virtualStrip) SetMc(val bool) {
	v.setter_bool("MC", val)
}

// GetComp logs a warning reason invalid parameter
// it always returns zero value
func (v *virtualStrip) GetComp() float64 {
	log.Warn("invalid parameter Comp for virtualStrip")
	return 0
}

// SetComp logs a warning reason invalid parameter
func (v *virtualStrip) SetComp(val float64) {
	log.Warn("invalid parameter Comp for virtualStrip")
}

// GetGate logs a warning reason invalid parameter
// it always returns zero value
func (v *virtualStrip) GetGate() float64 {
	log.Warn("invalid parameter Gate for virtualStrip")
	return 0
}

// SetGate logs a warning reason invalid parameter
func (v *virtualStrip) SetGate(val float64) {
	log.Warn("invalid parameter Gate for virtualStrip")
}

// GetAudibility logs a warning reason invalid parameter
// it always returns zero value
func (v *virtualStrip) GetAudibility() float64 {
	log.Warn("invalid parameter Audibility for virtualStrip")
	return 0
}

// SetAudibility logs a warning reason invalid parameter
func (v *virtualStrip) SetAudibility(val float64) {
	log.Warn("invalid parameter Audibility for virtualStrip")
}

// AppGain sets the gain in db by val for the app matching name.
func (v *strip) AppGain(name string, val float64) {
	v.setter_string("AppGain", fmt.Sprintf("(\"%s\", %f)", name, val))
}

// AppMute sets mute state as val for the app matching name.
func (v *strip) AppMute(name string, val bool) {
	var value int
	if val {
		value = 1
	} else {
		value = 0
	}
	v.setter_string("AppMute", fmt.Sprintf("(\"%s\", %f)", name, float64(value)))
}

// gainLayer represents the 8 gainlayers for a single strip
type gainLayer struct {
	iRemote
	index int
}

// newGainLayer returns a gainlayer struct
func newGainLayer(i, j int) gainLayer {
	return gainLayer{iRemote{fmt.Sprintf("strip[%d]", i), i}, j}
}

// Get gets the gain value for a single gainlayer
func (gl *gainLayer) Get() float64 {
	return gl.getter_float(fmt.Sprintf("gainlayer[%d]", gl.index))
}

// Set sets the gain value for a single gainlayer
func (gl *gainLayer) Set(val float64) {
	gl.setter_float(fmt.Sprintf("gainlayer[%d]", gl.index), val)
}

// newStripLevels returns a levels struct
func newStripLevels(i int, k *kind) levels {
	var init int
	var os int
	if i < k.PhysIn {
		init = i * 2
		os = 2
	} else {
		init = (k.PhysIn * 2) + ((i - k.PhysIn) * 8)
		os = 8
	}
	return levels{iRemote{fmt.Sprintf("strip[%d]", i), i}, k, init, os, "strip"}
}

// PreFader returns the level values for this strip, PREFADER mode
func (l *levels) PreFader() []float64 {
	_levelCache.stripMode = 0
	var levels []float64
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, convertLevel(_levelCache.stripLevels[i]))
	}
	return levels
}

// PostFader returns the level values for this strip, POSTFADER mode
func (l *levels) PostFader() []float64 {
	_levelCache.stripMode = 1
	var levels []float64
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, convertLevel(_levelCache.stripLevels[i]))
	}
	return levels
}

// PostMute returns the level values for this strip, POSTMUTE mode
func (l *levels) PostMute() []float64 {
	_levelCache.stripMode = 2
	var levels []float64
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, convertLevel(_levelCache.stripLevels[i]))
	}
	return levels
}
