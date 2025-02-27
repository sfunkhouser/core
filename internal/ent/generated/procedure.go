// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/procedure"
)

// Procedure is the model entity for the Procedure schema.
type Procedure struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// the name of the procedure
	Name string `json:"name,omitempty"`
	// description of the procedure
	Description string `json:"description,omitempty"`
	// status of the procedure
	Status string `json:"status,omitempty"`
	// type of the procedure
	ProcedureType string `json:"procedure_type,omitempty"`
	// version of the procedure
	Version string `json:"version,omitempty"`
	// purpose and scope
	PurposeAndScope string `json:"purpose_and_scope,omitempty"`
	// background of the procedure
	Background string `json:"background,omitempty"`
	// which controls are satisfied by the procedure
	Satisfies string `json:"satisfies,omitempty"`
	// json data for the procedure document
	Details map[string]interface{} `json:"details,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProcedureQuery when eager-loading is set.
	Edges                        ProcedureEdges `json:"edges"`
	control_objective_procedures *string
	standard_procedures          *string
	selectValues                 sql.SelectValues
}

// ProcedureEdges holds the relations/edges for other nodes in the graph.
type ProcedureEdges struct {
	// Control holds the value of the control edge.
	Control []*Control `json:"control,omitempty"`
	// Internalpolicy holds the value of the internalpolicy edge.
	Internalpolicy []*InternalPolicy `json:"internalpolicy,omitempty"`
	// Narratives holds the value of the narratives edge.
	Narratives []*Narrative `json:"narratives,omitempty"`
	// Risks holds the value of the risks edge.
	Risks []*Risk `json:"risks,omitempty"`
	// Tasks holds the value of the tasks edge.
	Tasks []*Task `json:"tasks,omitempty"`
	// Programs holds the value of the programs edge.
	Programs []*Program `json:"programs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedControl        map[string][]*Control
	namedInternalpolicy map[string][]*InternalPolicy
	namedNarratives     map[string][]*Narrative
	namedRisks          map[string][]*Risk
	namedTasks          map[string][]*Task
	namedPrograms       map[string][]*Program
}

// ControlOrErr returns the Control value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) ControlOrErr() ([]*Control, error) {
	if e.loadedTypes[0] {
		return e.Control, nil
	}
	return nil, &NotLoadedError{edge: "control"}
}

// InternalpolicyOrErr returns the Internalpolicy value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) InternalpolicyOrErr() ([]*InternalPolicy, error) {
	if e.loadedTypes[1] {
		return e.Internalpolicy, nil
	}
	return nil, &NotLoadedError{edge: "internalpolicy"}
}

// NarrativesOrErr returns the Narratives value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) NarrativesOrErr() ([]*Narrative, error) {
	if e.loadedTypes[2] {
		return e.Narratives, nil
	}
	return nil, &NotLoadedError{edge: "narratives"}
}

// RisksOrErr returns the Risks value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) RisksOrErr() ([]*Risk, error) {
	if e.loadedTypes[3] {
		return e.Risks, nil
	}
	return nil, &NotLoadedError{edge: "risks"}
}

// TasksOrErr returns the Tasks value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) TasksOrErr() ([]*Task, error) {
	if e.loadedTypes[4] {
		return e.Tasks, nil
	}
	return nil, &NotLoadedError{edge: "tasks"}
}

