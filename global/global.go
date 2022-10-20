package global

import (
	"gorm.io/gorm"
	"health/config"
)

var (
	Conf *config.ServerConf
	Db   *gorm.DB
)
