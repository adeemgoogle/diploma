package config

import "github.com/spf13/viper"

type Config struct {
	DBurl string `mapstructure:"DB_URL"`
	Port  string `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	config := new(Config)

	v := viper.New()
	v.AutomaticEnv()

	if err := v.BindEnv("DB_URL"); err != nil {
		return nil, err
	}
	if err := v.BindEnv("PORT"); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil

}
