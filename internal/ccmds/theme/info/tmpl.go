package info

import (
	"io"
	"strings"
	"text/template"

	"github.com/daved/shound/internal/config"
)

var infoTmplText = strings.TrimSpace(`
Name: {{.ThemeRepo}}
{{if .}}{{end}}
`)

func fprintInfo(w io.Writer, d *config.Config) error {
	aliasesTmpl, err := template.New("theme-info").Parse(infoTmplText)
	if err != nil {
		return err
	}

	return aliasesTmpl.Execute(w, d)
}
