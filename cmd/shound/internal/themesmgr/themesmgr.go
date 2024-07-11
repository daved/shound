package themesmgr

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/daved/shound/internal/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	cp "github.com/otiai10/copy"
)

var (
	tmpDirPrefix = "themes"
)

type ThemesMgr struct {
	appName       string
	out           io.Writer
	cnf           *config.Config
	tmpDirPrefix  string
	themesDir     string
	themeFileName string
}

func New(appName string, out io.Writer, cnf *config.Config, fileName string) *ThemesMgr {
	return &ThemesMgr{
		appName:       appName,
		out:           out,
		cnf:           cnf,
		tmpDirPrefix:  tmpDirPrefix,
		themesDir:     cnf.User.File.ThemesDir,
		themeFileName: fileName,
	}
}

func (i *ThemesMgr) Themes() ([]string, error) {
	themes, err := i.themes()
	if err != nil {
		return nil, fmt.Errorf("themes manager: %w", err)
	}

	return themes, nil
}

func (i *ThemesMgr) themes() ([]string, error) {
	var ts []string

	err := filepath.WalkDir(i.themesDir, func(path string, de fs.DirEntry, err error) error {
		if filepath.Base(path) == i.themeFileName {
			dir := filepath.Dir(path)
			relToThemesdir := dir[len(i.themesDir):]
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

func (i *ThemesMgr) SetTheme(theme string) error {
	eMsg := "themes manager: set theme: %w"

	cnfBytes, err := os.ReadFile(i.cnf.User.Flags.ConfFilePath)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	updCnfBytes, err := config.SetFileDataField(cnfBytes, config.FileFieldThemeRepo, theme)
	if err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := os.WriteFile(i.cnf.User.Flags.ConfFilePath, updCnfBytes, 0600); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := i.cnf.User.File.InitFromYAML(cnfBytes); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (i *ThemesMgr) stagingDirPath(theme string) string {
	return filepath.Join(os.TempDir(), i.appName, i.tmpDirPrefix, theme)
}

func (i *ThemesMgr) loadingDirPath(theme string) string {
	return filepath.Join(i.themesDir, theme)
}

func (i *ThemesMgr) prepareDir(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("prepare staging dir: %w", err)
	}

	return nil
}

func (i *ThemesMgr) ValidateThemeRemote(theme, hash string) error {
	stagingDirPath := i.stagingDirPath(theme)
	if err := i.validateThemeRemote(stagingDirPath, theme, hash); err != nil {
		return fmt.Errorf("themes manager: %w", err)
	}

	return nil
}

func (i *ThemesMgr) validateThemeRemote(tmpPath, theme, hash string) error {
	eMsg := "validate theme (remote): %w"

	if err := i.downloadTheme(tmpPath, theme); err != nil {
		if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return fmt.Errorf(eMsg, err)
		}
	}

	if err := i.checkoutInDir(tmpPath, hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := i.validateThemeDir(tmpPath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (i *ThemesMgr) ValidateThemeDir(dir string) error {
	if err := i.validateThemeDir(dir); err != nil {
		return fmt.Errorf("themes manager: %w", err)
	}

	return nil
}

func (i *ThemesMgr) validateThemeDir(dir string) error {
	// TODO: validateThemeDir()
	// ensure themeconfig file loads

	// check basic values

	// iterate over expected audio files and verify existence

	return nil
}

func (i *ThemesMgr) downloadTheme(dir, theme string) error {
	eMsg := "download theme: %w"

	if err := i.prepareDir(dir); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      fmt.Sprintf("https://%s", theme),
		Progress: i.out,
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

func (i *ThemesMgr) checkoutInDir(dir, hash string) error {
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

func (i *ThemesMgr) copyTheme(src, dst string) error {
	eMsg := "copy theme: %w"

	if err := i.prepareDir(dst); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := cp.Copy(src, dst); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (i *ThemesMgr) AddTheme(theme, hash string) error {
	// use stage theme and load theme here
	eMsg := "themes manager: add theme: %w"

	stagingDirPath := i.stagingDirPath(theme)
	if err := i.validateThemeRemote(stagingDirPath, theme, hash); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	loadingDirPath := i.loadingDirPath(theme)
	if err := i.deleteDir(loadingDirPath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	if err := i.copyTheme(stagingDirPath, loadingDirPath); err != nil {
		return fmt.Errorf(eMsg, err)
	}

	return nil
}

func (i *ThemesMgr) DeleteTheme(theme string) error {
	loadingDirPath := i.loadingDirPath(theme)
	if err := i.deleteDir(loadingDirPath); err != nil {
		return fmt.Errorf("themes manager: delete theme: %w", err)
	}

	return nil
}

func (i *ThemesMgr) deleteDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("delete dir: %w", err)
	}

	return nil
}

func (i *ThemesMgr) IsThemeInstalled(theme string) (bool, error) {
	dir := i.loadingDirPath(theme)
	isInstalled, err := i.isThemeInstalled(dir, theme)
	if err != nil {
		return isInstalled, fmt.Errorf("themes manager: %w", err)
	}

	return isInstalled, nil
}

func (i *ThemesMgr) isThemeInstalled(dir, theme string) (bool, error) {
	eMsg := "is theme installed: %w"
	themes, err := i.themes()
	if err != nil {
		return false, fmt.Errorf(eMsg, err)
	}

	inConfig := slices.Contains(themes, theme)
	if !inConfig {
		return false, nil
	}

	if err := i.validateThemeDir(dir); err != nil {
		return false, fmt.Errorf(eMsg, err)
	}

	return true, nil
}
