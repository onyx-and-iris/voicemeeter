package voicemeeter

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// iStrip defines the interface strip types must satisfy
type iStrip interface {
	String() string
	Mute() bool
	SetMute(val bool)
	Mono() bool
	SetMono(val bool)
	Solo() bool
	SetSolo(val bool)
	Limit() int
	SetLimit(val int)
	Label() string
	SetLabel(val string)
	Gain() float64
	SetGain(val float64)
	Mc() bool
	SetMc(val bool)
	Audibility() float64
	SetAudibility(val float64)
	PanX() float64
	SetPanX(val float64)
	PanY() float64
	SetPanY(val float64)
	ColorX() float64
	SetColorX(val float64)
	ColorY() float64
	SetColorY(val float64)
	FxX() float64
	SetFxX(val float64)
	FxY() float64
	SetFxY(val float64)
	FadeTo(target float64, time_ int)
	FadeBy(change float64, time_ int)
	AppGain(name string, gain float64)
	AppMute(name string, val bool)
	Eq() *eQ
	Comp() *comp
	Gate() *gate
	Denoiser() *denoiser
	GainLayer() []gainLayer
	Levels() *levels
	iOutputs
}

// strip represents a strip channel
type strip struct {
	iRemote
	outputs
	eQ        *eQ
	comp      *comp
	gate      *gate
	denoiser  *denoiser
	gainLayer []gainLayer
	levels    *levels
}

// Mute returns the value of the Mute parameter
func (s *strip) Mute() bool {
	return s.getter_bool("Mute")
}

// SetMute sets the value of the Mute parameter
func (s *strip) SetMute(val bool) {
	s.setter_bool("Mute", val)
}

// Mono returns the value of the Mono parameter
func (s *strip) Mono() bool {
	return s.getter_bool("Mono")
}

// SetMono sets the value of the Mono parameter
func (s *strip) SetMono(val bool) {
	s.setter_bool("Mono", val)
}

// Solo returns the value of the Solo parameter
func (s *strip) Solo() bool {
	return s.getter_bool("Solo")
}

// SetSolo sets the value of the Solo parameter
func (s *strip) SetSolo(val bool) {
	s.setter_bool("Solo", val)
}

// Limit returns the value of the Limit parameter
func (s *strip) Limit() int {
	return s.getter_int("Limit")
}

// SetLimit sets the value of the Limit parameter
func (s *strip) SetLimit(val int) {
	s.setter_int("Limit", val)
}

// Label returns the value of the Label parameter
func (s *strip) Label() string {
	return s.getter_string("Label")
}

// SetLabel sets the value of the Label parameter
func (s *strip) SetLabel(val string) {
	s.setter_string("Label", val)
}

// Gain returns the value of the Gain parameter
func (s *strip) Gain() float64 {
	return s.getter_float("Gain")
}

// SetGain sets the value of the Gain parameter
func (s *strip) SetGain(val float64) {
	s.setter_float("Gain", val)
}

// PanX returns the value of the Pan_X parameter
func (s *strip) PanX() float64 {
	return s.getter_float("Pan_x")
}

// SetPanX sets the value of the Pan_X parameter
func (s *strip) SetPanX(val float64) {
	s.setter_float("Pan_x", val)
}

// PanY returns the value of the Pan_Y parameter
func (s *strip) PanY() float64 {
	return s.getter_float("Pan_y")
}

// SetPanY sets the value of the Pan_Y parameter
func (s *strip) SetPanY(val float64) {
	s.setter_float("Pan_y", val)
}

// Eq returns the eQ field
func (s *strip) Eq() *eQ {
	return s.eQ
}

// GainLayer returns the gainlayer field
func (s *strip) GainLayer() []gainLayer {
	return s.gainLayer
}

