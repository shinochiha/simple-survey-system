package src

import (
	"github.com/survey-app/survey/app"
	"github.com/survey-app/survey/src/answer"
	"github.com/survey-app/survey/src/choice"
	"github.com/survey-app/survey/src/question"
	"github.com/survey-app/survey/src/response"
	"github.com/survey-app/survey/src/survey"
	// import : DONT REMOVE THIS COMMENT
)

func Migrator() *migratorUtil {
	if migrator == nil {
		migrator = &migratorUtil{}
		migrator.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			migrator.Run()
		}
		migrator.isConfigured = true
	}
	return migrator
}

var migrator *migratorUtil

type migratorUtil struct {
	isConfigured bool
}

func (*migratorUtil) Configure() {
	app.DB().RegisterTable("main", survey.Survey{})
	app.DB().RegisterTable("main", question.Question{})
	app.DB().RegisterTable("main", choice.Choice{})
	app.DB().RegisterTable("main", response.Response{})
	app.DB().RegisterTable("main", answer.Answer{})
	// RegisterTable : DONT REMOVE THIS COMMENT
}

func (*migratorUtil) Run() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	} else {
		err = app.DB().MigrateTable(tx, "main", app.Setting{})
	}
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
