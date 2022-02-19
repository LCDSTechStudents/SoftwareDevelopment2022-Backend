package login

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/services"
	"SoftwareDevelopment-Backend/server/services/authorize"
	"SoftwareDevelopment-Backend/server/services/authorize/crypto"
	"SoftwareDevelopment-Backend/server/services/authorize/io"
	"SoftwareDevelopment-Backend/server/services/authorize/tokenHandler"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func LoginHandler(content *content.Content, handler crypto.PasswordHandler, token tokenHandler.TokenHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login io.Login
		//parse user email and password
		ctx.BindJSON(&login)
		if !verify(login.Email, login.Password) {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("invalid email or password")))
			return
		}

		//query if exist and valid
		user, ok := query(login.Email, login.Password, handler, content.Db)

		switch ok {
		case authorize.WrongPassword:
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("wrong password")))
			return
		case authorize.NoSuchUser:
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("not registered")))
			return
		}

		//return JSON and generate token
		ctx.JSON(http.StatusOK, services.SuccessResponse(io.PostUser{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.NickName,
			Token:    token.GenerateToken(user.ID),
		}))
	}
}

func verify(email string, pw string) bool {
	if email != "" && pw != "" {
		return true
	}
	return false
}

func query(email string, password string, handler crypto.PasswordHandler, db *gorm.DB) (*authorize.User, int) {
	var user *authorize.User
	db.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return nil, authorize.NoSuchUser
	}
	if !handler.Check(password, user.Password) {
		return user, authorize.WrongPassword
	}
	return user, authorize.Success
}
