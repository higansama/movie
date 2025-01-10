package config

import (
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var Cfg Config
var singleton sync.Once

type Config struct {
	AppAttribute struct {
		Name string `mapstructure:"name"`
		Env  string `mapstructure:"env"` // dev/stage/prod
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"app_attribute"`
	MySqlConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
		ShowLog  bool   `mapstructure:"show_log"`
	} `mapstructure:"mysql_config"`
}

func LoadConfig(cfgPath string) error {
	var oerr error
	singleton.Do(func() {
		defaultPath := "config.yml"

		if cfgPath == "" {
			cfgPath = defaultPath
		}

		// Silently make config file if it doesn't exist
		_, err := os.Stat(cfgPath)
		if err != nil {
			if os.IsNotExist(err) {
				cfgFile, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_WRONLY, 0644)
				defer cfgFile.Close()
				if err == nil {
					expCfgFile, err := os.Open(defaultPath + ".example")
					defer expCfgFile.Close()
					if err == nil {
						_, _ = io.Copy(cfgFile, expCfgFile)
					}
				}
			}
		}

		viper.SetConfigFile(cfgPath)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		err = viper.ReadInConfig()
		if err != nil {
			oerr = err
			return
		}

		// create viper inititalization for config
		viper.SetDefault("app_attribute.name", "movie-app")
		viper.SetDefault("app_attribute.env", "dev")
		viper.SetDefault("app_attribute.port", "8008")
		viper.SetDefault("app_attribute.host", "localhost")
		// mysql config
		viper.SetDefault("mysql_config.host", "localhost")
		viper.SetDefault("mysql_config.port", "3306")
		viper.SetDefault("mysql_config.user", "root")
		viper.SetDefault("mysql_config.password", "")
		viper.SetDefault("mysql_config.database", "movies")

		err = viper.Unmarshal(&Cfg)
		if err != nil {
			log.Fatalln("cannot unmarshaling config")
		}
	})
	return oerr
}
