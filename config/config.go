package config

import (
	"bufio"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Port      string `json:"port"`
		AllowCors bool   `json:"allow_cors"`
	}
	Services struct {
		Auth struct {
			DB            DB     `json:"db"`
			JWTKey        string `json:"jwt_key"`
			PasswordCrypt string `json:"password_crypt"`
			SMTP          struct {
			} `json:"smtp"`
		} `json:"authorization"`
		FlashCards struct {
			DB DB
		}
	} `json:"services"`
}

func (c *Config) GetServiceDB(service int) DB {
	switch service {
	case AUTH:
		return c.Services.Auth.DB
		// can add more services here to get DB
	}
	return DB{}
}

func InitConfig(log *zap.Logger) *Config {
	config, ok := readFromFiles(log)
	log.Info("config initializing")
	if !ok {
		log.Fatal("invalid config")
	}
	return config
}

func readFromFiles(log *zap.Logger) (*Config, bool) {
	config := &Config{}
	path := "./files/"
	log.Info("file path: " + path)
	if !pathExists(path) {
		if err := os.MkdirAll(path, 0777); err != nil {
			log.Error("making path: ", zap.Error(err))
		}
	}
	if !pathExists(path + "config.json") {
		makeBlankConfig(path+"config.json", log)
		log.Info("made blank config")
		log.Info("generated new file, please fill in the config.")
		os.Exit(1)
		return nil, false
	}
	f, err := os.Open(path + "config.json")
	defer f.Close()
	if err != nil {

		log.Fatal("open file error")
	}
	raw, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("reading from file: ", zap.Error(err))
	}
	if err := json.Unmarshal(raw, &config); err != nil {
		log.Fatal("unmarshal from file: ", zap.Error(err))
		os.Exit(1)
	}
	return config, true
}

func makeBlankConfig(path string, log *zap.Logger) {
	f, err := os.Create(path)
	defer f.Close()

	if err != nil {
		return
	}
	writer := bufio.NewWriter(f)
	var c Config
	raw, err := json.Marshal(c)
	if err != nil {
		log.Fatal("writing config: ", zap.Error(err))
	}
	_, err = writer.Write(raw)
	if err != nil {
		log.Fatal("writing config: ", zap.Error(err))
	}
	writer.Flush()
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
