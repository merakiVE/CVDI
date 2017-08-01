package config

import (
	"bytes"
	"log"

	"github.com/spf13/viper"
	"github.com/merakiVE/CVDI/core/utils"
)

const (
	PATH_CONFIG = "cvdi.conf"
)

func init() {
	viper.SetConfigType("json")
}

func Load() () {

	data, err := utils.ReadBinaryFile(PATH_CONFIG)

	if err != nil {
		log.Fatal("Error reading config file")
		return
	}

	viper.ReadConfig(bytes.NewBuffer(data))
}

func Get(_key string) (interface{}) {
	return viper.Get(_key)
}

func GetString(_key string) (string) {
	return viper.GetString(_key)
}

func GetConfig() (*viper.Viper) {
	return viper.GetViper()
}
