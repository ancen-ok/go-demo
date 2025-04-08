package core

var Config Configuration

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
	RunModel       string   `yaml:"run_model" json:"runModel"`
	WhiteList      []string `yaml:"white_list" json:"whiteList"`
}

// Database 数据库配置
type Database struct {
	DBMode   string `yaml:"db_mode" json:"db_mode"`
	DBName   string `yaml:"db_name" json:"db_name"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}

// Redis 配置
type Redis struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
}

type JWTConfig struct {
	Issuer      string `yaml:"issuer" json:"issuer"`
	ExpiresTime int64  `yaml:"expires_time" json:"expires_time"`
	SecretKey   string `yaml:"secret_key" json:"secret_key"`
}

type Logger struct {
	Prefix string `yaml:"prefix" json:"prefix"`
}

func InitConfig() {
	/*var (
		config Configuration
	)*/

}
