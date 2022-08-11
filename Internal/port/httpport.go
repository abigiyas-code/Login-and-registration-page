package port

import (
	"github.com/gin-gonic/gin"
)

type HttpPort interface {
	SignUp(router gin.Engine) 
	RequestSendStart()
}
