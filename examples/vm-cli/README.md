## About

A simple voicemeeter-cli program. Offers ability to toggle, get and set parameters.

## Use

Toggle with `!` prefix, get by excluding `=` and set by including `=`. Mix and match arguments, for example:

`go run .\main.go strip[0].mute=0 strip[0].mute !strip[0].mute strip[0].mute bus[0].gain=-8.8`

Expected output:

```
Running command strip[0].mute=0
Value of strip[0].mute is: 0
Toggling strip[0].mute
Value of strip[0].mute is: 1
Running command bus[0].gain=-8.8
```
