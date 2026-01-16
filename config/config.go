package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer    string `json:"jwt_issuer"`

	UrlForgetPassword string `json:"url_forget_password"`
}

type MySql struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	User           string `json:"user"`
	Password       string `json:"password"`
	DBName         string `json:"db_name"`
	DBMaxOpenConns int    `json:"db_max_open_conns"`
	DBMaxIdleConns int    `json:"db_max_idle_conns"`
}

type Config struct {
	App   App   `json:"app"`
	MySql MySql `json:"mysql"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort:           viper.GetString("APP_PORT"),
			AppEnv:            viper.GetString("APP_ENV"),
			JwtSecretKey:      viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer:         viper.GetString("JWT_ISSUER"),
			UrlForgetPassword: viper.GetString("URL_FORGET_PASSWORD"),
		},
		MySql: MySql{
			Host:           viper.GetString("MYSQL_HOST"),
			Port:           viper.GetString("MYSQL_PORT"),
			User:           viper.GetString("MYSQL_USER"),
			Password:       viper.GetString("MYSQL_PASSWORD"),
			DBName:         viper.GetString("MYSQL_DB_NAME"),
			DBMaxOpenConns: viper.GetInt("MYSQL_DB_MAX_OPEN_CONNS"),
			DBMaxIdleConns: viper.GetInt("MYSQL_DB_MAX_IDLE_CONNS"),
		},
	}
}
