// Code generated by ent, DO NOT EDIT.

package taskhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/pkg/enums"
	"github.com/theopenlane/entx/history"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldRef, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// MappingID applies equality check predicate on the "mapping_id" field. It's identical to MappingIDEQ.
func MappingID(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldMappingID, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDeletedBy, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDescription, v))
}

// Due applies equality check predicate on the "due" field. It's identical to DueEQ.
func Due(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDue, v))
}

// Completed applies equality check predicate on the "completed" field. It's identical to CompletedEQ.
func Completed(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldCompleted, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldRef, v))
}

// RefContains applies the Contains predicate on the "ref" field.
func RefContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldRef, v))
}

// RefHasPrefix applies the HasPrefix predicate on the "ref" field.
func RefHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldRef, v))
}

// RefHasSuffix applies the HasSuffix predicate on the "ref" field.
func RefHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldRef))
}

// RefEqualFold applies the EqualFold predicate on the "ref" field.
func RefEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldRef, v))
}

// RefContainsFold applies the ContainsFold predicate on the "ref" field.
func RefContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldRef, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v history.OpType) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v history.OpType) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...history.OpType) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...history.OpType) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// MappingIDEQ applies the EQ predicate on the "mapping_id" field.
func MappingIDEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldMappingID, v))
}

// MappingIDNEQ applies the NEQ predicate on the "mapping_id" field.
func MappingIDNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldMappingID, v))
}

// MappingIDIn applies the In predicate on the "mapping_id" field.
func MappingIDIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldMappingID, vs...))
}

// MappingIDNotIn applies the NotIn predicate on the "mapping_id" field.
func MappingIDNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldMappingID, vs...))
}

// MappingIDGT applies the GT predicate on the "mapping_id" field.
func MappingIDGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldMappingID, v))
}

// MappingIDGTE applies the GTE predicate on the "mapping_id" field.
func MappingIDGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldMappingID, v))
}

// MappingIDLT applies the LT predicate on the "mapping_id" field.
func MappingIDLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldMappingID, v))
}

// MappingIDLTE applies the LTE predicate on the "mapping_id" field.
func MappingIDLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldMappingID, v))
}

// MappingIDContains applies the Contains predicate on the "mapping_id" field.
func MappingIDContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldMappingID, v))
}

// MappingIDHasPrefix applies the HasPrefix predicate on the "mapping_id" field.
func MappingIDHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldMappingID, v))
}

// MappingIDHasSuffix applies the HasSuffix predicate on the "mapping_id" field.
func MappingIDHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldMappingID, v))
}

// MappingIDEqualFold applies the EqualFold predicate on the "mapping_id" field.
func MappingIDEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldMappingID, v))
}

// MappingIDContainsFold applies the ContainsFold predicate on the "mapping_id" field.
func MappingIDContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldMappingID, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldDeletedBy, v))
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldTags))
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldTags))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldContainsFold(FieldDescription, v))
}

// DetailsIsNil applies the IsNil predicate on the "details" field.
func DetailsIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldDetails))
}

// DetailsNotNil applies the NotNil predicate on the "details" field.
func DetailsNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldDetails))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.TaskStatus) predicate.TaskHistory {
	vc := v
	return predicate.TaskHistory(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.TaskStatus) predicate.TaskHistory {
	vc := v
	return predicate.TaskHistory(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.TaskStatus) predicate.TaskHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskHistory(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.TaskStatus) predicate.TaskHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TaskHistory(sql.FieldNotIn(FieldStatus, v...))
}

// DueEQ applies the EQ predicate on the "due" field.
func DueEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldDue, v))
}

// DueNEQ applies the NEQ predicate on the "due" field.
func DueNEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldDue, v))
}

// DueIn applies the In predicate on the "due" field.
func DueIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldDue, vs...))
}

// DueNotIn applies the NotIn predicate on the "due" field.
func DueNotIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldDue, vs...))
}

// DueGT applies the GT predicate on the "due" field.
func DueGT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldDue, v))
}

// DueGTE applies the GTE predicate on the "due" field.
func DueGTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldDue, v))
}

// DueLT applies the LT predicate on the "due" field.
func DueLT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldDue, v))
}

// DueLTE applies the LTE predicate on the "due" field.
func DueLTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldDue, v))
}

// DueIsNil applies the IsNil predicate on the "due" field.
func DueIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldDue))
}

// DueNotNil applies the NotNil predicate on the "due" field.
func DueNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldDue))
}

// CompletedEQ applies the EQ predicate on the "completed" field.
func CompletedEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldEQ(FieldCompleted, v))
}

// CompletedNEQ applies the NEQ predicate on the "completed" field.
func CompletedNEQ(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNEQ(FieldCompleted, v))
}

// CompletedIn applies the In predicate on the "completed" field.
func CompletedIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIn(FieldCompleted, vs...))
}

// CompletedNotIn applies the NotIn predicate on the "completed" field.
func CompletedNotIn(vs ...time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotIn(FieldCompleted, vs...))
}

// CompletedGT applies the GT predicate on the "completed" field.
func CompletedGT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGT(FieldCompleted, v))
}

// CompletedGTE applies the GTE predicate on the "completed" field.
func CompletedGTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldGTE(FieldCompleted, v))
}

// CompletedLT applies the LT predicate on the "completed" field.
func CompletedLT(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLT(FieldCompleted, v))
}

// CompletedLTE applies the LTE predicate on the "completed" field.
func CompletedLTE(v time.Time) predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldLTE(FieldCompleted, v))
}

// CompletedIsNil applies the IsNil predicate on the "completed" field.
func CompletedIsNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldIsNull(FieldCompleted))
}

// CompletedNotNil applies the NotNil predicate on the "completed" field.
func CompletedNotNil() predicate.TaskHistory {
	return predicate.TaskHistory(sql.FieldNotNull(FieldCompleted))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TaskHistory) predicate.TaskHistory {
	return predicate.TaskHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TaskHistory) predicate.TaskHistory {
	return predicate.TaskHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TaskHistory) predicate.TaskHistory {
	return predicate.TaskHistory(sql.NotPredicates(p))
}
