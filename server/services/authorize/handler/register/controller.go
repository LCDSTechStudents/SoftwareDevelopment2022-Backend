package register

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize"
	io "SoftwareDevelopment-Backend/server/internalsvc/authorize/io"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/userpack"
	verifyCodeHandler "SoftwareDevelopment-Backend/server/internalsvc/authorize/verifyCodeHandler"
	"SoftwareDevelopment-Backend/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func RegHandler(content *content.Content) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reg io.Registration
		//parse user email and password
		auth := content.Data[authorize.AUTHORIZER].(*authorize.DefaultAuthorizer)

		err := ctx.ShouldBindJSON(&reg)
		if err != nil {
			auth.Error("registration bind json error", zap.Error(err))
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("error input json")))
			return
		}

		if !verify(reg.Email, reg.Password, reg.Nickname) {
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("invalid email or password")))
			return
		}

		//query if use is exists
		if isExist(auth.GetDB(), &reg) {
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("user already exists")))
			return
		}

		if !verifyCode(&reg, auth) {
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("incorrect verify code")))
			return
		}
		//verify code here
		switch auth.CheckCode(reg.Email, reg.VerifyCode) {
		case verifyCodeHandler.INCORRECTCODE:
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("incorrect verification code")))
			return
		case verifyCodeHandler.DIDNOTFINDEMAIL:
			ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("you did not send an verification code")))
			return
		}

		user, err := addUser(auth, reg)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("error while registering, please try again later")))
			return
		}

		ctx.JSON(http.StatusOK, services.SuccessResponse(io.NewUser{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
		}))
	}
}

func verify(email string, pw string, nickname string) bool {
	if email != "" && pw != "" && strings.Contains(email, "@") && len(nickname) < 20 {
		return true
	}
	return false
}

func isExist(db *gorm.DB, reg *io.Registration) bool {
	var user *userpack.User
	db.Where("email = ?", reg.Email).Find(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func verifyCode(reg *io.Registration, handler *authorize.DefaultAuthorizer) bool {
	if verifyCodeHandler.CORRECTCODE == handler.CheckCode(reg.Email, reg.VerifyCode) {
		return true
	}
	return false
}

func addUser(auth *authorize.DefaultAuthorizer, reg io.Registration) (*userpack.User, error) {

	user := userpack.User{
		ID:       <-auth.GetIDChan(),
		Email:    reg.Email,
		Nickname: reg.Nickname,
		Password: auth.HashPassword(reg.Password),
	}
	result := auth.GetDB().Create(&user)
	if result.Error != nil {
		auth.Error("making a new user: ", zap.Error(result.Error))
		return nil, result.Error
	}
	auth.Info("a new user created ")
	return &user, nil
	return nil, nil
}
