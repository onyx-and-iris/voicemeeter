package voicemeeter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// iVban defines the interface vban types must satisfy
type iVban interface {
	On() bool
	SetOn(val bool)
	Name() string
	SetName(val string)
	Ip() string
	SetIp(val string)
	Port() int
	SetPort(val int)
	Sr() int
	SetSr(val int)
	Channel() int
	SetChannel(val int)
	Bit() int
	SetBit(val int)
	Quality() int
	SetQuality(val int)
	Route() int
	SetRoute(val int)
}

type stream struct {
	iRemote
}

// On returns the value of the On parameter
func (v *stream) On() bool {
	return v.getter_bool("On")
}

// SetOn sets the value of the On parameter
func (v *stream) SetOn(val bool) {
	v.setter_bool("On", val)
}

// Name returns the value of the Name parameter
func (v *stream) Name() string {
	return v.getter_string("Name")
}

// SetLabel sets the value of the Name parameter
func (v *stream) SetName(val string) {
	v.setter_string("Name", val)
}

// Ip returns the value of the Ip parameter
func (v *stream) Ip() string {
	return v.getter_string("Ip")
}

// SetIp sets the value of the Ip parameter
func (v *stream) SetIp(val string) {
	v.setter_string("Ip", val)
}

// Port returns the value of the Port parameter
func (v *stream) Port() int {
	return v.getter_int("Port")
}

// SetPort sets the value of the Port parameter
func (v *stream) SetPort(val int) {
	v.setter_int("Port", val)
}

// Sr returns the value of the Sr parameter
func (v *stream) Sr() int {
	return v.getter_int("Sr")
}

// SetSr sets the value of the Sr parameter
func (v *stream) SetSr(val int) {
	v.setter_int("Sr", val)
}

// Channel returns the value of the Channel parameter
func (v *stream) Channel() int {
	return v.getter_int("Channel")
}

// SetChannel sets the value of the Channel parameter
func (v *stream) SetChannel(val int) {
	v.setter_int("Channel", val)
}

// Bit returns the value of the Bit parameter
func (v *stream) Bit() int {
	val := v.getter_int("Bit")
	if val == 1 {
		return 16
	}
	return 24
}

// SetBit sets the value of the Bit parameter
func (v *stream) SetBit(val int) {
	switch val {
	case 16:
		val = 1
	case 24:
		val = 2
	default:
		log.Warn("expected value 16 or 24")
		return
	}
	v.setter_int("Bit", val)
}

// Quality returns the value of the Quality parameter
func (v *stream) Quality() int {
	return v.getter_int("Quality")
}

// SetQuality sets the value of the Quality parameter
func (v *stream) SetQuality(val int) {
	v.setter_int("Quality", val)
}

// Route returns the value of the Route parameter
func (v *stream) Route() int {
	return v.getter_int("Route")
}

// SetRoute sets the value of the Route parameter
func (v *stream) SetRoute(val int) {
	v.setter_int("Route", val)
}

type VbanInstream struct {
	stream
}

func newVbanInStream(i int) iVban {
	vbi := VbanInstream{stream{iRemote{fmt.Sprintf("vban.instream[%d]", i), i}}}
	return &vbi
}

// SetSr logs a warning reason read only
func (vbi *VbanInstream) SetSr(val int) {
	log.Warn("SR is readonly for vban instreams")
}

// SetChannel logs a warning reason read only
func (vbi *VbanInstream) SetChannel(val int) {
	log.Warn("channel is readonly for vban instreams")
}

// SetBit logs a warning reason read only
func (vbi *VbanInstream) SetBit(val int) {
	log.Warn("bit is readonly for vban instreams")
}

type VbanOutStream struct {
	stream
}

func newVbanOutStream(i int) iVban {
	vbo := VbanOutStream{stream{iRemote{fmt.Sprintf("vban.outstream[%d]", i), i}}}
	return &vbo
}

type vban struct {
	InStream  []iVban
	OutStream []iVban
}

func newVban(k *kind) *vban {
	vbanIn := make([]iVban, k.VbanIn)
	for i := 0; i < k.VbanIn; i++ {
		vbanIn[i] = newVbanInStream(i)
	}
	vbanOut := make([]iVban, k.VbanOut)
	for i := 0; i < k.VbanOut; i++ {
		vbanOut[i] = newVbanOutStream(i)
	}
	return &vban{
		InStream:  vbanIn,
		OutStream: vbanOut,
	}
}

func (v *vban) Enable() {
	setParameterFloat("vban.Enable", 1)
}

func (v *vban) Disable() {
	setParameterFloat("vban.Enable", 0)
}
