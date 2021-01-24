package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// 起動環境切り分け用
	Env string `default:"local"`

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

func newConfig() config {
	var c config
	if err := envconfig.Process("KAUBANDUS", &c); err != nil {
		return c
	}
	log.Printf("config:%#+v", c)
	return c
}

func (c *config) IsCloud() bool {
	return c.Env == "local"
}

func (c *config) Dsn() string {
	if c.IsCloud() {
		return fmt.Sprintf("host=/cloudsql/%s user=%s password=%s dbname=%s sslmode=disable",
			c.DBHost, c.DBUser, c.DBPass, c.DBName)
	}
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBName, c.DBUser, c.DBPass)
}
