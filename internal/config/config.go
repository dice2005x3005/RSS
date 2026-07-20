package config

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	dbURL string `json:"db_url"`,
	currentUser string `json:"current_user_name"`
}

func (c Config) SetUser(name String) err {
	c.currentUser = name
	home := os.UserHomeDir()
	url := fmt.Sprintf(%s"/.gatorconfig.json", home)
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err := os.WriteFile(url, data)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error){
	home := os.UserHomeDir()
	url := fmt.Sprintf(%s"/.gatorconfig.json", home)
	data, err := os.ReadFile(url)
	var config Config
	if err != nil {
		return config, err
	}
	st, err := json.Unmarshal(data, &config)
	if err != nil {
		return Config{}{}, err
	}
	return config, nil
}