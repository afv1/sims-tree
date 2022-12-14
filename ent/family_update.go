// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/afv1/sims-tree/ent/family"
	"github.com/afv1/sims-tree/ent/predicate"
	"github.com/afv1/sims-tree/ent/relative"
	"github.com/afv1/sims-tree/ent/user"
)

// FamilyUpdate is the builder for updating Family entities.
type FamilyUpdate struct {
	config
	hooks    []Hook
	mutation *FamilyMutation
}

// Where appends a list predicates to the FamilyUpdate builder.
func (fu *FamilyUpdate) Where(ps ...predicate.Family) *FamilyUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetSurname sets the "surname" field.
func (fu *FamilyUpdate) SetSurname(s string) *FamilyUpdate {
	fu.mutation.SetSurname(s)
	return fu
}

// SetDescription sets the "description" field.
func (fu *FamilyUpdate) SetDescription(s string) *FamilyUpdate {
	fu.mutation.SetDescription(s)
	return fu
}

// SetLogoResource sets the "logo_resource" field.
func (fu *FamilyUpdate) SetLogoResource(s string) *FamilyUpdate {
	fu.mutation.SetLogoResource(s)
	return fu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fu *FamilyUpdate) SetOwnerID(id int) *FamilyUpdate {
	fu.mutation.SetOwnerID(id)
	return fu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fu *FamilyUpdate) SetNillableOwnerID(id *int) *FamilyUpdate {
	if id != nil {
		fu = fu.SetOwnerID(*id)
	}
	return fu
}

// SetOwner sets the "owner" edge to the User entity.
func (fu *FamilyUpdate) SetOwner(u *User) *FamilyUpdate {
	return fu.SetOwnerID(u.ID)
}

// AddRelativeIDs adds the "relatives" edge to the Relative entity by IDs.
func (fu *FamilyUpdate) AddRelativeIDs(ids ...int) *FamilyUpdate {
	fu.mutation.AddRelativeIDs(ids...)
	return fu
}

// AddRelatives adds the "relatives" edges to the Relative entity.
func (fu *FamilyUpdate) AddRelatives(r ...*Relative) *FamilyUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return fu.AddRelativeIDs(ids...)
}

