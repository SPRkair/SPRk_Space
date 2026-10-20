package config

import (
	"SPRk_Space/internal/utils"
	"fmt"
	"os"
)

type StorageDriverType string

const (
	DriverDB            StorageDriverType = "DB"
	DriverCache         StorageDriverType = "Cache"
	DriverFakeStorage   StorageDriverType = "FakeStorage"
	DriverHybridStorage StorageDriverType = "HybridStorage"
)

type Config struct {
	App         App         `json:"app" yaml:"app" ini:"app"`
	Logger      Logger      `json:"logger" yaml:"logger" ini:"logger"`
	Storage     Storage     `json:"storage" yaml:"storage" ini:"storage"`
	Services    Services    `json:"service" yaml:"service" ini:"service"`
	Controllers Controllers `json:"controllers" yaml:"controllers" ini:"controllers"`
}

type Logger struct {
	Write_file   bool   `json:"write_file" yaml:"write_file" ini:"write_file"`
	Log_dir      string `json:"log_dir" yaml:"log_dir" ini:"log_dir"`
	Log_filename string `json:"log_filename" yaml:"log_filename" ini:"log_filename"`
}

type Storage struct {
	StorageDriver StorageDriverType `json:"storagedriver" yaml:"storagedriver" ini:"storagedriver"`
	Cache         Cache             `json:"cache" yaml:"cache" ini:"cache"`
	DB            DB                `json:"db" yaml:"db" ini:"db"`
	FakeStorage   FakeStorage       `json:"fakestorage" yaml:"fakestorage" ini:"fakestorage"`
	HybridStorage HybridStorage     `json:"hybridstorage" yaml:"hybridstorage" ini:"hybridstorage"`
}

type DB struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type FakeStorage struct {
	DirectoryFakeStorage string `json:"DirectoryFakeStorage" yaml:"DirectoryFakeStorage" ini:"DirectoryFakeStorage"`
}

type App struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type Services struct {
	FakeService FakeService `json:"fakesrvice" yaml:"fakesrvice" ini:"fakesrvice"`
	Auth        Auth        `json:"auth" yaml:"auth" ini:"auth"`
	Executor    Executor    `json:"executor" yaml:"executor" ini:"executor"`
}

type FakeService struct {
}

type Auth struct {
	MinLoginRunes    int `json:"minloginrunes" yaml:"minloginrunes" ini:"minloginrunes"`
	MaxLoginRunes    int `json:"maxloginrunes" yaml:"maxloginrunes" ini:"maxloginrunes"`
	MinUserNameRunes int `json:"minusernamerunes" yaml:"minusernamerunes" ini:"minusernamerunes"`
	MaxUserNameRunes int `json:"maxusernamerunes" yaml:"maxusernamerunes" ini:"maxusernamerunes"`
	MinPassLen       int `json:"minpasslen" yaml:"minpasslen" ini:"minpasslen"`
}

type Controllers struct {
	HTTP     HTTP     `json:"http" yaml:"http" ini:"http"`
	Telegram Telegram `json:"telegram" yaml:"telegram" ini:"telegram"`
	Cli      Cli      `json:"cli" yaml:"cli" ini:"cli"`
}

type HTTP struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type Telegram struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type Cli struct {
	Field1 string `json:"field1" yaml:"field1" ini:"field1"`
	Field2 string `json:"field2" yaml:"field2" ini:"field2"`
}

type Cache struct {
	AutoDownloadFile bool `json:"auto_download_file" yaml:"auto_download_file" ini:"auto_download_file"`
}

type HybridStorage struct {
	AdditionalStorage StorageDriverType `json:"additionalstorage" yaml:"additionalstorage" ini:"additionalstorage"`
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
	err = c.validate()
	if err != nil {
		return &defaultConfig, err
	}
	return &c, nil
}

func getDefaultConfig() Config {
	return Config{
		App: App{
			Field1: "",
			Field2: "",
		},
		Logger: Logger{
			Log_dir:      "./log/",
			Log_filename: ".log",
			Write_file:   true,
		},
		Storage: Storage{
			StorageDriver: DriverHybridStorage,
			Cache: Cache{
				AutoDownloadFile: true,
			},
			DB: DB{
				Field1: "",
				Field2: "",
			},
			FakeStorage: FakeStorage{
				DirectoryFakeStorage: "./TestStorage/",
			},
			HybridStorage: HybridStorage{
				AdditionalStorage: DriverFakeStorage,
			},
		},
		Services: Services{
			Executor: Executor{
				CacheSyncWaitDuration: 60,
				TaskWaitDuration:      100,
				LogSaverWaitDuration:  100,
			},
			FakeService: FakeService{},
			Auth: Auth{
				MinLoginRunes: 3,
				MaxLoginRunes: 64,
				MinPassLen:    32,
			},
		},
		Controllers: Controllers{
			HTTP: HTTP{
				Field1: "",
				Field2: "",
			},
			Telegram: Telegram{
				Field1: "",
				Field2: "",
			},
			Cli: Cli{
				Field1: "",
				Field2: "",
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

func (cfg *Config) validate() error {
	if !cfg.Storage.StorageDriver.IsValid() {
		return fmt.Errorf("validate: unsupported storage driver %q", cfg.Storage.StorageDriver)
	}
	return nil
}

func (d StorageDriverType) IsValid() bool {
	switch d {
	case DriverDB, DriverCache, DriverFakeStorage:
		return true
	default:
		return false
	}
}