// Levels returns the levels field
func (s *strip) Levels() *levels {
	return s.levels
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

// PhysicalStrip represents a single physical strip
type PhysicalStrip struct {
	strip
}

// newPhysicalStrip returns a PhysicalStrip type
func newPhysicalStrip(i int, k *kind) iStrip {
	o := newOutputs(fmt.Sprintf("strip[%d]", i), i)
	e := newEq(fmt.Sprintf("strip[%d].EQ", i), i)
	c := newComp(i)
	g := newGate(i)
	d := newDenoiser(i)
	gl := make([]gainLayer, 8)
	for j := 0; j < 8; j++ {
		gl[j] = newGainLayer(i, j)
	}
	l := newStripLevels(i, k)
	ps := PhysicalStrip{strip{iRemote{fmt.Sprintf("strip[%d]", i), i}, o, e, c, g, d, gl, l}}
	return &ps
}

// String implements fmt.stringer interface
func (p *PhysicalStrip) String() string {
	return fmt.Sprintf("PhysicalStrip%d", p.index)
}

// Audibility returns the value of the Audibility parameter
func (p *PhysicalStrip) Audibility() float64 {
	return p.getter_float("Audibility")
}

// SetAudibility sets the value of the Audibility parameter
func (p *PhysicalStrip) SetAudibility(val float64) {
	p.setter_float("Audibility", val)
}

// Mc logs a warning reason invalid parameter
// it always returns zero value
func (p *PhysicalStrip) Mc() bool {
	log.Warn("invalid parameter MC for physicalStrip")
	return false
}

// SetMc logs a warning reason invalid parameter
func (p *PhysicalStrip) SetMc(val bool) {
	log.Warn("invalid parameter MC for physicalStrip")
}

// Comp returns the comp field
func (p *PhysicalStrip) Comp() *comp {
	return p.comp
}

// Gate returns the gate field
func (p *PhysicalStrip) Gate() *gate {
	return p.gate
}

// Denoiser returns the denoiser field
func (p *PhysicalStrip) Denoiser() *denoiser {
	return p.denoiser
}

// ColorX returns the value of the Color_X parameter
func (p *PhysicalStrip) ColorX() float64 {
	return p.getter_float("Color_x")
}

// SetColorX sets the value of the Color_X parameter
func (p *PhysicalStrip) SetColorX(val float64) {
	p.setter_float("Color_x", val)
}

// ColorY returns the value of the Color_Y parameter
func (p *PhysicalStrip) ColorY() float64 {
	return p.getter_float("Color_y")
}

// SetColorY sets the value of the Color_Y parameter
func (p *PhysicalStrip) SetColorY(val float64) {
	p.setter_float("Color_y", val)
}

// FxX returns the value of the Color_X parameter
func (p *PhysicalStrip) FxX() float64 {
	return p.getter_float("fx_x")
}

// SetFxX sets the value of the Color_X parameter
func (p *PhysicalStrip) SetFxX(val float64) {
	p.setter_float("fx_x", val)
}

// FxY returns the value of the Color_Y parameter
func (p *PhysicalStrip) FxY() float64 {
	return p.getter_float("fx_y")
}

// SetFxY sets the value of the Color_Y parameter
func (p *PhysicalStrip) SetFxY(val float64) {
	p.setter_float("fx_y", val)
}

// VirtualStrip represents a single virtual strip
type VirtualStrip struct {
	strip
}

// newVirtualStrip returns a VirtualStrip type
func newVirtualStrip(i int, k *kind) iStrip {
	o := newOutputs(fmt.Sprintf("strip[%d]", i), i)
	e := newEq(fmt.Sprintf("strip[%d].EQ", i), i)
	c := newComp(i)
	g := newGate(i)
	d := newDenoiser(i)
	gl := make([]gainLayer, 8)
	for j := 0; j < 8; j++ {
		gl[j] = newGainLayer(i, j)
	}
	l := newStripLevels(i, k)
	vs := VirtualStrip{strip{iRemote{fmt.Sprintf("strip[%d]", i), i}, o, e, c, g, d, gl, l}}
	return &vs
}

// String implements fmt.stringer interface
func (v *VirtualStrip) String() string {
	return fmt.Sprintf("VirtualStrip%d", v.index)
}

// Comp returns the comp field
func (v *VirtualStrip) Comp() *comp {
	return v.comp
}

// Gate returns the gate field
func (v *VirtualStrip) Gate() *gate {
	return v.gate
}

// Denoiser returns the denoiser field
func (v *VirtualStrip) Denoiser() *denoiser {
	return v.denoiser
}

// Mc returns the value of the MC parameter
func (v *VirtualStrip) Mc() bool {
	return v.getter_bool("MC")
}

// SetMc sets the value of the MC parameter
func (v *VirtualStrip) SetMc(val bool) {
	v.setter_bool("MC", val)
}

// ColorX logs a warning reason invalid parameter
// it always returns zero value
func (v *VirtualStrip) ColorX() float64 {
	log.Warn("invalid parameter ColorX for virtualStrip")
	return 0
}

// SetColorX logs a warning reason invalid parameter
func (v *VirtualStrip) SetColorX(val float64) {
	log.Warn("invalid parameter ColorX for virtualStrip")
}

// ColorY logs a warning reason invalid parameter
// it always returns zero value
func (v *VirtualStrip) ColorY() float64 {
	log.Warn("invalid parameter ColorY for virtualStrip")
	return 0
}

// SetColorY logs a warning reason invalid parameter
func (v *VirtualStrip) SetColorY(val float64) {
	log.Warn("invalid parameter ColorY for virtualStrip")
}

// FxX logs a warning reason invalid parameter
// it always returns zero value
func (v *VirtualStrip) FxX() float64 {
	log.Warn("invalid parameter FxX for virtualStrip")
	return 0
}

// SetFxX logs a warning reason invalid parameter
func (v *VirtualStrip) SetFxX(val float64) {
	log.Warn("invalid parameter SetFxX for virtualStrip")
}

// FxY logs a warning reason invalid parameter
// it always returns zero value
func (v *VirtualStrip) FxY() float64 {
	log.Warn("invalid parameter FxY for virtualStrip")
	return 0
}

// SetFxY logs a warning reason invalid parameter
func (v *VirtualStrip) SetFxY(val float64) {
	log.Warn("invalid parameter SetFxY for virtualStrip")
}

// Audibility logs a warning reason invalid parameter
// it always returns zero value
func (v *VirtualStrip) Audibility() float64 {
	log.Warn("invalid parameter Audibility for virtualStrip")
	return 0
}

// SetAudibility logs a warning reason invalid parameter
func (v *VirtualStrip) SetAudibility(val float64) {
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

type denoiser struct {
	iRemote
}

func newDenoiser(i int) *denoiser {
	return &denoiser{iRemote{fmt.Sprintf("strip[%d].denoiser", i), i}}
}

func (d *denoiser) Knob() float64 {
	return d.getter_float("")
}

func (d *denoiser) SetKnob(val float64) {
	d.setter_float("", val)
}

type comp struct {
	iRemote
}

func newComp(i int) *comp {
	return &comp{iRemote{fmt.Sprintf("strip[%d].comp", i), i}}
}

func (c *comp) Knob() float64 {
	return c.getter_float("")
}

func (c *comp) SetKnob(val float64) {
	c.setter_float("", val)
}

func (c *comp) GainIn() float64 {
	return c.getter_float("GainIn")
}

func (c *comp) SetGainIn(val float64) {
	c.setter_float("GainIn", val)
}

func (c *comp) Ratio() float64 {
	return c.getter_float("Ratio")
}

func (c *comp) SetRatio(val float64) {
	c.setter_float("Ratio", val)
}

func (c *comp) Threshold() float64 {
	return c.getter_float("Threshold")
}

func (c *comp) SetThreshold(val float64) {
	c.setter_float("Threshold", val)
}

func (c *comp) Attack() float64 {
	return c.getter_float("Attack")
}

func (c *comp) SetAttack(val float64) {
	c.setter_float("Attack", val)
}

func (c *comp) Release() float64 {
	return c.getter_float("Release")
}

func (c *comp) SetRelease(val float64) {
	c.setter_float("Release", val)
}

func (c *comp) Knee() float64 {
	return c.getter_float("Knee")
}

func (c *comp) SetKnee(val float64) {
	c.setter_float("Knee", val)
}

func (c *comp) GainOut() float64 {
	return c.getter_float("GainOut")
}

func (c *comp) SetGainOut(val float64) {
	c.setter_float("GainOut", val)
}

func (c *comp) MakeUp() bool {
	return c.getter_bool("MakeUp")
}

func (c *comp) SetMakeUp(val bool) {
	c.setter_bool("MakeUp", val)
}

type gate struct {
	iRemote
}

func newGate(i int) *gate {
	return &gate{iRemote{fmt.Sprintf("strip[%d].gate", i), i}}
}

func (g *gate) Knob() float64 {
	return g.getter_float("")
}

func (g *gate) SetKnob(val float64) {
	g.setter_float("", val)
}

func (g *gate) Threshold() float64 {
	return g.getter_float("Threshold")
}

func (g *gate) SetThreshold(val float64) {
	g.setter_float("Threshold", val)
}

func (g *gate) Damping() float64 {
	return g.getter_float("Damping")
}

func (g *gate) SetDamping(val float64) {
	g.setter_float("Damping", val)
}

func (g *gate) BPSidechain() int {
	return g.getter_int("BPSidechain")
}

func (g *gate) SetBPSidechain(val int) {
	g.setter_int("BPSidechain", val)
}

func (g *gate) Attack() float64 {
	return g.getter_float("Attack")
}

func (g *gate) SetAttack(val float64) {
	g.setter_float("Attack", val)
}

func (g *gate) Hold() float64 {
	return g.getter_float("Hold")
}

func (g *gate) SetHold(val float64) {
	g.setter_float("Hold", val)
}

func (g *gate) Release() float64 {
	return g.getter_float("Release")
}

func (g *gate) SetRelease(val float64) {
	g.setter_float("Release", val)
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
func newStripLevels(i int, k *kind) *levels {
	var init int
	var os int
	if i < k.PhysIn {
		init = i * 2
		os = 2
	} else {
		init = (k.PhysIn * 2) + ((i - k.PhysIn) * 8)
		os = 8
	}
	return &levels{iRemote{fmt.Sprintf("strip[%d]", i), i}, k, init, os, "strip"}
}

// PreFader returns the level values for this strip, PREFADER mode
func (l *levels) PreFader() []float64 {
	_levelCache.stripMode = 0
	levels := make([]float64, l.offset)
	for i := range levels {
		levels[i] = convertLevel(_levelCache.stripLevels[i+l.init])
	}
	return levels
}

// PostFader returns the level values for this strip, POSTFADER mode
func (l *levels) PostFader() []float64 {
	_levelCache.stripMode = 1
	levels := make([]float64, l.offset)
	for i := range levels {
		levels[i] = convertLevel(_levelCache.stripLevels[i+l.init])
	}
	return levels
}

// PostMute returns the level values for this strip, POSTMUTE mode
func (l *levels) PostMute() []float64 {
	_levelCache.stripMode = 2
	levels := make([]float64, l.offset)
	for i := range levels {
		levels[i] = convertLevel(_levelCache.stripLevels[i+l.init])
	}
	return levels
}
