package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Init() (*Cfg, error) {
	if err := load(); err != nil {
		return nil, fmt.Errorf("can't load config: %w", err)
	}

	cfg := &Cfg{}
	opts := env.Options{
		Prefix:                "APP_",
		UseFieldNameByDefault: true,
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, fmt.Errorf("can't parse config: %w", err)
	}

	return cfg, nil
}

func load() error {
	var err error
	cfgEnv := os.Getenv("ENV_FILE")
	if len(cfgEnv) > 0 {
		err = godotenv.Load(cfgEnv)
		if err != nil {
			return fmt.Errorf("can't parse config: %w", err)
		}
	} else {
		err = godotenv.Load(".env")
		if err != nil {
			return fmt.Errorf("can't parse config .env: %w", err)
		}
	}

	return nil
}
