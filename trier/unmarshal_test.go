package trier_test

import (
	"testing"

	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-marshaler/trier"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalFromPath_normal(t *testing.T) {
	type Type1 struct {
		Thing string `json:"thing" yaml:"thing"`
	}

	expected := map[string]marshaler.UnmarshaledContent[Type1]{
		"../test_data/concrete/test.json": {
			Value:    Type1{Thing: "a"},
			MimeType: mime_type.ApplicationJson,
		},
		"../test_data/concrete/test.yaml": {
			Value:    Type1{Thing: "b"},
			MimeType: mime_type.ApplicationXYaml},
		"../test_data/concrete/test.yml": {
			Value:    Type1{Thing: "c"},
			MimeType: mime_type.ApplicationXYaml,
		},
		"../test_data/concrete/dir1/test.json": {
			Value:    Type1{Thing: "a"},
			MimeType: mime_type.ApplicationJson,
		},
		"../test_data/concrete/dir1/test.yaml": {
			Value:    Type1{Thing: "b"},
			MimeType: mime_type.ApplicationXYaml,
		},
		"../test_data/concrete/dir1/test.yml": {
			Value:    Type1{Thing: "c"},
			MimeType: mime_type.ApplicationXYaml,
		},
	}

	actual := trier.UnmarshalFromPath[Type1]("../test_data/concrete")

	assert.Nil(t, actual.Error)
	assert.Equal(t, expected, actual.Result)
}

func Test_UnmarshalFromPath_generic(t *testing.T) {
	type Type1[T any] struct {
		Thing T `json:"thing" yaml:"thing"`
	}

	type Type2[T any] struct {
		Type1[T] `yaml:",inline"`
	}

	expected := map[string]marshaler.UnmarshaledContent[Type2[string]]{
		"../test_data/concrete/test.json": {
			Value:    Type2[string]{Type1[string]{Thing: "a"}},
			MimeType: mime_type.ApplicationJson,
		},
		"../test_data/concrete/test.yaml": {
			Value:    Type2[string]{Type1[string]{Thing: "b"}},
			MimeType: mime_type.ApplicationXYaml},
		"../test_data/concrete/test.yml": {
			Value:    Type2[string]{Type1[string]{Thing: "c"}},
			MimeType: mime_type.ApplicationXYaml,
		},
		"../test_data/concrete/dir1/test.json": {
			Value:    Type2[string]{Type1[string]{Thing: "a"}},
			MimeType: mime_type.ApplicationJson,
		},
		"../test_data/concrete/dir1/test.yaml": {
			Value:    Type2[string]{Type1[string]{Thing: "b"}},
			MimeType: mime_type.ApplicationXYaml,
		},
		"../test_data/concrete/dir1/test.yml": {
			Value:    Type2[string]{Type1[string]{Thing: "c"}},
			MimeType: mime_type.ApplicationXYaml,
		},
	}

	actual := trier.UnmarshalFromPath[Type2[string]]("../test_data/concrete")

	assert.Nil(t, actual.Error)
	assert.Equal(t, expected, actual.Result)
}
