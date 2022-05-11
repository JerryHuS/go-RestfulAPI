/**
 * @Author: alessonhu
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/5/5 11:04
 */

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	ModeDebug = "debug"
)

var (
	ServerConfig struct {
		HTTPPort string
		LogPath  string
		Mode     string
	}
	DBConfig struct {
		Driver      string
		Host        string
		Port        string
		User        string
		Password    string
		Name        string
		ConnectInfo string
	}
	RedisConfig struct {
		Address  []string
		Password string
	}
	ServiceConfig struct {
	}
)

// config ...
type config struct {
	Server   map[string]string `json:"server"`
	Redis    map[string]string `json:"redis"`
	Database map[string]string `json:"database"`
	Service  map[string]string `json:"service"`
}

// InitConfig init globalConfig
func InitConfig() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var c = &config{}
	err = json.Unmarshal(b, c)

	if err != nil {
		panic(err)
	}

	// Server
	ServerConfig.HTTPPort = c.Server["http_port"]
	ServerConfig.LogPath = c.Server["log_path"]
	ServerConfig.Mode = c.Server["mode"]

	// DB
	DBConfig.Driver = c.Database["driver"]
	DBConfig.Host = c.Database["host"]
	DBConfig.Port = c.Database["port"]
	DBConfig.User = c.Database["user"]
	DBConfig.Password = c.Database["password"]
	DBConfig.Name = c.Database["name"]
	DBConfig.ConnectInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Database["host"], c.Database["port"], c.Database["user"], c.Database["password"], c.Database["name"])

	// Redis
	hosts := strings.Split(c.Redis["hosts"], ";")
	ports := strings.Split(c.Redis["ports"], ";")
	for idx, host := range hosts {
		var port string
		if idx >= len(ports) {
			port = ports[0]
		} else {
			port = ports[idx]
		}
		RedisConfig.Address = append(RedisConfig.Address, fmt.Sprintf("%s:%s", host, port))
	}
	RedisConfig.Password = c.Redis["password"]
}
