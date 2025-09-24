package mp3file

import (
	"fmt"
	"lmrl/logic/types"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dhowden/tag"
	"github.com/hajimehoshi/go-mp3"
)

func ParseMP3File(filePath string) (*types.Sermon, error) {
	// 获取文件基本信息（大小、修改时间等）
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	fileName := filepath.Base(filePath)
	fileSize := formatFileSize(fileInfo.Size())
	modTime := fileInfo.ModTime().Format("2006-01-02") // 文件修改时间（YYYY-MM-DD）

	// 打开文件读取元数据（ID3标签）
	file, err := os.Open(filePath)
	if err != nil {
		return fallbackParse(filePath, fileInfo), nil // 打开失败时回退到基础解析
	}
	defer file.Close()

	// 解析MP3元数据（使用dhowden/tag库）
	meta, err := tag.ReadFrom(file)
	if err != nil {
		return fallbackParse(filePath, fileInfo), nil // 元数据读取失败时回退
	}

	// 1. 提取日期（优先级：标题 > 文件名 > 文件修改时间）
	dateStr := ""
	// fmt.Printf("Title:%s\n", meta.Title())
	// 从标题提取（如："旷野吗哪-20250910" → 2025-09-10）
	if titleDate := extractDateFromTitle(meta.Title()); titleDate != "" {
		dateStr = titleDate
		// fmt.Printf("TitleDate:%s\n", dateStr)
	} else if fileNameDate := extractDateFromFileName(fileName); fileNameDate != "" {
		// 从文件名提取（如："mw250910.mp3" → 2025-09-10）
		dateStr = fileNameDate
	} else {
		// 回退到文件修改时间
		dateStr = modTime
	}

	// 2. 提取讲道主题（优先级：注释 > 标题 > 文件名）
	mainTitle := ""
	// 从注释提取（如："对潜能的衡量在于神的同在与应许 (士师记6:11-18)" → 提取主题）
	// fmt.Printf("Comment:%s\n", meta.Comment())
	if comment := meta.Comment(); comment != "" {
		mainTitle = extractMainTitle(comment)
		// fmt.Printf("MainTitle:%s\n", mainTitle)
	}
	// 从标题提取（清理前缀和日期）
	if mainTitle == "" && meta.Title() != "" {
		mainTitle = cleanTitle(meta.Title())
	}
	// 回退到文件名（不含扩展名）
	if mainTitle == "" {
		mainTitle = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	}

	// 4. 构建Sermon对象（讲员固定为"孙大中"，因节目特性）
	return &types.Sermon{
		Filename: fileName,              // 文件名（如：mw250910.mp3）
		Title:    mainTitle,             // 讲道主题（如："对潜能的衡量在于神的同在与应许"）
		Date:     dateStr,               // 日期（如：2025-09-10）
		Speaker:  "孙大中",                 // 固定讲员（旷野吗哪节目主持）
		Duration: getDuration(filePath), // 时长（如：29:20）
		FileSize: fileSize,              // 文件大小（如：14.1 MB）
	}, nil
}

// 格式化文件大小为人类可读格式
func formatFileSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// fallbackParse 元数据读取失败时的回退解析
func fallbackParse(filePath string, fileInfo os.FileInfo) *types.Sermon {
	fileName := filepath.Base(filePath)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return &types.Sermon{
		Filename: fileName,
		Title:    baseName,
		Date:     fileInfo.ModTime().Format("2006-01-02"),
		Speaker:  "孙大中",
		Duration: "0分钟",
		FileSize: formatFileSize(fileInfo.Size()),
	}
}

// extractDateFromTitle 从标题提取日期（如："旷野吗哪-20250910" → 2025-09-10）
func extractDateFromTitle(title string) string {
	re := regexp.MustCompile(`(\d{4})(\d{2})(\d{2})`) // 匹配8位日期（YYYYMMDD）
	matches := re.FindStringSubmatch(title)
	if len(matches) == 4 {
		return fmt.Sprintf("%s-%s-%s", matches[1], matches[2], matches[3]) // 格式化为YYYY-MM-DD
	}
	return ""
}

// extractMainTitle 从注释提取主题（如："对潜能的衡量在于神的同在与应许 (士师记6:11-18)" → 提取括号外文本）
func extractMainTitle(comment string) string {
	re := regexp.MustCompile(`^([^()]+)`) // 匹配第一个"("前的文本
	matches := re.FindStringSubmatch(comment)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1]) // 去除首尾空格
	}
	return ""
}

// extractDateFromFileName 从文件名提取日期（如："mw250910.mp3" → 2025-09-10）
func extractDateFromFileName(fileName string) string {
	re := regexp.MustCompile(`mw(\d{2})(\d{2})(\d{2})`) // 匹配"mw+6位数字"格式
	matches := re.FindStringSubmatch(fileName)
	if len(matches) == 4 { // matches[1]=25, matches[2]=09, matches[3]=10
		return fmt.Sprintf("20%s-%s-%s", matches[1], matches[2], matches[3]) // 2025-09-10
	}
	return ""
}

// cleanTitle 清理标题（移除"旷野吗哪-"前缀和日期数字）
func cleanTitle(title string) string {
	title = strings.TrimPrefix(title, "旷野吗哪-")                      // 移除前缀
	title = regexp.MustCompile(`\d{8}`).ReplaceAllString(title, "") // 移除8位日期（如20250910）
	return strings.TrimSpace(title)                                 // 去除首尾空格
}

func getDuration(fileInfo string) string {
	duration := "0分钟" // 默认值
	file, err := os.Open(fileInfo)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return duration
	}
	defer file.Close()
	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		fmt.Println("Error creating decoder:", err)
		return duration
	}

	// 计算时长
	samples := decoder.Length() / 4 // 每个样本4字节
	return formatDuration(time.Duration(samples) * time.Second / time.Duration(decoder.SampleRate()))
}

func formatDuration(d time.Duration) string {
	// 创建一个基准时间加上duration
	t := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC).Add(d)
	// 格式化为分钟:秒
	return t.Format("04:05") // 注意这里的格式字符串
}
