package apiserver

type Config struct {
	BindAddr string
	LogLevel string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":80",
		LogLevel: "debug",
	}
}
