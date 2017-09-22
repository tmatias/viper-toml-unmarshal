package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host string `mapstructure:"hostname"`
	Port string
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type OutputConfig struct {
	File string
}

type Config struct {
	Db  DatabaseConfig `mapstructure:"database"`
	Out OutputConfig   `mapstructure:"output"`
}

func main() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		fmt.Printf("couldn't read config: %s", err)
	}
	fmt.Printf("host=%s port=%s user=%s pass=%s\n", c.Db.Host, c.Db.Port, c.Db.User, c.Db.Pass)
	fmt.Printf("output=%s\n", c.Out.File)
	fmt.Printf("%#v", c)
}
