package config

const (
	AUTH      = 1
	FLASHCARD = 2
)

type DB struct {
	URL      string `json:"URL"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	Protocol string `json:"Protocol"`
	DBName   string `json:"DBName"`
}
