# go 的学习demo

1.启动方式

### shell脚本启动
```
./start.sh start [port] //port默认值8080
```

### Dockfile 启动
~~Dockfile 二阶段构建 还没解决如何在msyql启动以后再执行启动文件~~

2.脚本说明

> [start.sh](start.sh)

    start 守护进程方式启动程序
    
    stop 退出守护进程及程序
    
    build 构建程序

> [app-daemon.sh](app-daemon.sh)

    守护进程脚本
