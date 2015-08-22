package wincolog

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	_ "github.com/mattn/go-isatty" // make cross-compiling easier
)

// WinColorTerminal wraps a colorable Writer
// with a ColorSupporter interface
type WinColorTerminal struct {
	output io.Writer
}

// Write implements the io.Writer interface
func (w *WinColorTerminal) Write(p []byte) (n int, err error) {
	return w.output.Write(p)
}

// ColorSupported implements the ColorSupporter interface
func (w *WinColorTerminal) ColorSupported() bool {
	return true
}

// Stdout returns a WinColorTerminal when stdout is a windows terminal
// returns the built-in os.Stdout otherwise
func Stdout() io.Writer {
	stdout := colorable.NewColorableStdout()
	if stdout == os.Stdout {
		return os.Stdout
	}

	return &WinColorTerminal{output: stdout}
}

// Stderr returns a WinColorTerminal when stderr is a windows terminal
// returns the built-in os.Stderr otherwise
func Stderr() io.Writer {
	stderr := colorable.NewColorableStderr()
	if stderr == os.Stderr {
		return os.Stderr
	}

	return &WinColorTerminal{output: stderr}
}
