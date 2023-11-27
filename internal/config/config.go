package config

import (
	env "github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"
)

var cfg Config

type Config struct {
	AppConfig
	DbConfig
	RedisConfig
	AuthKeys
	EmailConfig
}

type AppConfig struct {
	Host         string `env:"APP_HOST"`
	Port         string `env:"APP_PORT,required=true"`
	ReadTimeout  uint   `env:"APP_READTIMEOUT,required=true"`
	WriteTimeout uint   `env:"APP_WRITETIMEOUT,required=true"`
	IdleTimeout  uint   `env:"APP_IDLETIMEOUT,required=true"`
}
type RedisConfig struct {
	RedisAddr     string `env:"REDIS_ADDR,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDb       int    `env:"REDIS_DB,required=true"`
}
type DbConfig struct {
	DbHost     string `env:"POSTGRES_HOST,required=true"`
	DbUser     string `env:"POSTGRES_USER,required=true"`
	DbPassword string `env:"POSTGRES_PASSWORD,required=true"`
	DbName     string `env:"POSTGRES_DBNAME,required=true"`
	DbPort     string `env:"POSTGRES_PORT,required=true"`
	DbSSLMode  string `env:"POSTGRES_SSLMODE,required=true"`
	DbTimeZone string `env:"POSTGRES_TIMEZONE,required=true"`
}
type AuthKeys struct {
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
}
type EmailConfig struct {
	Port       int    `env:"EMAIL_PORT,required=true"`
	SenderMail string `env:"EMAIL_SENDERMAIL,required=true"`
	Password   string `env:"EMAIL_PASSWORD,required=true"`
}

func Init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("problem in config")
		return
	}
}

func GetConfig() Config {
	return cfg
}
