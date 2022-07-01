package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

type configFile struct {
	ApiVersion string `yaml:"apiVersion"`
	ApiKey     string `yaml:"apiKey"`
	ApiUrl     string `yaml:"apiUrl"`
}

var instance *Info
var once sync.Once

func GetConfigInfo() *Info {
	once.Do(func() {
		instance = &Info{}
		_ = instance.Complete()
	})
	return instance
}

func LoadConfigFile() (bool, error) {
	arvanConfig := GetConfigInfo()

	data, err := ioutil.ReadFile(arvanConfig.configFilePath)

	if err != nil {
		return false, err
	}

	configFileStruct := configFile{}

	err = yaml.Unmarshal(data, &configFileStruct)

	if err != nil {
		return false, err
	}

	arvanConfig.apiKey = configFileStruct.ApiKey
	arvanConfig.apiUrl = configFileStruct.ApiUrl

	return true, nil
}
