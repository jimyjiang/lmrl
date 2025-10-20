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
	t.Logf("bookAbbr: %s, chapterNum: %d, startVerse: %d, endVerse: %d", bookAbbr, chapterNum, startVerse, endVerse)
	result := Search(bibleData, "创1:2")
	for _, r := range result {
		t.Logf("result: %s", r.Text)
	}
	result = Search(bibleData, "渊面黑暗")
	for _, r := range result {
		t.Logf("result: %s", r.Text)
	}
}
