package verifyCode

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize"
	io2 "SoftwareDevelopment-Backend/server/internalsvc/authorize/io"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/userpack"
	"SoftwareDevelopment-Backend/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func VerifyCodeHandler(content *content.Content) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var verify io2.SendMail
		//parse user email and password
		auth := content.Data[authorize.AUTHORIZER].(*authorize.DefaultAuthorizer)
		ctx.BindJSON(&verify)

		//verify validation of email
		if !verifyRequest(verify) {
			ctx.JSON(http.StatusBadRequest, services.ErrorResponse(fmt.Errorf("invalid email")))
			return
		}

		if verify.Target == io2.FINDPW {
			//if user not found
			if !userExist(verify, content) {
				ctx.JSON(http.StatusOK, services.ErrorResponse(fmt.Errorf("user not registered")))
				return
			}

			verifyCode := auth.NewCode(verify.Email)
			if err := auth.SendCode(verify.Email, verifyCode); err != nil {
				ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("an error occur while sending email, please try again later")))
				return
			}

			ctx.JSON(http.StatusOK, services.SuccessResponse(io2.PostVerify{
				Email: verify.Email,
			}))

			return
		}

		if verify.Target == io2.REGISTER {
			//if user not found
			if userExist(verify, content) {
				ctx.JSON(http.StatusNotFound, services.ErrorResponse(fmt.Errorf("user already exists")))
				return
			}

			verifyCode := auth.NewCode(verify.Email)
			if err := auth.SendCode(verify.Email, verifyCode); err != nil {
				ctx.JSON(http.StatusInternalServerError, services.ErrorResponse(fmt.Errorf("an error occur while sending email, please try again later")))
				return
			}

			ctx.JSON(http.StatusOK, services.SuccessResponse(io2.PostVerify{
				Email: verify.Email,
				//VerifyCode: verifyCode,
			}))

			return
		}
	}

}

func verifyRequest(req io2.SendMail) bool {
	if req.Email != "" && strings.Contains(req.Email, "@") {
		return true
	}
	return false
}

func userExist(req io2.SendMail, ctn *content.Content) bool {
	var user userpack.User
	//TODO: adjust DB connection
	//ctn.Db.Where("email = ?", req.Email).Find(&user)
	if user.ID == 0 {
		return false
	}
	return true
}
