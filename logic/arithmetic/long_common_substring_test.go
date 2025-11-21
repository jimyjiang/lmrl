package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongCommonSubstring(t *testing.T) {
	testcase := []struct {
		s1       string
		s2       string
		expected float64
	}{
		{"", "", 0},
		{"abcdefghijklmnopqrstuvwxyz", "", 0},
		{"", "abcdefghijklmnopqrstuvwxyz", 0},
		{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz", 26},
		{"渊面黑暗", "地是空虚混沌。渊面黑暗。神的灵运行在水面上。", 4},
		{"做神百般恩的好的管家", "各人要照所得的恩赐彼此服事，作神百般恩赐的好管家。", 8},
	}

	for _, tc := range testcase {
		result := LongCommonSubstring(tc.s1, tc.s2)
		require.Equal(t, tc.expected, result)
	}
}
