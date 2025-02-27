// Code generated by ent, DO NOT EDIT.

package procedure

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the procedure type in the database.
	Label = "procedure"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldMappingID holds the string denoting the mapping_id field in the database.
	FieldMappingID = "mapping_id"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldProcedureType holds the string denoting the procedure_type field in the database.
	FieldProcedureType = "procedure_type"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldPurposeAndScope holds the string denoting the purpose_and_scope field in the database.
	FieldPurposeAndScope = "purpose_and_scope"
	// FieldBackground holds the string denoting the background field in the database.
	FieldBackground = "background"
	// FieldSatisfies holds the string denoting the satisfies field in the database.
	FieldSatisfies = "satisfies"
	// FieldDetails holds the string denoting the details field in the database.
	FieldDetails = "details"
	// EdgeControl holds the string denoting the control edge name in mutations.
	EdgeControl = "control"
	// EdgeInternalpolicy holds the string denoting the internalpolicy edge name in mutations.
	EdgeInternalpolicy = "internalpolicy"
	// EdgeNarratives holds the string denoting the narratives edge name in mutations.
	EdgeNarratives = "narratives"
	// EdgeRisks holds the string denoting the risks edge name in mutations.
	EdgeRisks = "risks"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// EdgePrograms holds the string denoting the programs edge name in mutations.
	EdgePrograms = "programs"
	// Table holds the table name of the procedure in the database.
	Table = "procedures"
	// ControlTable is the table that holds the control relation/edge. The primary key declared below.
	ControlTable = "control_procedures"
	// ControlInverseTable is the table name for the Control entity.
	// It exists in this package in order to avoid circular dependency with the "control" package.
	ControlInverseTable = "controls"
	// InternalpolicyTable is the table that holds the internalpolicy relation/edge. The primary key declared below.
	InternalpolicyTable = "internal_policy_procedures"
	// InternalpolicyInverseTable is the table name for the InternalPolicy entity.
	// It exists in this package in order to avoid circular dependency with the "internalpolicy" package.
	InternalpolicyInverseTable = "internal_policies"
	// NarrativesTable is the table that holds the narratives relation/edge. The primary key declared below.
	NarrativesTable = "procedure_narratives"
	// NarrativesInverseTable is the table name for the Narrative entity.
	// It exists in this package in order to avoid circular dependency with the "narrative" package.
	NarrativesInverseTable = "narratives"
	// RisksTable is the table that holds the risks relation/edge. The primary key declared below.
	RisksTable = "procedure_risks"
	// RisksInverseTable is the table name for the Risk entity.
	// It exists in this package in order to avoid circular dependency with the "risk" package.
	RisksInverseTable = "risks"
	// TasksTable is the table that holds the tasks relation/edge. The primary key declared below.
	TasksTable = "procedure_tasks"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "tasks"
	// ProgramsTable is the table that holds the programs relation/edge. The primary key declared below.
	ProgramsTable = "program_procedures"
	// ProgramsInverseTable is the table name for the Program entity.
	// It exists in this package in order to avoid circular dependency with the "program" package.
	ProgramsInverseTable = "programs"
)

// Columns holds all SQL columns for procedure fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldMappingID,
	FieldTags,
	FieldName,
	FieldDescription,
	FieldStatus,
	FieldProcedureType,
	FieldVersion,
	FieldPurposeAndScope,
	FieldBackground,
	FieldSatisfies,
	FieldDetails,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "procedures"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"control_objective_procedures",
	"standard_procedures",
}

var (
	// ControlPrimaryKey and ControlColumn2 are the table columns denoting the
	// primary key for the control relation (M2M).
	ControlPrimaryKey = []string{"control_id", "procedure_id"}
	// InternalpolicyPrimaryKey and InternalpolicyColumn2 are the table columns denoting the
	// primary key for the internalpolicy relation (M2M).
	InternalpolicyPrimaryKey = []string{"internal_policy_id", "procedure_id"}
	// NarrativesPrimaryKey and NarrativesColumn2 are the table columns denoting the
	// primary key for the narratives relation (M2M).
	NarrativesPrimaryKey = []string{"procedure_id", "narrative_id"}
	// RisksPrimaryKey and RisksColumn2 are the table columns denoting the
	// primary key for the risks relation (M2M).
	RisksPrimaryKey = []string{"procedure_id", "risk_id"}
	// TasksPrimaryKey and TasksColumn2 are the table columns denoting the
	// primary key for the tasks relation (M2M).
	TasksPrimaryKey = []string{"procedure_id", "task_id"}
	// ProgramsPrimaryKey and ProgramsColumn2 are the table columns denoting the
	// primary key for the programs relation (M2M).
	ProgramsPrimaryKey = []string{"program_id", "procedure_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/theopenlane/core/internal/ent/generated/runtime"
var (
	Hooks        [2]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultMappingID holds the default value on creation for the "mapping_id" field.
	DefaultMappingID func() string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the Procedure queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByMappingID orders the results by the mapping_id field.
func ByMappingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMappingID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByProcedureType orders the results by the procedure_type field.
func ByProcedureType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProcedureType, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByPurposeAndScope orders the results by the purpose_and_scope field.
func ByPurposeAndScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPurposeAndScope, opts...).ToFunc()
}

// ByBackground orders the results by the background field.
func ByBackground(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBackground, opts...).ToFunc()
}

// BySatisfies orders the results by the satisfies field.
func BySatisfies(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSatisfies, opts...).ToFunc()
}

// ByControlCount orders the results by control count.
func ByControlCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newControlStep(), opts...)
	}
}

// ByControl orders the results by control terms.
func ByControl(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newControlStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByInternalpolicyCount orders the results by internalpolicy count.
func ByInternalpolicyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInternalpolicyStep(), opts...)
	}
}

// ByInternalpolicy orders the results by internalpolicy terms.
func ByInternalpolicy(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInternalpolicyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNarrativesCount orders the results by narratives count.
func ByNarrativesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNarrativesStep(), opts...)
	}
}

// ByNarratives orders the results by narratives terms.
func ByNarratives(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNarrativesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRisksCount orders the results by risks count.
func ByRisksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRisksStep(), opts...)
	}
}

// ByRisks orders the results by risks terms.
func ByRisks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRisksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProgramsCount orders the results by programs count.
func ByProgramsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProgramsStep(), opts...)
	}
}

// ByPrograms orders the results by programs terms.
func ByPrograms(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProgramsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newControlStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ControlInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ControlTable, ControlPrimaryKey...),
	)
}
func newInternalpolicyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InternalpolicyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, InternalpolicyTable, InternalpolicyPrimaryKey...),
	)
}
func newNarrativesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NarrativesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NarrativesTable, NarrativesPrimaryKey...),
	)
}
func newRisksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RisksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, RisksTable, RisksPrimaryKey...),
	)
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TasksTable, TasksPrimaryKey...),
	)
}
func newProgramsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProgramsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProgramsTable, ProgramsPrimaryKey...),
	)
}
