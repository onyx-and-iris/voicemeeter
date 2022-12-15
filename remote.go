package voicemeeter

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Remote represents the API for a kind
type Remote struct {
	Kind     *kind
	Strip    []iStrip
	Bus      []iBus
	Button   []button
	Command  *command
	Vban     *vban
	Device   *device
	Recorder *recorder
	Midi     *midi_t

	pooler *pooler
}

// String implements the fmt.stringer interface
func (r *Remote) String() string {
	return fmt.Sprintf("Voicemeeter %s", r.Kind)
}

// Login logs into the API
// then it intializes the pooler
func (r *Remote) Login() error {
	err := login(r.Kind.Name)
	if err != nil {
		return err
	}
	r.InitPooler()
	return nil
}

// Logout logs out of the API
// it also terminates the pooler
func (r *Remote) Logout() error {
	r.pooler.run = false
	err := logout(r.Kind.Name)
	if err != nil {
		return err
	}
	return nil
}

// InitPooler initiates the Pooler
func (r *Remote) InitPooler() {
	r.pooler = newPooler(r.Kind)
}

// Run launches the Voicemeeter GUI for a kind.
func (r *Remote) Run(kindId string) error {
	err := runVoicemeeter(kindId)
	if err != nil {
		return err
	}
	time.Sleep(time.Second)
	clear()
	return nil
}

