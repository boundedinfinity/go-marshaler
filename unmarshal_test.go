package marshaler_test

import (
	"testing"

	"github.com/boundedinfinity/go-marshaler"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalFromPath(t *testing.T) {
	type Type1 struct {
		Thing string `json:"thing" yaml:"thing"`
	}

	expected := []Type1{
		{Thing: "a"}, {Thing: "b"}, {Thing: "c"},
		{Thing: "a"}, {Thing: "b"}, {Thing: "c"},
	}
	var actual []Type1

	err := marshaler.UnmarshalFromPath("./test_data", &actual)

	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, actual)
}

func Test_UnmarshalWithContextFromPath(t *testing.T) {
	type Type1 struct {
		Thing string `json:"thing" yaml:"thing"`
	}

	expected := map[string][]Type1{
		"test_data/test.json":      {{Thing: "a"}},
		"test_data/test.yaml":      {{Thing: "b"}},
		"test_data/test.yml":       {{Thing: "c"}},
		"test_data/dir1/test.json": {{Thing: "a"}},
		"test_data/dir1/test.yaml": {{Thing: "b"}},
		"test_data/dir1/test.yml":  {{Thing: "c"}},
	}

	actual := make(map[string][]Type1)
	err := marshaler.UnmarshalWithContextFromPath("./test_data", actual)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
