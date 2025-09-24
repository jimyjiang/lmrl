package jobs

import (
	"fmt"
	"io"
	"lmrl/logic"
	"lmrl/logic/cache"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	checkURL string
	maxFiles = 60 // 最大保留文件数
)

const (
	filePrefix = "mw"     // 文件前缀
	fileExt    = ".mp3"   // 文件扩展名
	timeFormat = "060102" // 年月日格式(YYMMDD)
)

func init() {

	checkURL = fmt.Sprintf("https://z.lydt.work/ly/audio/%d/mw/", time.Now().Year())
}

func RegisterDownloadMp3Job() {
	// 创建目录(如果不存在)
	// if err := os.MkdirAll(MP3_Files, 0755); err != nil {
	// 	fmt.Printf("创建目录失败: %v\n", err)
	// 	return
	// }
	Init()
	// 添加定时任务：每天8点到16点每小时执行一次（8,9,10,...,16）
	_, err := scheduler.AddFunc("0 8-16 * * *", func() {
		fmt.Printf("执行时间: %v\n", time.Now())
		if err := RunDownloadMp3Job(time.Now()); err != nil {
			fmt.Printf("任务执行失败: %v\n", err)
		}
	})
	if err != nil {
		fmt.Printf("创建定时任务失败: %v\n", err)
		return
	}
	fmt.Println("定时任务已注册，将在每天8点到16点每小时执行一次")
}

// RunDownloadMp3Job 执行检查下载任务
func RunDownloadMp3Job(t time.Time) error {
	fileName := generateFileName(t)
	filePath := filepath.Join(logic.MP3_DIR, fileName)

	// 检查文件是否已存在
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("文件已存在: %s\n", fileName)
		return nil
	}

	// 文件不存在，尝试下载
	fmt.Printf("尝试下载文件: %s\n", fileName)
	if err := downloadFile(fileName); err != nil {
		return fmt.Errorf("下载失败: %v", err)
	}

	// 清理旧文件
	if err := cleanupOldFiles(); err != nil {
		return fmt.Errorf("清理旧文件失败: %v", err)
	}
	if err := rebuildMp3Cache(); err != nil {
		return fmt.Errorf("重建MP3缓存失败: %v", err)
	}

	return nil
}

// generateFileName 生成当前日期的文件名
func generateFileName(t time.Time) string {
	return fmt.Sprintf("%s%s%s", filePrefix, t.Format(timeFormat), fileExt)
}

// downloadFile 下载文件
func downloadFile(fileName string) error {
	fileURL := checkURL + fileName
	resp, err := http.Get(fileURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("文件不存在(404)")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP请求失败: %s", resp.Status)
	}

	filePath := filepath.Join(logic.MP3_DIR, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// cleanupOldFiles 清理旧文件，保留最新的maxFiles个文件
func cleanupOldFiles() error {
	files, err := filepath.Glob(filepath.Join(logic.MP3_DIR, filePrefix+"*"+fileExt))
	if err != nil {
		return err
	}

	if len(files) <= maxFiles {
		return nil
	}

	// 按文件名排序(最早的在前)
	sort.Strings(files)

	// 删除最早的文件
	for i := 0; i < len(files)-maxFiles; i++ {
		fmt.Printf("删除旧文件: %s\n", filepath.Base(files[i]))
		if err := os.Remove(files[i]); err != nil {
			return err
		}
	}

	return nil
}
func rebuildMp3Cache() error {
	m := map[cache.FileName]*logic.Sermon{}

	// 读取目录
	entries, err := os.ReadDir(logic.MP3_DIR)
	if err != nil {
		return fmt.Errorf("读取目录失败: %v", err)
	}

	// 收集MP3文件信息
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".mp3") {
			continue
		}

		filePath := filepath.Join(logic.MP3_DIR, entry.Name())
		if _, ok := m[filePath]; ok {
			continue
		}
		sermon, err := logic.ParseMP3File(filePath)
		if err != nil {
			fmt.Printf("解析文件 %s 失败: %v\n", entry.Name(), err)
			continue
		}
		m[filePath] = sermon
	}
	cache.GetMp3Cache().ReBuild(m)
	return nil
}
