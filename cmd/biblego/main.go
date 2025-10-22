package main

import (
	"fmt"
	"lmrl/logic/bible"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("参数不正确")
		return
	}
	data, err := bible.LoadFromCompressedProtobuf()
	if err != nil {
		fmt.Println("Error loading bible data: " + err.Error())
	}
	argq := args[1]
	queries := bible.ParseBibleVerses(bible.Clean(argq))
	results := []*bible.SearchResult{}
	if len(queries) == 0 {
		results = bible.Search(data, argq)
	} else {
		for _, query := range queries {
			results = append(results, bible.SearchByRange(data, query.BookAbbr, query.ChapterNum, query.StartVerseNum, query.EndVerseNum)...)
		}
	}

	for _, raw := range results {
		fmt.Printf("%s %s\n", raw.Reference, raw.Text)
	}
}
