package environment

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

type RedisConfig struct {
	Host     string
	Port     string
	CacheTTL int `default:"300"`
}

type DBConfig struct {
	Type string
	Host string
	Port string
	DB   string
	User string
	Pass string

	Debug bool
}

type OtherConfig struct{}

type ServerConfig struct {
	Host string
	Port string

	StopTimeout int `json:"stop_timeout" default:"60"`
}

type config struct {
	Server ServerConfig `json:"server_config"`

	DB    DBConfig    `json:"db_config"`
	Redis RedisConfig `json:"redis_config"`

	Other OtherConfig `json:"other_config"`
}

func (env *Env) parseConfig() (err error) {
	var configFile *os.File

	configFile, err = os.Open(*configFilePath)
	if env.Check(err) {
		return
	}

	configBytes, err := ioutil.ReadAll(configFile)
	if env.Check(err) {
		return
	}

	conf := config{}
	err = json.Unmarshal(configBytes, &conf)
	if env.Check(err) {
		return
	}

	env.Conf = &conf
	env.Log.Debugf("%+v\n", &conf)

	return nil
}
