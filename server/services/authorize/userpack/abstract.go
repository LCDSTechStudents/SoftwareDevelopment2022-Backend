package userpack

const (
	Success       = 1
	WrongPassword = 2
	NoSuchUser    = 3
)

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
