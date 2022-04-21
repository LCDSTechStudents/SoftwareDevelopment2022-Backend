package verifyCodeHandler

import (
	"SoftwareDevelopment-Backend/config"
	"fmt"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	INCORRECTCODE   = 0
	CORRECTCODE     = 1
	DIDNOTFINDEMAIL = 2
)

var randomNumberCandidate = [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

type DefaultCodeHandler struct {
	randomCandidate       [10]byte
	randomCandidateLength int
	log                   *zap.Logger
	config                *config.Config
	cache                 *cache.Cache
}

func (d *DefaultCodeHandler) NewCode(email string) int {
	code := d.generateCode(d.config.Services.Auth.VerifyCodeLength)
	d.cache.Set(email, code, cache.DefaultExpiration)
	return code
}

func (d *DefaultCodeHandler) CheckCode(email string, code int) int {
	c, found := d.cache.Get(email)
	if !found {
		return DIDNOTFINDEMAIL
	}
	if c == code {
		return CORRECTCODE
	}
	return INCORRECTCODE
}

func (d *DefaultCodeHandler) generateCode(length int) int {
	var sb strings.Builder

	for i := 0; i < length; i++ {
		_, _ = fmt.Fprint(&sb, d.randomCandidate[rand.Intn(d.randomCandidateLength-1)])
	}

	r, _ := strconv.Atoi(sb.String())
	return r
}

func InitDefaultCodeHandler(logger *zap.Logger, config *config.Config) *DefaultCodeHandler {
	return &DefaultCodeHandler{
		//tw : timingwheel.NewTimingWheel(time.Millisecond, 20),
		randomCandidate:       randomNumberCandidate,
		randomCandidateLength: len(randomNumberCandidate),
		cache:                 cache.New(5*time.Minute, 10*time.Minute),
		config:                config,
		log:                   logger,
	}
}
