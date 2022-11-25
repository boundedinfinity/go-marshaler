package marshaler

import (
	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrMimeTypeUnsupported  = errorer.Errorf("unknown MIME type")
	ErrMimeTypeUnsupportedV = errorer.Errorfn(ErrMimeTypeUnsupported)
	ErrNotFile              = errorer.Errorf("not a file")
	ErrNotFileV             = errorer.Errorfn(ErrNotFile)
)
