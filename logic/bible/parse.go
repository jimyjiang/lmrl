package bible

import (
	"regexp"
	"strconv"
	"strings"
)

var AbbrTable = [][2]string{
	{"创世记", "创"}, {"出埃及记", "出"}, {"利未记", "利"}, {"民数记", "民"}, {"申命记", "申"},
	{"约书亚记", "书"}, {"士师记", "士"}, {"路得记", "得"}, {"撒母耳记上", "撒上"}, {"撒母耳记下", "撒下"},
	{"列王纪上", "王上"}, {"列王纪下", "王下"}, {"历代志上", "代上"}, {"历代志下", "代下"}, {"以斯拉记", "拉"},
	{"尼希米记", "尼"}, {"以斯帖记", "斯"}, {"约伯记", "伯"}, {"诗篇", "诗"}, {"箴言", "箴"},
	{"传道书", "传"}, {"雅歌", "歌"}, {"以赛亚书", "赛"}, {"耶利米书", "耶"}, {"耶利米哀歌", "哀"},
	{"以西结书", "结"}, {"但以理书", "但"}, {"何西阿书", "何"}, {"约珥书", "珥"}, {"阿摩司书", "摩"},
	{"俄巴底亚书", "俄"}, {"约拿书", "拿"}, {"弥迦书", "弥"}, {"那鸿书", "鸿"}, {"哈巴谷书", "哈"},
	{"西番雅书", "番"}, {"哈该书", "该"}, {"撒迦利亚书", "亚"}, {"玛拉基书", "玛"}, {"马太福音", "太"},
	{"马可福音", "可"}, {"路加福音", "路"}, {"约翰福音", "约"}, {"使徒行传", "徒"}, {"罗马书", "罗"},
	{"哥林多前书", "林前"}, {"哥林多后书", "林后"}, {"加拉太书", "加"}, {"以弗所书", "弗"}, {"腓立比书", "腓"},
	{"歌罗西书", "西"}, {"帖撒罗尼迦前书", "帖前"}, {"帖撒罗尼迦后书", "帖后"}, {"提摩太前书", "提前"},
	{"提摩太后书", "提后"}, {"提多书", "多"}, {"腓利门书", "门"}, {"希伯来书", "来"}, {"雅各书", "雅"},
	{"彼得前书", "彼前"}, {"彼得后书", "彼后"}, {"约翰一书", "约壹"}, {"约翰二书", "约贰"}, {"约翰三书", "约叁"},
	{"犹大书", "犹"}, {"启示录", "启"},
}

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
			if lastChar == "篇" {
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
