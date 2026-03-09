package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	filename string = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	newConfig := Config{}
	homeDir, _ := os.UserHomeDir()
	path := filepath.Join(homeDir, filename)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Read failed.")
		return newConfig
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	json.Unmarshal(data, &newConfig)
	return newConfig
}

func (c Config) SetUser(user string) {
	c.CurrentUserName = user
	homeDir, _ := os.UserHomeDir()
	path := filepath.Join(homeDir, filename)
	data, err := json.Marshal(c)

	if err != nil {
		fmt.Println("Marshaling data failed.")
		return
	}
	err = os.WriteFile(path, data, 0755)
	if err != nil {
		fmt.Println("Write failed.")
		return
	}
}
