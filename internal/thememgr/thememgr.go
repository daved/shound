package thememgr

type ThemeMgr struct {
	themes []string
}

func NewThemeMgr(themes []string) *ThemeMgr {
	return &ThemeMgr{
		themes: themes,
	}
}

func (m *ThemeMgr) Themes() []string {
	return m.themes
}
