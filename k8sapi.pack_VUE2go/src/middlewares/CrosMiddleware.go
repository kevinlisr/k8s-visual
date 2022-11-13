package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CrosMiddleware struct {

}

func NewCrosMiddleware() *CrosMiddleware {
	return &CrosMiddleware{}
}

func (*CrosMiddleware)OnRequest(c *gin.Context) error {
	method := c.Request.Method
	if method != "" {
		c.Header("Access-Control-Allow-Origin", "*") // ke jiang * ti huan wei zhi ding de yu ming
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
	}
		//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	return nil
}

func (*CrosMiddleware)OnResponse(result interface{}) (interface{},error) {
	return result, nil
}



