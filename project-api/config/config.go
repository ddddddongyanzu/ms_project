package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"test.com/project-common/logs"
)

var C = InitConfig()

type Config struct {
	viper      *viper.Viper
	SC         *ServerConfig
	GC         *GrpcConfig
	EtcdConfig *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
}

func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	workDir, _ := os.Getwd()
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath("/etc/ms_project/user")
	conf.viper.AddConfigPath(workDir + "/config")

	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	conf.ReadServerConfig()
	conf.InitZapLog()
	conf.ReadEtcdConfig()
	return conf
}

func (c *Config) InitZapLog() {
	lc := &logs.LogConfig{
		DebugFileName: c.viper.GetString("zap.debugFileName"),
		InfoFileName:  c.viper.GetString("zap.infoFileName"),
		WarnFileName:  c.viper.GetString("zap.warnFileName"),
		MaxSize:       c.viper.GetInt("zap.maxSize"),
		MaxBackups:    c.viper.GetInt("zap.maxBackups"),
		MaxAge:        c.viper.GetInt("zap.maxAge"),
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *Config) ReadServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.SC = sc
}

func (c *Config) ReadEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln(err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec
}
