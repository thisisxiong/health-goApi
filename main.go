package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"health/initialize"
	"health/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	initialize.InitConfig()

	initialize.Initdb()

	r := gin.Default()

	router.InitRouter(r)

	port := flag.Int("port", 8080, "端口号 默认8080")
	flag.Parse()
	go func() {
		r.Run(fmt.Sprintf(":%d", *port))
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("注销成功")

}
