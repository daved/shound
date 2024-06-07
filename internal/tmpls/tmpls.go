package tmpls

import (
	"io"
	"strings"
	"text/template"
)

var aliasesTmplText = strings.TrimSpace(`
{{range $alias, $sanitized := .Aliases -}}
if ! (alias {{$alias}} 2>/dev/null | grep "shound identify" &>/dev/null); then
	_shound_{{$sanitized}}="{{$alias}}"
	alias {{$alias}} &>/dev/null && _shound_{{$sanitized}}="$(alias {{$alias}} | cut -d "=" -f2-)" && _shound_{{$sanitized}}="${_shound_{{$sanitized}}:1:${#_shound_{{$sanitized}}}-2}"
	alias {{$alias}}="(\$(shound identify --playcmd {{$alias}}) &) && $_shound_{{$sanitized}}"
fi
{{end}}

{{if .NotFoundSound -}}
function command_not_found_handle() {
	($(shound identify --playcmd {{.NotFoundKey}}) &)
	printf "%s: command not found\n" "$1" >&2
	return 127
}
{{end}}
`)

type AliasesData struct {
	Aliases       map[string]string // map[Alias]SaneAlias
	NotFoundKey   string
	NotFoundSound string
}

func MakeAliasesData(notFoundKey, notFoundSound string, aliases []string) AliasesData {
	aliasMap := make(map[string]string)
	for _, alias := range aliases {
		aliasMap[alias] = strings.ReplaceAll(alias, "-", "__")
	}

	return AliasesData{
		Aliases:       aliasMap,
		NotFoundKey:   notFoundKey,
		NotFoundSound: notFoundSound,
	}
}

type Tmpls struct {
	FprintAliases func(io.Writer, AliasesData) error
}

func NewTmpls() (*Tmpls, error) {
	aliasesTmpl, err := template.New("aliases").Parse(aliasesTmplText)
	if err != nil {
		return nil, err
	}

	ts := &Tmpls{
		FprintAliases: func(w io.Writer, d AliasesData) error {
			return aliasesTmpl.Execute(w, d)
		},
	}

	return ts, nil
}
