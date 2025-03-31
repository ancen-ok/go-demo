package core

var Config Configuration

type Configuration struct {
	Web    Server    `yaml:"web"`
	Db     Database  `yaml:"db"`
	Redis  Redis     `yaml:"redis"`
	Jwt    JWTConfig `yaml:"jwt"`
	Logger Logger    `yaml:"logger"`
}

type Server struct {
	Port           int64    `yaml:"port"`
	ContextPath    string   `yaml:"context_path"`
	ReadTimeout    int64    `yaml:"read_timeout"`
	WriteTimeout   int64    `yaml:"write_timeout"`
	MaxHeaderBytes int      `yaml:"max_header_bytes"`
	RunModel       string   `yaml:"run_model"`
	WhiteList      []string `yaml:"white_list"`
}

// Database 数据库配置
type Database struct {
	DBMode   string `yaml:"db_mode"`
	DBName   string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Redis 配置
type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Issuer      string `yaml:"issuer"`
	ExpiresTime int64  `yaml:"expires_time"`
	SecretKey   string `yaml:"secret_key"`
}

type Logger struct {
	Prefix string `yaml:"prefix"`
}

func InitConfig() {
	/*var (
		config Configuration
	)*/

}
