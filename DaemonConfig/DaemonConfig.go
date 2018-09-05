package DaemonConfig

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
)

type Daemon struct {
	Name        string `json:"name"`
	Cmd         string `json:"cmd"`
	Enable      bool   `json:enable`
	LogFile     string `json:"log_file"`
	WorkerCount int    `json:"worker_count"`
	StartDelay  int    `json:"start_delay"`
}

type Config struct {
	ready              bool
	Daemons            []Daemon `json:"daemons"`
	CheckConfigChanges bool     `json:"check_config_changes"`
}

var configData Config
var configHash string
var configPath string

func GetConfig(inputConfigPath string) Config {
	if configData.ready == false {
		configPath = inputConfigPath
		config := Config{}
		text := ReadConfig()
		configHash = makeConfigHash(text)
		err := json.Unmarshal(text, &config)
		if err != nil {
			panic("Can't parse config json")
		}
		configData = config
		configData.ready = true
	}
	return configData
}

func ReadConfig() []byte {
	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("Can't read config file")
	}
	return dat
}

func makeConfigHash(text []byte) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IsConfigChanged() bool {
	return configHash != makeConfigHash(ReadConfig())
}
