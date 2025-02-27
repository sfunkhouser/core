// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/internalpolicyhistory"
	"github.com/theopenlane/entx/history"
)

// InternalPolicyHistory is the model entity for the InternalPolicyHistory schema.
type InternalPolicyHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref string `json:"ref,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation history.OpType `json:"operation,omitempty"`
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
	// the name of the policy
	Name string `json:"name,omitempty"`
	// description of the policy
	Description string `json:"description,omitempty"`
	// status of the policy
	Status string `json:"status,omitempty"`
	// type of the policy
	PolicyType string `json:"policy_type,omitempty"`
	// version of the policy
	Version string `json:"version,omitempty"`
	// purpose and scope
	PurposeAndScope string `json:"purpose_and_scope,omitempty"`
	// background of the policy
	Background string `json:"background,omitempty"`
	// json data for the policy document
	Details      map[string]interface{} `json:"details,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*InternalPolicyHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case internalpolicyhistory.FieldTags, internalpolicyhistory.FieldDetails:
			values[i] = new([]byte)
		case internalpolicyhistory.FieldOperation:
			values[i] = new(history.OpType)
		case internalpolicyhistory.FieldID, internalpolicyhistory.FieldRef, internalpolicyhistory.FieldCreatedBy, internalpolicyhistory.FieldUpdatedBy, internalpolicyhistory.FieldDeletedBy, internalpolicyhistory.FieldMappingID, internalpolicyhistory.FieldName, internalpolicyhistory.FieldDescription, internalpolicyhistory.FieldStatus, internalpolicyhistory.FieldPolicyType, internalpolicyhistory.FieldVersion, internalpolicyhistory.FieldPurposeAndScope, internalpolicyhistory.FieldBackground:
			values[i] = new(sql.NullString)
		case internalpolicyhistory.FieldHistoryTime, internalpolicyhistory.FieldCreatedAt, internalpolicyhistory.FieldUpdatedAt, internalpolicyhistory.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the InternalPolicyHistory fields.
func (iph *InternalPolicyHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case internalpolicyhistory.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				iph.ID = value.String
			}
		case internalpolicyhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				iph.HistoryTime = value.Time
			}
		case internalpolicyhistory.FieldRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value.Valid {
				iph.Ref = value.String
			}
		case internalpolicyhistory.FieldOperation:
			if value, ok := values[i].(*history.OpType); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value != nil {
				iph.Operation = *value
			}
		case internalpolicyhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				iph.CreatedAt = value.Time
			}
		case internalpolicyhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				iph.UpdatedAt = value.Time
			}
		case internalpolicyhistory.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				iph.CreatedBy = value.String
			}
		case internalpolicyhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				iph.UpdatedBy = value.String
			}
		case internalpolicyhistory.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				iph.DeletedAt = value.Time
			}
		case internalpolicyhistory.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				iph.DeletedBy = value.String
			}
		case internalpolicyhistory.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				iph.MappingID = value.String
			}
		case internalpolicyhistory.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &iph.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case internalpolicyhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				iph.Name = value.String
			}
		case internalpolicyhistory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				iph.Description = value.String
			}
		case internalpolicyhistory.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				iph.Status = value.String
			}
		case internalpolicyhistory.FieldPolicyType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field policy_type", values[i])
			} else if value.Valid {
				iph.PolicyType = value.String
			}
		case internalpolicyhistory.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				iph.Version = value.String
			}
		case internalpolicyhistory.FieldPurposeAndScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purpose_and_scope", values[i])
			} else if value.Valid {
				iph.PurposeAndScope = value.String
			}
		case internalpolicyhistory.FieldBackground:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field background", values[i])
			} else if value.Valid {
				iph.Background = value.String
			}
		case internalpolicyhistory.FieldDetails:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field details", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &iph.Details); err != nil {
					return fmt.Errorf("unmarshal field details: %w", err)
				}
			}
		default:
			iph.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the InternalPolicyHistory.
// This includes values selected through modifiers, order, etc.
func (iph *InternalPolicyHistory) Value(name string) (ent.Value, error) {
	return iph.selectValues.Get(name)
}

// Update returns a builder for updating this InternalPolicyHistory.
// Note that you need to call InternalPolicyHistory.Unwrap() before calling this method if this InternalPolicyHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (iph *InternalPolicyHistory) Update() *InternalPolicyHistoryUpdateOne {
	return NewInternalPolicyHistoryClient(iph.config).UpdateOne(iph)
}

// Unwrap unwraps the InternalPolicyHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (iph *InternalPolicyHistory) Unwrap() *InternalPolicyHistory {
	_tx, ok := iph.config.driver.(*txDriver)
	if !ok {
		panic("generated: InternalPolicyHistory is not a transactional entity")
	}
	iph.config.driver = _tx.drv
	return iph
}

// String implements the fmt.Stringer.
func (iph *InternalPolicyHistory) String() string {
	var builder strings.Builder
	builder.WriteString("InternalPolicyHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", iph.ID))
	builder.WriteString("history_time=")
	builder.WriteString(iph.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(iph.Ref)
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", iph.Operation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(iph.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(iph.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(iph.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(iph.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(iph.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(iph.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(iph.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", iph.Tags))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(iph.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(iph.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(iph.Status)
	builder.WriteString(", ")
	builder.WriteString("policy_type=")
	builder.WriteString(iph.PolicyType)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(iph.Version)
	builder.WriteString(", ")
	builder.WriteString("purpose_and_scope=")
	builder.WriteString(iph.PurposeAndScope)
	builder.WriteString(", ")
	builder.WriteString("background=")
	builder.WriteString(iph.Background)
	builder.WriteString(", ")
	builder.WriteString("details=")
	builder.WriteString(fmt.Sprintf("%v", iph.Details))
	builder.WriteByte(')')
	return builder.String()
}

// InternalPolicyHistories is a parsable slice of InternalPolicyHistory.
type InternalPolicyHistories []*InternalPolicyHistory
