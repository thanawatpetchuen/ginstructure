package config

import "github.com/spf13/viper"

func LoadConfig(v interface{}, fileName string) error {
	vp := viper.New()
	vp.AddConfigPath("config")
	vp.SetConfigType("yaml")
	vp.SetConfigName(fileName)

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	return vp.Unmarshal(v)
}
