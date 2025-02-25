package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const Fullpath = "/home/iain/.gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) { // reads the json file of my computer
	configstruct := Config{}
	homedir, err := os.UserHomeDir()
	if err != nil {
		return configstruct, err
	}
	path := filepath.Join(homedir, ".gatorconfig.json")

	file, err := os.Open(path) // need to OPEN the json file before we can read it

	if err != nil {
		return configstruct, err
	}
	defer file.Close() // need to close the file after os.Open

	data, err := io.ReadAll(file)
	if err != nil {
		return configstruct, err
	}

	err = json.Unmarshal(data, &configstruct)
	if err != nil {
		return configstruct, err
	}

	return configstruct, nil
}

func Write(cfg Config) error {

	jsonData, err := json.MarshalIndent(cfg, "", "  ") // marshals the information ALREADY in cfg struct
	if err != nil {
		return err
	}

	err = os.WriteFile(Fullpath, jsonData, 0644) // path where data goes, data im sending, permission code (read+write for owner)
	if err != nil {
		return err
	}
	return nil
}

func (c Config) SetUser(username string) error {

	c.CurrentUserName = username
	Write(c)

	return nil
}
