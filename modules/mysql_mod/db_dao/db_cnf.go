package db_dao

type DB struct {
	Ip              string     `yaml:"ip"`
	Port            int        `yaml:"port"`
	Database        string     `yaml:"database"`
	User            string     `yaml:"user"`
	Password        string     `yaml:"password"`
	Params          string     `yaml:"params"`
	LogMode         int        `yaml:"log_mode"`
	MaxConnections  int        `yaml:"max_connections"`   //超出连接限制
	MaxOpenConns    int        `yaml:"max_open_conns"`    //最大连接数,开发越多越好,则mysql 连接数而定
	MaxIdleConns    int        `yaml:"max_idle_conns"`    //最多保留空闲连接, 允许更多的空闲连接将提高性能，因为这样可以减少从头开始建立新连接的可能性，从而有助于节省资源,MaxIdleConns应该始终小于或等于MaxOpenConns
	ConnMaxLifetime int        `yaml:"conn_max_lifetime"` //ConnMaxLifetime越短，从零开始创建连接的频率就越高
	SlaveArray      []MysqlCnf `yaml:"slave_array"`
}

type MysqlCnf struct {
	Ip       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Params   string `yaml:"params"`
	LogMode  int    `yaml:"log_mode"`
}
