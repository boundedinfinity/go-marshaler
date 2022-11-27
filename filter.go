package marshaler

import (
	"os"

	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

var (
	SUPPORTED_TYPES = []mime_type.MimeType{
		mime_type.ApplicationYaml,
		mime_type.ApplicationJson,
	}
)

func filterByMimeType(path string, info os.FileInfo) bool {
	supportedExts, _ := file_extention.GetExts(SUPPORTED_TYPES...)
	ext := pather.Ext(path)

	if slicer.Contains(supportedExts, file_extention.FileExtention(ext)) {
		return true
	}

	return false
}
