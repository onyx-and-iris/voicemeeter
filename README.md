[![Go Reference](https://pkg.go.dev/badge/github.com/onyx-and-iris/voicemeeter-api-go.svg)](https://pkg.go.dev/github.com/onyx-and-iris/voicemeeter-api-go)

# A Go Wrapper for Voicemeeter API

This package offers a Go interface for the Voicemeeter Remote C API.

For an outline of past/future changes refer to: [CHANGELOG](CHANGELOG.md)

## Tested against

-   Basic 1.0.8.4
-   Banana 2.0.6.4
-   Potato 3.0.2.4

## Requirements

-   [Voicemeeter](https://voicemeeter.com/)
-   Go 1.18 or greater

## Installation

#### GO.MOD

Add to your `go.mod` file:

`require github.com/onyx-and-iris/voicemeeter-api-go vX.X.X`

where `vX.X.X` is the version you require.

#### GO GET

Install voicemeeter-api-go package from your console to download the latest version.

`go get github.com/onyx-and-iris/voicemeeter-api-go`

## `Use`

#### `main.go`

```go
package main

import (
	"fmt"
	"log"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

func main() {
	kindId := "banana"
	vm, err := voicemeeter.NewRemote(kindId)
	if err != nil {
		log.Fatal(err)
	}

	err = vm.Login()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	vm.Strip[0].SetLabel("rode podmic")
	vm.Strip[0].SetMute(true)
	fmt.Printf("Strip 0 (%s) mute was set to %v\n", vm.Strip[0].GetLabel(), vm.Strip[0].GetMute())
}
```

## `kindId`

Pass the kind of Voicemeeter as an argument. kindId may be:

-   `basic`
-   `banana`
-   `potato`

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

register an object as an observer

#### `vm.Deregister(o observer)`

deregister an object as an observer

#### `vm.EventAdd(<event>)`

adds an event to the pooler eg. vm.EventAdd("ldirty")

#### `vm.EventRemove(<event>)`

removes an event to the pooler eg. vm.EventRemove("pdirty")

#### `vm.Pdirty()`

returns True iff a GUI parameter has changed

#### `vm.Mdirty()`

returns True iff a macrobutton parameter has changed

## `Available commands`

### Strip

The following methods are available

-   `GetMute() bool`
-   `SetMute(val bool)`
-   `GetMono() bool`
-   `SetMono(val bool)`
-   `GetSolo() bool`
-   `SetSolo(val bool)`
-   `GetLimit() int`
-   `SetLimit(val int)` from -40 to 12
-   `GetLabel() string`
-   `SetLabel(val string)`
-   `GetGain() float64`
-   `SetGain(val float32)` from -60.0 to 12.0
-   `GetMc() bool`
-   `SetMc(val bool)`
-   `GetComp() float64`
-   `SetComp(val float32)` from 0.0 to 10.0
-   `GetGate() float64`
-   `SetGate(val float32)` from 0.0 to 10.0
-   `GetAudibility() float64`
-   `SetAudibility(val float32)` from 0.0 to 10.0
-   `GetA1() bool - GetA5() bool`
-   `SetA1(val bool) - SetA5(val bool)`

example:

```go
vm.Strip[3].SetGain(3.7)
fmt.Println(vm.Strip[0].GetLabel())
vm.Strip[4].SetA1(true)
```

##### Gainlayers

-   `vm.Strip[i].GainLayer()[j]`

The following methods are available

-   `Get() float64`
-   `Set(val float32)`

example:

```go
vm.Strip[6].GainLayer()[3].Set(-13.6)
```

##### Levels

-   `vm.Strip[i].Levels()`

The following methods are available

-   `PreFader() []float32`
-   `PostFader() []float32`
-   `PostMute() []float32`

example:

```go
fmt.Println(vm.Strip[5].Levels().PreFader())
```

### Bus

The following methods are available

-   `String() string`
-   `GetMute() bool`
-   `SetMute(val bool)`
-   `GetEq() bool`
-   `SetEq(val bool)`
-   `GetMono() bool`
-   `SetMono(val bool)`
-   `GetLabel() string`
-   `SetLabel(val string)`
-   `GetGain() float64`
-   `SetGain(val float32)` from -60.0 to 12.0

```go
vm.Bus[3].SetEq(true)
fmt.Println(vm.Bus[0].GetLabel())
```

##### Modes

-   `vm.Bus[i].Mode()`

The following methods are available

-   `SetNormal(val bool)`
-   `GetNormal() bool`
-   `SetAmix(val bool)`
-   `GetAmix() bool`
-   `SetBmix(val bool)`
-   `GetBmix() bool`
-   `SetRepeat(val bool)`
-   `GetRepeat() bool`
-   `SetComposite(val bool)`
-   `GetComposite() bool`
-   `SetTvMix(val bool)`
-   `GetTvMix() bool`
-   `SetUpMix21(val bool)`
-   `GetUpMix21() bool`
-   `SetUpMix41(val bool)`
-   `GetUpMix41() bool`
-   `SetUpMix61(val bool)`
-   `GetUpMix61() bool`
-   `SetCenterOnly(val bool)`
-   `GetCenterOnly() bool`
-   `SetLfeOnly(val bool)`
-   `GetLfeOnly() bool`
-   `SetRearOnly(val bool)`
-   `GetRearOnly() bool`

example:

```go
vm.Bus[3].Mode().SetAmix(true)
vm.Bus[4].Mode().SetCenterOnly(true)
```

##### Levels

-   `vm.Bus[i].Levels()`

The following methods are available

-   `All() []float32`

example:

```go
fmt.Println(vm.Bus[1].Levels().All())
```

### Button

The following methods are available

-   `GetState() bool`
-   `SetState(val bool)`
-   `GetStateOnly() bool`
-   `SetStateOnly(val bool)`
-   `GetTrigger() bool`
-   `SetTrigger(val bool)`

example:

```go
vm.Button[37].SetState(true)
fmt.Println(vm.Button[64].GetStateOnly())
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

-   `vm.Vban.InStream` `vm.Vban.OutStream`

The following methods are available

-   `GetOn() bool`
-   `SetOn(val bool)`
-   `GetName() string`
-   `SetName(val string)`
-   `GetIp() string`
-   `SetIp(val string)`
-   `GetPort() int`
-   `SetPort(val int)` from 1024 to 65535
-   `GetSr() int`
-   `SetSr(val int)` (11025, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000)
-   `GetChannel() int`
-   `SetChannel(val int)` from 1 to 8
-   `GetBit() int`
-   `SetBit(val int)` 16 or 24
-   `GetQuality() int`
-   `SetQuality(val int)` from 0 to 4
-   `GetRoute() int`
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
vm.EventAdd("ldirty")

vm.EventRemove("pdirty")
```

### Run tests

To run all tests:

```
go test ./...
```

### Official Documentation

-   [Voicemeeter Remote C API](https://github.com/onyx-and-iris/Voicemeeter-SDK/blob/main/VoicemeeterRemoteAPI.pdf)
