package api

import (
	"lmrl/logic/mp3file"
	"lmrl/logic/types"
	"sort"

	"github.com/gin-gonic/gin"
)

func init() {
	go mp3file.GetSermonsFromDir(types.MP3_DIR, true)
}
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

func ListSermon(c *gin.Context) {
	list, err := GetSermonsFromDir(types.MP3_DIR)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to retrieve sermons",
		})
		return
	}
	c.JSON(200, gin.H{
		"list": list,
	})
}
