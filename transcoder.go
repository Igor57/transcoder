package transcoder

import (
	"io"
	"os/exec"
)

// Transcoder ...
type Transcoder interface {
	Start(optsBeforeInput Options, opts Options) (<-chan Progress, error)
	StartAndReturnCmd(optsBeforeInput Options, opts Options) (<-chan Progress, *exec.Cmd, error)
	Input(i string) Transcoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	Output(o string) Transcoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	WithOptions(opts Options) Transcoder
	WithAdditionalOptions(opts Options) Transcoder
	GetMetadata() (Metadata, error)
}
