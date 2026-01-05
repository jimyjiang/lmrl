package jobs

import (
	"fmt"
	"io"
	"lmrl/logic/cache"
	"lmrl/logic/mp3file"
	"lmrl/logic/types"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var (
	maxFiles = 60 // 最大保留文件数
)

const (
	filePrefix = "mw"     // 文件前缀
	fileExt    = ".mp3"   // 文件扩展名
	timeFormat = "060102" // 年月日格式(YYMMDD)
)

func RegisterDownloadMp3Job() {
	Init()
	// 添加定时任务：每天5点到10点每小时执行一次（5,6,7,...,10）
	_, err := scheduler.AddFunc("0 5-10 * * *", func() {
		fmt.Printf("执行时间: %v\n", time.Now())
		if err := RunDownloadMp3Job(time.Now()); err != nil {
			fmt.Printf("任务执行失败: %v\n", err)
		}
	})
	if err != nil {
		fmt.Printf("创建定时任务失败: %v\n", err)
		return
	}
	fmt.Println("定时任务已注册，将在每天5点到10点每小时执行一次")
}

// RunDownloadMp3Job 执行检查下载任务
func RunDownloadMp3Job(t time.Time) error {
	fileName := generateFileName(t)
	filePath := filepath.Join(types.MP3_DIR, fileName)

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

func getCheckURL() string {
	return fmt.Sprintf("https://x.lydt.work/storage/ly/audio/%d/mw/", time.Now().Year())
}

// downloadFile 下载文件
func downloadFile(fileName string) error {
	fileURL := getCheckURL() + fileName
	fmt.Printf("HTTP请求: %v\n", fileURL)
	resp, err := http.Get(fileURL)
	if err != nil {
		fmt.Printf("HTTP请求失败: %v\n", fileURL)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("文件不存在(404)")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP请求失败: %s", resp.Status)
	}

	filePath := filepath.Join(types.MP3_DIR, fileName)
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
	files, err := filepath.Glob(filepath.Join(types.MP3_DIR, filePrefix+"*"+fileExt))
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
	mapSermons, err := mp3file.GetSermonsFromDir(types.MP3_DIR, false)
	if err != nil {
		return err
	}
	cache.GetMp3Cache().ReBuild(mapSermons)
	return nil
}
