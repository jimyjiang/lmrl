package mp3file

import (
	"context"
	"fmt"
	"lmrl/logic/cache"
	"lmrl/logic/types"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var ch = make(chan func(), 100)

func worker(ctx context.Context) {
	for {
		select {
		case f := <-ch:
			f()
		case <-ctx.Done():
			return
		}
	}
}
func StartWorker(ctx context.Context) {
	for i := 0; i < 4; i++ {
		go worker(ctx)
	}
}
func GetSermonsFromDir(dirPath string, setCache bool) (map[types.FileName]*types.Sermon, error) {
	mp3Cache := cache.GetMp3Cache()
	// 读取目录
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("读取目录失败: %v", err)
	}
	m := make(map[types.FileName]*types.Sermon)
	mlock := sync.Mutex{}
	setMap := func(fileName string, sermon *types.Sermon) {
		mlock.Lock()
		defer mlock.Unlock()
		m[fileName] = sermon
	}
	wg := sync.WaitGroup{}
	wg.Add(len(entries))
	// 收集MP3文件信息
	for _, entry := range entries {
		ch <- func(entry os.DirEntry) func() {
			return func() {
				defer wg.Done()
				if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".mp3") {
					return
				}
				filePath := filepath.Join(dirPath, entry.Name())
				if cachedSermon, ok := mp3Cache.Get(filePath); ok {
					setMap(filePath, cachedSermon)
					return
				}
				sermon, err := ParseMP3File(filePath)
				if err != nil {
					fmt.Printf("解析文件 %s 失败: %v\n", entry.Name(), err)
					return
				}
				if setCache {
					mp3Cache.Set(filePath, sermon)
				}
				setMap(filePath, sermon)
			}
		}(entry)
	}
	wg.Wait()

	return m, nil
}
