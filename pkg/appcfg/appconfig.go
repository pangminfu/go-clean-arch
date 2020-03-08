package appcfg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Load() (*viper.Viper, error) {
	config := viper.New()
	config.AutomaticEnv()
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.SetConfigName("application")
	config.SetConfigType("toml")
	config.AddConfigPath("./configs")
	err := config.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error on parsing config file: %s", err)
	}

	return config, nil
}
