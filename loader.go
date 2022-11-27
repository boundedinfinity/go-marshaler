package marshaler

import (
	"io/fs"
	"io/ioutil"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

func MimeTypeFromFile(path string) (mime_type.MimeType, error) {
	if ok, err := pather.IsFile(path); err != nil {
		return mime_type.Unkown, err
	} else {
		if !ok {
			return mime_type.Unkown, ErrNotFileV(path)
		}
	}

	ext := pather.Ext(path)
	mt, err := file_extention.GetMimeType(ext)

	return mt, err
}

func MimeTypeFromPath(root string) (map[string]mime_type.MimeType, error) {
	var m map[string]mime_type.MimeType
	var paths []string

	pather.WalkFiles(root, filterByMimeType, func(path string, info fs.FileInfo) error {
		paths = append(paths, path)
		return nil
	})

	for _, path := range paths {
		mt, err := MimeTypeFromFile(path)

		if err != nil {
			return m, err
		}

		m[path] = mt
	}

	return m, nil
}

func ReadFromFile(path string) ([]byte, mime_type.MimeType, error) {
	mt, err := MimeTypeFromFile(path)

	if err != nil {
		return nil, mime_type.Unkown, err
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, mime_type.Unkown, err
	}

	return data, mt, nil
}

type ReadContent struct {
	Data     []byte
	MimeType mime_type.MimeType
}

func ReadFromPath(root string) (map[string]ReadContent, error) {
	m := make(map[string]ReadContent)
	var paths []string

	pather.WalkFiles(root, filterByMimeType, func(path string, info fs.FileInfo) error {
		paths = append(paths, path)
		return nil
	})

	for _, path := range paths {
		data, mt, err := ReadFromFile(path)

		if err != nil {
			return m, err
		}

		m[path] = ReadContent{
			Data:     data,
			MimeType: mt,
		}
	}

	return m, nil
}
