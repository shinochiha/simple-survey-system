package answer

import "github.com/survey-app/survey/app"

// Answer is the main model of Answer data. It provides a convenient interface for app.ModelInterface
type Answer struct {
	app.Model
	ID         app.NullUUID     `json:"id"          db:"m.id"          gorm:"column:id;primaryKey"`
	ResponseId app.NullUUID     `json:"response_id" db:"m.response_id" gorm:"column:response_id"`
	QuestionId app.NullUUID     `json:"question_id" db:"m.question_id" gorm:"column:question_id"`
	ChoiseId   app.NullUUID     `json:"choise_id"   db:"m.choise_id"   gorm:"column:choise_id"`
	AnswerText app.NullText     `json:"answer_text" db:"m.answer_text" gorm:"column:answer_text"`
	CreatedAt  app.NullDateTime `json:"created_at"  db:"m.created_at"  gorm:"column:created_at"`
	UpdatedAt  app.NullDateTime `json:"updated_at"  db:"m.updated_at"  gorm:"column:updated_at"`
	DeletedAt  app.NullDateTime `json:"deleted_at"  db:"m.deleted_at"  gorm:"column:deleted_at"`
}

// EndPoint returns the Answer end point, it used for cache key, etc.
func (Answer) EndPoint() string {
	return "answers"
}

// TableVersion returns the versions of the Answer table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Answer) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Answer table in the database.
func (Answer) TableName() string {
	return "answers"
}

// TableAliasName returns the table alias name of the Answer table, used for querying.
func (Answer) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Answer data in the database, used for querying.
func (m *Answer) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Answer data in the database, used for querying.
func (m *Answer) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Answer data in the database, used for querying.
func (m *Answer) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Answer data in the database, used for querying.
func (m *Answer) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Answer schema, used for querying.
func (m *Answer) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Answer schema in the open api documentation.
func (Answer) OpenAPISchemaName() string {
	return "Answer"
}

// ParamCreate is the expected parameters for create a new Answer data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Answer data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Answer data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Answer data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
