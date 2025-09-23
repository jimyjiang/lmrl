package bible

import (
	"regexp"
	"strconv"
	"strings"
)

func replaceAbbr(rangeStr string) string {
	for _, replaceStr := range AbbrTable {
		rangeStr = strings.Replace(rangeStr, replaceStr[0], replaceStr[1], -1)
	}
	return rangeStr
}

type BibelVerse struct {
	BookAbbr      string
	ChapterNum    int
	StartVerseNum int
	EndVerseNum   int
}

func ParseBibleVerses(input string) []BibelVerse {
	var result = []BibelVerse{}
	segments := regexp.MustCompile(`[、，；,;]\s*`).Split(replaceAbbr(input), -1)

	var bookAbbr string
	var chapterNum int
	var startVerseNum int
	var endVerseNum int

	push := func(bookAbbr string, chapterNum, startVerseNum, endVerseNum int) {
		if bookAbbr == "" {
			return
		}
		result = append(result, BibelVerse{
			BookAbbr:      bookAbbr,
			ChapterNum:    chapterNum,
			StartVerseNum: startVerseNum,
			EndVerseNum:   endVerseNum,
		})
	}

	for _, segment := range segments {
		startVerseNum = 1
		endVerseNum = -1

		// fmt.Printf("%v", string(segment[len(segment)-1]))
		// if len(segment) > 0 && string(segment[len(segment)-1]) == "篇" {
		// 	segment = segment[:len(segment)-1]
		// }
		if len(segment) > 0 {
			// 将字符串转换为rune切片，正确处理Unicode字符
			runes := []rune(segment)
			lastChar := string(runes[len(runes)-1])
			if lastChar == "篇" || lastChar == "章" {
				segment = string(runes[:len(runes)-1])
			}
		}

		// 处理 "创1:1" 或 "创1:1-2" 或 "创1:1-2:3"
		re := regexp.MustCompile(`^([^\d]+)(\d+):(\d+)-?(\d+)?:?(\d+)?$`)
		match := re.FindStringSubmatch(segment)
		if match != nil {
			if match[3] != "" && match[4] == "" {
				bookAbbr = strings.TrimSpace(match[1])
				chapterNum, _ = strconv.Atoi(match[2])
				startVerseNum, _ = strconv.Atoi(match[3])
				endVerseNum = startVerseNum
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
				continue
			}
			if match[4] != "" && match[5] == "" {
				bookAbbr = strings.TrimSpace(match[1])
				chapterNum, _ = strconv.Atoi(match[2])
				startVerseNum, _ = strconv.Atoi(match[3])
				endVerseNum, _ = strconv.Atoi(match[4])
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
				continue
			}
			if match[5] != "" {
				bookAbbr = strings.TrimSpace(match[1])
				chapterNum, _ = strconv.Atoi(match[2])
				startVerseNum, _ = strconv.Atoi(match[3])
				endVerseNum = -1
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)

				chapterNumMin := chapterNum + 1
				chapterNumMax, _ := strconv.Atoi(match[4])
				for chapterNum = chapterNumMin; chapterNum < chapterNumMax; chapterNum++ {
					startVerseNum = 1
					endVerseNum = -1
					push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
				}
				chapterNum = chapterNumMax
				startVerseNum = 1
				endVerseNum, _ = strconv.Atoi(match[5])
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
				continue
			}
		}

		// 处理 "诗篇4"
		re = regexp.MustCompile(`^([^\d]+)(\d+)$`)
		match = re.FindStringSubmatch(segment)
		if match != nil {
			bookAbbr = strings.TrimSpace(match[1])
			chapterNum, _ = strconv.Atoi(match[2])
			startVerseNum = 1
			endVerseNum = -1
			push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			continue
		}
		// 处理 "4:1-5:3"
		re = regexp.MustCompile(`^(\d+):(\d+)-(\d+):(\d+)$`)
		match = re.FindStringSubmatch(segment)
		if match != nil {
			chapterNum, _ = strconv.Atoi(match[1])
			startVerseNum, _ = strconv.Atoi(match[2])
			endVerseNum = -1
			push(bookAbbr, chapterNum, startVerseNum, endVerseNum)

			chapterNumMin := chapterNum + 1
			chapterNumMax, _ := strconv.Atoi(match[3])
			for chapterNum = chapterNumMin; chapterNum < chapterNumMax; chapterNum++ {
				startVerseNum = 1
				endVerseNum = -1
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			}
			chapterNum = chapterNumMax
			startVerseNum = 1
			endVerseNum, _ = strconv.Atoi(match[4])
			push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			continue
		}
		// 处理 "3:1-2"
		re = regexp.MustCompile(`^(\d+):(\d+)(?:-(\d+))?$`)
		match = re.FindStringSubmatch(segment)
		if match != nil {
			if match[2] != "" && match[3] == "" {
				chapterNum, _ = strconv.Atoi(match[1])
				startVerseNum, _ = strconv.Atoi(match[2])
				endVerseNum = startVerseNum
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			}
			if match[3] != "" {
				chapterNum, _ = strconv.Atoi(match[1])
				startVerseNum, _ = strconv.Atoi(match[2])
				endVerseNum, _ = strconv.Atoi(match[3])
				push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			}
			continue
		}

		// 处理 "9-18"
		re = regexp.MustCompile(`^(\d+)-(\d+)$`)
		match = re.FindStringSubmatch(segment)
		if match != nil {
			startVerseNum, _ = strconv.Atoi(match[1])
			endVerseNum, _ = strconv.Atoi(match[2])
			push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			continue
		}

		// 处理 "6"
		re = regexp.MustCompile(`^(\d+)$`)
		match = re.FindStringSubmatch(segment)
		if match != nil {
			chapterNum, _ = strconv.Atoi(match[1])
			startVerseNum = 1
			endVerseNum = -1
			push(bookAbbr, chapterNum, startVerseNum, endVerseNum)
			continue
		}
	}

	return result
}
