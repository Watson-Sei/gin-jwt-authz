package ginjwtauthz

import (
	"net/http"

	"github.com/Watson-Sei/gin-jwt-authz/utils"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

type Options struct {
	failWithError  bool
	checkAllScopes bool
}

func DefaultOptions() Options {
	return Options{
		failWithError:  true,
		checkAllScopes: true,
	}
}

func CheckPermissions(expectedScopes []string, options Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Request.Context().Value("user")

		if len(expectedScopes) == 0 {
			c.Next()
		}

		var allowed bool = false
		if options.checkAllScopes {
			allowed = utils.Every(expectedScopes, utils.InterfaceSliceConversion(user.(*jwt.Token).Claims.(jwt.MapClaims)["permissions"].([]interface{})))
		} else {
			allowed = utils.Some(expectedScopes, utils.InterfaceSliceConversion(user.(*jwt.Token).Claims.(jwt.MapClaims)["permissions"].([]interface{})))
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Forbidden",
				"message": "Insufficient scope",
			})
			c.Abort()
		}

		c.Next()
	}
}
