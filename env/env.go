package env

import "github.com/spf13/viper"

var JwtSecretKey string

func InitEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	JwtSecretKey = viper.GetString("JWT_SECRET_KEY")
}
