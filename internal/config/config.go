package config

import (
	"io/ioutil"

	"github.com/gookit/validate"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	defaultAppPort     = 4001
	defaultAppHost     = "localhost"
	defaultAppEnv      = "development"
	defaultAppName     = "Go Rest Boilerplate"
	defaultSessionName = "rfid"
)

type Config struct {
	AppPort int    `yaml:"app_port" env:"APP_PORT"`
	AppHost string `yaml:"app_host" env:"APP_HOST"`
	AppEnv  string `yaml:"app_env" env:"APP_ENV"`
	AppName string `yaml:"app_name" env:"APP_NAME"`

	GithubClientId     string `yaml:"github_client_id" env:"GITHUB_CLIENT_ID"`
	GithubClientSecret string `yaml:"github_client_secret" env:"GITHUB_CLIENT_SECRET,secret"`
	GithubCallbackUrl  string `yaml:"github_callback_url" env:"GITHUB_CALLBACK_URL"`

	SessionSecret string `validate:"required|string" yaml:"session_secret" env:"SESSION_SECRET,secret"`
	SessionName   string `validate:"required|string" yaml:"session_name" env:"SESSION_NAME"`

	PostgresUser     string `validate:"required|string" yaml:"postgres_user" env:"POSTGRES_USER"`
	PostgresPassword string `yaml:"postgres_password" env:"POSTGRES_PASSWORD,secret"`
	PostgresHost     string `validate:"required|string" yaml:"postgres_host" env:"POSTGRES_HOST"`
	PostgresPort     int    `yaml:"postgres_port" env:"POSTGRES_PORT"`
	PostgresDb       string `validate:"required|string" yaml:"postgres_db" env:"POSTGRES_DB"`
	PostgresSslMode  string `yaml:"postgres_ssl_mode" env:"POSTGRES_SSL_MODE"`
}

func (c *Config) Validate() error {
	v := validate.Struct(&c)
	if v.Validate() {
		return nil
	} else {
		return v.Errors
	}
}

func Load(file string, logger *zap.SugaredLogger) (*Config, error) {
	c := Config{
		AppPort:     defaultAppPort,
		AppHost:     defaultAppHost,
		AppEnv:      defaultAppEnv,
		AppName:     defaultAppName,
		SessionName: defaultSessionName,
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// TODO: Load env vars

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil
}
