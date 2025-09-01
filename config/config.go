package config

import "github.com/spf13/viper"

type App struct {
	AppPort string `json"app_port"`
	AppEnv string `json"app_env"`

	JwtSecretKey string `json:"jwt_secret_key"`
	JwtIssuer string `json:"jwt_issuer"`
}

type PsqlDB struct {
	Host	string `json:"host"`
	Port	string `json:"port"`
	User	string `json:"user"`
	Password	string `json:"password"`
	DBName	string `json:"db_name"`
	DBMaxOpen	string `json:"db_max_open"`
	DBMaxIdle	string `json:"db_max_idle"`
}

type Config struct {
	App App
	Psql PsqlDB
}

func NewConfig() *Config{
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv: viper.GetString("APP_PORT"),

			JwtSecretKey: viper.GetString("JWT_SECRET_KEY"),
			JwtIssuer: viper.GetString("JWT_ISSUER"),
		},
		Psql: PsqlDB{
			Host: viper.GetString("DATABASE_HOST"),
			Port: viper.GetString("DATABASE_HOST"),
			User: viper.GetString("DATABASE_HOST"),
			Password: viper.GetString("DATABASE_HOST"),
			DBName: viper.GetString("DATABASE_HOST"),
			DBMaxOpen: viper.GetString("DATABASE_HOST"),
			DBMaxIdle: viper.GetString("DATABASE_HOST"),
		},
	}
}