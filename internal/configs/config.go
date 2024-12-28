package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type resource struct {
	Name            string
	Endpoint        string
	Destination_url string
}

type configuration struct {
	Server struct {
		Host        string
		Listen_port string
	}
	Resources []resource
}

var Config *configuration

func NewConfiguration() (*configuration, error) {
	viper.AddConfigPath("data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config file: %s", err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}
	return Config, nil
}
