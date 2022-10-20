package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"health/config"
	"health/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config/conf.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf := config.ServerConf{}
	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	global.Conf = &conf
	if conf.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}
