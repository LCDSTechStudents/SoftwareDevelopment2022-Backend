package register

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/services"
	"SoftwareDevelopment-Backend/server/services/authorize/crypto"
	"SoftwareDevelopment-Backend/server/services/authorize/idGenerator"
	"SoftwareDevelopment-Backend/server/services/authorize/io"
	"SoftwareDevelopment-Backend/server/services/authorize/userpack"
	"SoftwareDevelopment-Backend/server/services/authorize/verifyCodeHandler"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func RegHandler(content *content.Content, handler crypto.PasswordHandler, code verifyCodeHandler.VerifyCodeHandler, id idGenerator.IDGenerator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reg io.Registration
		//parse user email and password

		ctx.BindJSON(&reg)
		if !verify(reg.Email, reg.Password, reg.Nickname) {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("invalid email or password")))
			return
		}

		//query if use is exists
		if isExist(content.Db, &reg) {
			ctx.JSON(http.StatusNotAcceptable, services.ErrorResponse(fmt.Errorf("user already exists")))
			return
		}

		//if !verifyCode(&reg, code){
		//	ctx.JSON(http.StatusNotAcceptable, services.ErrorResponse(fmt.Errorf("incorrect verify code")))
		//	return
		//}
		//verify code here
		switch code.CheckCode(reg.Email, reg.VerifyCode) {
		case verifyCodeHandler.INCORRECTCODE:
			ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("incorrect verification code")))
			return
		case verifyCodeHandler.DIDNOTFINDEMAIL:
			ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("you did not send an verification code")))
			return
		}

		user, err := addUser(content, reg, handler, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("error while registering, please try again later")))
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

//func verifyCode(reg *io.Registration, handler verifyCodeHandler.VerifyCodeHandler) bool{
//	if reg.VerifyCode == handler.CheckCode(reg.Email, reg.VerifyCode){
//		return true
//	}
//	return false
//}

func addUser(content *content.Content, reg io.Registration, pw crypto.PasswordHandler, generator idGenerator.IDGenerator) (*userpack.User, error) {

	user := userpack.User{
		ID:       <-generator.GetIDChan(),
		Email:    reg.Email,
		Nickname: reg.Nickname,
		Password: pw.HashPassword(reg.Password),
	}
	result := content.Db.Create(&user)
	if result.Error != nil {
		content.Log.Error("making a new user: ", zap.Error(result.Error))
		return nil, result.Error
	}
	content.Log.Info("a new user created ")
	return &user, nil
}
