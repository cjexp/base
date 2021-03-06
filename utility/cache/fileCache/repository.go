package fileCache

import (
	"fmt"
	"time"

	"github.com/cjtoolkit/ctx/v2/ctxHttp"

	"github.com/cjexp/base/utility/cache"
	"github.com/cjexp/base/utility/cache/internal"
	"github.com/cjexp/base/utility/loggers"
	"github.com/cjtoolkit/ctx/v2"
)

func GetCacheRepository(context ctx.Context) cache.Repository {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initCacheRepository(context), nil
	}).(cache.Repository)
}

type cacheRepository struct {
	prefix       string
	core         Core
	errorService loggers.ErrorService
}

func initCacheRepository(context ctx.Context) cache.Repository {
	return &cacheRepository{
		prefix:       cache.GetSettings(context).CachePrefix,
		core:         GetCore(context),
		errorService: loggers.GetErrorService(context),
	}
}

func (c *cacheRepository) Persist(name string, expiration time.Duration, miss cache.Miss, hit cache.Hit) interface{} {
	name = fmt.Sprintf(c.prefix, name)

	var (
		data interface{}
		b    []byte
		err  error
	)

	if b, err = c.core.GetBytesCheck(name, expiration); err == nil {
		data, err = hit(b)
		c.errorService.CheckErrorAndPanic(err)
	} else {
		data, b, err = miss()
		c.errorService.CheckErrorAndPanic(err)
		c.core.SetBytes(name, b, expiration)
	}

	return data
}

func GetModifiedRepository(context ctx.Context) cache.ModifiedRepository {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initCacheModifiedRepository(context), nil
	}).(cache.ModifiedRepository)
}

type cacheModifiedRepository struct {
	prefix          string
	core            Core
	cacheRepository cache.Repository
	errorService    loggers.ErrorService
}

func initCacheModifiedRepository(context ctx.Context) cache.ModifiedRepository {
	return &cacheModifiedRepository{
		prefix:          cache.GetSettings(context).CachePrefix,
		core:            GetCore(context),
		cacheRepository: GetCacheRepository(context),
		errorService:    loggers.GetErrorService(context),
	}
}

func (c *cacheModifiedRepository) Persist(context ctx.Context, name string, expiration time.Duration, miss cache.Miss, hit cache.Hit) interface{} {
	modifiedName := fmt.Sprintf(name, c.prefix)
	modifiedTime := c.getModifiedTime(modifiedName, expiration, context)

	data := c.cacheRepository.Persist(name, expiration, func() (data interface{}, b []byte, err error) {
		data, b, err = miss()
		c.errorService.CheckErrorAndPanic(err)

		stat, err := c.core.Stat(modifiedName)
		c.errorService.CheckErrorAndPanic(err)
		modifiedTime = stat.ModTime()

		return
	}, hit)

	internal.CheckModifiedTime(modifiedTime, context)

	return data
}

func (c *cacheModifiedRepository) getModifiedTime(name string, expiration time.Duration, context ctx.Context) time.Time {
	var modifiedTime time.Time

	stat, err := c.core.Stat(name)
	if err != nil {
		return modifiedTime
	}
	if time.Now().After(stat.ModTime().Add(expiration)) {
		return modifiedTime
	}
	modifiedTime = stat.ModTime()

	internal.CheckIfModifiedSince(ctxHttp.Request(context), modifiedTime)

	return modifiedTime
}