// ProgramsOrErr returns the Programs value or an error if the edge
// was not loaded in eager-loading.
func (e ProcedureEdges) ProgramsOrErr() ([]*Program, error) {
	if e.loadedTypes[5] {
		return e.Programs, nil
	}
	return nil, &NotLoadedError{edge: "programs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Procedure) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case procedure.FieldTags, procedure.FieldDetails:
			values[i] = new([]byte)
		case procedure.FieldID, procedure.FieldCreatedBy, procedure.FieldUpdatedBy, procedure.FieldDeletedBy, procedure.FieldMappingID, procedure.FieldName, procedure.FieldDescription, procedure.FieldStatus, procedure.FieldProcedureType, procedure.FieldVersion, procedure.FieldPurposeAndScope, procedure.FieldBackground, procedure.FieldSatisfies:
			values[i] = new(sql.NullString)
		case procedure.FieldCreatedAt, procedure.FieldUpdatedAt, procedure.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case procedure.ForeignKeys[0]: // control_objective_procedures
			values[i] = new(sql.NullString)
		case procedure.ForeignKeys[1]: // standard_procedures
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Procedure fields.
func (pr *Procedure) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case procedure.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case procedure.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case procedure.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = value.Time
			}
		case procedure.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pr.CreatedBy = value.String
			}
		case procedure.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pr.UpdatedBy = value.String
			}
		case procedure.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = value.Time
			}
		case procedure.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				pr.DeletedBy = value.String
			}
		case procedure.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				pr.MappingID = value.String
			}
		case procedure.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case procedure.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case procedure.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case procedure.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pr.Status = value.String
			}
		case procedure.FieldProcedureType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field procedure_type", values[i])
			} else if value.Valid {
				pr.ProcedureType = value.String
			}
		case procedure.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				pr.Version = value.String
			}
		case procedure.FieldPurposeAndScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purpose_and_scope", values[i])
			} else if value.Valid {
				pr.PurposeAndScope = value.String
			}
		case procedure.FieldBackground:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field background", values[i])
			} else if value.Valid {
				pr.Background = value.String
			}
		case procedure.FieldSatisfies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field satisfies", values[i])
			} else if value.Valid {
				pr.Satisfies = value.String
			}
		case procedure.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pr.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		case procedure.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field control_objective_procedures", values[i])
			} else if value.Valid {
				pr.control_objective_procedures = new(string)
				*pr.control_objective_procedures = value.String
			}
		case procedure.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field standard_procedures", values[i])
			} else if value.Valid {
				pr.standard_procedures = new(string)
				*pr.standard_procedures = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Procedure.
// This includes values selected through modifiers, order, etc.
func (pr *Procedure) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryControl queries the "control" edge of the Procedure entity.
func (pr *Procedure) QueryControl() *ControlQuery {
	return NewProcedureClient(pr.config).QueryControl(pr)
}

// QueryInternalpolicy queries the "internalpolicy" edge of the Procedure entity.
func (pr *Procedure) QueryInternalpolicy() *InternalPolicyQuery {
	return NewProcedureClient(pr.config).QueryInternalpolicy(pr)
}

// QueryNarratives queries the "narratives" edge of the Procedure entity.
func (pr *Procedure) QueryNarratives() *NarrativeQuery {
	return NewProcedureClient(pr.config).QueryNarratives(pr)
}

// QueryRisks queries the "risks" edge of the Procedure entity.
func (pr *Procedure) QueryRisks() *RiskQuery {
	return NewProcedureClient(pr.config).QueryRisks(pr)
}

// QueryTasks queries the "tasks" edge of the Procedure entity.
func (pr *Procedure) QueryTasks() *TaskQuery {
	return NewProcedureClient(pr.config).QueryTasks(pr)
}

// QueryPrograms queries the "programs" edge of the Procedure entity.
func (pr *Procedure) QueryPrograms() *ProgramQuery {
	return NewProcedureClient(pr.config).QueryPrograms(pr)
}

// Update returns a builder for updating this Procedure.
// Note that you need to call Procedure.Unwrap() before calling this method if this Procedure
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Procedure) Update() *ProcedureUpdateOne {
	return NewProcedureClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Procedure entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Procedure) Unwrap() *Procedure {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("generated: Procedure is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Procedure) String() string {
	var builder strings.Builder
	builder.WriteString("Procedure(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(pr.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(pr.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pr.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(pr.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(pr.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", pr.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(pr.Status)
	builder.WriteString(", ")
	builder.WriteString("procedure_type=")
	builder.WriteString(pr.ProcedureType)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(pr.Version)
	builder.WriteString(", ")
	builder.WriteString("purpose_and_scope=")
	builder.WriteString(pr.PurposeAndScope)
	builder.WriteString(", ")
	builder.WriteString("background=")
	builder.WriteString(pr.Background)
	builder.WriteString(", ")
	builder.WriteString("satisfies=")
	builder.WriteString(pr.Satisfies)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", pr.Details))
	builder.WriteByte(')')
	return builder.String()
}

// NamedControl returns the Control named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Procedure) NamedControl(name string) ([]*Control, error) {
	if pr.Edges.namedControl == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedControl[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Procedure) appendNamedControl(name string, edges ...*Control) {
	if pr.Edges.namedControl == nil {
		pr.Edges.namedControl = make(map[string][]*Control)
	}
	if len(edges) == 0 {
		pr.Edges.namedControl[name] = []*Control{}
	} else {
		pr.Edges.namedControl[name] = append(pr.Edges.namedControl[name], edges...)
	}
}

// NamedInternalpolicy returns the Internalpolicy named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Procedure) NamedInternalpolicy(name string) ([]*InternalPolicy, error) {
	if pr.Edges.namedInternalpolicy == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedInternalpolicy[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Procedure) appendNamedInternalpolicy(name string, edges ...*InternalPolicy) {
	if pr.Edges.namedInternalpolicy == nil {
		pr.Edges.namedInternalpolicy = make(map[string][]*InternalPolicy)
	}
	if len(edges) == 0 {
		pr.Edges.namedInternalpolicy[name] = []*InternalPolicy{}
	} else {
		pr.Edges.namedInternalpolicy[name] = append(pr.Edges.namedInternalpolicy[name], edges...)
	}
}

// NamedNarratives returns the Narratives named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Procedure) NamedNarratives(name string) ([]*Narrative, error) {
	if pr.Edges.namedNarratives == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedNarratives[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Procedure) appendNamedNarratives(name string, edges ...*Narrative) {
	if pr.Edges.namedNarratives == nil {
		pr.Edges.namedNarratives = make(map[string][]*Narrative)
	}
	if len(edges) == 0 {
		pr.Edges.namedNarratives[name] = []*Narrative{}
	} else {
		pr.Edges.namedNarratives[name] = append(pr.Edges.namedNarratives[name], edges...)
	}
}

// NamedRisks returns the Risks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Procedure) NamedRisks(name string) ([]*Risk, error) {
	if pr.Edges.namedRisks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedRisks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Procedure) appendNamedRisks(name string, edges ...*Risk) {
	if pr.Edges.namedRisks == nil {
		pr.Edges.namedRisks = make(map[string][]*Risk)
	}
	if len(edges) == 0 {
		pr.Edges.namedRisks[name] = []*Risk{}
	} else {
		pr.Edges.namedRisks[name] = append(pr.Edges.namedRisks[name], edges...)
	}
}

// NamedTasks returns the Tasks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Procedure) NamedTasks(name string) ([]*Task, error) {
	if pr.Edges.namedTasks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedTasks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Procedure) appendNamedTasks(name string, edges ...*Task) {
	if pr.Edges.namedTasks == nil {
		pr.Edges.namedTasks = make(map[string][]*Task)
	}
	if len(edges) == 0 {
		pr.Edges.namedTasks[name] = []*Task{}
	} else {
		pr.Edges.namedTasks[name] = append(pr.Edges.namedTasks[name], edges...)
	}
}

// NamedPrograms returns the Programs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Procedure) NamedPrograms(name string) ([]*Program, error) {
	if pr.Edges.namedPrograms == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedPrograms[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Procedure) appendNamedPrograms(name string, edges ...*Program) {
	if pr.Edges.namedPrograms == nil {
		pr.Edges.namedPrograms = make(map[string][]*Program)
	}
	if len(edges) == 0 {
		pr.Edges.namedPrograms[name] = []*Program{}
	} else {
		pr.Edges.namedPrograms[name] = append(pr.Edges.namedPrograms[name], edges...)
	}
}

// Procedures is a parsable slice of Procedure.
type Procedures []*Procedure
