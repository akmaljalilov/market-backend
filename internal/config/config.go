package config

import (
	"io"
	"market/internal/utils"
	"os"

	"github.com/caarlos0/env/v11"
)

func InitConfig() (*Config, error) {
	cfg := &Config{}
	cfgModePath := "./config.yml"
	if _, err := os.Stat(cfgModePath); err != nil {
		return nil, err
	}
	files := []string{cfgModePath}
	for _, file := range files {
		if len(file) > 0 {
			err := utils.ReadYml(file, cfg)
			if err != nil && io.EOF != err {
				return nil, err
			}
		}
	}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
