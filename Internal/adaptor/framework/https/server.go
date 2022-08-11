package https

import (
	"github.com/abigiya/internal/Internal/adaptor/application/account/register"
	"github.com/gin-gonic/gin"
)

type ServerAdaptor struct {
	Register register.ConnectandCheck
	router   gin.Engine
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func NewDbAdaptor(dbDriver, dbSource string) *ServerAdaptor {
	createAccountAdaptor := register.NewUser(dbDriver, dbSource)

	serveAdaptor := &ServerAdaptor{
		Register: createAccountAdaptor,
		router:   gin.Engine{},
	}

	router := gin.New()
	router.POST("/Register", serveAdaptor.SignUp)

	return serveAdaptor
}

func (serve *ServerAdaptor) RequestSendStart(address string) error {
	return serve.router.Run(address)
}
