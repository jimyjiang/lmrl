package bible

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	bibleData, err := LoadFromCompressedProtobuf()
	require.NoError(t, err)

	bookAbbr, chapterNum, startVerse, endVerse, ok := parseReference("创1:2")
	require.True(t, ok)
	require.Equal(t, "创", bookAbbr)
	require.Equal(t, 1, chapterNum)
	require.Equal(t, 2, startVerse)
	require.Equal(t, 2, endVerse)

	result := Search(bibleData, "创1:1")
	require.NotNil(t, result)
	require.Equal(t, 1, len(result))
	require.Equal(t, "创1:1", result[0].Reference)
	require.Equal(t, "起初神创造天地。", result[0].Text)

	result = Search(bibleData, "渊面黑暗")
	require.NotNil(t, result)
	require.Equal(t, 1, len(result))
	require.Equal(t, "创1:2", result[0].Reference)
	require.Equal(t, "地是空虚混沌。渊面黑暗。神的灵运行在水面上。", result[0].Text)

	result = Search(bibleData, "来13:25")
	require.NotNil(t, result)
	require.Equal(t, 1, len(result))
	require.Equal(t, "来13:25", result[0].Reference)
	require.Equal(t, "愿恩惠常与你们众人同在。阿们。", result[0].Text)

	result = Search(bibleData, "来14:12")
	require.Nil(t, result)
}

func TestFullTextSearchByCommonSubstring(t *testing.T) {
	bibleData, err := LoadFromCompressedProtobuf()
	require.NoError(t, err)

	result := FullTextSearchByCommonSubstring(bibleData, "渊面黑暗")
	// t.Logf("result: %v", result)
	require.NotNil(t, result)
	require.Equal(t, 1, len(result))
	require.Equal(t, "创1:2", result[0].Reference)
	require.Equal(t, "地是空虚混沌。渊面黑暗。神的灵运行在水面上。", result[0].Text)

	result = FullTextSearchByCommonSubstring(bibleData, "做神百般恩的好的管家")
	require.NotNil(t, result)
	require.Equal(t, 1, len(result))
	require.Equal(t, "彼前4:10", result[0].Reference)
	require.Equal(t, "各人要照所得的恩赐彼此服事，作神百般恩赐的好管家。", result[0].Text)

}
