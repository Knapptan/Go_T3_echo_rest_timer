package models

import (
	"time"

	"github.com/labstack/gommon/log"
)

type Config struct {
	Port        string
	Environment EnvironmentType
	User        string
	Date        time.Time
}

type EnvironmentType string

const (
	EnvDevelopment EnvironmentType = "dev"
	EnvStaging     EnvironmentType = "stag"
	EnvProduction  EnvironmentType = "prod"
)

func (e EnvironmentType) IsValid() bool {
	switch e {
	case EnvDevelopment, EnvStaging, EnvProduction:
		return true
	}
	return false
}

// 0 = DEBUG, 1 = INFO, 2 = WARN, 3 = ERROR
func (e EnvironmentType) LogLevel() log.Lvl {
	switch e {
	case EnvDevelopment:
		return log.DEBUG
	case EnvStaging:
		return log.INFO
	case EnvProduction:
		return log.WARN
	default:
		return log.ERROR
	}
}
