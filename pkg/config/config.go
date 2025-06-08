package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	SUSCRIPTION_MAX_MESSAGES int    `json:"SUSCRIPTION_MAX_MESSAGES"`
	GCP_SERVICE_ACCOUNT_FILE string `json:"GCP_SERVICE_ACCOUNT_FILE"`
	GCP_BASE_URL             string `json:"GCP_BASE_URL"`
}

var AppConfig Config

func Load() error {
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &AppConfig)
	if err != nil {
		return err
	}

	return nil
}
