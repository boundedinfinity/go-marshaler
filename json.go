package marshaler

import (
	"encoding/json"
	"fmt"
	"strings"
)

func unmarshalJson[T any](ss *[]T, bs []byte) error {
	s := string(bs)
	s = strings.TrimSpace(s)
	f := s[0:1]

	switch f {
	case "{":
		var x T

		if err := json.Unmarshal(bs, &s); err != nil {
			return err
		}

		*ss = append(*ss, x)
	case "[":
		var xs []T

		if err := json.Unmarshal(bs, &xs); err != nil {
			return err
		}

		*ss = append(*ss, xs...)
	default:
		return fmt.Errorf("unsupported JSON")
	}

	return nil
}
