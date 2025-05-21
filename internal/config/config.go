package config

import (
	"fmt"
	"os"
	"time"

	"echo_rest_timer/internal/models"

	"github.com/joho/godotenv"
)

func Load() (*models.Config, error) {

	if err := godotenv.Load("ini.env"); err != nil {
		return nil, fmt.Errorf("error loading ini.env: %w", err)
	}

	environment := models.EnvironmentType(os.Getenv("ENVIRONMENT"))
	if environment == "" {
	} else if !environment.IsValid() {
		return nil, fmt.Errorf("environment variable ENVIRONMENT wrong format set dev/stag/prod")
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("environment variable PORT is not set")
	}

	user := os.Getenv("RBUSER")
	if user == "" {
		return nil, fmt.Errorf("environment variable RBUSER is not set")
	}

	datestr := os.Getenv("DATE")
	if user == "" {
		return nil, fmt.Errorf("environment variable DATE is not set")
	}

	date, err := time.Parse("02.01.2006", datestr)
	if err != nil {
		return nil, fmt.Errorf("environment variable DATE wrong format %v", err)
	}

	cfg := &models.Config{
		Port:        port,
		User:        user,
		Environment: environment,
		Date:        date,
	}

	return cfg, nil
}
