package trier

import (
	"github.com/boundedinfinity/go-commoner/trier"
	"github.com/boundedinfinity/go-marshaler"
)

func UnmarshalFromPath[T any](root string) trier.Try[map[string]marshaler.UnmarshaledContent[T]] {
	m := make(map[string]marshaler.UnmarshaledContent[T])
	err := marshaler.UnmarshalFromPath(root, m)
	return trier.Complete(m, err)
}
