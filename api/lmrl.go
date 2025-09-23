package api

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/dhowden/tag"
	"github.com/gin-gonic/gin"
)

type Sermon struct {
	Filename string // 文件名
	Title    string // 讲道主题
	Date     string // 日期，格式如"2025-09-10"
	Speaker  string // 讲员姓名
	Duration string // 时长，如"45分钟"
	FileSize string // 文件大小，如"12MB"
}

func GetSermonsFromDir(dirPath string) ([]Sermon, error) {
	var sermons []Sermon

	// 读取目录
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录失败: %v", err)
	}

	// 收集MP3文件信息
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".mp3") {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())
		sermon, err := parseMP3File(filePath)
		if err != nil {
			fmt.Printf("解析文件 %s 失败: %v\n", entry.Name(), err)
			continue
		}

		sermons = append(sermons, *sermon)
	}

	// 按日期倒序排列
	sort.Slice(sermons, func(i, j int) bool {
		return sermons[i].Date > sermons[j].Date
	})

	return sermons, nil
}

func parseMP3File(filePath string) (*Sermon, error) {
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
	// 从标题提取（如："旷野吗哪-20250910" → 2025-09-10）
	if titleDate := extractDateFromTitle(meta.Title()); titleDate != "" {
		dateStr = titleDate
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
	if comment := meta.Comment(); comment != "" {
		mainTitle = extractMainTitle(comment)
	}
	// 从标题提取（清理前缀和日期）
	if mainTitle == "" && meta.Title() != "" {
		mainTitle = cleanTitle(meta.Title())
	}
	// 回退到文件名（不含扩展名）
	if mainTitle == "" {
		mainTitle = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	}

	// 3. 提取时长（从元数据的"持续时间"字段，如："29:20"）
	duration := GetDuration(meta, fileInfo)
	// duration := meta.Duration() // 直接获取元数据中的时长字符串（如"29:20"）
	// if duration == "" {
	// 	duration = "0分钟" // 若元数据无时长，显示默认值
	// }

	// 4. 构建Sermon对象（讲员固定为"孙大中"，因节目特性）
	return &Sermon{
		Filename: fileName,  // 文件名（如：mw250910.mp3）
		Title:    mainTitle, // 讲道主题（如："对潜能的衡量在于神的同在与应许"）
		Date:     dateStr,   // 日期（如：2025-09-10）
		Speaker:  "孙大中",     // 固定讲员（旷野吗哪节目主持）
		Duration: duration,  // 时长（如：29:20）
		FileSize: fileSize,  // 文件大小（如：14.1 MB）
	}, nil
}

func GetDuration(meta tag.Metadata, fileInfo os.FileInfo) string {
	// 获取或估算音频时长（兼容所有tag.Metadata实现）
	duration := "0分钟" // 默认值

	// 方法1：尝试从注释中提取时长（如注释中包含"29:20"）
	if comment := meta.Comment(); comment != "" {
		if dur := extractDurationFromComment(comment); dur != "" {
			duration = dur
		}
	}

	// 方法2：根据文件大小估算（假设平均比特率为128kbps）
	if duration == "0分钟" {
		estimated := estimateDurationFromFileSize(fileInfo.Size())
		if estimated > 0 {
			duration = formatDuration(estimated)
		}
	}
	return duration
}

// 辅助函数：根据文件大小估算时长（秒）
func estimateDurationFromFileSize(size int64) int {
	const bitrate = 128000 // 假设平均比特率为128kbps
	if size <= 0 || bitrate <= 0 {
		return 0
	}
	return int(float64(size) * 8 / float64(bitrate))
}

// 辅助函数：格式化秒数为"MM:SS"或"X分钟"
func formatDuration(seconds int) string {
	mins := seconds / 60
	secs := seconds % 60
	if secs == 0 {
		return fmt.Sprintf("%d分钟", mins)
	}
	return fmt.Sprintf("%d:%02d", mins, secs)
}

// 辅助函数：从注释文本提取时长（如"时长：29:20"）
func extractDurationFromComment(comment string) string {
	re := regexp.MustCompile(`(?:时长|时间|duration)[:：\s]*(\d+[:分]\d+|[^\s]+)`)
	matches := re.FindStringSubmatch(comment)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
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

// extractDateFromFileName 从文件名提取日期（如："mw250910.mp3" → 2025-09-10）
func extractDateFromFileName(fileName string) string {
	re := regexp.MustCompile(`mw(\d{2})(\d{2})(\d{2})`) // 匹配"mw+6位数字"格式
	matches := re.FindStringSubmatch(fileName)
	if len(matches) == 4 { // matches[1]=25, matches[2]=09, matches[3]=10
		return fmt.Sprintf("20%s-%s-%s", matches[1], matches[2], matches[3]) // 2025-09-10
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

// cleanTitle 清理标题（移除"旷野吗哪-"前缀和日期数字）
func cleanTitle(title string) string {
	title = strings.TrimPrefix(title, "旷野吗哪-")                      // 移除前缀
	title = regexp.MustCompile(`\d{8}`).ReplaceAllString(title, "") // 移除8位日期（如20250910）
	return strings.TrimSpace(title)                                 // 去除首尾空格
}

// fallbackParse 元数据读取失败时的回退解析
func fallbackParse(filePath string, fileInfo os.FileInfo) *Sermon {
	fileName := filepath.Base(filePath)
	baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	return &Sermon{
		Filename: fileName,
		Title:    baseName,
		Date:     fileInfo.ModTime().Format("2006-01-02"),
		Speaker:  "孙大中",
		Duration: "0分钟",
		FileSize: formatFileSize(fileInfo.Size()),
	}
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

func LMRL(c *gin.Context) {
	list, err := GetSermonsFromDir("/root/web/灵命日粮")
	// list, err := GetSermonsFromDir("/Users/jimmy.jiang/doc/基督/灵命日粮/202509")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to retrieve sermons",
		})
		return
	}
	c.HTML(200, "index.tpl", gin.H{
		"SermonList": list,
	})
}
