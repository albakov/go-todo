package server

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	AllowOrigin string `toml:"allow_origin"`
	Path        string `toml:"path_to_json"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		AllowOrigin: "http://js-todo.ru.local",
	}
}
