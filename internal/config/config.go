package config

import (
	"io/ioutil"

	"github.com/gookit/validate"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	defaultAppPort     = 4000
	defaultAppHost     = "localhost"
	defaultAppEnv      = "development"
	defaultAppName     = "Go Rest Boilerplate"
	defaultSessionName = "session"
	defaultSslMode     = "disable"
	defaultDb          = "goboilerplate"
	defaultUser        = "postgres"
)

type Config struct {
	AppPort int    `yaml:"app_port" env:"APP_PORT" validate:"required|numeric"`
	AppHost string `yaml:"app_host" env:"APP_HOST" validate:"required|string"`
	AppEnv  string `yaml:"app_env" env:"APP_ENV" validate:"required|string"`
	AppName string `yaml:"app_name" env:"APP_NAME" validate:"required|string"`

	// // Will add support for later :)
	// GithubClientId     string `yaml:"github_client_id" env:"GITHUB_CLIENT_ID"`
	// GithubClientSecret string `yaml:"github_client_secret" env:"GITHUB_CLIENT_SECRET,secret"`
	// GithubCallbackUrl  string `yaml:"github_callback_url" env:"GITHUB_CALLBACK_URL"`

	SessionSecret string `yaml:"session_secret" env:"SESSION_SECRET,secret" validate:"required|string"`
	SessionName   string `yaml:"session_name" env:"SESSION_NAME" validate:"required|string"`

	PostgresUser     string `yaml:"postgres_user" env:"POSTGRES_USER" validate:"required|string"`
	PostgresPassword string `yaml:"postgres_password" env:"POSTGRES_PASSWORD,secret" validate:"string"`
	PostgresHost     string `yaml:"postgres_host" env:"POSTGRES_HOST" validate:"string"`
	PostgresPort     int    `yaml:"postgres_port" env:"POSTGRES_PORT" validate:"numeric"`
	PostgresDb       string `yaml:"postgres_db" env:"POSTGRES_DB" validate:"required|string"`
	PostgresSslMode  string `yaml:"postgres_ssl_mode" env:"POSTGRES_SSL_MODE" validate:"string"`
}

func (c *Config) Validate() error {
	v := validate.Struct(c)
	if v.Validate() {
		return nil
	}
	return v.Errors
}

func Load(file string, logger *zap.SugaredLogger) (*Config, error) {
	c := Config{
		AppPort:         defaultAppPort,
		AppHost:         defaultAppHost,
		AppEnv:          defaultAppEnv,
		AppName:         defaultAppName,
		SessionName:     defaultSessionName,
		PostgresSslMode: defaultSslMode,
		PostgresDb:      defaultDb,
		PostgresUser:    defaultUser,
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
