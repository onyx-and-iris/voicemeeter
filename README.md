[![Go Reference](https://pkg.go.dev/badge/github.com/onyx-and-iris/voicemeeter-api-go.svg)](https://pkg.go.dev/github.com/onyx-and-iris/voicemeeter-api-go)

# A Go Wrapper for Voicemeeter API

This package offers a Go interface for the Voicemeeter Remote C API.

For an outline of past/future changes refer to: [CHANGELOG](CHANGELOG.md)

## Tested against

-   Basic 1.0.8.2
-   Banana 2.0.6.2
-   Potato 3.0.2.2

## Requirements

-   [Voicemeeter](https://voicemeeter.com/)
-   Go 1.18 or greater

## Installation

Add to your `go.mod` file:

`require github.com/onyx-and-iris/voicemeeter-api-go v1.0.0`

Install voicemeeter-api-go package from your console

`go get github.com/onyx-and-iris/voicemeeter-api-go`

## `Use`

#### `main.go`

```go
package main

import (
	"fmt"

	"github.com/onyx-and-iris/voicemeeter-api-go/voicemeeter"
)

func main() {
	kindId := "banana"
	vmRem := voicemeeter.GetRemote(kindId)

	vmRem.Login()

	vmRem.Strip[0].SetLabel("rode podmic")
	vmRem.Strip[0].SetMute(true)
	fmt.Printf("Strip 0 (%s) mute was set to %v\n", vmRem.Strip[0].GetLabel(), vmRem.Strip[0].GetMute())

	vmRem.Logout()
}
```


## `kindId`

Pass the kind of Voicemeeter as an argument. kindId may be:

-   `basic`
-   `banana`
-   `potato`

## `Remote Type`
#### `vmRem.Strip`
[]t_strip slice containing both physicalStrip and virtualStrip types
#### `vmRem.Bus`
[]t_bus slice containing both physicalBus and virtualBus types
#### `vmRem.Button`
[]button slice containing button types, one for each macrobutton
#### `vmRem.Command`
pointer to command type, represents action type functions
#### `vmRem.Vban`
pointer to vban type, containing both vbanInStream and vbanOutStream slices
#### `vmRem.Device`
pointer to device type, represents physical input/output hardware devices
#### `vmRem.Recorder`
pointer to recorder type, represents the recorder

#### `vmRem.Type()`
returns the type of Voicemeeter as a string
#### `vmRem.Version()`
returns the version of Voicemeeter as a string
#### `vmRem.SendText(<script>)`
sets many parameters in script format eg. ("Strip[0].Mute=1;Bus[3].Gain=3.6")
#### `vmRem.Register(o observer)`
register an object as an observer
#### `vmRem.Deregister(o observer)`
deregister an object as an observer
#### `vmRem.Pdirty()`
returns True iff a GUI parameter has changed
#### `vmRem.Mdirty()`
returns True iff a macrobutton paramter has changed

## `Available commands`

### Strip

The following functions are available

-	`GetMute() bool`
-	`SetMute(val bool)`
-	`GetMono() bool`
-	`SetMono(val bool)`
-	`GetSolo() bool`
-	`SetSolo(val bool)`
-	`GetLimit() int`
-	`SetLimit(val int)` from -40 to 12
-	`GetLabel() string`
-	`SetLabel(val string)`
-	`GetGain() float64`
-	`SetGain(val float32)` from -60.0 to 12.0
-	`GetMc() bool`
-	`SetMc(val bool)`
-	`GetComp() float64`
-	`SetComp(val float32)` from 0.0 to 10.0
-	`GetGate() float64`
-	`SetGate(val float32)` from 0.0 to 10.0
-	`GetAudibility() float64`
-	`SetAudibility(val float32)` from 0.0 to 10.0
-   `GetA1() bool - GetA5() bool`
-   `SetA1(val bool) - SetA5(val bool)`

example:

```go
vmRem.Strip[3].SetGain(3.7)
fmt.Println(vmRem.Strip[0].GetLabel())
```

### Bus

The following functions are available

-	`String() string`
-	`GetMute() bool`
-	`SetMute(val bool)`
-	`GetEq() bool`
-	`SetEq(val bool)`
-	`GetMono() bool`
-	`SetMono(val bool)`
-	`GetLabel() string`
-	`SetLabel(val string)`
-	`GetGain() float64`
-	`SetGain(val float32)` from -60.0 to 12.0

### Button

The following functions are available

-	`GetState() bool`
-	`SetState(val bool)`
-	`GetStateOnly() bool`
-	`SetStateOnly(val bool)`
-	`GetTrigger() bool`
-	`SetTrigger(val bool)`

### Command

The following functions are available

-	`Show()`
-	`Hide()`
-	`Shutdown()`
-	`Restart()`
-	`Lock(val bool)`

### VBAN

-   `vmRem.Vban.Enable()` `vmRem.Vban.Disable()` Turn VBAN on or off

##### Instream | Outstream

-   `vmRem.Vban.InStream` `vmRem.Vban.OutStream`

The following functions are available

-	`GetOn() bool`
-	`SetOn(val bool)`
-	`GetName() string`
-	`SetName(val string)`
-	`GetIp() string`
-	`SetIp(val string)`
-	`GetPort() int`
-	`SetPort(val int)` from 1024 to 65535
-	`GetSr() int`
-	`SetSr(val int)` (11025, 16000, 22050, 24000, 32000, 44100, 48000, 64000, 88200, 96000)
-	`GetChannel() int`
-	`SetChannel(val int)` from 1 to 8
-	`GetBit() int`
-	`SetBit(val int)` 16 or 24
-	`GetQuality() int`
-	`SetQuality(val int)` from 0 to 4
-	`GetRoute() int`
-	`SetRoute(val int)` from 0 to 8

### Device

The following functions are available

-	`Ins`
-	`Outs`
-	`Input(val int)`
-	`Output(val int)`

### Recorder

The following functions are available

-	`Play()`
-	`Stop()`
-	`Pause()`
-	`Replay()`
-	`Record()`
-	`Ff()`
-	`Rew()`

### Run tests

To run all tests:

```
go run test ./...
```

### Official Documentation

-   [Voicemeeter Remote C API](https://github.com/onyx-and-iris/Voicemeeter-SDK/blob/main/VoicemeeterRemoteAPI.pdf)
