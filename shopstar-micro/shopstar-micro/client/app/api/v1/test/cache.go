package tests

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"sixstar/shopstar-micro/client/core/cache"
)

type TestCache struct {

}
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
func (*TestCache) GetCache(ctx *gin.Context)  {
	//panic("ooo err ")

	fmt.Println(cache.CacheManager.Get("name"))
	ctx.String(200, "ok set ")
}

func (*TestCache) SetCache(ctx *gin.Context)  {
	fmt.Println(cache.CacheManager.Set("name", []byte("shineyork"), &cache.Options{
		Expiration: time.Duration(30) * time.Second,
	}))

	fmt.Println(cache.CacheManager.Get("name"))
	ctx.String(200, "okok" )
}

