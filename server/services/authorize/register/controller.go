package register

import (
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/services"
	"SoftwareDevelopment-Backend/server/services/authorize"
	"SoftwareDevelopment-Backend/server/services/authorize/crypto"
	"SoftwareDevelopment-Backend/server/services/authorize/io"
	"SoftwareDevelopment-Backend/server/services/authorize/tokenHandler"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func RegHandler(content *content.Content, handler crypto.PasswordHandler, token tokenHandler.TokenHandler) gin.HandlerFunc {
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
		}

		//send verify code

		//return JSON and generate token

	}
}

func verify(email string, pw string, nickname string) bool {
	if email != "" && pw != "" && strings.Contains(email, "@") && len(nickname) < 20 {
		return true
	}
	return false
}

func isExist(db *gorm.DB, reg *io.Registration) bool {
	var user *authorize.User
	db.Where("email = ?", reg.Email).Find(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func randToken(num int) string {
	b := make([]byte, num)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
