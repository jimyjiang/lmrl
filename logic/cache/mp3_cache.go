package cache

import (
	"encoding/json"
	"lmrl/logic"
	"sync"
)

var mp3Cache = &Mp3Cache{}

type FileName = string
type Mp3Cache struct {
	m map[FileName]*logic.Sermon
	sync.RWMutex
}

func GetMp3Cache() *Mp3Cache {
	if mp3Cache == nil {
		mp3Cache = &Mp3Cache{}
	}
	return mp3Cache
}

func (cache *Mp3Cache) Get(fileName string) (*logic.Sermon, bool) {
	cache.RLock()
	defer cache.RUnlock()

	s, ok := cache.m[fileName]
	return s, ok
}
func (cache *Mp3Cache) Set(fileName string, sermon *logic.Sermon) {
	cache.Lock()
	defer cache.Unlock()

	cache.m[fileName] = sermon
}
func (cache *Mp3Cache) ReBuild(m map[FileName]*logic.Sermon) {
	cache.Lock()
	defer cache.Unlock()

	cache.m = m
}

func (cache *Mp3Cache) String() string {
	cache.RLock()
	defer cache.RUnlock()

	bufs, _ := json.MarshalIndent(cache.m, "", "  ")
	return string(bufs)
}
