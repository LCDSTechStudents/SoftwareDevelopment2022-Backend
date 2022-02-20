package config

const (
	AUTH      = 1
	FLASHCARD = 2
)

type DB struct {
	URL      string `json:"url"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	DBName   string `json:"db_name"`
}
