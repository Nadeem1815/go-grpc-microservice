package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	AuthSvcUrl  string `mapstructure:"PORT"`
	ProductSvcl string `mapstructure:"PORT"`
	OrderSvcUrl string `mapstructure:"PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return

	}
	err = viper.Unmarshal(&c)
	return
}
