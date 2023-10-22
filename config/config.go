package app_config

import (
	"github.com/spf13/viper"
)

const (
	varNameEnv                    = "APP_ENV"
	varNamePrometheusSubsystemEnv = "PROMETHEUS_SUBSYSTEM"
	varNameLogLevelEnv            = "LOG_LEVEL"
	varNameDevelModeEnv           = "DEVEL_MODE"
)

type Config struct {
	env                 string
	prometheusSubsystem string
	logLevel            string
	ginForceDebug       bool
	develMode           bool
}

func CreateConfig() *Config {
	viper.AutomaticEnv()
	env := viper.GetString(varNameEnv)
	logLevel := viper.GetString(varNameLogLevelEnv)
	prometheusSubsystem := viper.GetString(varNamePrometheusSubsystemEnv)
	develMode := viper.GetBool(varNameDevelModeEnv)

	cfg := Config{
		env:                 env,
		prometheusSubsystem: prometheusSubsystem,
		logLevel:            logLevel,
		develMode:           develMode,
	}

	return &cfg
}

func (c *Config) Env() string {
	return c.env
}

func (c *Config) PrometheusSubsystem() string {
	return c.prometheusSubsystem
}

func (c *Config) LogLevel() string {
	return c.logLevel
}

func (c *Config) DevelMode() bool {
	return c.develMode
}
