package verifyCode

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/services"
	"SoftwareDevelopment-Backend/server/services/authorize"
	"SoftwareDevelopment-Backend/server/services/authorize/io"
	"SoftwareDevelopment-Backend/server/services/authorize/smtp"
	"SoftwareDevelopment-Backend/server/services/authorize/verifyCodeHandler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func VerifyCodeHandler(content *content.Content, code verifyCodeHandler.VerifyCodeHandler, smtp smtp.EmailHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var verify io.SendMail
		//parse user email and password
		ctx.BindJSON(&verify)

		//verify validation of email
		if !verifyRequest(verify) {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("invalid email")))
			return
		}

		if verify.Target == io.FINDPW {
			//if user not found
			if !userExist(verify, content) {
				ctx.JSON(http.StatusNotFound, services.ErrorResponse(fmt.Errorf("user not registered")))
				return
			}

			verifyCode := code.NewCode(verify.Email)
			if err := smtp.SendCode(verify.Email, verifyCode); err != nil {
				ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("an error occur while sending email, please try again later")))
				return
			}

			ctx.JSON(http.StatusOK, services.SuccessResponse(io.PostVerify{
				Email:      verify.Email,
				VerifyCode: verifyCode,
			}))

			return
		}

		if verify.Target == io.REGISTER {
			//if user not found
			if userExist(verify, content) {
				ctx.JSON(http.StatusNotFound, services.ErrorResponse(fmt.Errorf("user already exists")))
				return
			}

			verifyCode := code.NewCode(verify.Email)
			if err := smtp.SendCode(verify.Email, verifyCode); err != nil {
				ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("an error occur while sending email, please try again later")))
				return
			}

			ctx.JSON(http.StatusOK, services.SuccessResponse(io.PostVerify{
				Email:      verify.Email,
				VerifyCode: verifyCode,
			}))

			return
		}
	}

}

func verifyRequest(req io.SendMail) bool {
	if req.Email != "" && strings.Contains(req.Email, "@") {
		return true
	}
	return false
}

func userExist(req io.SendMail, ctn *content.Content) bool {
	var user authorize.User
	ctn.Db.Where("email = ?", req.Email).Find(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
