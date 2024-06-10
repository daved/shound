package thememgr

import "fmt"

type ThemesProvider interface {
	Themes() ([]string, error)
}

type ThemeMgr struct {
	ts ThemesProvider
}

func NewThemeMgr(lp ThemesProvider) *ThemeMgr {
	return &ThemeMgr{
		ts: lp,
	}
}

func (m *ThemeMgr) Themes() ([]string, error) {
	ts, err := m.ts.Themes()
	if err != nil {
		return nil, fmt.Errorf("theme manager: themes: %w", err)
	}

	return ts, nil
}
