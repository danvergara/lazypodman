package app

import (
	"github.com/danvergara/lazypodman/pkg/config"
	"github.com/danvergara/lazypodman/pkg/gui"
	"github.com/danvergara/lazypodman/pkg/logger"
	"github.com/sirupsen/logrus"
)

// App struct
type App struct {
	Config *config.Config
	Log    *logrus.Entry
	Gui    *gui.Gui
}

// NewApp boostrap a new application
func NewApp(config *config.Config, composeFile string) (*App, error) {
	app := &App{Config: config}
	app.Log = logger.NewLogger(config)

	projectName, err := config.ProjectName()
	if err != nil {
		return nil, err
	}

	app.Gui, err = gui.NewGui(composeFile, projectName)
	if err != nil {
		return nil, err
	}

	return app, nil

}

// Run the application
func (app *App) Run() error {
	err := app.Gui.RunWithSubprocesses()
	return err
}
