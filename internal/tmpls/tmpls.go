package main

import (
	"io"
	"strings"
	"text/template"
)

var x = strings.TrimSpace(`
{{$soundDir := .SoundDir -}} 

alias _shound="{{.PlayCmd}}"
{{range $alias, $sound := .CmdsSounds -}}
if ! (alias {{$alias}} 2>/dev/null | grep "_shound" &>/dev/null); then
	_shound_{{$alias}}={{$alias}}
	alias {{$alias}} &>/dev/null && _shound_{{$alias}}="$(alias {{$alias}} | cut -d "=" -f2-)" && _shound_{{$alias}}="${_shound_{{$alias}}:1:${#_shound_{{$alias}}}-2}"
	alias {{$alias}}="(_shound \"{{$soundDir}}/{{$sound}}\" &) && $_shound_{{$alias}}"
fi
{{end}}

{{if .NoCmdSound}}
function command_not_found_handle() {
	(_shound "{{$soundDir}}/{{.NoCmdSound}}" &)
	printf "%s: command not found\n" "$1" >&2
	return 127
}
{{end}}
`)

type AliasesData struct {
	CmdsSounds CmdsSounds
	NoCmdSound string
	PlayCmd    string
	SoundDir   string
}

type Tmpls struct {
	Aliases func(io.Writer, AliasesData) error
}

func NewTmpls() (*Tmpls, error) {
	aliasesTmpl, err := template.New("aliases").Parse(x)
	if err != nil {
		return nil, err
	}

	ts := &Tmpls{
		Aliases: func(w io.Writer, d AliasesData) error {
			return aliasesTmpl.Execute(w, d)
		},
	}

	return ts, nil
}
