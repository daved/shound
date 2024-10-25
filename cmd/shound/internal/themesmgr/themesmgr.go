package themesmgr

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/daved/shound/cmd/shound/internal/config"
	"github.com/daved/shound/cmd/shound/internal/fs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	cp "github.com/otiai10/copy"
)

var tmpDirPrefix = "themes"

type ThemesMgr struct {
	fs            fs.FS
	out           io.Writer
	appName       string
	themeFileName string
	cnf           *config.Sourced
	tmpDirPrefix  string
}

func New(fs fs.FS, out io.Writer, appName, fileName string, cnf *config.Sourced) *ThemesMgr {
	return &ThemesMgr{
		fs:            fs,
		out:           out,
		appName:       appName,
		themeFileName: fileName,
		cnf:           cnf,
		tmpDirPrefix:  tmpDirPrefix,
	}
}

func (m *ThemesMgr) Themes() ([]string, error) {
	themes, err := m.themes()
	if err != nil {
		return nil, fmt.Errorf("themes manager: %w", err)
	}

	return themes, nil
}

func (m *ThemesMgr) themes() ([]string, error) {
	var ts []string

	themesDir := m.cnf.Resolved.ThemesDir

	err := filepath.WalkDir(themesDir, func(path string, de fs.DirEntry, err error) error {
		if filepath.Base(path) == m.themeFileName {
			dir := filepath.Dir(path)
			relToThemesdir := dir[len(themesDir):]
			noSeps := strings.Trim(relToThemesdir, string(os.PathSeparator))
			ts = append(ts, noSeps)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("themes: %w", err)
	}

	return ts, nil
}

func (m *ThemesMgr) SetTheme(theme string) error {
	eMsg := "themes manager: set theme: %w"

	cnfBytes, err := m.fs.ReadFile(m.cnf.Resolved.ConfFilePath)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	updCnfBytes, err := config.SetFileDataField(cnfBytes, config.FileFieldThemeRepo, theme)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	cnf := &config.File{}
	if err := cnf.InitFromYAML(updCnfBytes); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := m.fs.WriteFile(m.cnf.Resolved.ConfFilePath, updCnfBytes, 0600); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (m *ThemesMgr) stagingDirPath(theme string) string {
	return filepath.Join(os.TempDir(), m.appName, m.tmpDirPrefix, theme)
}

func (m *ThemesMgr) loadingDirPath(theme string) string {
	return filepath.Join(m.cnf.Resolved.ThemesDir, theme)
}

func (m *ThemesMgr) prepareDir(path string) error {
	if err := m.fs.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("prepare staging dir: %w", err)
	}

	return nil
}

func (m *ThemesMgr) ValidateThemeRemote(theme, hash string) error {
	stagingDirPath := m.stagingDirPath(theme)
	if err := m.validateThemeRemote(stagingDirPath, theme, hash); err != nil {
		return fmt.Errorf("themes manager: %w", err)
	}

	return nil
}

func (m *ThemesMgr) validateThemeRemote(tmpPath, theme, hash string) error {
	eMsg := "validate theme (remote): %w"

	if err := m.downloadTheme(tmpPath, theme); err != nil {
		if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return fmt.Errorf(eMsg, err)
		}
	}

	if err := m.checkoutInDir(tmpPath, hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := m.validateThemeDir(tmpPath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (m *ThemesMgr) ValidateThemeDir(dir string) error {
	if err := m.validateThemeDir(dir); err != nil {
		return fmt.Errorf("themes manager: %w", err)
	}

	return nil
}

func (m *ThemesMgr) validateThemeDir(dir string) error {
	eMsg := "validate theme directory: %w"

	cnfFilePath := filepath.Join(dir, m.themeFileName)
	cnfBytes, err := m.fs.ReadFile(cnfFilePath)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	cnf := &config.ThemeFile{}
	if err := cnf.InitFromYAML(cnfBytes); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if len(cnf.CmdSounds) == 0 {
		return fmt.Errorf(eMsg, errors.New("theme file is missing CmdSounds entry"))
	}

	for cmd, snd := range cnf.CmdSounds {
		sndFilePath := filepath.Join(dir, snd)
		if _, err := m.fs.Stat(sndFilePath); err != nil {
			return fmt.Errorf(eMsg, fmt.Errorf("verify CmdSounds: %s: %w", cmd, err))
		}
	}

	return nil
}

func (m *ThemesMgr) downloadTheme(dir, theme string) error {
	eMsg := "download theme: %w"

	if err := m.prepareDir(dir); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://%s", theme),
		Progress: m.out,
	})
	if err != nil {
		if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return fmt.Errorf(eMsg, err)
		}

		repo, err := git.PlainOpen(dir)
		if err != nil {
			return fmt.Errorf(eMsg, err)
		}

		if err := repo.Fetch(&git.FetchOptions{}); err != nil {
			if !errors.Is(err, git.NoErrAlreadyUpToDate) {
				return fmt.Errorf(eMsg, err)
			}
		}
	}

	return nil
}

func treeCheckout(t *git.Worktree, hash string) error {
	eMsg := "tree checkout: %w"

	ckOutOpts := &git.CheckoutOptions{}
	ckOutOpts.Hash = plumbing.NewHash(hash)
	if err := t.Checkout(ckOutOpts); err != nil {
		if hash != "" {
			return fmt.Errorf(eMsg, err)
		}

		ckOutOpts.Branch = plumbing.Main
		if err := t.Checkout(ckOutOpts); err != nil {
			return fmt.Errorf(eMsg, err)
		}
	}

	return nil
}

func (m *ThemesMgr) checkoutInDir(dir, hash string) error {
	eMsg := "checkout in dir: %w"

	repo, err := git.PlainOpen(dir)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	tree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := treeCheckout(tree, hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (m *ThemesMgr) copyTheme(src, dst string) error {
	eMsg := "copy theme: %w"

	if err := m.prepareDir(dst); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := cp.Copy(src, dst); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (m *ThemesMgr) AddTheme(theme, hash string) error {
	// use stage theme and load theme here
	eMsg := "themes manager: add theme: %w"

	stagingDirPath := m.stagingDirPath(theme)
	if err := m.validateThemeRemote(stagingDirPath, theme, hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	loadingDirPath := m.loadingDirPath(theme)
	if err := m.deleteDir(loadingDirPath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := m.copyTheme(stagingDirPath, loadingDirPath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (m *ThemesMgr) DeleteTheme(theme string) error {
	loadingDirPath := m.loadingDirPath(theme)
	if err := m.deleteDir(loadingDirPath); err != nil {
		return fmt.Errorf("themes manager: delete theme: %w", err)
	}

	return nil
}

func (m *ThemesMgr) deleteDir(dir string) error {
	if err := m.fs.RemoveAll(dir); err != nil {
		return fmt.Errorf("delete dir: %w", err)
	}

	return nil
}

func (m *ThemesMgr) IsThemeInstalled(theme string) (bool, error) {
	dir := m.loadingDirPath(theme)
	isInstalled, err := m.isThemeInstalled(dir, theme)
	if err != nil {
		return isInstalled, fmt.Errorf("themes manager: %w", err)
	}

	return isInstalled, nil
}

func (m *ThemesMgr) isThemeInstalled(dir, theme string) (bool, error) {
	eMsg := "is theme installed: %w"
	themes, err := m.themes()
	if err != nil {
		return false, fmt.Errorf(eMsg, err)
	}

	inConfig := slices.Contains(themes, theme)
	if !inConfig {
		return false, nil
	}

	if err := m.validateThemeDir(dir); err != nil {
		return false, fmt.Errorf(eMsg, err)
	}

	return true, nil
}
