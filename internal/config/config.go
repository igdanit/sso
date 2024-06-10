package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	ENV_CONFIG_PATH = "CONFIG_PATH"
)

var (
	cfg *AppConfig
)

type AppConfig struct {
	Env   string      `yaml:"env" env-required:"true"`
	GRPC  GrpcConfig  `yaml:"grpc"`
	Token TokenConfig `yaml:"token"`
}

type GrpcConfig struct {
	Port    int           `yaml:"port" env-default:"7200"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type TokenConfig struct {
	TTL time.Duration `yaml:"ttl" env-required:"true"`
}

// Get singleton config
func GetConfig() *AppConfig {
	if cfg != nil {
		return cfg
	}
	cfg = loadConfig()
	return cfg
}

func loadConfig() *AppConfig {
	configPath := *fetchConfigPath()
	if configPath == "" {
		panic(fmt.Sprint("config path is not provided. Provide it through --config flag or $", ENV_CONFIG_PATH))
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("%s | config file don't exist", configPath))
	}
	tmpCfg := new(AppConfig)
	if err := cleanenv.ReadConfig(configPath, tmpCfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	cfg = tmpCfg
	return cfg
}

// Read config path from startup flags or environment variables
func fetchConfigPath() *string {
	// --config=path/to/config.yml
	configPath := flag.String("config", "", "provide path to config file")
	flag.Parse()

	if *configPath == "" {
		*configPath = os.Getenv(ENV_CONFIG_PATH)
	}

	return configPath
}
