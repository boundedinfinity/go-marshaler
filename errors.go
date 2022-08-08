package marshaler

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

var (
	ErrMimeTypeUnsupported  = errors.New("unknown MIME type")
	ErrMimeTypeUnsupportedV = func(v mime_type.MimeType) error {
		return errV(v.String(), ErrMimeTypeUnsupported)
	}
)

func errV(v string, err error) error {
	return fmt.Errorf("%v : %w", v, err)
}
