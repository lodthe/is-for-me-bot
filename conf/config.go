package conf

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Telegram Telegram
	DB       DB
}

type Telegram struct {
	BotToken string `env:"TELEGRAM_BOT_TOKEN,required"`
}

type DB struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`

	MaxConnections int `env:"DB_MAX_CONNECTIONS" envDefault:"10"`

	GORMDebug bool `env:"DB_GORM_DEBUG" envDefault:"false"`
}

func Read() Config {
	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		log.WithError(err).Fatal("failed to read the config")
	}

	return conf
}
