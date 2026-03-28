package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimitMiddleware() gin.HandlerFunc {
	rate, _ := limiter.NewRateFromFormatted("100-M")
	store := memory.NewStore()
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	return middleware
}
