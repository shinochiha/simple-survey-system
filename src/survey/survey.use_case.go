package survey

import (
	"net/http"
	"net/url"
	"time"

	"github.com/survey-app/survey/app"
)

// UseCase returns a UseCaseHandler for expected use case functional.
func UseCase(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	u := UseCaseHandler{
		Ctx:   &ctx,
		Query: url.Values{},
	}
	if len(query) > 0 {
		u.Query = query[0]
	}
	return u
}

// UseCaseHandler provides a convenient interface for Survey use case, use UseCase to access UseCaseHandler.
type UseCaseHandler struct {
	Survey

	// injectable dependencies
	Ctx   *app.Ctx   `json:"-" db:"-" gorm:"-"`
	Query url.Values `json:"-" db:"-" gorm:"-"`
}

// Async return UseCaseHandler with async process.
func (u UseCaseHandler) Async(ctx app.Ctx, query ...url.Values) UseCaseHandler {
	ctx.IsAsync = true
	return UseCase(ctx, query...)
}

// GetByID returns the Survey data for the specified ID.
func (u UseCaseHandler) GetByID(id string) (Survey, error) {
	res := Survey{}

	// check permission
	err := u.Ctx.ValidatePermission("surveys.detail")
	if err != nil {
		return res, err
	}

	// get from cache and return if exists
	cacheKey := u.EndPoint() + "." + id
	app.Cache().Get(cacheKey, &res)
	if res.ID.Valid {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.NewError(http.StatusInternalServerError, err.Error())
	}

	// get from db
	key := "id"
	if !app.Validator().IsValid(id, "uuid") {
		key = "code"
	}
	u.Query.Add(key, id)
	err = app.First(tx, &res, u.Query)
	if err != nil {
		return res, u.Ctx.NotFoundError(err, u.EndPoint(), key, id)
	}

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Get returns the list of Survey data.
func (u UseCaseHandler) Get() (app.ListModel, error) {
	res := app.ListModel{}

	// check permission
	err := u.Ctx.ValidatePermission("surveys.list")
	if err != nil {
		return res, err
	}
	// get from cache and return if exists
	cacheKey := u.EndPoint() + "?" + u.Query.Encode()
	err = app.Cache().Get(cacheKey, &res)
	if err == nil {
		return res, err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return res, app.NewError(http.StatusInternalServerError, err.Error())
	}

	// set pagination info
	res.Count,
		res.PageContext.Page,
		res.PageContext.PerPage,
		res.PageContext.PageCount,
		err = app.PaginationInfo(tx, &Survey{}, u.Query)
	if err != nil {
		return res, app.NewError(http.StatusInternalServerError, err.Error())
	}
	// return data count if $per_page set to 0
	if res.PageContext.PerPage == 0 {
		return res, err
	}

	// find data
	data, err := app.Find(tx, &Survey{}, u.Query)
	if err != nil {
		return res, app.NewError(http.StatusInternalServerError, err.Error())
	}
	res.SetData(data, u.Query)

	// save to cache and return if exists
	app.Cache().Set(cacheKey, res)
	return res, err
}

// Create creates a new data Survey with specified parameters.
func (u UseCaseHandler) Create(p *ParamCreate) error {

	// check permission
	err := u.Ctx.ValidatePermission("surveys.create")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}
	p.Ctx = u.Ctx
	// set default value for undefined field
	err = p.setDefaultValue(Survey{})
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// save data to db
	err = tx.Model(&p).Create(&p).Error
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	err = p.ProcessArray(Survey{})
	if err != nil {
		return err
	}
	// invalidate cache
	app.Cache().Invalidate(u.EndPoint())

	// Array Relation

	// save history (user activity), send webhook, etc
	// go u.Ctx.Hook("POST", "create", p.ID.String, p)
	return nil
}

// UpdateByID updates the Survey data for the specified ID with specified parameters.
func (u UseCaseHandler) UpdateByID(id string, p *ParamUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("surveys.edit")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PUT", p.Reason.String, old.ID.String, old)
	return nil
}

// PartiallyUpdateByID updates the Survey data for the specified ID with specified parameters.
func (u UseCaseHandler) PartiallyUpdateByID(id string, p *ParamPartiallyUpdate) error {

	// check permission
	err := u.Ctx.ValidatePermission("surveys.edit")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	// set default value for undefined field
	err = p.setDefaultValue(old)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id = ?", old.ID).Updates(p).Error
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("PATCH", p.Reason.String, old.ID.String, old)
	return nil
}

// DeleteByID deletes the Survey data for the specified ID.
func (u UseCaseHandler) DeleteByID(id string, p *ParamDelete) error {

	// check permission
	err := u.Ctx.ValidatePermission("surveys.delete")
	if err != nil {
		return err
	}

	// validate param
	err = u.Ctx.ValidateParam(p)
	if err != nil {
		return err
	}

	// get previous data
	old, err := u.GetByID(id)
	if err != nil {
		return err
	}

	// prepare db for current ctx
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// update data on the db
	err = tx.Model(&p).Where("id = ?", old.ID).Update("deleted_at", time.Now().UTC()).Error
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// invalidate cache
	app.Cache().Invalidate(u.EndPoint(), old.ID.String)

	// save history (user activity), send webhook, etc
	go u.Ctx.Hook("DELETE", p.Reason.String, old.ID.String, old)
	return nil
}

// setDefaultValue set default value of undefined field when create or update Survey data.
func (u *UseCaseHandler) setDefaultValue(old Survey) error {
	if !old.ID.Valid {
		u.ID = app.NewNullUUID()
	} else {
		u.ID = old.ID
	}

	if !u.IsActive.Valid {
		u.IsActive.Set(true)
	}

	return nil
}

func (u *UseCaseHandler) ProcessArray(old Survey) error {
	tx, err := u.Ctx.DB()
	if err != nil {
		return app.NewError(http.StatusInternalServerError, err.Error())
	}

	// questions one to many
	if len(old.Questions) > 0 {
		if u.Ctx.Action.Method == "PUT" {
			//delete old survey data
			err = tx.Delete(&Question{}, "survey_id = ?", old.ID.String).Error
			if err != nil {
				return err
			}
		}
	}

	if len(u.Questions) > 0 {
		questions := []Question{}
		for _, q := range u.Questions {
			question := Question{}
			question.ID = app.NewNullUUID()
			question.SurveyID.Set(u.ID.String)
			question.QuestionText.Set(q.QuestionText.String)

			if len(q.Choises) > 0 {
				choises := []Choise{}
				for _, c := range q.Choises {
					choise := Choise{}
					choise.ID = app.NewNullUUID()
					choise.QuestionID.Set(question.ID.String)
					choise.ChoiseText.Set(c.ChoiseText.String)

					choises = append(choises, choise)
				}
				err = tx.Create(&choises).Error
				if err != nil {
					return app.NewError(http.StatusInternalServerError, err.Error())
				}
			}
			questions = append(questions, question)
		}
		err = tx.Create(&questions).Error
		if err != nil {
			return app.NewError(http.StatusInternalServerError, err.Error())
		}
	}

	return nil
}
