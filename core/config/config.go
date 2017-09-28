package config

import (
	"bytes"
	"log"
	"strings"
	"errors"

	"github.com/spf13/viper"
	"github.com/merakiVE/CVDI/core/utils"

)

var PATH_FILES_CONFIG []string

type Configuration struct {
	instanceConfig *viper.Viper
	hasLoaded      bool
}

func init() {
	PATH_FILES_CONFIG = make([]string, 0)
	PATH_FILES_CONFIG = []string{"cvdi.conf", "/etc/cvdi/cvdi.conf", "/opt/cvdi/cvdi.conf"}

	viper.SetConfigType("json")
}

func (this *Configuration) verifyExistFileConfig() (string, error) {
	for _, path_config := range PATH_FILES_CONFIG {
		if utils.Exists(path_config) {
			return path_config, nil
		}
	}
	return "", errors.New("No exist file config, create config file in " + strings.Join(PATH_FILES_CONFIG, ","))
}

func (this *Configuration) Load() () {
	path_config, err := this.verifyExistFileConfig()

	if err != nil {
		panic(err)
	}

	data, err := utils.ReadBinaryFile(path_config)

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
