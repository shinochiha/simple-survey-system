package question

import "github.com/survey-app/survey/app"

// Question is the main model of Question data. It provides a convenient interface for app.ModelInterface
type Question struct {
	app.Model
	ID           app.NullUUID     `json:"id"            db:"m.id"            gorm:"column:id;primaryKey"`
	SurveyId     app.NullUUID     `json:"survey_id"     db:"m.survey_id"     gorm:"column:survey_id"`
	QuestionText app.NullText     `json:"question_text" db:"m.question_text" gorm:"column:question_text"`
	IsActive     app.NullBool     `json:"is_active"     db:"m.is_active"     gorm:"column:is_active"`
	CreatedAt    app.NullDateTime `json:"created_at"    db:"m.created_at"    gorm:"column:created_at"`
	UpdatedAt    app.NullDateTime `json:"updated_at"    db:"m.updated_at"    gorm:"column:updated_at"`
	DeletedAt    app.NullDateTime `json:"deleted_at"    db:"m.deleted_at"    gorm:"column:deleted_at"`
}

// EndPoint returns the Question end point, it used for cache key, etc.
func (Question) EndPoint() string {
	return "questions"
}

// TableVersion returns the versions of the Question table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Question) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Question table in the database.
func (Question) TableName() string {
	return "questions"
}

// TableAliasName returns the table alias name of the Question table, used for querying.
func (Question) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Question data in the database, used for querying.
func (m *Question) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Question data in the database, used for querying.
func (m *Question) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Question data in the database, used for querying.
func (m *Question) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Question data in the database, used for querying.
func (m *Question) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Question schema, used for querying.
func (m *Question) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Question schema in the open api documentation.
func (Question) OpenAPISchemaName() string {
	return "Question"
}

// ParamCreate is the expected parameters for create a new Question data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Question data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Question data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Question data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
