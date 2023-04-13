package setup

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// 起動環境切り分け用
	Env Env `default:"local"`

	// DB接続設定用
	DBHost string `split_words:"true" default:"localhost"`
	DBPort string `split_words:"true" default:"11111"`
	DBName string `split_words:"true" default:"kaubandusdb"`
	DBUser string `split_words:"true" default:"postgres"`
	DBPass string `split_words:"true" default:"yuckyjuice"`

	// Webサーバ設定用
	WebPort string `split_words:"true" default:"8080"`

	// トレース設定用
	Trace bool `split_words:"true" default:"false"`
}

func ReadConfig() Config {
	var c Config
	if err := envconfig.Process("KAUBANDUS", &c); err != nil {
		return c
	}
	log.Printf("config:%#+v", c)
	return c
}

func (c *Config) Dsn() string {
	if c.Env.IsGCP() {
		return fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable",
			c.DBHost, c.DBUser, c.DBPass, c.DBName)
	}
	if c.Env.IsLocal() {
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			c.DBHost, c.DBPort, c.DBName, c.DBUser, c.DBPass)
	}
	return ""
}
