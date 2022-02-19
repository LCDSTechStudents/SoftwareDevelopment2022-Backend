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
