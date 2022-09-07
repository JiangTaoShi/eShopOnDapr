package configs

import (
	"fmt"
	"github.com/JiangTaoShi/eShopOnDapr/pkg/env"
	"github.com/fsnotify/fsnotify"
	"time"
)
import "github.com/spf13/viper"

var config = new(Config)

func Get() Config {
	return *config
}

type Config struct {
	MySQL struct {
		Read struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"read"`
		Write struct {
			Addr string `toml:"addr"`
			User string `toml:"user"`
			Pass string `toml:"pass"`
			Name string `toml:"name"`
		} `toml:"write"`
		Base struct {
			MaxOpenConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
		} `toml:"base"`
	} `toml:"mysql"`
	Server struct {
		RunMode      string        `toml:"runMode"`
		HttpPort     int           `toml:"httpPort"`
		ReadTimeout  time.Duration `toml:"readTimeout"`
		WriteTimeout time.Duration `toml:"writeTimeout"`
	} `toml:"mysql"`
	Language struct {
		Local string `toml:"local"`
	} `toml:"language"`
}

func init() {
	viper.SetConfigName(env.Active().Value() + "-configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(config)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}
