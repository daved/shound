package list

import (
	"io"
	"strings"
	"text/template"

	"github.com/daved/shound/internal/config"
)

var listTmplText = strings.TrimSpace(`
Themes Directory: {{.ThemesDir}}
{{if .}}{{end}}
`)

func fprintList(w io.Writer, d *config.Config) error {
	aliasesTmpl, err := template.New("theme-info").Parse(listTmplText)
	if err != nil {
		return err
	}

	return aliasesTmpl.Execute(w, d)
}
