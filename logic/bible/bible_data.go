package bible

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"embed"
	"io"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strconv"

	"google.golang.org/protobuf/proto"
)

//go:embed all:resources/*.pb.gz
var embeddedResources embed.FS

// LoadFromCompressedProtobuf 从压缩的 protobuf 文件加载圣经数据
// 直接返回 protobuf 生成的 BibleData 对象
func LoadFromCompressedProtobuf() (*BibleData, error) {
	// 读取压缩文件
	compressed, err := fs.ReadFile(embeddedResources, "resources/bible-data.pb.gz")
	if err != nil {
		return nil, err
	}

	// 解压数据
	gzReader, err := gzip.NewReader(bytes.NewReader(compressed))
	if err != nil {
		return nil, err
	}
	defer gzReader.Close()

	buffer, err := io.ReadAll(gzReader)
	if err != nil {
		return nil, err
	}

	// 解码 protobuf 消息
	var bibleData BibleData
	if err := proto.Unmarshal(buffer, &bibleData); err != nil {
		return nil, err
	}

	return &bibleData, nil
}

// Bible 表示完整的圣经数据结构
type Bible struct {
	Books map[string]*BibleBook
}

// Book 表示圣经中的一卷书
type BibleBook struct {
	Name         string
	Abbreviation string
	Chapters     [][]string // 每个章节是一个字符串数组，索引0不使用（1-based）
}

// NewBible 创建一个新的 Bible 实例
func NewBible() *Bible {
	return &Bible{
		Books: make(map[string]*BibleBook),
	}
}

// NewBook 创建一个新的 Book 实例
func NewBook(name, abbreviation string) *BibleBook {
	return &BibleBook{
		Name:         name,
		Abbreviation: abbreviation,
		Chapters:     make([][]string, 1), // 第0章不使用
	}
}

// LoadBibleData 从文本文件加载圣经数据
func LoadBibleData(textFilePath string) (*Bible, error) {
	bible := NewBible()

	file, err := os.Open(textFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`^(\S+?)(\d+):(\d+) (.+)?$`)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}

		bookAbbr := match[1]
		chapterNum, _ := strconv.Atoi(match[2])
		verseNum, _ := strconv.Atoi(match[3])
		text := match[4]

		// 获取或创建书卷
		book, exists := bible.Books[bookAbbr]
		if !exists {
			book = NewBook(bookAbbr, bookAbbr)
			bible.Books[bookAbbr] = book
		}

		// 确保有足够的章节
		for len(book.Chapters) <= chapterNum {
			book.Chapters = append(book.Chapters, make([]string, 0))
		}

		// 确保有足够的经文位置 (1-based 索引)
		chapter := book.Chapters[chapterNum]
		for len(chapter) <= verseNum {
			chapter = append(chapter, "")
		}

		// 设置经文文本
		chapter[verseNum] = text
		book.Chapters[chapterNum] = chapter
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bible, nil
}

// SaveToCompressedProtobuf 将圣经数据保存为压缩的 protobuf 格式
func SaveToCompressedProtobuf(bible *Bible, filePath string) error {
	// 转换为 protobuf 结构
	var books []*Book
	bookAbbrevs := make([]string, 0, len(bible.Books))
	for _, book := range bible.Books {
		bookAbbrevs = append(bookAbbrevs, book.Abbreviation)
	}
	sort.Strings(bookAbbrevs)
	for _, bookAbbrev := range bookAbbrevs {
		book := bible.Books[bookAbbrev]
		chapters := make([]*Chapter, 1)

		// 跳过第0章（因为章节是1-based）
		for i := 1; i < len(book.Chapters); i++ {
			chapter := book.Chapters[i]
			var verses []string

			// 跳过第0节（因为经文是1-based）
			for j := 1; j < len(chapter); j++ {
				verses = append(verses, chapter[j])
			}

			chapters = append(chapters, &Chapter{
				Verses: verses,
			})
		}

		books = append(books, &Book{
			Name:         book.Name,
			Abbreviation: book.Abbreviation,
			Chapters:     chapters,
		})
	}

	// 创建 protobuf 消息
	bibleData := &BibleData{
		Books: books,
	}

	// 编码为二进制
	buffer, err := proto.Marshal(bibleData)
	if err != nil {
		return err
	}

	// 压缩数据
	var buf bytes.Buffer
	gzWriter := gzip.NewWriter(&buf)
	if _, err := gzWriter.Write(buffer); err != nil {
		return err
	}
	if err := gzWriter.Close(); err != nil {
		return err
	}

	// 写入文件
	if err := os.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
