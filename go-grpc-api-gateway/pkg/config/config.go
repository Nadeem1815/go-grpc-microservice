package config

type Config struct {
	Port        string `mapstructure:"PORT"`
	AuthSvcUrl  string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcl string `mapstructure:"PORT"`
	OrderSvcUrl string `mapstructure:"PORT"`
}

func LoadConfig() (Config, error) {
	// viper.AddConfigPath("./pkg/config/envs")
	// viper.SetConfigName("dev")
	// viper.SetConfigType("env")

	// viper.AutomaticEnv()
	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return

	// }
	// err = viper.Unmarshal(&c)
	// return
	return Config{
		Port:        ":3000",
		AuthSvcUrl:  "auth-svc:50051",
		ProductSvcl: "localhost:50051",
		OrderSvcUrl: "localhost:50053",
	}, nil
}
