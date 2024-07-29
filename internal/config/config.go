package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

//Структура соответствующая файлу конфигурации local.yaml

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

// Функция MustLoad возвращает объект конфига
func MustLoad() *Config {
	path := fetchConfigPatch()

	if path == "" {
		panic("config patch is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) { // IsNotExist возвращает true т.к файл не найден

		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("filed to read config: " + err.Error())
	}

	return &cfg
}

// Функция читает и возвращает путь конфигурационнога файла либо из флага, либо из переменной окружения
func fetchConfigPatch() string {

	var res string

	flag.StringVar(&res, "config", "", "patch to config file") // не понятно какой файл передать в config
	flag.Parse()

	if res == "" {

		res = os.Getenv("CONFIG_PATH") // или задать в пременную окружения

	}

	return res

}
