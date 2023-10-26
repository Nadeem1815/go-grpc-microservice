package config

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
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
	// err = viper.Unmarshal(&Config)
	// if err != nil {
	// 	return
	// }
	// return
	return Config{
		Port:         ":50051",
		DBUrl:        "postgres://nadeem:181511@auth-db:5432/auth_svc",
		JWTSecretKey: "r43t18sc",
	}, nil
}
