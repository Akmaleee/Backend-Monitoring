package helper

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	APP_NAME                 string `mapstructure:"APP_NAME"`
	APP_ENVIRONMENT          string `mapstructure:"APP_ENVIRONMENT"`
	APP_URL                  string `mapstructure:"APP_URL"`
	API_KEY                  string `mapstructure:"API_KEY"`
	JWT_SECRET               string `mapstructure:"JWT_SECRET"`
	JWT_EXPIRE               int    `mapstructure:"JWT_EXPIRE"`
	INFRASTRUCTURE_MYSQL_DSN string `mapstructure:"INFRASTRUCTURE_MYSQL_DSN"`
	LDAP_SERVER              string `mapstructure:"LDAP_SERVER"`
	LDAP_PORT                int    `mapstructure:"LDAP_PORT"`
	LDAP_BASE_DN             string `mapstructure:"LDAP_BASE_DN"`
	LDAP_USER_DN             string `mapstructure:"LDAP_USER_DN"`
	LDAP_USE_TLS             bool   `mapstructure:"LDAP_USE_TLS"`
}

var (
	once     sync.Once
	instance *Config
	err      error
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		instance, err = loadConfig()
	})
	return instance, err
}

func loadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	return &config, err
}
