package svc

import (
	"strings"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// LoadConfig - функция для загрузки конфига.
// сначала смотрим файл <name>.yaml, затем оверрайдим из consul и в конце из env.
// example: envPrefix = "MS_" or ""
func LoadConfig(envPrefix, name string) (*viper.Viper, error) {
	var v = viper.New()

	v.SetConfigName(name + ".yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			return nil, err
		}
	}

	v.AddRemoteProvider("consul", "localhost:8500", "configuration/"+name+"/data")
	if err := v.ReadRemoteConfig(); err != nil {
		if strings.Contains(err.Error(), "No Files Found") {
			// Config not found; ignore error
		} else {
			return nil, err
		}
	}

	r := strings.NewReplacer(".", "_")

	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(r)
	v.AutomaticEnv()

	return v, nil
}
