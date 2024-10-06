package app

import "gorm.io/gorm"

type App struct {
	Conf *Config
	DB   *gorm.DB
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

	return app, nil
}
