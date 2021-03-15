package main

import "sync"

var (
	flyweightMapFactory *FlyweightMapFactory
	once                sync.Once
)

// FlyweightMapFactory 享元工厂
type FlyweightMapFactory struct {
	maps map[int]*FlyweightMap
	mu   sync.Mutex
}

func GetFlyweightMapFactory() *FlyweightMapFactory {
	once.Do(func() {
		maps := make(map[int]*FlyweightMap)
		flyweightMapFactory = &FlyweightMapFactory{maps: maps}
	})

	return flyweightMapFactory
}

func (fmf *FlyweightMapFactory) GetFlyweightMap(userId int) *FlyweightMap {
	fmf.mu.Lock()
	defer fmf.mu.Unlock()
	if _, ok := fmf.maps[userId]; !ok {
		fmf.maps[userId] = &FlyweightMap{mapSource: make(map[int]bool)}
	}

	return fmf.maps[userId]
}

// FlyweightMap 原神中房主的地图资源
type FlyweightMap struct {
	mapSource map[int]bool
}

func (fm *FlyweightMap) GetSource(sourceId int) bool {
	return fm.mapSource[sourceId]
}
