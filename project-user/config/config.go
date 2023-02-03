package config

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var AppConfig = InitConfig()

type Config struct {
	viper *viper.Viper
	SC    *ServerConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

func (c *Config) ReadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.viper.GetString("redis.host"), c.viper.GetString("redis.port")),
		Password: c.viper.GetString("redis.password"),
		DB:       c.viper.GetInt("redis.db"),
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func InitConfig() *Config {
	dir, _ := os.Getwd()
	conf := &Config{
		viper: viper.New(),
	}
	conf.viper.SetConfigName("config")        // name of config file (without extension)
	conf.viper.SetConfigType("yml")           // REQUIRED if the config file does not have the extension in the name
	conf.viper.AddConfigPath(dir + "/config") // optionally look for config in the working directory
	err := conf.viper.ReadInConfig()          // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 读取服务配置
	conf.ReadServerConfig()

	return conf
}
