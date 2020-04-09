package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	Runmode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

var cfg *ini.File

// 初始化配置文件
func Setup() {
	var err error
	cfg, err = ini.Load("conf/gravitation.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/gravitation.ini': %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// 匹配配置选项
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.mapTo %s err: %v", section, err)
	}
}