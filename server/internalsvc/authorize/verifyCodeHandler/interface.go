package verifyCodeHandler

type VerifyCodeHandler interface {
	NewCode(email string) int
	CheckCode(email string, code int) int
}
