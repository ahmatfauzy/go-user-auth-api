package middleware

import (
	"auth-api/internal/config"
	"auth-api/pkg/jwt"
	"auth-api/pkg/redisclient"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(cfg *config.Config, redis *redisclient.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ValidateToken(tokenStr, cfg.JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		// Cek blacklist di Redis
		_, err = redis.Get("blacklist:" + tokenStr)
		if err == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token revoked"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}