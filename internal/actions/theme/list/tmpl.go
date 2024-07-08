package list

import (
	"io"
	"strings"
	"text/template"
)

var listTmplText = strings.TrimSpace(`
Themes Directory: {{.ThemesDir}}
Available Themes:{{range $theme := .Themes}}
  {{$theme -}}
{{end}}
{{if .}}{{end}}
`)

type listData struct {
	ThemesDir string
	Themes    []string
}

func makeListData(dir string, themes []string) listData {
	return listData{
		ThemesDir: dir,
		Themes:    themes,
	}
}

func fprintList(w io.Writer, d listData) error {
	aliasesTmpl, err := template.New("theme-info").Parse(listTmplText)
	if err != nil {
		return err
	}

	return aliasesTmpl.Execute(w, d)
}
