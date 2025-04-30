package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var Config *Configuration

type Configuration struct {
	Web    Server    `yaml:"web" json:"web"`
	Db     Database  `yaml:"db" json:"db"`
	Redis  Redis     `yaml:"redis" json:"redis"`
	Jwt    JWTConfig `yaml:"jwt" json:"jwt"`
	Logger Logger    `yaml:"logger" json:"logger"`
}

type Server struct {
	Port           int64    `yaml:"port" json:"port"`
	ContextPath    string   `yaml:"context_path" json:"context_path"`
	ReadTimeout    int64    `yaml:"read_timeout" json:"read_timeout"`
	WriteTimeout   int64    `yaml:"write_timeout" json:"write_timeout"`
	MaxHeaderBytes int      `yaml:"max_header_bytes" json:"max_header_bytes"`
	RunModel       string   `yaml:"model" json:"model"`
	WhiteList      []string `yaml:"white_list" json:"white_list"`
}

func (s *Server) Whites() []string {
	result := make([]string, 0)
	for _, item := range s.WhiteList {
		if s.ContextPath == "" {
			result = append(result, item)
		} else {
			if strings.HasPrefix(s.ContextPath, "/") {
				result = append(result, s.ContextPath+item)
			} else {
				result = append(result, "/"+s.ContextPath+item)
			}
		}
	}
	return s.WhiteList
}

// Database 数据库配置
type Database struct {
	DBMode   string `yaml:"db_mode" json:"db_mode"`
	DbName   string `yaml:"db_name" json:"db_name"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

// Redis 配置
type Redis struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
}

func (d Database) Link() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.DbName)
}

// JwtConfig jwt配置
type JWTConfig struct {
	Issuer      string `yaml:"issuer" json:"issuer"`
	ExpiresTime int64  `yaml:"expires_time" json:"expires_time"`
	SecretKey   string `yaml:"secret_key" json:"secret_key"`
}

type Logger struct {
	Prefix string `yaml:"prefix" json:"prefix"`
}

func InitConfig() {
	var (
		config Configuration
		bytes  []byte
		err    error
	)
	if bytes, err = os.ReadFile("config.yaml"); err != nil {
		panic("读取配置文件失败=>" + err.Error())
	}
	if err = yaml.Unmarshal(bytes, &config); err != nil {
		panic("解析配置文件失败=>" + err.Error())
	}
	Config = &config

}
