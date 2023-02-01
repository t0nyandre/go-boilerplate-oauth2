package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gookit/validate"
	"go.uber.org/zap"
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
	AppPort int    `json:"app_port" env:"APP_PORT" validate:"required|numeric"`
	AppHost string `json:"app_host" env:"APP_HOST" validate:"required|string"`
	AppEnv  string `json:"app_env" env:"APP_ENV" validate:"required|string"`
	AppName string `json:"app_name" env:"APP_NAME" validate:"required|string"`

	// // Will add support for later :)
	// GithubClientId     string `json:"github_client_id" env:"GITHUB_CLIENT_ID"`
	// GithubClientSecret string `json:"github_client_secret" env:"GITHUB_CLIENT_SECRET,secret"`
	// GithubCallbackUrl  string `json:"github_callback_url" env:"GITHUB_CALLBACK_URL"`

	SessionSecret string `json:"session_secret" env:"SESSION_SECRET,secret" validate:"required|string"`
	SessionName   string `json:"session_name" env:"SESSION_NAME" validate:"required|string"`

	PostgresUser     string `json:"postgres_user" env:"POSTGRES_USER" validate:"required|string"`
	PostgresPassword string `json:"postgres_password" env:"POSTGRES_PASSWORD,secret" validate:"string"`
	PostgresHost     string `json:"postgres_host" env:"POSTGRES_HOST" validate:"string"`
	PostgresPort     int    `json:"postgres_port" env:"POSTGRES_PORT" validate:"numeric"`
	PostgresDb       string `json:"postgres_db" env:"POSTGRES_DB" validate:"required|string"`
	PostgresSslMode  string `json:"postgres_ssl_mode" env:"POSTGRES_SSL_MODE" validate:"string"`
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
	if err = json.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// TODO: Load env vars

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil
}
