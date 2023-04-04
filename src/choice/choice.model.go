package choice

import "github.com/survey-app/survey/app"

// Choice is the main model of Choice data. It provides a convenient interface for app.ModelInterface
type Choice struct {
	app.Model
	ID         app.NullUUID     `json:"id"          db:"m.id"          gorm:"column:id;primaryKey"`
	QuestionId app.NullUUID     `json:"question_id" db:"m.question_id" gorm:"column:question_id"`
	ChoiseText app.NullText     `json:"choise_text" db:"m.choise_text" gorm:"column:choise_text"`
	CreatedAt  app.NullDateTime `json:"created_at"  db:"m.created_at"  gorm:"column:created_at"`
	UpdatedAt  app.NullDateTime `json:"updated_at"  db:"m.updated_at"  gorm:"column:updated_at"`
	DeletedAt  app.NullDateTime `json:"deleted_at"  db:"m.deleted_at"  gorm:"column:deleted_at"`
}

// EndPoint returns the Choice end point, it used for cache key, etc.
func (Choice) EndPoint() string {
	return "choices"
}

// TableVersion returns the versions of the Choice table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Choice) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Choice table in the database.
func (Choice) TableName() string {
	return "choices"
}

// TableAliasName returns the table alias name of the Choice table, used for querying.
func (Choice) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Choice data in the database, used for querying.
func (m *Choice) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Choice data in the database, used for querying.
func (m *Choice) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Choice data in the database, used for querying.
func (m *Choice) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Choice data in the database, used for querying.
func (m *Choice) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Choice schema, used for querying.
func (m *Choice) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Choice schema in the open api documentation.
func (Choice) OpenAPISchemaName() string {
	return "Choice"
}

// ParamCreate is the expected parameters for create a new Choice data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Choice data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Choice data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Choice data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
