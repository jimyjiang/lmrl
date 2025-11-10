package main

import (
	"fmt"
	"lmrl/logic/bible"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Print("参数不正确")
		return
	}
	data, err := bible.LoadFromCompressedProtobuf()
	if err != nil {
		fmt.Print("Error loading bible data: " + err.Error())
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

	segments := make([]string, 0, len(results))
	for _, raw := range results {
		segments = append(segments, fmt.Sprintf("%s %s", raw.Reference, raw.Text))
	}
	if len(segments) > 0 {
		fmt.Print(strings.Join(segments, "\n"))
	} else {
		fmt.Println("未找到结果")
	}
}
