package initializing

import (
	"github.com/caarlos0/env/v6"
)

// EnvVars представляет структуру для хранения параметров окружения.
type EnvVars struct {
	ConnectionStringDB string `env:"DATABASE_DSN"`
	MainPath           string `env:"MAIN_PATH"`
}

// ReadEnv считывает параметры из переменных окружения и возвращает объект EnvVars.
func ReadEnv() (*EnvVars, error) {
	envs := EnvVars{}

	// Считывание значений переменных окружения в объект EnvVars
	if err := env.Parse(&envs); err != nil {
		return nil, err
	}

	return &envs, nil
}
