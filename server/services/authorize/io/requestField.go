package io

const (
	REGISTER = "register"
	FINDPW   = "find_password"
)

type Login struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type Registration struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SendMail struct {
	Email  string `json:"email" binding:"required"`
	Target string `json:"target" binding:"required"`
}
