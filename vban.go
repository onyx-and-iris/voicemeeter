package voicemeeter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// iVban defines the interface vban types must satisfy
type iVban interface {
	GetOn() bool
	SetOn(val bool)
	GetName() string
	SetName(val string)
	GetIp() string
	SetIp(val string)
	GetPort() int
	SetPort(val int)
	GetSr() int
	SetSr(val int)
	GetChannel() int
	SetChannel(val int)
	GetBit() int
	SetBit(val int)
	GetQuality() int
	SetQuality(val int)
	GetRoute() int
	SetRoute(val int)
}

type vbanStream struct {
	iRemote
}

// GetOn returns the value of the On parameter
func (v *vbanStream) GetOn() bool {
	return v.getter_bool("On")
}

// SetOn sets the value of the On parameter
func (v *vbanStream) SetOn(val bool) {
	v.setter_bool("On", val)
}

// GetName returns the value of the Name parameter
func (v *vbanStream) GetName() string {
	return v.getter_string("Name")
}

// SetLabel sets the value of the Name parameter
func (v *vbanStream) SetName(val string) {
	v.setter_string("Name", val)
}

// GetIp returns the value of the Ip parameter
func (v *vbanStream) GetIp() string {
	return v.getter_string("Ip")
}

// SetIp sets the value of the Ip parameter
func (v *vbanStream) SetIp(val string) {
	v.setter_string("Ip", val)
}

// GetPort returns the value of the Port parameter
func (v *vbanStream) GetPort() int {
	return v.getter_int("Port")
}

// SetPort sets the value of the Port parameter
func (v *vbanStream) SetPort(val int) {
	v.setter_int("Port", val)
}

// GetSr returns the value of the Sr parameter
func (v *vbanStream) GetSr() int {
	return v.getter_int("Sr")
}

// SetSr sets the value of the Sr parameter
func (v *vbanStream) SetSr(val int) {
	v.setter_int("Sr", val)
}

// GetChannel returns the value of the Channel parameter
func (v *vbanStream) GetChannel() int {
	return v.getter_int("Channel")
}

// SetChannel sets the value of the Channel parameter
func (v *vbanStream) SetChannel(val int) {
	v.setter_int("Channel", val)
}

// GetBit returns the value of the Bit parameter
func (v *vbanStream) GetBit() int {
	val := v.getter_int("Bit")
	if val == 1 {
		return 16
	}
	return 24
}

// SetBit sets the value of the Bit parameter
func (v *vbanStream) SetBit(val int) {
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

// GetQuality returns the value of the Quality parameter
func (v *vbanStream) GetQuality() int {
	return v.getter_int("Quality")
}

// SetQuality sets the value of the Quality parameter
func (v *vbanStream) SetQuality(val int) {
	v.setter_int("Quality", val)
}

// GetRoute returns the value of the Route parameter
func (v *vbanStream) GetRoute() int {
	return v.getter_int("Route")
}

// SetRoute sets the value of the Route parameter
func (v *vbanStream) SetRoute(val int) {
	v.setter_int("Route", val)
}

type vbanInStream struct {
	vbanStream
}

func newVbanInStream(i int) iVban {
	vbi := vbanInStream{vbanStream{iRemote{fmt.Sprintf("vban.instream[%d]", i), i}}}
	return iVban(&vbi)
}

// SetSr logs a warning reason read only
func (vbi *vbanInStream) SetSr(val int) {
	log.Warn("SR is readonly for vban instreams")
}

// SetChannel logs a warning reason read only
func (vbi *vbanInStream) SetChannel(val int) {
	log.Warn("channel is readonly for vban instreams")
}

// SetBit logs a warning reason read only
func (vbi *vbanInStream) SetBit(val int) {
	log.Warn("bit is readonly for vban instreams")
}

type vbanOutStream struct {
	vbanStream
}

func newVbanOutStream(i int) iVban {
	vbo := vbanOutStream{vbanStream{iRemote{fmt.Sprintf("vban.outstream[%d]", i), i}}}
	return iVban(&vbo)
}

type vban struct {
	InStream  []iVban
	OutStream []iVban
}

func newVban(k *kind) *vban {
	_vbanIn := make([]iVban, k.VbanIn)
	for i := 0; i < k.VbanIn; i++ {
		_vbanIn[i] = newVbanInStream(i)
	}
	_vbanOut := make([]iVban, k.VbanOut)
	for i := 0; i < k.VbanOut; i++ {
		_vbanOut[i] = newVbanOutStream(i)
	}
	return &vban{
		InStream:  _vbanIn,
		OutStream: _vbanOut,
	}
}

func (v *vban) Enable() {
	setParameterFloat("vban.Enable", 1)
}

func (v *vban) Disable() {
	setParameterFloat("vban.Enable", 0)
}
