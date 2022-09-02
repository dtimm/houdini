package ptyutil

import (
	"os"

	"github.com/pkg/term/termios"
	"golang.org/x/sys/unix"
)

func SetRaw(tty *os.File) error {
	var attr unix.Termios

	_, err := termios.Tcgetattr(tty.Fd())
	if err != nil {
		return err
	}

	termios.Cfmakeraw(&attr)

	return termios.Tcsetattr(tty.Fd(), termios.TCSANOW, &attr)
}
