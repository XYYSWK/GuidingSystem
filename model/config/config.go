package config

import "time"

type PrivateConfig struct {
	Mysql MysqlConfig `yaml:"Mysql"`
	Redis RedisConfig `yaml:"Redis"`
	Token Token       `yaml:"Token"`
}

type MysqlConfig struct {
	DBType      string `yaml:"DBType"`
	Username    string `yaml:"Username"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	DBName      string `yaml:"DBName"`
	TablePrefix string `yaml:"TablePrefix"`
	Charset     string `yaml:"Charset"`
	ParseTime   bool   `yaml:"ParseTime"`
	MaxIdle     int    `yaml:"MaxIdle"`
	MaxOpen     int    `yaml:"MaxOpen"`
}

type RedisConfig struct {
	Address   string        `yaml:"Address"`
	Password  string        `yaml:"Password"`
	DB        int           `yaml:"DB"`        // 数据库索引
	PoolSize  int           `yaml:"PoolSize"`  // Redis 连接池大小
	CacheTime time.Duration `yaml:"CacheTime"` // 缓存时间
}

type Token struct {
	Key                  string        `yaml:"Key"`                  // 生成 token 的密钥
	UserTokenDuration    time.Duration `yaml:"UserTokenDuration"`    // 用户 token 的有效期限
	AccountTokenDuration time.Duration `yaml:"AccountTokenDuration"` // 账户 token 的有效期限
	AuthorizationKey     string        `yaml:"AuthorizationKey"`     // 授权密钥，用于进行授权验证
	AuthorizationType    string        `yaml:"AuthorizationType"`    // 授权类型，指定授权的具体方式或策略
}

type PublicConfig struct {
	Server Server    `yaml:"Server"`
	App    AppConfig `yaml:"App"`
	Log    LogConfig `yaml:"Log"`
}

type Server struct {
	RunMode               string        `yaml:"RunMode"`               // gin 工作模式
	HttpPort              string        `yaml:"HttpPort"`              // 默认的 HTTP 监听端口号
	ReadTimeout           time.Duration `yaml:"ReadTimeout"`           // 允许读取的最大持续时间
	WriteTimeout          time.Duration `yaml:"WriteTimeout"`          // 允许写入的最大持续时间
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"` // 默认上下文超时
}

type AppConfig struct {
	Name      string `yaml:"Name"`
	Version   string `yaml:"Version"`
	StartTime string `yaml:"StartTime"` // 启动时间
	MachineID int64  `yaml:"MachineID"` // 机器ID
}

type LogConfig struct {
	Level         string `yaml:"Level"`         // 日志级别
	LogSavePath   string `yaml:"LogSavePath"`   // 日志保存路径
	LowLevelFile  string `yaml:"LowLevelFile"`  // 低级别日志文件名
	LogFileExt    string `yaml:"LogFileExt"`    // 日志文件扩展名
	HighLevelFile string `yaml:"HighLevelFile"` // 高级别日志文件名
	MaxSize       int    `yaml:"MaxSize"`       // 每个日志文件的最大尺寸
	MaxAge        int    `yaml:"MaxAge"`        // 保留的最大天数
	MaxBackups    int    `yaml:"MaxBackups"`    // 保留的最大备份数量
	Compress      bool   `yaml:"Compress"`      // 是否压缩日志文件
}
