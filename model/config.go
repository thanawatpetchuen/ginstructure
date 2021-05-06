package model

type Config struct {
	Env         string
	ServiceName string `mapstructure:"service-name"`
	Http        struct {
		Port int
	}
}
