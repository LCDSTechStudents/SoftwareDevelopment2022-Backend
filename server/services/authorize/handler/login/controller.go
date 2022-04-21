package login

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize"
	io2 "SoftwareDevelopment-Backend/server/internalsvc/authorize/io"
	"SoftwareDevelopment-Backend/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHandler(content *content.Content) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login io2.Login
		auth := content.Data[authorize.AUTHORIZER].(*authorize.DefaultAuthorizer)
		//parse user email and password
		ctx.BindJSON(&login)
		if !verify(login.Email, login.Password) {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("invalid email or password")))
			return
		}

		//query if exist and valid
		user, ok := auth.VerifyLoginInfo(login.Email, login.Password)
		switch ok {
		case authorize.WrongPassword:
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("wrong password")))
			return
		case authorize.UserNotFound:
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("not registered")))
			return
		}

		//return JSON and generate token
		ctx.JSON(http.StatusOK, services.SuccessResponse(io2.PostUser{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
			Token:    auth.GenerateToken(user.ID),
		}))
	}
}

func verify(email string, pw string) bool {
	if email != "" && pw != "" {
		return true
	}
	return false
}
