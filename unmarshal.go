package marshaler

import (
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/trier"
	"github.com/boundedinfinity/mimetyper/file_extention"
	"github.com/boundedinfinity/mimetyper/mime_type"
)

func Unmarshal[T any](ss *[]T, mt mime_type.MimeType, bs []byte) error {
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

func UnmarshalFromFile[T any](path string, out *[]T) error {
	if !pather.IsFile(path) {
		return fmt.Errorf("%v is not a file", path)
	}

	ext := pather.Ext(path)
	mt, err := file_extention.GetMimeType(ext)

	if err != nil {
		return err
	}

	bs, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	var items []T

	if err := Unmarshal(&items, mt, bs); err != nil {
		return err
	}

	*out = append(*out, items...)

	return nil
}

func UnmarshalFromPath[T any](root string, out *[]T) error {
	m := make(map[string][]T)

	if err := UnmarshalWithContextFromPath(root, m); err != nil {
		return err
	}

	for _, items := range m {
		*out = append(*out, items...)
	}

	return nil
}

func UnmarshalWithContextFromPath[T any](root string, out map[string][]T) error {
	var paths []string

	pather.WalkFiles(root, filterByMimeType, func(path string, info fs.FileInfo) error {
		paths = append(paths, path)
		return nil
	})

	for _, path := range paths {
		var items []T

		if err := UnmarshalFromFile(path, &items); err != nil {
			return err
		}

		out[path] = items[:]
	}

	return nil
}

func TryUnmarshalFromPath[T any](root string) trier.Try[[]T] {
	var items []T
	err := UnmarshalFromPath(root, &items)

	return trier.Complete(items, err)
}
