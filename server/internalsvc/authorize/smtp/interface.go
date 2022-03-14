package smtp

type EmailHandler interface {
	SendCode(email string, code int) error
}
