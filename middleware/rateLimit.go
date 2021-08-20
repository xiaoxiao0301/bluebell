package middleware

import (
	"go_web/web_app/controller"
	"go_web/web_app/dict"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func ApiRateLimit(fillInterval time.Duration, cap int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(context *gin.Context) {
		//if bucket.TakeAvailable(1) == 0 {
		//	controller.ReturnErr(context, dict.CodeRateLimit)
		//	context.Abort()
		//	return
		//}
		//context.Next()
		if bucket.TakeAvailable(1) == 1 {
			context.Next()
			return
		}
		controller.ReturnErr(context, dict.CodeRateLimit)
		context.Abort()
	}
}
