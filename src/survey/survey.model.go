package survey

import "github.com/survey-app/survey/app"

// Survey is the main model of Survey data. It provides a convenient interface for app.ModelInterface
type Survey struct {
	app.Model
	ID          app.NullUUID     `json:"id"          db:"m.id"              gorm:"column:id;primaryKey"`
	Title       app.NullString   `json:"title"       db:"m.title"           gorm:"column:title"`
	Description app.NullString   `json:"description" db:"m.description"     gorm:"column:description"`
	IsActive    app.NullBool     `json:"is_active"   db:"m.is_active"       gorm:"column:is_active"`
	CreatedAt   app.NullDateTime `json:"created_at"  db:"m.created_at"      gorm:"column:created_at"`
	UpdatedAt   app.NullDateTime `json:"updated_at"  db:"m.updated_at"      gorm:"column:updated_at"`
	DeletedAt   app.NullDateTime `json:"deleted_at"  db:"m.deleted_at,hide" gorm:"column:deleted_at"`
	Questions   []Question       `json:"questions"   db:"survey.id={id}"    gorm:"-"`
}

// EndPoint returns the Survey end point, it used for cache key, etc.
func (Survey) EndPoint() string {
	return "surveys"
}

// TableVersion returns the versions of the Survey table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Survey) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the Survey table in the database.
func (Survey) TableName() string {
	return "surveys"
}

// TableAliasName returns the table alias name of the Survey table, used for querying.
func (Survey) TableAliasName() string {
	return "m"
}

// GetRelations returns the relations of the Survey data in the database, used for querying.
func (m *Survey) GetRelations() map[string]map[string]any {
	// m.AddRelation("left", "users", "cu", []map[string]any{{"column1": "cu.id", "column2": "m.created_by_user_id"}})
	// m.AddRelation("left", "users", "uu", []map[string]any{{"column1": "uu.id", "column2": "m.updated_by_user_id"}})
	return m.Relations
}

// GetFilters returns the filter of the Survey data in the database, used for querying.
func (m *Survey) GetFilters() []map[string]any {
	m.AddFilter(map[string]any{"column1": "m.deleted_at", "operator": "=", "value": nil})
	return m.Filters
}

// GetSorts returns the default sort of the Survey data in the database, used for querying.
func (m *Survey) GetSorts() []map[string]any {
	m.AddSort(map[string]any{"column": "m.updated_at", "direction": "desc"})
	return m.Sorts
}

// GetFields returns list of the field of the Survey data in the database, used for querying.
func (m *Survey) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the Survey schema, used for querying.
func (m *Survey) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// OpenAPISchemaName returns the name of the Survey schema in the open api documentation.
func (Survey) OpenAPISchemaName() string {
	return "Survey"
}

type Question struct {
	app.Model
	ID           app.NullUUID `json:"id"            db:"q.id"             gorm:"column:id"`
	SurveyID     app.NullUUID `json:"survey.id"     db:"q.survey_id,hide" gorm:"column:survey_id"`
	QuestionText app.NullText `json:"question_text" db:"q.question_text"  gorm:"column:question_text"`
	Choises      []Choise     `json:"choices"       db:"question.id={id}" gorm:"-"`
}

// TableVersion returns the versions of the questions table in the database.
// Change this value with date format YY.MM.DDHHii when any table structure changes.
func (Question) TableVersion() string {
	return "28.06.291152"
}

// TableName returns the name of the questions table in the database.
func (Question) TableName() string {
	return "questions"
}

// TableAliasName returns the table alias name of the questions table, used for querying.
func (Question) TableAliasName() string {
	return "q"
}

// GetRelations returns the relations of the questions data in the database, used for querying.
func (m *Question) GetRelations() map[string]map[string]any {
	return m.Relations
}

// GetFilters returns the filter of the questions data in the database, used for querying.
func (m *Question) GetFilters() []map[string]any {
	return m.Filters
}

// GetFields returns list of the field of the questions data in the database, used for querying.
func (m *Question) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

// GetSchema returns the questions schema, used for querying.
func (m *Question) GetSchema() map[string]any {
	return m.SetSchema(m)
}

type Choise struct {
	app.Model
	ID         app.NullUUID `json:"id"          db:"c.id"               gorm:"column:id"`
	QuestionID app.NullUUID `json:"question.id" db:"c.question_id,hide" gorm:"column:question_id"`
	ChoiseText app.NullText `json:"choise_text" db:"c.choise_text"      gorm:"column:choise_text"`
}

func (Choise) TableVersion() string {
	return "28.06.291152"
}

func (Choise) TableName() string {
	return "choices"
}

func (Choise) TableAliasName() string {
	return "c"
}

func (m *Choise) GetRelations() map[string]map[string]any {
	return m.Relations
}

func (m *Choise) GetFilters() []map[string]any {
	return m.Filters
}

func (m *Choise) GetFields() map[string]map[string]any {
	m.SetFields(m)
	return m.Fields
}

func (m *Choise) GetSchema() map[string]any {
	return m.SetSchema(m)
}

// ParamCreate is the expected parameters for create a new Survey data.
type ParamCreate struct {
	UseCaseHandler
}

// ParamUpdate is the expected parameters for update the Survey data.
type ParamUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamPartiallyUpdate is the expected parameters for partially update the Survey data.
type ParamPartiallyUpdate struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}

// ParamDelete is the expected parameters for delete the Survey data.
type ParamDelete struct {
	UseCaseHandler
	Reason app.NullString `json:"reason" gorm:"-" validate:"required"`
}
