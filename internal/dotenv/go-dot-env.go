package dotenv

import "github.com/joho/godotenv"

type GoDotEnv struct {
	DotEnver
}

func NewGoDotEnv() *GoDotEnv {
	return &GoDotEnv{}
}

func (e *GoDotEnv) Load() (err error) {
	return godotenv.Load()
}
