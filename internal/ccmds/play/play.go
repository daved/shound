package play

import (
	"context"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/daved/clic"
	"github.com/daved/flagset"
	"github.com/daved/shound/internal/ccmds/ccmd"
	"github.com/daved/shound/internal/config"
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/flac"
	"github.com/gopxl/beep/speaker"
)

type Play struct {
	out io.Writer

	fs  *flagset.FlagSet
	cnf *config.Config
}

func New(out io.Writer, name string, cnf *config.Config) *Play {
	fs := flagset.New(name)

	c := Play{
		out: out,
		fs:  fs,
		cnf: cnf,
	}

	return &c
}

func (c *Play) AsClic(subs ...*clic.Clic) *clic.Clic {
	cmd := clic.New(c, subs...)
	cmd.Meta[clic.MetaKeyCmdDesc] = "Play audio file associated with the provided command"
	cmd.Meta[clic.MetaKeyArgsHint] = "<command_name>"

	return cmd
}

func (c *Play) FlagSet() *flagset.FlagSet {
	return c.fs
}

func (c *Play) HandleCommand(ctx context.Context, cmd *clic.Clic) error {
	if err := ccmd.HandleHelpFlag(c.out, cmd, c.cnf.Help); err != nil {
		return err
	}

	args := c.fs.Args()
	if len(args) == 0 {
		return errors.New("play: no command name")
	}
	arg := args[0]

	sound, ok := c.cnf.CmdSounds[arg]
	if !ok && arg == c.cnf.NotFoundKey {
		sound = c.cnf.NotFoundSound
	}

	soundPath := filepath.Join(c.cnf.ThemeDir, sound)

	f, err := os.Open(soundPath)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := flac.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	return nil
}
