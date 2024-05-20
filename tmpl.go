package main

import "strings"

var x = strings.TrimSpace(`
alias _shound="pw-cat --playback"
alias hw="(_shound \$HOME/.cache/shound/star_trek/alert19.flac &) && echo hello world"
`) + "\n"

// TODO: setup templating properly
