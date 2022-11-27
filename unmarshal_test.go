package marshaler_test

import (
	"testing"

	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalFromPath(t *testing.T) {
	type Type1 struct {
		Thing string `json:"thing" yaml:"thing"`
	}

	expected := map[string]marshaler.UnmarshaledContent[Type1]{
		"test_data/concrete/test.json": {
			Value:    Type1{Thing: "a"},
			MimeType: mime_type.ApplicationJson,
		},
		"test_data/concrete/test.yaml": {
			Value:    Type1{Thing: "b"},
			MimeType: mime_type.ApplicationXYaml},
		"test_data/concrete/test.yml": {
			Value:    Type1{Thing: "c"},
			MimeType: mime_type.ApplicationXYaml,
		},
		"test_data/concrete/dir1/test.json": {
			Value:    Type1{Thing: "a"},
			MimeType: mime_type.ApplicationJson,
		},
		"test_data/concrete/dir1/test.yaml": {
			Value:    Type1{Thing: "b"},
			MimeType: mime_type.ApplicationXYaml,
		},
		"test_data/concrete/dir1/test.yml": {
			Value:    Type1{Thing: "c"},
			MimeType: mime_type.ApplicationXYaml,
		},
	}

	actual := make(map[string]marshaler.UnmarshaledContent[Type1])
	err := marshaler.UnmarshalFromPath("./test_data/concrete", actual)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
