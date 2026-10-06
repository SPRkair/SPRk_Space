package config

import (
	"SPRk_Space/internal/utils"
	"fmt"
	"os"
)

type Config struct {
	App     App     `json:"app" yaml:"app" ini:"app"`
	HTTP    HTTP    `json:"http" yaml:"http" ini:"http"`
	Storage Storage `json:"storage" yaml:"storage" ini:"storage"`
	Service Service `json:"service" yaml:"service" ini:"service"`
}

type Storage struct {
	Cache Cache `json:"cache" yaml:"cache" ini:"cache"`
	DB    DB    `json:"db" yaml:"db" ini:"db"`
}

type HTTP struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type DB struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type App struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type Service struct {
	Executor Executor `json:"executor" yaml:"executor" ini:"executor"`
}

type Cache struct {
	AutoDownloadFile bool `json:"auto_download_file" yaml:"auto_download_file" ini:"auto_download_file"`
}

type Executor struct {
	CacheSyncWaitDuration int64 `json:"cache_sync_wait_duration" yaml:"cache_sync_wait_duration" ini:"cache_sync_wait_duration"`
	TaskWaitDuration      int64 `json:"task_wait_duration" yaml:"task_wait_duration" ini:"task_wait_duration"`
	LogSaverWaitDuration  int64 `json:"log_saver_wait_duration" yaml:"log_saver_wait_duration" ini:"log_saver_wait_duration"`
}

func NewConfig() (conf *Config, err error) {
	defaultConfig := getDefaultConfig()
	c := defaultConfig
	err = c.loadConfigFromFile(os.Getenv("CONFIG_FILE_PATH"))
	if err != nil {
		return &defaultConfig, fmt.Errorf("error: config: NewConfig: loadConfigFromFile: %w", err)
	}
	return &c, nil
}

func getDefaultConfig() Config {
	return Config{
		App: App{
			Field1: "",
			Field2: "",
		},
		HTTP: HTTP{
			Field1: "",
			Field2: "",
		},
		Storage: Storage{
			Cache: Cache{
				AutoDownloadFile: true,
			},
			DB: DB{
				Field1: "",
				Field2: "",
			},
		},
		Service: Service{
			Executor: Executor{
				CacheSyncWaitDuration: 60,
				TaskWaitDuration:      100,
				LogSaverWaitDuration:  100,
			},
		},
	}
}

func (c *Config) loadConfigFromFile(filename string) (err error) {
	err = utils.LoadFrom(filename, c)
	if err != nil {
		return err
	}
	return nil
}
