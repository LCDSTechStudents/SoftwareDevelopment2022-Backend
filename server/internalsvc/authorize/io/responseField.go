package io

type PostUser struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Token    string `json:"token"`
}

type PostVerify struct {
	Email string `json:"email"`
	//VerifyCode int    `json:"verify_code"`
}

type NewUser struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}
