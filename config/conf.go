package config

// 声明Config全局变量
var Config *Conf

type Conf struct {
	Mysql        MysqlConfig
	Redis        RedisConfig
	Env          EnvConfig
	Jwt          JwtConfig
	DefaultAdmin DefaultAdminConfig
	TestNet      TestNetConfig
	MainNet      MainNetConfig
	Token        TokenConfig
	Threshold    ThresholdConfig
	Email        EmailConfig
}

type MysqlConfig struct {
	Host         string `toml:"host"`
	Port         string `toml:"port"`
	DbName       string `toml:"db_name"`
	UserName     string `toml:"user_name"`
	Password     string `toml:"password"`
	MaxOpenConns int    `toml:"max_open_coons"`
	MaxIdelConns int    `toml:"max_idel_conns"`
	MaxLifeTime  int    `toml:"max_life_time"`
}

type RedisConfig struct {
	Host        string `toml:"host"`
	Port        string `toml:"port"`
	Db          int    `toml:"db"`
	Password    string `toml:"password"`
	MaxIdle     int    `toml:"max_idel"`
	MaxActive   int    `toml:"max_active"`
	IdleTimeOut int    `toml:"idle_timeout"`
}

type EnvConfig struct {
	Port               string `toml:"port"`
	Version            string `toml:"version"`
	Protocol           string `toml:"protocol"`
	DomainName         string `toml:"domain_name"`
	TaskDuration       int64  `toml:"task_duration"`
	WssTimeOutDuration int64  `toml:"wss_timeout_duration"`
	TaskExtendDuration int64  `toml:"task_extend_duration"`
}

type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`
	ExpireTime int    `toml:"expire_time"`
}

type DefaultAdminConfig struct {
	UserName string `toml:"userName"`
	Password string `toml:"password"`
}

type TestNetConfig struct {
	ChainId              string `toml:"chain_id"`
	NetUrl               string `toml:"net_url"`
	PlgrAddress          string `toml:"plgr_address"`
	PledgePoolToken      string `toml:"pledge_pool_token"`
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"`
}

type MainNetConfig struct {
	ChainId              string `toml:"chain_id"`
	NetUrl               string `toml:"net_url"`
	PlgrAddress          string `toml:"plgr_address"`
	PledgePoolToken      string `toml:"pledge_pool_token"`
	BscPledgeOracleToken string `toml:"bsc_pledge_oracle_token"`
}

type TokenConfig struct {
	LogoUrl string `toml:"logo_url"`
}

type ThresholdConfig struct {
	PledgePoolTokenThresholdBnb string `toml:"pledge_pool_token_threshold_bnb"`
}

type EmailConfig struct {
	Username string   `toml:"username"`
	Pwd      string   `toml:"pwd"`
	Host     string   `toml:"host"`
	Port     string   `toml:"port"`
	From     string   `toml:"from"`
	Subject  string   `toml:"subject"`
	To       []string `toml:"to"`
	Cc       []string `toml:"cc"`
}
