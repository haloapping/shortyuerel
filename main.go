package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	// load env variable
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("dev.config")
	err := viper.ReadInConfig()
	if err != nil {
		zlog.Error().Msg(err.Error())
	}
	viper.AutomaticEnv()

	// setup database confing
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSLMODE"),
	)
	dbconfig, err := NewDBConfig(connStr)
	if err != nil {
		panic(err)
	}

	// initiate database pooling
	pool, err := NewDBPool(dbconfig)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	urlRepository := NewReposity(pool)
	urlService := NewService(urlRepository)
	urlHandler := NewHandler(urlService)

	r := echo.New()
	r.POST("/", urlHandler.CreateShortUrl)
	r.GET("/{longUrl}", urlHandler.GetShortUrlByLongUrl)

	r.Start(fmt.Sprintf(":%s", viper.GetString("APP_PORT")))
}
