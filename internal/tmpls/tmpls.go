package tmpls

import (
	"io"
	"strings"
	"text/template"

	"github.com/daved/shound/internal/config"
)

var x = strings.TrimSpace(`
{{range $alias, $sound := .CmdsSounds -}}
if ! (alias {{$alias}} 2>/dev/null | grep "_shound" &>/dev/null); then
	_shound_{{$alias}}={{$alias}}
	alias {{$alias}} &>/dev/null && _shound_{{$alias}}="$(alias {{$alias}} | cut -d "=" -f2-)" && _shound_{{$alias}}="${_shound_{{$alias}}:1:${#_shound_{{$alias}}}-2}"
	alias {{$alias}}="(\$(shound identify --playcmd {{$alias}}) &) && $_shound_{{$alias}}"
fi
{{end}}

{{if .NotFoundSound}}
function command_not_found_handle() {
	($(shound identify --playcmd {{.NotFoundKey}}) &)
	printf "%s: command not found\n" "$1" >&2
	return 127
}
{{end}}
`)

type AliasesData struct {
	CmdsSounds    config.CmdsSounds
	NotFoundKey   string
	NotFoundSound string
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
