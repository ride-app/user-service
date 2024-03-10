package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Production bool   `env:"PRODUCTION" env-description:"dev or prod"               env-default:"true"`
	LogDebug   bool   `env:"LOG_DEBUG"  env-description:"should log at debug level" env-default:"false"`
	Port       int32  `env:"PORT"       env-description:"server port"               env-default:"50051"`
	Project_Id string `env:"PROJECT_ID" env-description:"firebase project id"       env-default:"NO_PROJECT"`
}

func New() (*Config, error) {
	config := Config{
		Production: true,
		LogDebug:   false,
		Port:       50051,
		Project_Id: "NO_PROJECT",
	}

	if err := cleanenv.ReadEnv(&config); err != nil {
		return &config, err
	}

	return &config, nil
}
