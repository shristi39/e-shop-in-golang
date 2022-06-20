package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(jwt *JwtWrapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, error(err))
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, error(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s ", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, error(err))
			return
		}
		accessToken := fields[1]
		payload, err := jwt.ValidateToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, error(err))
			return
		}
		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
