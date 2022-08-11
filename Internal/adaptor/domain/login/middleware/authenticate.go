package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)
// an interface send valid data to http response
type ValidData interface {
	Validate() // use databinding concept
}

func Authenticate() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{})
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}
