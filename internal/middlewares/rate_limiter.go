package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	lastSeen   time.Time
	tokens     int
	lastRefill time.Time
	mu         sync.Mutex
}

var (
	visitors      = make(map[string]*visitor)
	visitorsMutex sync.Mutex

	rateLimit      = 5
	refillDuration = 10 * time.Second
)

func getVisitor(ip string) *visitor {
	visitorsMutex.Lock()
	defer visitorsMutex.Unlock()

	v, exists := visitors[ip]
	if !exists {
		v = &visitor{
			tokens:     rateLimit,
			lastRefill: time.Now(),
		}
		visitors[ip] = v
	}
	v.lastSeen = time.Now()
	return v
}

func cleanupVisitors() {
	for {
		time.Sleep(1 * time.Minute)
		visitorsMutex.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 5*time.Minute {
				delete(visitors, ip)
			}
		}
		visitorsMutex.Unlock()
	}
}

func RateLimiterMiddleware() gin.HandlerFunc {
	go cleanupVisitors()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		v := getVisitor(ip)

		v.mu.Lock()
		defer v.mu.Unlock()

		now := time.Now()
		elapsed := now.Sub(v.lastRefill)

		if elapsed >= refillDuration {
			v.tokens = rateLimit
			v.lastRefill = now
		}

		if v.tokens > 0 {
			v.tokens--
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"error": "Too many requests. Try again later.",
		})
	}
}
