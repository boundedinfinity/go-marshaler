package marshaler

import (
	"bytes"
	"errors"
	"io"

	"gopkg.in/yaml.v2"
)

func unmarshalYaml[T any](ss *[]T, bs []byte) error {
	d := yaml.NewDecoder(bytes.NewReader(bs))

	for {
		var s *T
		err := d.Decode(s)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		if s == nil {
			continue
		}

		*ss = append(*ss, *s)
	}

	return nil
}
