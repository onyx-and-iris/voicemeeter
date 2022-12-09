[![Go Reference](https://pkg.go.dev/badge/github.com/onyx-and-iris/voicemeeter.svg)](https://pkg.go.dev/github.com/onyx-and-iris/voicemeeter)

# A Go Wrapper for Voicemeeter API

This package offers a Go interface for the Voicemeeter Remote C API.

For an outline of past/future changes refer to: [CHANGELOG](CHANGELOG.md)

## Tested against

-   Basic 1.0.8.8
-   Banana 2.0.6.8
-   Potato 3.0.2.8

## Requirements

-   [Voicemeeter](https://voicemeeter.com/)
-   Go 1.18 or greater

## Installation

Initialize your own module then `go get`

```
go mod init github.com/x/y
go get github.com/onyx-and-iris/voicemeeter
```

## `Use`

#### `main.go`

```go
package main

import (
	"fmt"
	"log"

	"github.com/onyx-and-iris/voicemeeter/v2"
)

func main() {
	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	vm.Strip[0].SetLabel("rode podmic")
	vm.Strip[0].SetMute(true)
	fmt.Printf("Strip 0 (%s) mute was set to %v\n", vm.Strip[0].Label(), vm.Strip[0].Mute())
}

func vmConnect() (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote("banana", 20)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}

	return vm, nil
}
```

## `voicemeeter.NewRemote(<kindId>, <delay>)`

### `kindId`

Pass the kind of Voicemeeter as an argument. kindId may be:

-   `basic`
-   `banana`
-   `potato`

### `delay`

Pass a delay in milliseconds to force the getters to wait for dirty parameters to clear.

Useful if not listening for event updates.

## `Remote Type`

#### `vm.Strip`

[]t_strip slice containing both physicalStrip and virtualStrip types

#### `vm.Bus`

[]t_bus slice containing both physicalBus and virtualBus types

#### `vm.Button`

[]button slice containing button types, one for each macrobutton

#### `vm.Command`

pointer to command type, represents action type functions

#### `vm.Vban`

pointer to vban type, containing both vbanInStream and vbanOutStream slices

#### `vm.Device`

pointer to device type, represents physical input/output hardware devices

#### `vm.Recorder`

pointer to recorder type, represents the recorder

#### `vm.Midi`

pointer to midi type, represents a connected midi device

#### `vm.Type()`

returns the type of Voicemeeter as a string

#### `vm.Version()`

returns the version of Voicemeeter as a string

#### `vm.GetFloat(<param>)`

gets a float parameter value

#### `vm.SetFloat(<param>, <value>)`

sets a float parameter value eg. vm.SetFloat("strip[0].mute", 1)

#### `vm.GetString(<param>)`

gets a string parameter value

#### `vm.SetString(<param>, <value>)`

sets a string parameter value eg. vm.SetString("strip[0].label", "podmic")

#### `vm.SendText(<script>)`

sets many parameters in script format eg. ("Strip[0].Mute=1;Bus[3].Gain=3.6")

#### `vm.Register(o observer)`

register an observer type as an observer

#### `vm.Deregister(o observer)`

deregister an observer type as an observer

#### `vm.EventAdd(<events>)`

adds a single or multiple events to the pooler. Accepts a string or slice of strings.

#### `vm.EventRemove(<events>)`

removes a single or multiple events from the pooler. Accepts a string or slice of strings.

#### `vm.Pdirty()`

returns True iff a GUI parameter has changed

#### `vm.Mdirty()`

returns True iff a macrobutton parameter has changed

#### `vm.Sync()`

Use this to force dirty parameters to clear after a delay in milliseconds.

## `Available commands`

### Strip

The following methods are available

-   `Mute() bool`
-   `SetMute(val bool)`
-   `Mono() bool`
-   `SetMono(val bool)`
-   `Solo() bool`
-   `SetSolo(val bool)`
-   `Limit() int`
-   `SetLimit(val int)` from -40 to 12
-   `Label() string`
-   `SetLabel(val string)`
-   `Gain() float64`
-   `SetGain(val float64)` from -60.0 to 12.0
-   `Mc() bool`
-   `SetMc(val bool)`
-   `Audibility() float64`
-   `SetAudibility(val float64)` from 0.0 to 10.0
-   `A1() bool - A5() bool`
-   `SetA1(val bool) - SetA5(val bool)`
-   `B1() bool - B3() bool`
-   `SetB1(val bool) bool - SetB3(val bool) bool`
-   `AppGain(name string, gain float64)`
-   `AppMute(name string, val bool)`

example:

```go
vm.Strip[3].SetGain(3.7)
fmt.Println(vm.Strip[0].Label())
vm.Strip[4].SetA1(true)

vm.Strip[5].AppGain("Spotify", 0.5)
vm.Strip[5].AppMute("Spotify", true)
```

##### Comp

-   `vm.Strip[i].Comp()`

The following methods are available

-   `Knob() float64`
-   `SetKnob(val float64)` from 0.0 to 10.0
-   `GainIn() float64`
-   `SetGainIn(val float64)` from -24.0 to 24.0
-   `Ratio() float64`
-   `SetRatio(val float64)` from 1.0 to 8.0
-   `Threshold() float64`
-   `SetThreshold(val float64)` from -40.0 to -3.0
-   `Attack() float64`
-   `SetAttack(val float64)` from 0.0 to 200.0
-   `Release() float64`
-   `SetRelease(val float64)` from 0.0 to 5000.0
-   `Knee() float64`
-   `SetKnee(val float64)` from 0.0 to 1.0
-   `GainOut() float64`
-   `SetGainOut(val float64)` from -24.0 to 24.0
-   `MakeUp() bool`
-   `SetMakeUp(val bool)`

example:

```go
vm.Strip[3].Comp().SetRatio(3.5)
```

##### Gate

-   `vm.Strip[i].Gate()`

The following methods are available

-   `Knob() float64`
-   `SetKnob(val float64)` from 0.0 to 10.0
-   `Threshold() float64`
-   `SetThreshold(val float64)` from -60.0 to -10.0
-   `Damping() float64`
-   `SetDamping(val float64)` from -60.0 to -10.0
-   `BPSidechain() int`
-   `SetBPSidechain(val int)` from 100 to 4000
-   `Attack() float64`
-   `SetAttack(val float64)` from 0.0 to 1000.0
-   `Hold() float64`
-   `SetHold(val float64)` from 0.0 to 5000.0
-   `Release() float64`
-   `SetRelease(val float64)` from 0.0 to 5000.0

example:

```go
fmt.Println(vm.Strip[4].Gate().Attack())
```

##### Denoiser

-   `vm.Strip[i].Denoiser()`

The following methods are available

-   `Knob() float64`
-   `SetKnob(val float64)` from 0.0 to 10.0

example:

```go
vm.Strip[1].Denoiser().SetKnob(4.2)
```

##### Gainlayer

-   `vm.Strip[i].Gainlayer()[j]`

The following methods are available

-   `Get() float64`
-   `Set(val float64)`

example:

```go
vm.Strip[6].GainLayer()[3].Set(-13.6)
```

##### Levels

-   `vm.Strip[i].Levels()`

The following methods are available

-   `PreFader() []float64`
-   `PostFader() []float64`
-   `PostMute() []float64`

example:

```go
fmt.Println(vm.Strip[5].Levels().PreFader())
```

### Bus

The following methods are available

-   `String() string`
-   `Mute() bool`
-   `SetMute(val bool)`
-   `Mono() bool`
-   `SetMono(val bool)`
-   `Label() string`
-   `SetLabel(val string)`
-   `Gain() float64`
-   `SetGain(val float64)` from -60.0 to 12.0

example:

```go
vm.Bus[3].SetEq(true)
fmt.Println(vm.Bus[0].Label())
```

##### Modes

-   `vm.Bus[i].Mode()`

The following methods are available

-   `SetNormal(val bool)`
-   `Normal() bool`
-   `SetAmix(val bool)`
-   `Amix() bool`
-   `SetBmix(val bool)`
-   `Bmix() bool`
-   `SetRepeat(val bool)`
-   `Repeat() bool`
-   `SetComposite(val bool)`
-   `Composite() bool`
-   `SetTvMix(val bool)`
-   `TvMix() bool`
-   `SetUpMix21(val bool)`
-   `UpMix21() bool`
-   `SetUpMix41(val bool)`
-   `UpMix41() bool`
-   `SetUpMix61(val bool)`
-   `UpMix61() bool`
-   `SetCenterOnly(val bool)`
-   `CenterOnly() bool`
-   `SetLfeOnly(val bool)`
-   `LfeOnly() bool`
-   `SetRearOnly(val bool)`
-   `RearOnly() bool`

example:

```go
vm.Bus[3].Mode().SetAmix(true)
vm.Bus[4].Mode().SetCenterOnly(true)
```

##### Levels

-   `vm.Bus[i].Levels()`

The following methods are available

-   `All() []float64`

example:

```go
fmt.Println(vm.Bus[1].Levels().All())
```

### Strip | Bus

##### EQ

-   `vm.Strip[i].Eq()` `vm.Bus[i].Eq()`

The following methods are available.

-   `On() bool`
-   `SetOn(val bool)`
-   `Ab() bool`
-   `SetAb(val bool)`

example:

```go
vm.Strip[1].Eq().SetOn(true)
fmt.Println(vm.Bus[3].Eq().Ab())
```

The following methods are available.

-   `FadeTo(target float64, time_ int)`: float, int
-   `FadeBy(change float64, time_ int)`: float, int

Modify gain to or by the selected amount in db over a time interval in ms.

example:

```go
vm.Strip[3].FadeBy(-8.3, 500)
vm.Bus[3].FadeTo(-12.8, 500)
```

### Button

The following methods are available

-   `State() bool`
-   `SetState(val bool)`
-   `StateOnly() bool`
-   `SetStateOnly(val bool)`
-   `Trigger() bool`
-   `SetTrigger(val bool)`

example:

```go
vm.Button[37].SetState(true)
fmt.Println(vm.Button[64].StateOnly())
```

### Command

The following methods are available

-   `Show()` Show Voicemeeter GUI if it's hidden
-   `Hide()` Hide Voicemeeter GUI if it's shown
-   `Shutdown()` Shuts down the GUI
-   `Restart()` Restart the audio engine
-   `Lock(val bool)` Lock the Voicemeeter GUI

example:

```go
vm.Command.Restart()
vm.Command.Show()
```

### VBAN

-   `vm.Vban.Enable()` `vm.Vban.Disable()` Turn VBAN on or off

##### Instream | Outstream

-   `vm.Vban.InStream[i]` `vm.Vban.OutStream[i]`

The following methods are available

-   `On() bool`
-   `SetOn(val bool)`
-   `Name() string`
-   `SetName(val string)`
-   `Ip() string`
-   `SetIp(val string)`
-   `Port() int`
-   `SetPort(val int)` from 1024 to 65535
-   `Sr() int`
-   `SetSr(val int)` (11025, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000)
-   `Channel() int`
-   `SetChannel(val int)` from 1 to 8
-   `Bit() int`
-   `SetBit(val int)` 16 or 24
-   `Quality() int`
-   `SetQuality(val int)` from 0 to 4
-   `Route() int`
-   `SetRoute(val int)` from 0 to 8

example:

```go
# turn VBAN on
vm.Vban.Enable()

// turn on vban instream 0
vm.Vban.InStream[0].SetOn(true)

// set bit property for outstream 3 to 24
vm.Vban.OutStream[3].SetBit(24)
```

### Device

The following methods are available

-   `Ins()`
-   `Outs()`
-   `Input(val int)`
-   `Output(val int)`

example:

```go
for i := 0; i < int(vm.Device.Ins()); i++ {
	fmt.Println(vm.Device.Input(i))
}
```

### Recorder

The following methods are available

-   `Play()`
-   `Stop()`
-   `Pause()`
-   `Replay()`
-   `Record()`
-   `Ff()`
-   `Rew()`

example:

```go
vm.Recorder.Play()
vm.Recorder.Stop()

# Enable loop play
vm.Recorder.Loop(true)

# Disable recorder out channel B2
vm.Recorder.SetB2(false)
```

### Midi

The following methods are available

-   `Channel()` returns the current midi channel
-   `Current()` returns the most recently pressed midi button
-   `Get(<button>)` returns the value in cache for the midi button

example:

```go
var current = vm.Midi.Current()
var val = vm.Midi.Get(current)
```

### Events

By default level updates are disabled. Any event may be enabled or disabled. The following events exist:

-   `pdirty` parameter updates
-   `mdirty` macrobutton updates
-   `midi` midi updates
-   `ldirty` level updates

example:

```go
events := []string{"ldirty", "mdirty", "pdirty"}

vm.EventAdd(events...)

vm.EventRemove(events...)
```

### Run tests

To run all tests:

```
go test ./...
```

### Official Documentation

-   [Voicemeeter Remote C API](https://github.com/onyx-and-iris/Voicemeeter-SDK/blob/main/VoicemeeterRemoteAPI.pdf)
