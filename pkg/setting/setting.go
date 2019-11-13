package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int

	HTTPProto string

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	PageSize int

	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatalf("Failed to parse 'conf/app.ini': %v", err)
		return
	}
	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Failed to get section 'server': %v", err)
		return
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	HTTPProto = sec.Key("HTTP_PROTO").MustString("http")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60))
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60))
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Faile to get section 'app': %v", err)
		return
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}