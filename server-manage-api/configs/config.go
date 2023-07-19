package configs

import "github.com/spf13/viper"

type Config struct {
	Env  string
	Port int

	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbSsl      string

	SwaggerHost string
}

func Load() Config {
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	conf := Config{}
	conf.Env = viper.GetString("ENV")
	conf.Port = viper.GetInt("PORT")

	conf.DbHost = viper.GetString("POSTGRES_HOST")
	conf.DbPort = viper.GetInt("POSTGRES_PORT")
	conf.DbUser = viper.GetString("POSTGRES_USER")
	conf.DbPassword = viper.GetString("POSTGRES_PASSWORD")
	conf.DbName = viper.GetString("POSTGRES_DB")
	conf.DbSsl = viper.GetString("POSTGRES_SSL")

	conf.SwaggerHost = viper.GetString("SWAGGER_HOST")

	return conf
}
