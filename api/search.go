package api

import (
	"lmrl/logic/bible"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bibleData *bible.BibleData

func init() {
	var err error
	bibleData, err = bible.LoadFromCompressedProtobuf()
	if err != nil {
		panic(err)
	}
}

type SearchResult struct {
	Reference string `json:"reference"`
	Text      string `json:"text"`
}

func Search(c *gin.Context) {
	q := c.Query("q")
	q = bible.Clean(q)
	queries := bible.ParseBibleVerses(q)
	data := []*bible.SearchResult{}
	if len(queries) == 0 {
		data = bible.Search(bibleData, q)
	} else {
		for _, query := range queries {
			data = append(data, bible.SearchByRange(bibleData, query.BookAbbr, query.ChapterNum, query.StartVerseNum, query.EndVerseNum)...)
		}
	}

	results := []SearchResult{}
	for _, result := range data {
		results = append(results, SearchResult{
			Reference: result.Reference,
			Text:      result.Text,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}
