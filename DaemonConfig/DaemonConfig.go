package DaemonConfig

import (
	"encoding/json"
	"io/ioutil"
)

type Daemon struct {
	Name        string `json:"name"`
	Cmd         string `json:"cmd"`
	Enable      bool   `json:enable`
	WorkerCount int    `json:"worker_count"`
}

type Config struct {
	ready   bool
	Daemons []Daemon `json:"daemons"`
}

var configData Config
var configHash string

func GetConfig(configPath string) Config {
	if configData.ready == false {
		config := Config{}
		text := ReadConfig(configPath)
		err := json.Unmarshal(text, &config)
		if err != nil {
			panic("Can't parse config json")
		}
		configData = config
		configData.ready = true
	}
	return configData
}

func ReadConfig(configPath string) []byte {
	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("Can't read config file")
	}
	return dat
}
