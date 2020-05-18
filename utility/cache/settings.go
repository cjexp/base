package cache

import (
	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/constant"
)

const (
	cachePrefix         = constant.CachePrefix
	cachePrefixModified = constant.CachePrefixModified
	cacheFileFolderName = constant.CacheFileFolderName
)

func GetSettings(context ctx.BackgroundContext) *Settings {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return &Settings{
			CachePrefix:         cachePrefix,
			CachePrefixModified: cachePrefixModified,
			CacheFileFolderName: cacheFileFolderName,
		}, nil
	}).(*Settings)
}

type Settings struct {
	CachePrefix         string
	CachePrefixModified string
	CacheFileFolderName string
}
