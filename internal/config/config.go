package config

// ! 这里的yaml标签是必须的，并且tag名称需要和配置文件中的配置项名称一致，否则无法解析
type Config struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
	Mode string `yaml:"Mode"`

	Database struct {
		DriverName string `yaml:"DriverName"`
		DataSource string `yaml:"DataSource"`
	} `yaml:"Database"`

	Redis struct {
		Addr     string `yaml:"Addr"`
		Password string `yaml:"Password"`
		DB       int    `yaml:"DB"`
	} `yaml:"Redis"`

	Auth struct { // jwt鉴权配置
		AccessSecret string `yaml:"AccessSecret"` // jwt密钥
		AccessExpire int64  `yaml:"AccessExpire"` // 有效期，单位：秒
	} `yaml:"Auth"`

	Log struct { // 配置中心配置
		Dir    string `yaml:"Dir"`
		Prefix string `yaml:"Prefix"`
	} `yaml:"Log"`
}
