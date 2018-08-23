package DaemonConfig

import (
	//	"crypto/md5"
	//	"encoding/hex"
	"encoding/json"
	"io/ioutil"
)

const (
	configPath = "config/config.json"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetConfig() Config {
	if configData.ready == false {
		config := Config{}

		text := ReadConfig()
		json.Unmarshal(text, &config)

		configData = config
		configData.ready = true
	}
	return configData
}

func ReadConfig() []byte {
	dat, err := ioutil.ReadFile(configPath)
	check(err)
	return dat
}

// func CheckConfigChanges(){
// 	currentConfig = ReadConfig();
//     currentConfigHash = MakeConfigHash(currentConfig)

//     if (currentConfigHash)
// }

// func GetConfigHash() string {
// 	return configHash
// }

// func SetConfigHash(hash string) {
// 	configHash = hash
// }

// func MakeConfigHash(text string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(text))

// 	return hex.EncodeToString(hasher.Sum(nil))
// }
