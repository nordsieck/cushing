package baton

import (
	"io"
	"os"
	"time"

	isatty "github.com/mattn/go-isatty"
)

var states = []byte(`-/|\`)

const backspace = byte(8)

type Twirler interface {
	Twirl() error
}

// Show progress in a tty
type Baton struct {
	W     io.Writer
	Start time.Time
	Count int
}

func NewTwirler() Twirler {
	if isatty.IsTerminal(os.Stderr.Fd()) || isatty.IsCygwinTerminal(os.Stderr.Fd()) {
		os.Stderr.Write([]byte{backspace}) // best effort
		return &Baton{Start: time.Now()}
	}
	return &Phantom{Start: time.Now()}
}

func (b *Baton) Twirl() error {
	w := b.W
	if w == nil {
		w = os.Stderr
	}
	c := b.Count % len(states)
	if _, err := w.Write([]byte{states[c], backspace}); err != nil {
		return err
	}
	b.Count += 1
	return nil
}

// Stub for non-ttyl use
type Phantom struct{ Start time.Time }

func (p *Phantom) Twirl() error { return nil }
