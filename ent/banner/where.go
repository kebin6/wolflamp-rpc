// Code generated by ent, DO NOT EDIT.

package banner

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kebin6/wolflamp-rpc/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldStatus, v))
}

// Endpoint applies equality check predicate on the "endpoint" field. It's identical to EndpointEQ.
func Endpoint(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldEndpoint, v))
}

// Module applies equality check predicate on the "module" field. It's identical to ModuleEQ.
func Module(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldModule, v))
}

// FileType applies equality check predicate on the "file_type" field. It's identical to FileTypeEQ.
func FileType(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldFileType, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldPath, v))
}

// JumpURL applies equality check predicate on the "jump_url" field. It's identical to JumpURLEQ.
func JumpURL(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldJumpURL, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldStatus))
}

// EndpointEQ applies the EQ predicate on the "endpoint" field.
func EndpointEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldEndpoint, v))
}

// EndpointNEQ applies the NEQ predicate on the "endpoint" field.
func EndpointNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldEndpoint, v))
}

// EndpointIn applies the In predicate on the "endpoint" field.
func EndpointIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldEndpoint, vs...))
}

// EndpointNotIn applies the NotIn predicate on the "endpoint" field.
func EndpointNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldEndpoint, vs...))
}

// EndpointGT applies the GT predicate on the "endpoint" field.
func EndpointGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldEndpoint, v))
}

// EndpointGTE applies the GTE predicate on the "endpoint" field.
func EndpointGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldEndpoint, v))
}

// EndpointLT applies the LT predicate on the "endpoint" field.
func EndpointLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldEndpoint, v))
}

// EndpointLTE applies the LTE predicate on the "endpoint" field.
func EndpointLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldEndpoint, v))
}

// EndpointContains applies the Contains predicate on the "endpoint" field.
func EndpointContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldEndpoint, v))
}

// EndpointHasPrefix applies the HasPrefix predicate on the "endpoint" field.
func EndpointHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldEndpoint, v))
}

// EndpointHasSuffix applies the HasSuffix predicate on the "endpoint" field.
func EndpointHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldEndpoint, v))
}

// EndpointEqualFold applies the EqualFold predicate on the "endpoint" field.
func EndpointEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldEndpoint, v))
}

// EndpointContainsFold applies the ContainsFold predicate on the "endpoint" field.
func EndpointContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldEndpoint, v))
}

// ModuleEQ applies the EQ predicate on the "module" field.
func ModuleEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldModule, v))
}

// ModuleNEQ applies the NEQ predicate on the "module" field.
func ModuleNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldModule, v))
}

// ModuleIn applies the In predicate on the "module" field.
func ModuleIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldModule, vs...))
}

// ModuleNotIn applies the NotIn predicate on the "module" field.
func ModuleNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldModule, vs...))
}

// ModuleGT applies the GT predicate on the "module" field.
func ModuleGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldModule, v))
}

// ModuleGTE applies the GTE predicate on the "module" field.
func ModuleGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldModule, v))
}

// ModuleLT applies the LT predicate on the "module" field.
func ModuleLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldModule, v))
}

// ModuleLTE applies the LTE predicate on the "module" field.
func ModuleLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldModule, v))
}

// ModuleContains applies the Contains predicate on the "module" field.
func ModuleContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldModule, v))
}

// ModuleHasPrefix applies the HasPrefix predicate on the "module" field.
func ModuleHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldModule, v))
}

// ModuleHasSuffix applies the HasSuffix predicate on the "module" field.
func ModuleHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldModule, v))
}

// ModuleEqualFold applies the EqualFold predicate on the "module" field.
func ModuleEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldModule, v))
}

// ModuleContainsFold applies the ContainsFold predicate on the "module" field.
func ModuleContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldModule, v))
}

// FileTypeEQ applies the EQ predicate on the "file_type" field.
func FileTypeEQ(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldFileType, v))
}

// FileTypeNEQ applies the NEQ predicate on the "file_type" field.
func FileTypeNEQ(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldFileType, v))
}

// FileTypeIn applies the In predicate on the "file_type" field.
func FileTypeIn(vs ...uint8) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldFileType, vs...))
}

// FileTypeNotIn applies the NotIn predicate on the "file_type" field.
func FileTypeNotIn(vs ...uint8) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldFileType, vs...))
}

// FileTypeGT applies the GT predicate on the "file_type" field.
func FileTypeGT(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldFileType, v))
}

// FileTypeGTE applies the GTE predicate on the "file_type" field.
func FileTypeGTE(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldFileType, v))
}

// FileTypeLT applies the LT predicate on the "file_type" field.
func FileTypeLT(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldFileType, v))
}

// FileTypeLTE applies the LTE predicate on the "file_type" field.
func FileTypeLTE(v uint8) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldFileType, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldPath, v))
}

// JumpURLEQ applies the EQ predicate on the "jump_url" field.
func JumpURLEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldJumpURL, v))
}

// JumpURLNEQ applies the NEQ predicate on the "jump_url" field.
func JumpURLNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldJumpURL, v))
}

// JumpURLIn applies the In predicate on the "jump_url" field.
func JumpURLIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldJumpURL, vs...))
}

// JumpURLNotIn applies the NotIn predicate on the "jump_url" field.
func JumpURLNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldJumpURL, vs...))
}

// JumpURLGT applies the GT predicate on the "jump_url" field.
func JumpURLGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldJumpURL, v))
}

// JumpURLGTE applies the GTE predicate on the "jump_url" field.
func JumpURLGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldJumpURL, v))
}

// JumpURLLT applies the LT predicate on the "jump_url" field.
func JumpURLLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldJumpURL, v))
}

// JumpURLLTE applies the LTE predicate on the "jump_url" field.
func JumpURLLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldJumpURL, v))
}

// JumpURLContains applies the Contains predicate on the "jump_url" field.
func JumpURLContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldJumpURL, v))
}

// JumpURLHasPrefix applies the HasPrefix predicate on the "jump_url" field.
func JumpURLHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldJumpURL, v))
}

// JumpURLHasSuffix applies the HasSuffix predicate on the "jump_url" field.
func JumpURLHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldJumpURL, v))
}

// JumpURLIsNil applies the IsNil predicate on the "jump_url" field.
func JumpURLIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldJumpURL))
}

// JumpURLNotNil applies the NotNil predicate on the "jump_url" field.
func JumpURLNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldJumpURL))
}

// JumpURLEqualFold applies the EqualFold predicate on the "jump_url" field.
func JumpURLEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldJumpURL, v))
}

// JumpURLContainsFold applies the ContainsFold predicate on the "jump_url" field.
func JumpURLContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldJumpURL, v))
}

// HasFile applies the HasEdge predicate on the "file" edge.
func HasFile() predicate.Banner {
	return predicate.Banner(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFileWith applies the HasEdge predicate on the "file" edge with a given conditions (other predicates).
func HasFileWith(preds ...predicate.File) predicate.Banner {
	return predicate.Banner(func(s *sql.Selector) {
		step := newFileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Banner) predicate.Banner {
	return predicate.Banner(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Banner) predicate.Banner {
	return predicate.Banner(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Banner) predicate.Banner {
	return predicate.Banner(sql.NotPredicates(p))
}