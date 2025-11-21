package bible

import (
	"fmt"
	"lmrl/logic/arithmetic"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

// SearchResult 表示搜索结果
type SearchResult struct {
	Reference string
	Text      string
}

// Search 统一检索入口
func Search(bible *BibleData, query string) []*SearchResult {
	// 尝试作为引用查询
	if refResults := SearchByReference(bible, query); refResults != nil {
		return refResults
	}

	// 否则作为全文精确搜索
	if refResults := FullTextSearch(bible, query); refResults != nil {
		return refResults
	}
	// 否则作为全文模糊搜索
	return FullTextSearchByCommonSubstring(bible, query)
}

// parseReference 解析引用格式（如"约4:43"或"约4:43-54"）
func parseReference(ref string) (bookAbbr string, chapterNum, startVerse, endVerse int, ok bool) {
	re := regexp.MustCompile(`^(\S+?)(\d+):(\d+)(?:-(\d+))?$`)
	match := re.FindStringSubmatch(ref)
	if match == nil {
		return "", 0, 0, 0, false
	}

	bookAbbr = match[1]
	chapterNum, _ = strconv.Atoi(match[2])
	startVerse, _ = strconv.Atoi(match[3])

	if match[4] != "" {
		endVerse, _ = strconv.Atoi(match[4])
	} else {
		endVerse = startVerse
	}

	return bookAbbr, chapterNum, startVerse, endVerse, true
}

// SearchByReference 精确/范围检索
func SearchByReference(bible *BibleData, ref string) []*SearchResult {
	bookAbbr, chapterNum, startVerse, endVerse, ok := parseReference(ref)
	if !ok {
		return nil
	}

	return SearchByRange(bible, bookAbbr, chapterNum, startVerse, endVerse)
}

// SearchByRange 范围检索
func SearchByRange(bible *BibleData, bookAbbr string, chapterNum, startVerse, endVerse int) []*SearchResult {
	// 查找书卷
	var book *Book
	for _, b := range bible.GetBooks() {
		if b.GetAbbreviation() == bookAbbr {
			book = b
			break
		}
	}
	if book == nil {
		return nil
	}

	// 检查章节是否存在
	if chapterNum <= 0 || chapterNum >= len(book.GetChapters()) {
		return nil
	}

	chapter := book.GetChapters()[chapterNum]
	// fmt.Printf("chapter: %v\n", chapter)
	verses := chapter.GetVerses()

	// 处理 endVerse 为 -1 的情况（表示章节末尾）
	if endVerse == -1 {
		endVerse = len(verses)
	}

	// 收集结果
	var results []*SearchResult
	for verseNum := startVerse; verseNum <= endVerse; verseNum++ {
		if verseNum > 0 && verseNum <= len(verses) {
			text := verses[verseNum-1]
			// if text != "" {
			results = append(results, &SearchResult{
				Reference: fmt.Sprintf("%s%d:%d", bookAbbr, chapterNum, verseNum),
				Text:      text,
			})
			// }
		}
	}

	return results
}

// FullTextSearch 全文检索
func FullTextSearch(bible *BibleData, query string) []*SearchResult {
	if query == "" {
		return nil
	}

	queryArr := strings.Split(query, " ")
	queryArr = lo.Map(queryArr, func(item string, index int) string {
		return strings.Trim(item, " ")
	})
	queryArr = lo.Filter(queryArr, func(item string, index int) bool {
		return item != ""
	})

	var results []*SearchResult

	// 遍历所有书卷
	for _, book := range bible.GetBooks() {
		bookAbbr := book.GetAbbreviation()

		// 遍历所有章节
		for chapterIdx, chapter := range book.GetChapters() {
			chapterNum := chapterIdx

			// 遍历所有经文
			for verseIdx, verse := range chapter.GetVerses() {
				verseNum := verseIdx + 1

				if lo.EveryBy(queryArr, func(item string) bool {
					return strings.Contains(verse, item)
				}) {
					results = append(results, &SearchResult{
						Reference: fmt.Sprintf("%s%d:%d", bookAbbr, chapterNum, verseNum),
						Text:      verse,
					})
				}
				// 检查是否包含查询词
				// if strings.Contains(verse, query) {
				// 	results = append(results, &SearchResult{
				// 		Reference: fmt.Sprintf("%s%d:%d", bookAbbr, chapterNum, verseNum),
				// 		Text:      verse,
				// 	})
				// }
			}
		}
	}

	return results
}

// FullTextSearch 全文检索
func FullTextSearchByCommonSubstring(bible *BibleData, query string) []*SearchResult {
	query = strings.ReplaceAll(query, " ", "")
	if query == "" {
		return nil
	}

	var results []*SearchResult

	// 遍历所有书卷
	for _, book := range bible.GetBooks() {
		bookAbbr := book.GetAbbreviation()

		// 遍历所有章节
		for chapterIdx, chapter := range book.GetChapters() {
			chapterNum := chapterIdx

			// 遍历所有经文
			for verseIdx, verse := range chapter.GetVerses() {
				verseNum := verseIdx + 1

				length := arithmetic.LongCommonSubstring(query, verse)
				queryLength := math.Ceil(float64(len([]rune(query))) * 0.8)
				if length >= queryLength {
					results = append(results, &SearchResult{
						Reference: fmt.Sprintf("%s%d:%d", bookAbbr, chapterNum, verseNum),
						Text:      verse,
					})
				}

			}
		}
	}

	return results
}
