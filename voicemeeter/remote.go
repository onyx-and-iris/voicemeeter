package voicemeeter

import (
	"fmt"
	"os"
)

// A remote type represents the API for a kind,
// comprised of slices representing each member
type remote struct {
	kind    *kind
	Strip   []t_strip
	Bus     []t_bus
	Button  []button
	Command *command
	Vban    *vban
}

// String implements the stringer interface
func (r *remote) String() string {
	return fmt.Sprintf("Voicemeeter %s", r.kind)
}

func (r *remote) Login() {
	login(r.kind.name)
}

func (r *remote) Logout() {
	logout()
}

func (r *remote) Type() string {
	return getVMType()
}

func (r *remote) Version() string {
	return getVersion()
}

func (r *remote) SendText(script string) {
	setParametersMulti(script)
}

// NewRemote returns a remote type of a kind,
// this is the interface entry point.
func NewRemote(kindId string) *remote {
	_kind, ok := kindMap[kindId]
	if !ok {
		err := fmt.Errorf("unknown Voicemeeter kind '%s'", kindId)
		fmt.Println(err)
		os.Exit(1)
	}

	_strip := make([]t_strip, _kind.numStrip())
	for i := 0; i < _kind.numStrip(); i++ {
		if i < _kind.physIn {
			_strip[i] = newPhysicalStrip(i)
		} else {
			_strip[i] = newVirtualStrip(i)
		}
	}
	_bus := make([]t_bus, _kind.numBus())
	for i := 0; i < _kind.numBus(); i++ {
		if i < _kind.physOut {
			_bus[i] = newPhysicalBus(i)
		} else {
			_bus[i] = newVirtualBus(i)
		}
	}
	_button := make([]button, 80)
	for i := 0; i < 80; i++ {
		_button[i] = newButton(i)
	}
	_command := newCommand()
	_vban := newVban(_kind)

	return &remote{
		kind:    _kind,
		Strip:   _strip,
		Bus:     _bus,
		Button:  _button,
		Command: _command,
		Vban:    _vban,
	}
}
