package authorize

const (
	Success       = 1
	WrongPassword = 2
	NoSuchUser    = 3
)

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

type PostUser struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

type Login struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
