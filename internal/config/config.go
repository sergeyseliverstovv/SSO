package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTl    time.Duration `yaml:"token_ttl" env-requiret:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPatch()

	if path == "" {
		panic("config patch is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {

		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("filed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPatch() string {

	var res string

	flag.StringVar(&res, "config", "", "patch to config file")
	flag.Parse()

	if res == "" {

		res = os.Getenv("CONFIG_PATCH")

	}

	return res

}
