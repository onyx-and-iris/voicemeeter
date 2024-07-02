## About

A Voicemeeter CLI, offers ability to toggle, get and set parameters.

## Install

First build and install it with `go install` (skip this step if using binary from [Releases](https://github.com/onyx-and-iris/voicemeeter/releases))

## Use

Commands that begin with `!` will toggle a parameter, commands that contain `=` will set a parameter, all other commands will get a value.

You may pass the following optional flags:

- -h: Print the help dialogue
- -i: Enable interactive mode
- -k: The kind of Voicemeeter GUI to launch, defaults to Banana
- -l: Log level (0 up to 6), defaults to 3, Warn Level
- -d: Set the delay between commands, defaults to 20ms
- -v: Enable extra console output (toggle and set messages).

for example:

`vm-cli.exe -v -l=4 -k=potato strip[0].mute=0 strip[0].mute !strip[0].mute strip[0].mute bus[0].gain=-8.8 command.lock=1`

Expected output:

```
time="<timestamp>" level=info msg="Logged into Voicemeeter Potato v3.1.1.1"
Setting strip[0].mute=0
strip[0].mute: 0.00
Toggling strip[0].mute
strip[0].mute: 1.00
Setting bus[0].gain=-8.8
Setting command.lock=1
time="<timestamp>" level=info msg="Logged out of Voicemeeter Potato"
```

If running in interactive mode enter `Q`, to exit.
