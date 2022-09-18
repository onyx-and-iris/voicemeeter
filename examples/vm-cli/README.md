## About

A simple voicemeeter-cli program. Offers ability to toggle, get and set parameters.

## Use

Toggle with `!` prefix, get by excluding `=` and set by including `=`. Mix and match arguments.

You may pass an optional flag -kind to set the kind of Voicemeeter. Defaults to banana.

You may pass an optional flag -delay to set a delay on the getters. Defaults to 15ms.

for example:

`go run .\main.go -kind=potato -delay=18 strip[0].mute=0 strip[0].mute !strip[0].mute strip[0].mute bus[0].gain=-8.8 command.lock=1`

Expected output:

```
Logged into Voicemeeter potato
Running command strip[0].mute=0
Value of strip[0].mute is: 0
Toggling strip[0].mute
Value of strip[0].mute is: 1
Running command bus[0].gain=-8.8
Running command command.lock=1
```
