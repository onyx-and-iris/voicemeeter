package voicemeeter

import (
	"fmt"
	"os"
)

// A remote type represents the API for a kind,
// comprised of slices representing each member
type remote struct {
	kind   *kind
	Strip  []strip
	Bus    []bus
	Button []button
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
// this exported method is the interface entry point.
func NewRemote(kind_id string) remote {
	kindMap := map[string]*kind{
		"basic":  newBasicKind(),
		"banana": newBananaKind(),
		"potato": newPotatoKind(),
	}

	_kind, ok := kindMap[kind_id]
	if !ok {
		err := fmt.Errorf("unknown Voicemeeter kind '%s'", kind_id)
		fmt.Println(err)
		os.Exit(1)
	}

	_strip := make([]strip, _kind.numStrip())
	for i := 0; i < _kind.physIn+_kind.virtIn; i++ {
		_strip[i] = newStrip(i, _kind)
	}
	_bus := make([]bus, _kind.numBus())
	for i := 0; i < _kind.physOut+_kind.virtOut; i++ {
		_bus[i] = newBus(i, _kind)
	}
	_button := make([]button, 80)
	for i := 0; i < 80; i++ {
		_button[i] = newButton(i)
	}

	return remote{
		kind:   _kind,
		Strip:  _strip,
		Bus:    _bus,
		Button: _button,
	}
}