// Mutation returns the FamilyMutation object of the builder.
func (fu *FamilyUpdate) Mutation() *FamilyMutation {
	return fu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (fu *FamilyUpdate) ClearOwner() *FamilyUpdate {
	fu.mutation.ClearOwner()
	return fu
}

// ClearRelatives clears all "relatives" edges to the Relative entity.
func (fu *FamilyUpdate) ClearRelatives() *FamilyUpdate {
	fu.mutation.ClearRelatives()
	return fu
}

// RemoveRelativeIDs removes the "relatives" edge to Relative entities by IDs.
func (fu *FamilyUpdate) RemoveRelativeIDs(ids ...int) *FamilyUpdate {
	fu.mutation.RemoveRelativeIDs(ids...)
	return fu
}

// RemoveRelatives removes "relatives" edges to Relative entities.
func (fu *FamilyUpdate) RemoveRelatives(r ...*Relative) *FamilyUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return fu.RemoveRelativeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FamilyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FamilyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FamilyUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FamilyUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FamilyUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FamilyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   family.Table,
			Columns: family.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: family.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Surname(); ok {
		_spec.SetField(family.FieldSurname, field.TypeString, value)
	}
	if value, ok := fu.mutation.Description(); ok {
		_spec.SetField(family.FieldDescription, field.TypeString, value)
	}
	if value, ok := fu.mutation.LogoResource(); ok {
		_spec.SetField(family.FieldLogoResource, field.TypeString, value)
	}
	if fu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   family.OwnerTable,
			Columns: []string{family.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   family.OwnerTable,
			Columns: []string{family.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.RelativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   family.RelativesTable,
			Columns: []string{family.RelativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedRelativesIDs(); len(nodes) > 0 && !fu.mutation.RelativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   family.RelativesTable,
			Columns: []string{family.RelativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RelativesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   family.RelativesTable,
			Columns: []string{family.RelativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{family.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FamilyUpdateOne is the builder for updating a single Family entity.
type FamilyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FamilyMutation
}

// SetSurname sets the "surname" field.
func (fuo *FamilyUpdateOne) SetSurname(s string) *FamilyUpdateOne {
	fuo.mutation.SetSurname(s)
	return fuo
}

// SetDescription sets the "description" field.
func (fuo *FamilyUpdateOne) SetDescription(s string) *FamilyUpdateOne {
	fuo.mutation.SetDescription(s)
	return fuo
}

// SetLogoResource sets the "logo_resource" field.
func (fuo *FamilyUpdateOne) SetLogoResource(s string) *FamilyUpdateOne {
	fuo.mutation.SetLogoResource(s)
	return fuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fuo *FamilyUpdateOne) SetOwnerID(id int) *FamilyUpdateOne {
	fuo.mutation.SetOwnerID(id)
	return fuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fuo *FamilyUpdateOne) SetNillableOwnerID(id *int) *FamilyUpdateOne {
	if id != nil {
		fuo = fuo.SetOwnerID(*id)
	}
	return fuo
}

// SetOwner sets the "owner" edge to the User entity.
func (fuo *FamilyUpdateOne) SetOwner(u *User) *FamilyUpdateOne {
	return fuo.SetOwnerID(u.ID)
}

// AddRelativeIDs adds the "relatives" edge to the Relative entity by IDs.
func (fuo *FamilyUpdateOne) AddRelativeIDs(ids ...int) *FamilyUpdateOne {
	fuo.mutation.AddRelativeIDs(ids...)
	return fuo
}

// AddRelatives adds the "relatives" edges to the Relative entity.
func (fuo *FamilyUpdateOne) AddRelatives(r ...*Relative) *FamilyUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return fuo.AddRelativeIDs(ids...)
}

// Mutation returns the FamilyMutation object of the builder.
func (fuo *FamilyUpdateOne) Mutation() *FamilyMutation {
	return fuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (fuo *FamilyUpdateOne) ClearOwner() *FamilyUpdateOne {
	fuo.mutation.ClearOwner()
	return fuo
}

// ClearRelatives clears all "relatives" edges to the Relative entity.
func (fuo *FamilyUpdateOne) ClearRelatives() *FamilyUpdateOne {
	fuo.mutation.ClearRelatives()
	return fuo
}

// RemoveRelativeIDs removes the "relatives" edge to Relative entities by IDs.
func (fuo *FamilyUpdateOne) RemoveRelativeIDs(ids ...int) *FamilyUpdateOne {
	fuo.mutation.RemoveRelativeIDs(ids...)
	return fuo
}

// RemoveRelatives removes "relatives" edges to Relative entities.
func (fuo *FamilyUpdateOne) RemoveRelatives(r ...*Relative) *FamilyUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return fuo.RemoveRelativeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FamilyUpdateOne) Select(field string, fields ...string) *FamilyUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Family entity.
func (fuo *FamilyUpdateOne) Save(ctx context.Context) (*Family, error) {
	var (
		err  error
		node *Family
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FamilyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Family)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FamilyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FamilyUpdateOne) SaveX(ctx context.Context) *Family {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FamilyUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FamilyUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FamilyUpdateOne) sqlSave(ctx context.Context) (_node *Family, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   family.Table,
			Columns: family.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: family.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Family.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, family.FieldID)
		for _, f := range fields {
			if !family.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != family.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Surname(); ok {
		_spec.SetField(family.FieldSurname, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Description(); ok {
		_spec.SetField(family.FieldDescription, field.TypeString, value)
	}
	if value, ok := fuo.mutation.LogoResource(); ok {
		_spec.SetField(family.FieldLogoResource, field.TypeString, value)
	}
	if fuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   family.OwnerTable,
			Columns: []string{family.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   family.OwnerTable,
			Columns: []string{family.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.RelativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   family.RelativesTable,
			Columns: []string{family.RelativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedRelativesIDs(); len(nodes) > 0 && !fuo.mutation.RelativesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   family.RelativesTable,
			Columns: []string{family.RelativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RelativesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   family.RelativesTable,
			Columns: []string{family.RelativesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relative.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Family{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{family.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
