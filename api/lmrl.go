package api

import (
	"lmrl/logic/mp3file"
	"lmrl/logic/types"
	"sort"

	"github.com/gin-gonic/gin"
)

// func GetSermonsFromDir(dirPath string) ([]*types.Sermon, error) {
// 	var sermons []*types.Sermon

// 	mp3Cache := cache.GetMp3Cache()

// 	// 读取目录
// 	entries, err := os.ReadDir(dirPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("读取目录失败: %v", err)
// 	}

// 	// 收集MP3文件信息
// 	for _, entry := range entries {
// 		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".mp3") {
// 			continue
// 		}

// 		filePath := filepath.Join(dirPath, entry.Name())
// 		if cachedSermon, ok := mp3Cache.Get(filePath); ok {
// 			sermons = append(sermons, cachedSermon)
// 			continue
// 		}
// 		sermon, err := mp3file.ParseMP3File(filePath)
// 		if err != nil {
// 			fmt.Printf("解析文件 %s 失败: %v\n", entry.Name(), err)
// 			continue
// 		}
// 		mp3Cache.Set(filePath, sermon)

// 		sermons = append(sermons, sermon)
// 	}

// 	// 按日期倒序排列
// 	sort.Slice(sermons, func(i, j int) bool {
// 		return sermons[i].Date > sermons[j].Date
// 	})

// 	return sermons, nil
// }

func GetSermonsFromDir(dirPath string) ([]*types.Sermon, error) {
	mapSermons, err := mp3file.GetSermonsFromDir(dirPath, true)
	if err != nil {
		return nil, err
	}
	sermons := make([]*types.Sermon, 0, len(mapSermons))
	for _, sermon := range mapSermons {
		sermons = append(sermons, sermon)
	}
	// 按日期倒序排列
	sort.Slice(sermons, func(i, j int) bool {
		return sermons[i].Date > sermons[j].Date
	})

	return sermons, nil
}

func LMRL(c *gin.Context) {
	list, err := GetSermonsFromDir(types.MP3_DIR)
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
