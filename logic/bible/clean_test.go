package bible

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClean(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{" ", " "},
		{"提摩太后书2:22-3:9", "提摩太后书2:22-3:9"},
		{"激励思考的谈话（提摩太后书2:22-3:9）", "提摩太后书2:22-3:9"},
		{"激励思考的谈话（提摩太后书2:22-3:9", "提摩太后书2:22-3:9"},
		{"提摩太后书2:22-3:9）", "提摩太后书2:22-3:9"},
		{"激励思考的谈话（提摩太后书2:22-3:9）ska", "提摩太后书2:22-3:9"},
		{"激励思考的谈话）（提摩太后书2:22-3:9）", "提摩太后书2:22-3:9"},
		{"激励思考的谈话）（提摩太后书2:22-3:9", "激励思考的谈话）（提摩太后书2:22-3:9"},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output := Clean(test.input)
			require.Equal(t, test.output, output)
		})
	}
}
