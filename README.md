# shound

    go install github.com/daved/shound/cmd/shound

## Additional Installation Steps

### Setup a Config File

```
mkdir ~/.config/shound
curl -o ~/.config/shound/config.yaml https://raw.githubusercontent.com/daved/shound/main/cmd/shound/example.config.yaml
```

### Update Shell Config
```
# add the following or equivalent to the shell rc file (e.g. ~/.bashrc)
type shound &>/dev/null && eval "$(shound export)"

```

## Usage

```
shound -h       # shows all top-level subcommands
shound -h theme # shows subcommands for the theme top-level subcommand
```

```
shound theme install github.com/daved/shound-star_trek
shound theme set github.com/daved/shound-star_trek
```

Then modify the config file (~/.config/shound/config.yaml) as preferred. The value that will most likely require a change is the command to use for playing audio. It may end up being something like aplay or paplay. Some modifications of the config file may require the shell to be restarted.

## Note

This readme was written up from memory, so my apologies if something is out of order. There's simply not a lot of time available for me to test this at the moment.
