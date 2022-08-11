package https

import (
	"context"
	"net/http"

	"github.com/abigiya/internal/Internal/adaptor/framework/dbservice/sqlc"
	"github.com/gin-gonic/gin"
)

func (serve *ServerAdaptor) SignUp(ctx *gin.Context) {
	var request *SendAccountRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusPreconditionFailed, errorResponse(err))
	}

	agree := sqlc.CreateAccountParams{
		Firstname:    request.Firstname,
		Lastname:     request.Lastname,
		Emailaddress: request.Emailaddress,
		Username:     request.Username,
		Password:     request.Password,
	}

	response, err := serve.Register.ConfirmAccount(context.Background(), agree.Firstname, agree.Lastname, agree.Emailaddress, agree.Username, int64(agree.Password))
	if err != nil {
		ctx.JSON(http.StatusPreconditionRequired, gin.H{
			"error format": response.Status,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status pass, Registered!": response.Status,
		"account":                  response.Store,
	})
}
