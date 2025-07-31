package config

import "os"

var DBPATH = ConfigFilePath() + "/todo-db.todo"

func ConfigFilePath() string {
	configDir, _ := os.UserConfigDir()
	return configDir + "/todo-cli"
}

func CheckDatabase() bool {
	filepath := DBPATH
	if _, err := os.ReadDir(ConfigFilePath()); err != nil {
		os.Mkdir(ConfigFilePath(), 0777)
	}
	_, err := os.ReadFile(filepath)
	return err == nil
}