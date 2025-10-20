package bible

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadFromCompressedProtobuf(t *testing.T) {
	data, err := LoadFromCompressedProtobuf()
	require.NoError(t, err)
	require.NotNil(t, data)
}