// Type returns the type of Voicemeeter (basic, banana, potato)
func (r *Remote) Type() string {
	val, err := getVMType()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// Version returns the version of Voicemeeter as a string
func (r *Remote) Version() string {
	val, err := getVersion()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// Pdirty returns true iff a parameter value has changed
func (r *Remote) Pdirty() (bool, error) {
	pdirty, err := pdirty()
	return pdirty, err
}

// Mdirty returns true iff a macrobutton value has changed
func (r *Remote) Mdirty() (bool, error) {
	mdirty, err := mdirty()
	return mdirty, err
}

// Sync is a helper method that waits for dirty parameters to clear
func (r *Remote) Sync() {
	time.Sleep(time.Duration(vmdelay) * time.Millisecond)
	clear()
}

// GetFloat gets a float parameter value
func (r *Remote) GetFloat(name string) (float64, error) {
	val, err := getParameterFloat(name)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// SetFloat sets a float paramter value
func (r *Remote) SetFloat(name string, value float64) error {
	err := setParameterFloat(name, value)
	if err != nil {
		return err
	}
	return nil
}

// GetString gets a string parameter value
func (r *Remote) GetString(name string) (string, error) {
	val, err := getParameterString(name)
	if err != nil {
		return "", err
	}
	return val, nil
}

// SetString sets a string parameter value
func (r *Remote) SetString(name, value string) error {
	err := setParameterString(name, value)
	if err != nil {
		return err
	}
	return nil
}

// SendText sets multiple parameters by script
func (r *Remote) SendText(script string) error {
	err := setParametersMulti(script)
	if err != nil {
		return err
	}
	return nil
}

// Register forwards the register method to Pooler
func (r *Remote) Register(channel chan string) {
	r.pooler.Register(channel)
}

// EventAdd adds events to the Pooler
func (r *Remote) EventAdd(events ...string) {
	r.pooler.event.Add(events...)
}

// EventRemove removes events from the Pooler
func (r *Remote) EventRemove(events ...string) {
	r.pooler.event.Remove(events...)
}

// remoteBuilder defines the interface builder types must satisfy
type remoteBuilder interface {
	setKind() remoteBuilder
	makeStrip() remoteBuilder
	makeBus() remoteBuilder
	makeButton() remoteBuilder
	makeCommand() remoteBuilder
	makeVban() remoteBuilder
	makeDevice() remoteBuilder
	makeRecorder() remoteBuilder
	makeMidi() remoteBuilder
	Build() remoteBuilder
	Get() *Remote
}

// director is responsible for directing the genericBuilder
type director struct {
	builder remoteBuilder
}

// SetBuilder sets the appropriate builder type for a kind
func (d *director) SetBuilder(b remoteBuilder) {
	d.builder = b
}

// Construct calls the build function for the specific builder
func (d *director) Construct() {
	d.builder.Build()
}

// Get forwards the Get method to the builder
func (d *director) Get() *Remote {
	return d.builder.Get()
}

// genericBuilder represents a generic builder type
type genericBuilder struct {
	k *kind
	r Remote
}

// setKind sets the kind for a builder of a kind
func (b *genericBuilder) setKind() remoteBuilder {
	b.r.Kind = b.k
	return b
}

// makeStrip makes a strip slice and assigns it to remote.Strip
// []iStrip comprises of both physical and virtual strip types
func (b *genericBuilder) makeStrip() remoteBuilder {
	log.Info("building strip")
	strip := make([]iStrip, b.k.NumStrip())
	for i := 0; i < b.k.NumStrip(); i++ {
		if i < b.k.PhysIn {
			strip[i] = newPhysicalStrip(i, b.k)
		} else {
			strip[i] = newVirtualStrip(i, b.k)
		}
	}
	b.r.Strip = strip
	return b
}

// makeBus makes a bus slice and assigns it to remote.Bus
// []t_bus comprises of both physical and virtual bus types
func (b *genericBuilder) makeBus() remoteBuilder {
	log.Info("building bus")
	bus := make([]iBus, b.k.NumBus())
	for i := 0; i < b.k.NumBus(); i++ {
		if i < b.k.PhysOut {
			bus[i] = newPhysicalBus(i, b.k)
		} else {
			bus[i] = newVirtualBus(i, b.k)
		}
	}
	b.r.Bus = bus
	return b
}

// makeButton makes a button slice and assigns it to remote.Button
func (b *genericBuilder) makeButton() remoteBuilder {
	log.Info("building button")
	button := make([]button, 80)
	for i := 0; i < 80; i++ {
		button[i] = newButton(i)
	}
	b.r.Button = button
	return b
}

// makeCommand makes a command type and assigns it to remote.Command
func (b *genericBuilder) makeCommand() remoteBuilder {
	log.Info("building command")
	b.r.Command = newCommand()
	return b
}

// makeVban makes a vban type and assigns it to remote.Vban
func (b *genericBuilder) makeVban() remoteBuilder {
	log.Info("building vban")
	b.r.Vban = newVban(b.k)
	return b
}

// makeDevice makes a device type and assigns it to remote.Device
func (b *genericBuilder) makeDevice() remoteBuilder {
	log.Info("building device")
	b.r.Device = newDevice()
	return b
}

// makeRecorder makes a recorder type and assigns it to remote.Recorder
func (b *genericBuilder) makeRecorder() remoteBuilder {
	log.Info("building recorder")
	b.r.Recorder = newRecorder()
	return b
}

// makeMidi makes a midi type and assigns it to remote.Midi
func (b *genericBuilder) makeMidi() remoteBuilder {
	log.Info("building midi")
	b.r.Midi = newMidi()
	return b
}

// Get returns a fully constructed remote type for a kind
func (b *genericBuilder) Get() *Remote {
	return &b.r
}

// basicBuilder represents a builder specific to basic type
type basicBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a basic type
func (basb *genericBuilder) Build() remoteBuilder {
	return basb.setKind().
		makeStrip().
		makeBus().
		makeButton().
		makeCommand().
		makeVban().
		makeDevice().
		makeMidi()
}

// bananaBuilder represents a builder specific to banana type
type bananaBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a banana type
func (banb *bananaBuilder) Build() remoteBuilder {
	return banb.setKind().
		makeStrip().
		makeBus().
		makeButton().
		makeCommand().
		makeVban().
		makeDevice().
		makeRecorder().
		makeMidi()
}

// potatoBuilder represents a builder specific to potato type
type potatoBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a potato type
func (potb *potatoBuilder) Build() remoteBuilder {
	return potb.setKind().
		makeStrip().
		makeBus().
		makeButton().
		makeCommand().
		makeVban().
		makeDevice().
		makeRecorder().
		makeMidi()
}

var (
	vmsync  bool
	vmdelay int
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

// NewRemote returns a Remote type for a kind
// this is the interface entry point
func NewRemote(kindId string, delay int) (*Remote, error) {
	kind, ok := kindMap[kindId]
	if !ok {
		err := fmt.Errorf("unknown Voicemeeter kind '%s'", kindId)
		return nil, err
	}
	if delay < 0 {
		err := fmt.Errorf("invalid delay value. should be >= 0")
		return nil, err
	}
	vmsync = delay > 0
	vmdelay = delay

	director := director{}
	switch kind.Name {
	case "basic":
		director.SetBuilder(&basicBuilder{genericBuilder{kind, Remote{}}})
	case "banana":
		director.SetBuilder(&bananaBuilder{genericBuilder{kind, Remote{}}})
	case "potato":
		director.SetBuilder(&potatoBuilder{genericBuilder{kind, Remote{}}})
	}
	director.Construct()
	return director.Get(), nil
}
