package voicemeeter

type t_vban interface {
	GetOn() bool
	SetOn(val bool)
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

type vbanInStream struct {
	vbanStream
}

func newVbanInStream(i int) t_vban {
	vbi := vbanInStream{vbanStream{iRemote{"vban.instream", i}}}
	return t_vban(&vbi)
}

type vbanOutStream struct {
	vbanStream
}

func newVbanOutStream(i int) t_vban {
	vbo := vbanOutStream{vbanStream{iRemote{"vban.outstream", i}}}
	return t_vban(&vbo)
}

type vban struct {
	InStream  []t_vban
	OutStream []t_vban
}

func newVban(k *kind) *vban {
	_vbanIn := make([]t_vban, k.vbanIn)
	for i := 0; i < k.vbanIn; i++ {
		_vbanIn[i] = newVbanInStream(i)
	}
	_vbanOut := make([]t_vban, k.vbanOut)
	for i := 0; i < k.vbanOut; i++ {
		_vbanOut[i] = newVbanOutStream(i)
	}
	return &vban{
		InStream:  _vbanIn,
		OutStream: _vbanOut,
	}
}
