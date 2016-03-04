package config

import (
  "github.com/spf13/viper"
  "fmt"
)
//https://golang.org/doc/code.html
func load_config() {
  viper.SetConfigName("config") // name of config file (without extension)
  viper.AddConfigPath(".")   // path to look for the config file in
  viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
  viper.ReadInConfig() // Find and read the config file

  viper.GetString("host") // case insensitive Setting & Getting
  fmt.Printf(viper.GetString("host"))
}
