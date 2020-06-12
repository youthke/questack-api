package conf

import (
	"github.com/BurntSushi/toml"
	"log"
	"runtime"
)

var conf Config

type (
	Config struct {
		DB DBConfig
	}

	DBConfig struct {
		User string
		Password string
		IP string
	}

)

func Setup(file string) error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	_, err := toml.DecodeFile(file, &conf)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}

