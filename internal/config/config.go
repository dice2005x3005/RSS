package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	DbURL string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func (c Config) SetUser(name string) error {
	c.CurrentUser = name
	home, err := os.UserHomeDir()
	url := fmt.Sprintf("%s/.gatorconfig.json", home)
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(url, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error){
	home, err := os.UserHomeDir()
	url := fmt.Sprintf("%s/.gatorconfig.json", home)
	data, err := os.ReadFile(url)
	var config Config
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}