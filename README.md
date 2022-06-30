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

Add to your go.mod file:

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

#### `vmRem.Type`
returns the type of Voicemeeter as a string
#### `vmRem.Version`
returns the version of Voicemeeter as a string
#### `vmRem.SendText`
sets many parameters in script format ("Strip[0].Mute=1;Bus[3].Gain=3.6")
#### `vmRem.Register`
register an object as an observer
#### `vmRem.Deregister`
deregister an object as an observer


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



### Run tests

To run all tests:

```
go run test ./...
```

### Official Documentation

-   [Voicemeeter Remote C API](https://github.com/onyx-and-iris/Voicemeeter-SDK/blob/main/VoicemeeterRemoteAPI.pdf)
