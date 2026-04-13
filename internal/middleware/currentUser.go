package middleware

import (
	"errors"
	"real-estate-app/internal/service/auth"

	"github.com/gin-gonic/gin"
)

var ErrUnauthorized = errors.New("unauthorized")

func CurrentUser(c *gin.Context) (*auth.UserClaims, error) {
	v, ok := c.Get(CurrentUserKey)
	if !ok {
		return nil, ErrUnauthorized
	}

	claims, ok := v.(*auth.UserClaims)
	if !ok {
		return nil, ErrUnauthorized
	}

	return claims, nil
}
