package setting

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var (
	//Cfg is the service configuration.
	Cfg *Config
	//RunMode is the service running mode tag configuration.
	RunMode string
	//HTTPPort is the service listening port configuration.
	HTTPPort int
	//HTTPProto is the http proto
	HTTPProto string
	//ReadTimeout is the http service read time-out limit configuration.
	ReadTimeout time.Duration
	//WriteTimeout is the http serice write time-out limit configuration.
	WriteTimeout time.Duration
	//PageSize is the service default pagging number while execute sql queries.
	PageSize int
	//JwtSecret is the AUTH JWT secret configuration
	JwtSecret string
	//ServiceName is the service's name
	ServiceName string
)

//Load loads configuration from a yaml file.
func Load(confPath string) {
	if Cfg != nil {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	// yamlFile, err := ioutil.ReadFile("./config/srv.yaml")
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(fmt.Sprintf("load yaml configuration file error: %v", err))
	}
	conf := &Config{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		panic(fmt.Sprintf("Unmarshall yaml configuration file failed: %v", err))
	}
	Cfg = conf
	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.RunMode
}

func loadServer() {
	HTTPPort = Cfg.Server.HTTPPort
	HTTPProto = Cfg.Server.HTTPProto
	ReadTimeout = time.Duration(Cfg.Server.ReadTimeout) * time.Second
	WriteTimeout = time.Duration(Cfg.Server.WriteTimeout) * time.Second
}

func loadApp() {
	ServiceName = Cfg.App.ServiceName
	JwtSecret = Cfg.App.JwtSecret
	PageSize = Cfg.App.PageSize
}
