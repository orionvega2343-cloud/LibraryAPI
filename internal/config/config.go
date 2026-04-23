package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Database config
type DBConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `env:"DB_PASSWORD"`
	Name string `yaml:"name"`
}

// Server config
type ServerConfig struct {
	Port int `yaml:"port"`
}

// Jwt config
type JWTConfig struct {
	Secret string `env:"JWT_SECRET"`
}

// Config
type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	JWT    JWTConfig    `yaml:"jwt"`
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil { // .env not found
		panic(err)
	}
	c := Config{}
	ok := cleanenv.ReadConfig("config.yml", &c) // read config into struct
	if ok != nil {
		panic(ok)
	}
	return &c
}
