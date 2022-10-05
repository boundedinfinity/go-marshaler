package marshaler

import "github.com/boundedinfinity/mimetyper/mime_type"

func Marshal[T any](ss *[]T, mt mime_type.MimeType, bs []byte) error {
	realmt := mime_type.ResolveMimeType(mt)

	switch realmt {
	case mime_type.ApplicationJson:
		return unmarshalJson(ss, bs)
	case mime_type.ApplicationXYaml:
		return unmarshalYaml(ss, bs)
	default:
		return ErrMimeTypeUnsupportedV(mt)
	}
}
