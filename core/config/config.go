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

type Configuration struct {
	instanceConfig *viper.Viper
	hasLoaded      bool
}

func init() {
	viper.SetConfigType("json")
}

func (this *Configuration) Load() () {

	data, err := utils.ReadBinaryFile(PATH_CONFIG)

	if err != nil {
		log.Fatal("Error reading config file")
		return
	}

	this.hasLoaded = true

	viper.ReadConfig(bytes.NewBuffer(data))
}

func (this *Configuration) Get(_key string) (interface{}) {
	return viper.Get(_key)
}

func (this *Configuration) GetString(_key string) (string) {
	return viper.GetString(_key)
}

func (this *Configuration) GetNativeConfig() (*viper.Viper) {
	return this.instanceConfig
}
