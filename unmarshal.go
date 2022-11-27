package marshaler

import (
	"encoding/json"
	"io/fs"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"gopkg.in/yaml.v3"
)

func UnmarshalFromBytes[T any](data []byte, v T, mt mime_type.MimeType) error {
	realmt := mime_type.ResolveMimeType(mt)
	var err error

	switch realmt {
	case mime_type.ApplicationJson:
		err = json.Unmarshal(data, v)
	case mime_type.ApplicationXYaml:
		err = yaml.Unmarshal(data, v)
	default:
		err = ErrMimeTypeUnsupportedV(mt)
	}

	return err
}

func UnmarshalFromFile[T any](path string) (T, mime_type.MimeType, error) {
	var v T

	data, mt, err := ReadFromFile(path)

	if err != nil {
		return v, mime_type.Unkown, err
	}

	if err := UnmarshalFromBytes(data, &v, mt); err != nil {
		return v, mime_type.Unkown, err
	}

	return v, mt, nil
}

type UnmarshaledContent[V any] struct {
	Value    V
	MimeType mime_type.MimeType
}

func UnmarshalFromPath[T any](root string, m map[string]UnmarshaledContent[T]) error {
	var paths []string

	pather.WalkFiles(root, filterByMimeType, func(path string, info fs.FileInfo) error {
		paths = append(paths, path)
		return nil
	})

	for _, path := range paths {
		v, mt, err := UnmarshalFromFile[T](path)

		if err != nil {
			return err
		}

		m[path] = UnmarshaledContent[T]{
			Value:    v,
			MimeType: mt,
		}
	}

	return nil
}
