package middleware

import (
	"github.com/happilymarrieddad/nats-api-playground/api/internal/repos"

	"github.com/gin-gonic/gin"
)

func SetGlobalRepoToContext(gr repos.GlobalRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(GlobalRepoContextKey.String(), gr)
		c.Next()
	}
}

func GetGlobalRepoFromContext(c *gin.Context) repos.GlobalRepo {
	if iface, exists := c.Get(GlobalRepoContextKey.String()); exists {
		if gr, ok := iface.(repos.GlobalRepo); ok {
			return gr
		}
	}
	// This will never happen...
	return nil
}
