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

func Router() *routerUtil {
	if router == nil {
		router = &routerUtil{}
		router.Configure()
		router.isConfigured = true
	}
	return router
}

var router *routerUtil

type routerUtil struct {
	isConfigured bool
}

func (r *routerUtil) Configure() {
	app.Server().AddRoute("/api/version", "GET", app.VersionHandler, nil)

	app.Server().AddRoute("/api/v1/surveys", "POST", survey.REST().Create, survey.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/surveys", "GET", survey.REST().Get, survey.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/surveys/{id}", "GET", survey.REST().GetByID, survey.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/surveys/{id}", "PUT", survey.REST().UpdateByID, survey.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/surveys/{id}", "PATCH", survey.REST().PartiallyUpdateByID, survey.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/surveys/{id}", "DELETE", survey.REST().DeleteByID, survey.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/questions", "POST", question.REST().Create, question.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/questions", "GET", question.REST().Get, question.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/questions/{id}", "GET", question.REST().GetByID, question.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/questions/{id}", "PUT", question.REST().UpdateByID, question.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/questions/{id}", "PATCH", question.REST().PartiallyUpdateByID, question.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/questions/{id}", "DELETE", question.REST().DeleteByID, question.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/choices", "POST", choice.REST().Create, choice.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/choices", "GET", choice.REST().Get, choice.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/choices/{id}", "GET", choice.REST().GetByID, choice.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/choices/{id}", "PUT", choice.REST().UpdateByID, choice.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/choices/{id}", "PATCH", choice.REST().PartiallyUpdateByID, choice.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/choices/{id}", "DELETE", choice.REST().DeleteByID, choice.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/responses", "POST", response.REST().Create, response.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/responses", "GET", response.REST().Get, response.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/responses/{id}", "GET", response.REST().GetByID, response.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/responses/{id}", "PUT", response.REST().UpdateByID, response.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/responses/{id}", "PATCH", response.REST().PartiallyUpdateByID, response.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/responses/{id}", "DELETE", response.REST().DeleteByID, response.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/v1/answers", "POST", answer.REST().Create, answer.OpenAPI().Create())
	app.Server().AddRoute("/api/v1/answers", "GET", answer.REST().Get, answer.OpenAPI().Get())
	app.Server().AddRoute("/api/v1/answers/{id}", "GET", answer.REST().GetByID, answer.OpenAPI().GetByID())
	app.Server().AddRoute("/api/v1/answers/{id}", "PUT", answer.REST().UpdateByID, answer.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/v1/answers/{id}", "PATCH", answer.REST().PartiallyUpdateByID, answer.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/v1/answers/{id}", "DELETE", answer.REST().DeleteByID, answer.OpenAPI().DeleteByID())

	// AddRoute : DONT REMOVE THIS COMMENT
}
