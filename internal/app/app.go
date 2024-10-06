package app

type App struct {
	Conf *Config
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

	return app, nil
}
