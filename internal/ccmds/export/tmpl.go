package export

import (
	"io"
	"strings"
	"text/template"
)

var aliasesTmplText = strings.TrimSpace(`
{{range $alias, $sanitized := .Aliases -}}
if ! (alias {{$alias}} 2>/dev/null | grep "shound play" &>/dev/null); then
	_shound_{{$sanitized}}="{{$alias}}"
	alias {{$alias}} &>/dev/null && _shound_{{$sanitized}}="$(alias {{$alias}} | cut -d "=" -f2-)" && _shound_{{$sanitized}}="${_shound_{{$sanitized}}:1:${#_shound_{{$sanitized}}}-2}"
	alias {{$alias}}="(shound play {{$alias}} &) && $_shound_{{$sanitized}}"
fi
{{end}}

{{if .NotFoundSound -}}
function command_not_found_handle() {
	(shound play {{.NotFoundKey}} &)
	printf "%s: command not found\n" "$1" >&2
	return 127
}
{{end}}
`)

type aliasesData struct {
	Aliases       map[string]string // map[Alias]SaneAlias
	NotFoundKey   string
	NotFoundSound string
}

func makeAliasesData(notFoundKey, notFoundSound string, aliases []string) aliasesData {
	aliasMap := make(map[string]string)
	for _, alias := range aliases {
		aliasMap[alias] = strings.ReplaceAll(alias, "-", "__")
	}

	return aliasesData{
		Aliases:       aliasMap,
		NotFoundKey:   notFoundKey,
		NotFoundSound: notFoundSound,
	}
}

func fprintAliases(w io.Writer, d aliasesData) error {
	aliasesTmpl, err := template.New("aliases").Parse(aliasesTmplText)
	if err != nil {
		return err
	}

	return aliasesTmpl.Execute(w, d)
}
