package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

var (
	ipBucketMap = make(map[string]*ratelimit.Bucket)
	mu          sync.Mutex
)

func RateLimiterPerIP(rate, burst int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		mu.Lock()
		bucket, exists := ipBucketMap[ip]
		if !exists {
			bucket = ratelimit.NewBucketWithQuantum(time.Second, int64(rate), int64(burst))
			ipBucketMap[ip] = bucket
		}
		mu.Unlock()

		if bucket.TakeAvailable(1) < 1 {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests, please try again later.",
			})
			return
		}
		ctx.Next()
	}
}
