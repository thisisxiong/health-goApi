package config

type MysqlConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Level    int
}

type ServerConf struct {
	Mysql   MysqlConf
	Debug   bool
	LogPath string
}
