## Requirements

-   [OBS Studio](https://obsproject.com/)
-   [GOOBS Go Client for Websocket v5](https://github.com/andreykaipov/goobs)

## About

A simple demonstration showing how to sync OBS scene switches to Voicemeeter states. The script assumes you have connection info saved in
a config file named `config.toml` placed next to `main.go`. It also assumes you have scenes named `START` `BRB` `END` and `LIVE`.

A valid `config.toml` file might look like this:

```toml
[connection]
Host="localhost"
Port=4455
Password="mystrongpass"

```
