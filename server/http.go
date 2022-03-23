package server

import (
	"SoftwareDevelopment-Backend/config"
	"SoftwareDevelopment-Backend/server/content"
	"SoftwareDevelopment-Backend/server/internalsvc"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/crypto"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/idGenerator"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/smtp"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/tokenHandler"
	"SoftwareDevelopment-Backend/server/internalsvc/authorize/verifyCodeHandler"
	"SoftwareDevelopment-Backend/server/services/authorize/handler/login"
	"SoftwareDevelopment-Backend/server/services/authorize/handler/register"
	"SoftwareDevelopment-Backend/server/services/authorize/handler/verifyCode"
	"github.com/RussellLuo/timingwheel"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HTTPServer struct {
	config    *config.Config
	log       *zap.Logger
	engine    *gin.Engine
	ctn       map[int]*content.Content
	tw        *timingwheel.TimingWheel
	internals map[string]internalsvc.Internal
}

func (s *HTTPServer) Run() error {
	if err := s.engine.Run(s.config.Server.Port); err != nil {
		return err
	}

	return nil
}

func (s *HTTPServer) Stop() {
	s.log.Sync()
}

func InitHTTPServer(config *config.Config, logger *zap.Logger) Server {
	s := &HTTPServer{
		config: config,
		log:    logger,
		engine: gin.New(),
		ctn:    make(map[int]*content.Content),
		tw:     timingwheel.NewTimingWheel(time.Millisecond, 20),
	}
	//use zap to substitute original logger
	s.engine.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	if config.Server.AllowCors {
		logger.Info("Server allow cors enabled")
		s.engine.Use(Cors())
	} else {
		logger.Info("Server allow cors disabled")
	}

	//init content services

	s.InitInternalSvc()

	s.initContent()
	//set mode
	gin.SetMode(gin.DebugMode)

	//init internal dependencies

	//init handlers
	s.regHandlers()

	//allow cors

	return s
}

func (s *HTTPServer) initContent() {
	s.ctn[config.AUTH] = content.InitContent(s.config, s.log, config.AUTH, nil)
}

//router initialize
func (s *HTTPServer) regHandlers() {

	password := crypto.InitPasswordHandler(s.config)
	token := tokenHandler.InitTokenHandler(s.config, s.log)
	code := verifyCodeHandler.InitDefaultCodeHandler(s.log, s.config)
	smtp := smtp.InitDefaultSMTP(s.log, s.config)
	idGen := idGenerator.InitDefaultIDGenerator()

	s.engine.POST("/v1/auth/login", login.LoginHandler(s.ctn[config.AUTH], password, token))
	s.engine.POST("/v1/auth/reg", register.RegHandler(s.ctn[config.AUTH], password, code, idGen))
	s.engine.POST("/v1/auth/send_verify", verifyCode.VerifyCodeHandler(s.ctn[config.AUTH], code, smtp))
}

//Cors management

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token,AccessToken,Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.Header("Access-Control-Max-Age", "600")
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func (s *HTTPServer) InitInternalSvc() {
	//s.internals[]
}
