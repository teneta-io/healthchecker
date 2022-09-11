package config

import "time"

// Log describes settings for application logger.
type Log struct {
	Level string `toml:"level"`
	Mode  string `toml:"mode"`
}

// Server represents settings for creating http server.
type Server struct {
	ShutdownTimeout Duration `toml:"shutdown_timeout"`
	ReadTimeout     Duration `toml:"read_timeout"`
	WriteTimeout    Duration `toml:"write_timeout"`
	IdleTimeout     Duration `toml:"idle_timeout"`
}

type Config struct {
	Log    Log
	Server Server
}

func NewConfig() *Config {
	return &Config{
		Log: Log{
			Level: "info",
			Mode:  "dev",
		},
		Server: Server{
			ShutdownTimeout: Duration{
				Duration: time.Second * 30,
			},
			ReadTimeout: Duration{
				Duration: time.Second * 30,
			},
			WriteTimeout: Duration{
				Duration: time.Second * 30,
			},
			IdleTimeout: Duration{
				Duration: time.Second * 30,
			},
		},
	}
}
