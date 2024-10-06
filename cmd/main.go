package main

import (
	"flag"
	"log/slog"
	a "mine/internal/app"
	"mine/internal/web"
)

func main() {
	configFilePath := flag.String("config", "etc/config.yaml", "config file path")
	// 获取参数
	flag.Parse()
	if app, err := a.NewApp(*configFilePath); nil != err {
		slog.Error("应用启动失败：" + err.Error())
	} else {
		slog.Info("hello world:" + app.Conf.APP.Mode)
		web.InitRouter(app)
		if err := app.LinkStart(); nil != err {
			slog.Error("应用运行失败：" + err.Error())
		}
	}
}
