package login

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/crypto"
	io2 "SoftwareDevelopment-Backend/server/internalsvc/authorize/io"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/tokenHandler"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/userpack"
	"SoftwareDevelopment-Backend/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func LoginHandler(content *content.Content, handler crypto.PasswordHandler, token tokenHandler.TokenHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login io2.Login
		//parse user email and password
		ctx.BindJSON(&login)
		if !verify(login.Email, login.Password) {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("invalid email or password")))
			return
		}

		//query if exist and valid
		//TODO: Verify
		user, ok := query(login.Email, login.Password, handler, content.Data["DB"].(*gorm.DB))

		switch ok {
		case userpack.WrongPassword:
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("wrong password")))
			return
		case userpack.NoSuchUser:
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("not registered")))
			return
		}

		//return JSON and generate token
		ctx.JSON(http.StatusOK, services.SuccessResponse(io2.PostUser{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
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

func query(email string, password string, handler crypto.PasswordHandler, db *gorm.DB) (*userpack.User, int) {
	var user *userpack.User
	db.Where("email = ?", email).Find(&user)
	if user.ID == 0 {
		return nil, userpack.NoSuchUser
	}
	if !handler.Check(password, user.Password) {
		return user, userpack.WrongPassword
	}
	return user, userpack.Success
}
