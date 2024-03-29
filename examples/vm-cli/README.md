## About

A simple voicemeeter-cli program. Offers ability to toggle, get and set parameters.

## Install

First build and install it with `go install` (skip this step if using binary from [Releases](https://github.com/onyx-and-iris/voicemeeter/releases))

## Use

Toggle with `!` prefix, get by excluding `=` and set by including `=`. Mix and match arguments.

You may pass the following optional flags:

-   -v: (-verbose) to toggle console output.
-   -i: (-interactive) to toggle interactive mode.
-   -k: (-kind) to set the kind of Voicemeeter. Defaults to banana.
-   -d: (-delay) to set a delay on the getters. Defaults to 20ms.

for example:

`vm-cli.exe -v -k=potato -d=25 strip[0].mute=0 strip[0].mute !strip[0].mute strip[0].mute bus[0].gain=-8.8 command.lock=1`

Expected output:

```
Running command strip[0].mute=0
Value of strip[0].mute is: 0
Toggling strip[0].mute
Value of strip[0].mute is: 1
Running command bus[0].gain=-8.8
Running command command.lock=1
```

If running in interactive mode enter `q`, `quit` or `<Enter>` to exit.
