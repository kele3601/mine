package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Conf *Config
	DB   *gorm.DB
	Web  *gin.Engine
}

var app *App

func NewApp(configFilePath string) (*App, error) {
	if nil != app {
		return app, nil
	}
	app = &App{}
	if conf, err := NewConfig(configFilePath); nil != err {
		return nil, err
	} else {
		app.Conf = conf
	}

	if db, err := NewDB(*app.Conf); nil != err {
		return nil, err
	} else {
		app.DB = db
	}
	app.Web = gin.Default()

	return app, nil
}

func (a *App) LinkStart() error {
	return a.Web.Run(a.Conf.WEB.Port)
}
