package response

import "github.com/survey-app/survey/app"

// Response is the main model of Response data. It provides a convenient interface for app.ModelInterface
type Response struct {
	app.Model
	ID              app.NullUUID     `json:"id"               db:"m.id"               gorm:"column:id;primaryKey"`
	SurveyId        app.NullUUID     `json:"survey_id"        db:"m.survey_id"        gorm:"column:survey_id"`
	RespondentName  app.NullString   `json:"respondent_name"  db:"m.respondent_name"  gorm:"column:respondent_name"`
	RespondentEmail app.NullString   `json:"respondent_email" db:"m.respondent_email" gorm:"column:respondent_email"`
	IsActive        app.NullBool     `json:"is_active"        db:"m.is_active"        gorm:"column:is_active"`
	CreatedAt       app.NullDateTime `json:"created_at"       db:"m.created_at"       gorm:"column:created_at"`
	UpdatedAt       app.NullDateTime `json:"updated_at"       db:"m.updated_at"       gorm:"column:updated_at"`
	DeletedAt       app.NullDateTime `json:"deleted_at"       db:"m.deleted_at"       gorm:"column:deleted_at"`
}

// EndPoint returns the Response end point, it used for cache key, etc.
func (Response) EndPoint() string {
	return "responses"
}

// TableVersion returns the versions of the Response table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Response) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Response table in the database.
func (Response) TableName() string {
	return "responses"
}

// TableAliasName returns the table alias name of the Response table, used for querying.
func (Response) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Response data in the database, used for querying.
func (m *Response) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Response data in the database, used for querying.
func (m *Response) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Response data in the database, used for querying.
func (m *Response) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Response data in the database, used for querying.
func (m *Response) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Response schema, used for querying.
func (m *Response) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Response schema in the open api documentation.
func (Response) OpenAPISchemaName() string {
	return "Response"
}

// ParamCreate is the expected parameters for create a new Response data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Response data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Response data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Response data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
